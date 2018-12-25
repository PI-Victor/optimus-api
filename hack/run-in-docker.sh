#!/bin/bash
echo "utility needed: go get -u github.com/golang/dep/cmd/dep"
echo "CERT_PATH=\"/go/src/github.com/cloudflavor/optimus-api/certificate.pem\" KEY_PATH=\"/go/src/github.com/cloudflavor/optimus-api/key.pem\" go run cmd/api/main.go"
echo "openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem"
docker run --rm --name optimus-app -ti --link optimus-db -v $GOPATH:/go -w /go/src/github.com/cloudflavor/optimus-api -p 8000:8000 golang:1.10-stretch /bin/bash
