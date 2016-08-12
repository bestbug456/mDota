# Go parameters
GOCMD=go
GOINSTALL=go build

# Avoid problem with the gopath adding
# temporarely ourself to it
GOPATH := $(CURDIR)

# Create a stand-alone repository
# in order to avoid to download
# untested 3° party dependences

GOPATH := $(CURDIR)/src/_vendor:$(GOPATH)

all:
	cd ./src && $(GOINSTALL) 
	cd ./src && mv src mDota