#!/bin/bash -e

oapi-codegen --config ../openapi-codegen-config.yml ../schema/openapi.yml > ../src/rest/generated.go
