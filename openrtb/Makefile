GO_IMAGE=vungle/golang
DOCKER_GOPATH = $(shell docker run --rm $(GO_IMAGE) /bin/bash -c 'echo $$GOPATH')
GO_WORKDIR = $(DOCKER_GOPATH)/src/github.com/Vungle/vungo/openrtb

genfiles:
	@find . -name "*_easyjson.go" -type f -delete
	@docker run --rm \
		-v $$(pwd):$(GO_WORKDIR) \
		-w $(GO_WORKDIR) \
		$(GO_IMAGE) \
		/bin/bash -c 'go get -u github.com/mailru/easyjson/... && go generate'
