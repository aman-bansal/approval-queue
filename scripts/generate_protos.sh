#!/bin/bash

for proto in `ls $GOPATH/src/github.com/aman-bansal/approval-queue/pkg/proto | tr -s '\n'`
do
    echo "$proto"
    docker run --rm -v $(pwd):$(pwd) -w $(pwd)/proto_files znly/protoc:0.2.0 --grpc-gateway_out=logtostderr=true:$GOPATH/src/. --go_out=plugins=grpc:$GOPATH/src/. -I$GOPATH/src/github.com/aman-bansal/approval-queue/pkg/proto $proto
done
