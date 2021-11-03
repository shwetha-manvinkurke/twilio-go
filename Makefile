.PHONY: githooks install test goimports govet golint docker-build docker-push cover

githooks:
	ln -sf ../../githooks/pre-commit .git/hooks/pre-commit

install: githooks
	go build -v ./...

test:
	go test ./...

cluster-test:
	go test --tags=cluster

goimports:
	go get golang.org/x/tools/cmd/goimports
	goimports -w .

govet: goimports
	go vet ./...

golint: govet
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0
	golangci-lint run
	go mod tidy

export API_DEFINITIONS_SHA=$(shell git log --oneline | grep Regenerated | head -n1 | cut -d ' ' -f 5)
docker-build:
	export API_DEFINITIONS_SHA=$(git log --oneline | grep Regenerated | head -n1 | cut -d ' ' -f 5)
	echo ${API_DEFINITIONS_SHA}
	docker build -t twilio/twilio-go .
	docker tag twilio/twilio-go smanuvin/twilio-go:${GITHUB_TAG}
	docker tag twilio/twilio-go smanuvin/twilio-go:apidefs-${API_DEFINITIONS_SHA}
	docker tag twilio/twilio-go smanuvin/twilio-go:latest

docker-push:
	export API_DEFINITIONS_SHA=$(shell git log --oneline | grep Regenerated | head -n1 | cut -d ' ' -f 5)
	docker push smanuvin/twilio-go:${GITHUB_TAG}
	docker push smanuvin/twilio-go:apidefs-${API_DEFINITIONS_SHA}
	docker push smanuvin/twilio-go:latest

GO_DIRS = $(shell go list ./... | grep -v /rest/ | grep -v /form )
cover:
	go test ${GO_DIRS} -coverprofile coverage.out
	go test ${GO_DIRS} -json > test-report.out


testme:
	echo ${GITHUB_TAG}
	echo ${API_DEFINITIONS_SHA}