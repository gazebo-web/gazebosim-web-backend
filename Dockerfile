# Idea taken from here:
# https://github.com/docker-library/docs/blob/6b6b3f34023ab90821453ed1e88e7e9165c6b0d1/.template-helpers/variant-onbuild.md

FROM golang:1.14.2

RUN apt-get update && apt-get install -y nano vim &&  \
  git config --global user.name "gz-webserver"  &&  \
  git config --global user.email "gz-webserver@test.org"

COPY . /root/gazebosim-web-backend
WORKDIR /root/gazebosim-web-backend

# Install documentation
RUN git clone https://github.com/ignitionrobotics/docs

# Build app
RUN go build
CMD ["./gazebosim-web-backend"]

EXPOSE 8000
