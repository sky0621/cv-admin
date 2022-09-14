#!/bin/bash -e

oapi-codegen --config ../openapi-codegen-config.yml ../schema/openapi.yml > ../src/interfaces/swagger/generated.go
