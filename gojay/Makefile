.PHONY: build
build: 
	go build ./

.PHONY: test
test: 
	go test -race -v .

.PHONY: cover
cover: 
	go test -tags test -coverprofile=coverage.out

.PHONY: coverhtml
coverhtml: 
	go tool cover -html=coverage.out