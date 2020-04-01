# Idea taken from here:
# https://github.com/docker-library/docs/blob/6b6b3f34023ab90821453ed1e88e7e9165c6b0d1/.template-helpers/variant-onbuild.md

FROM golang:1.9.4

RUN apt-get update && apt-get install -y nano vim &&  \
  git config --global user.name "ign-webserver"  &&  \
  git config --global user.email "ign-webserver@test.org"

RUN mkdir -p /go/src/bitbucket.org/ignitionrobotics/ign-webserver
COPY . /go/src/bitbucket.org/ignitionrobotics/ign-webserver
WORKDIR /go/src/bitbucket.org/ignitionrobotics/ign-webserver

# Install go dep
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep
# install the dependencies without checking for go code
RUN dep ensure -vendor-only

# Install documentation
RUN hg clone https://bitbucket.org/ignitionrobotics/docs -b default 

# Build app
RUN go install
CMD ["/go/bin/ign-webserver"]

EXPOSE 8000
