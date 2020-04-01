#!/bin/bash

echo "FROM golang:1.8.3-onbuild" > Dockerfile
echo "EXPOSE 8000" >> Dockerfile

if [[ $# -lt 1 ]]; then
  echo "Creating production Dockerfile"
elif [[ $1 == "-d" ]]; then
  echo "Creating development Dockerfile"
  zip /tmp/ignfuelserver.zip -r *
fi
