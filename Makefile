SHELL = /bin/bash 

TARGET = lot
VERSION = v0.0.1
GIT_REPO_PATH = $(shell git remote -v | grep fetch | awk '{print $$2}'| sed -e 's/^git@//' | sed -e 's/.git$$//' | sed -e 's/:/\//')

DOCKER_BIN = $(shell which docker) 
DOCKER_RUN = $(DOCKER_BIN) run --rm
DOCKER_FLAGS = --volume $(GO_SRC):/go/src \
		   	   --volume $(shell pwd):/go/src/$(GIT_REPO_PATH) \
		       --workdir /go/src/$(GIT_REPO_PATH)

GO_IMAGE = golang
GO_SRC = $(shell [ -n $$GOPATH/src ] && echo $$GOPATH/src || $(DOCKER_BIN) volume create go_src)
GO_FLAGS = -v
GO_OS = darwin linux
GO_ARCH = 386 amd64 arm arm64

build: install
	$(DOCKER_RUN) \
		$(DOCKER_FLAGS) \
		$(GO_IMAGE) \
			go build $(GO_FLAGS) -o main .

release: install
	$(foreach GOOS,$(GO_OS),\
		$(foreach GOARCH,$(GO_ARCH),\
			$(DOCKER_RUN) \
				$(DOCKER_FLAGS) \
				-e GOOS=$(GOOS) \
				-e GOARCH=$(GOARCH) \
				$(GO_IMAGE) \
					go build $(GO_FLAGS) -o build/$(TARGET)-$(VERSION)-$(GOOS)-$(GOARCH) .;\
		)\
	)
	file build/*
	
install: 
	$(DOCKER_RUN) $(DOCKER_FLAGS) $(GO_IMAGE) go get	

shell:
	$(DOCKER_RUN) $(DOCKER_FLAGS) -it $(GO_IMAGE) bash
