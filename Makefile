build:
	go install flatsharing/...

dep:
	go get -v flatsharing/...

test:
	go test flatsharing/...

lint:
	golint

.PHONY: dep lint test db