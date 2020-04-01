.PHONY: all

GOPATH:=${PWD}

APP:=bitbucket.org/ignitionrobotics/ign-fuelserver

default: all

all:
	GOPATH=${GOPATH} go get golang.org/x/tools/cmd/cover
	GOPATH=${GOPATH} go get github.com/golang/lint/golint
	GOPATH=${GOPATH} go get -t -v ${APP}
	GOPATH=${GOPATH} go install ${APP}

lint:
	${GOPATH}/bin/golint ${APP}

test:
	GOPATH=${GOPATH} go test ${APP}
