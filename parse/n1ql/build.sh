#! /bin/sh

echo Running nex...
nex n1ql.nex
#echo Running goyacc...
go tool yacc n1ql.y
echo Running go build...
go build
