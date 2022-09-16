#!/bin/bash -e

oapi-codegen --config ../openapi-codegen-config.yml ../schema/openapi.yml > ../src/swagger/generated.go
