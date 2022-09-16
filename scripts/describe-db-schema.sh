#!/bin/bash -e

cd ../src/adapter/gateway
go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
