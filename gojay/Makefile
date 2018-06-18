.PHONY: build
build: 
	go build ./

.PHONY: test
test: 
	go test -race .

.PHONY: cover
cover: 
	go test -tags test -coverprofile=profile.out