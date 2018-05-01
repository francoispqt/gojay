.PHONY: test
test:
	go test -race -run=^Test -v

.PHONY: cover
cover: 
	go test -coverprofile=coverage.out 

.PHONY: coverhtml
coverhtml: 
	go tool cover -html=coverage.out