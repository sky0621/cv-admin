.PHONY: redoc-lint
redoc-lint:
	npx @redocly/cli lint ./schema/root.yml

.PHONY: redoc
redoc: redoc-lint
	npx @redocly/cli bundle ./schema/root.yml -o ./schema/openapi.yml

.PHONY: gen
gen: redoc
	go tool oapi-codegen --config=openapi-codegen-config.yml ./schema/openapi.yml > ./src/rest/generated.go

.PHONY: build
build:
	go build -o ./bin/cvadmin ./src/main.go

.PHONY: serve
serve: build
	./bin/cvadmin server

.PHONY: migrate
migrate: build
	./bin/cvadmin migrate

.PHONY: output
output: build
	./bin/cvadmin export --userid 1 --dir /Users/sky0621/work/github.com/sky0621/cv/app2/public/data

.PHONY: submit
submit: build
	./bin/cvadmin submission --userid 1
