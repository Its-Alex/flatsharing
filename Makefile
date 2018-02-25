dep:
	go get -u github.com/mattes/migrate
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	cd src/flatsharing && dep ensure -v -vendor-only

build:
	go install flatsharing/...

test:
	go test flatsharing/...

lint:
	gometalinter --config=.gometalinter.json src/flatsharing/...

migrate-db:


.PHONY: build dep test lint