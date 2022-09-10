#!/bin/bash -e

oapi-codegen --config ../openapi-codegen-config.yml ../definitions/openapi/openapi.yml > ../generated/swagger/gen.go
