#!/bin/bash -e

cd ../src
go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
