#!/bin/bash

#docker run --rm -v $(pwd):$(pwd) -w $(pwd)/proto_files znly/protoc:0.2.0 --grpc-gateway_out=logtostderr=true:$GOPATH/src/. --go_out=plugins=grpc:$GOPATH/src/. -I$GOPATH/src/gitlab.com/machodev/user_authorization_manager/pkg/proto_files  user_manager.proto
for proto in `ls $GOPATH/src/github.com/aman-bansal/approval-queue/pkg/proto | tr -s '\n'`
do
    echo "$proto"
    docker run --rm -v $GOPATH/src/github.com/aman-bansal/approval-queue/pkg:$GOPATH/src/github.com/aman-bansal/approval-queue/pkg -w $GOPATH/src/github.com/aman-bansal/approval-queue/pkg/proto znly/protoc:0.2.0 --grpc-gateway_out=logtostderr=true:$GOPATH/src/. --go_out=plugins=grpc:$GOPATH/src/. -I$GOPATH/src/github.com/aman-bansal/approval-queue/pkg/proto $proto
done
