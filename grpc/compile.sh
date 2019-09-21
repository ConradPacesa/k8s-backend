#! /bin/bash

## This file compiles the protocol buffers for each 

protoc -I item/ item/item.proto --go_out=plugins=grpc:item

cp ./item/item.pb.go ../api-gateway/grpc/item.pb.go
cp ./item/item.pb.go ../item-service/grpc/item.pb.go