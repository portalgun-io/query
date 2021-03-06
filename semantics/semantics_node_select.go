//  Copyright (c) 2018 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package semantics

import (
	"github.com/couchbase/query/algebra"
)

func (this *SemChecker) VisitSelectTerm(node *algebra.SelectTerm) (interface{}, error) {
	return node.Select().Accept(this)
}

func (this *SemChecker) VisitSubselect(node *algebra.Subselect) (r interface{}, err error) {
	saveSemFlag := this.semFlag
	defer func() { this.semFlag = saveSemFlag }()
	this.unsetSemFlag(_SEM_WHERE | _SEM_ON)
	if node.With() != nil {
		if err = node.With().MapExpressions(this); err != nil {
			return nil, err
		}
	}

	if node.From() != nil {
		if r, err = node.From().Accept(this); err != nil {
			return r, err
		}
	}

	if node.Let() != nil {
		if err = node.Let().MapExpressions(this); err != nil {
			return nil, err
		}
	}

	if node.Where() != nil {
		this.setSemFlag(_SEM_WHERE)
		_, err = this.Map(node.Where())
		this.unsetSemFlag(_SEM_WHERE)
		if err != nil {
			return nil, err
		}
	}

	if node.Group() != nil {
		if err = node.Group().MapExpressions(this); err != nil {
			return nil, err
		}
	}

	err = node.Projection().MapExpressions(this)

	return nil, err
}
