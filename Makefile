NAME     := custom-emoji-cloner
VERSION  := v0.0.2

REVISION := $(shell git rev-parse --short HEAD)
SRCS     := $(shell find src -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
GOBUILD  = go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o ../$@/$(NAME)-$$GOOS-$$GOARCH

dist: $(SRCS)
	mkdir -p dist
	GOOS=darwin  GOARCH=amd64 sh -ec 'cd src; $(GOBUILD)'
	GOOS=linux   GOARCH=amd64 sh -ec 'cd src; $(GOBUILD)'
	GOOS=linux   GOARCH=arm   sh -ec 'cd src; $(GOBUILD)'
	GOOS=windows GOARCH=amd64 sh -ec 'cd src; $(GOBUILD)'

clean:
	rm -rf dist

get:
	cd src; go get -v
