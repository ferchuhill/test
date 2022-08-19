test:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -v ./client/testing

testEnd:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -cover -v ./client/testing/endtoend

testDocker:
	docker-compose up --build | grep -e form3-tests_1 -e form3-tests-e2e_1