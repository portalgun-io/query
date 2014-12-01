//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package http

import (
	"net/http"
	"time"

	"github.com/couchbaselabs/query/accounting"
	"github.com/couchbaselabs/query/server"
)

type HttpEndpoint struct {
	server  *server.Server
	metrics bool
	httpsrv http.Server
	bufpool BufferPool
}

func NewHttpEndpoint(server *server.Server, metrics bool, addr string) *HttpEndpoint {
	rv := &HttpEndpoint{
		server:  server,
		metrics: metrics,
		bufpool: NewSyncPool(server.KeepAlive()),
	}

	rv.httpsrv.Addr = addr
	rv.httpsrv.Handler = rv

	// Bind HttpEndpoint object to /query/service endpoint; use default Server Mux
	http.Handle("/query/service", rv)

	// TODO: Deprecate (remove) this binding after QE has migrated to /query/service
	http.Handle("/query", rv)

	return rv
}

func (this *HttpEndpoint) ListenAndServe() error {
	return http.ListenAndServe(this.httpsrv.Addr, nil)
}

// If the server channel is full and we are unable to queue a request,
// we respond with a timeout status.
func (this *HttpEndpoint) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	request := newHttpRequest(resp, req, this.bufpool)

	defer this.doStats(request)

	if request.State() == server.FATAL {
		// There was problems creating the request: Fail it and return
		request.Failed(this.server)
		return
	}

	select {
	case this.server.Channel() <- request:
		// Wait until the request exits.
		<-request.CloseNotify()
	default:
		// Timeout.
		resp.WriteHeader(http.StatusServiceUnavailable)
	}

}

func (this *HttpEndpoint) doStats(request *httpRequest) {
	// Update metrics:
	service_time := time.Since(request.ServiceTime())
	request_time := time.Since(request.RequestTime())
	acctstore := this.server.AccountingStore()
	accounting.RecordMetrics(acctstore, request_time, service_time, request.resultCount,
		request.resultSize, request.errorCount, request.warningCount, request.Statement())
}
