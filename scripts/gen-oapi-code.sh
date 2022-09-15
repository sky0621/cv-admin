#!/bin/bash -e

oapi-codegen --config ../openapi-codegen-config.yml ../schema/openapi.yml > ../src/adapter/controller/swagger/generated.go
