# Idea taken from here:
# https://github.com/docker-library/docs/blob/6b6b3f34023ab90821453ed1e88e7e9165c6b0d1/.template-helpers/variant-onbuild.md

FROM golang:1.14.2

RUN apt-get update && apt-get install -y nano vim &&  \
  git config --global user.name "ign-webserver"  &&  \
  git config --global user.email "ign-webserver@test.org"

COPY . /root/web-server
WORKDIR /root/web-server

# Install documentation
RUN git clone https://github.com/ignitionrobotics/docs -b master

# Build app
RUN go build
CMD ["./web-server"]

EXPOSE 8000
