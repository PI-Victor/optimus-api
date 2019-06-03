Optimus API
---
[![Build Status](https://travis-ci.org/cloudflavor/optimus-api.svg?branch=master)](https://travis-ci.org/cloudflavor/optimus-api)
[![GolangCI](https://golangci.com/badges/github.com/cloudflavor/optimus-api.svg)](https://golangci.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflavor/optimus-api)](https://goreportcard.com/report/github.com/cloudflavor/optimus-api)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/cloudflavor/optimus-api/master/LICENSE)
[![codecov](https://codecov.io/gh/cloudflavor/optimus-api/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudflavor/optimus-api)  
This is the REST API that powers the Optimus CI/CD platform.  

Running the API:

First run a Postgres docker container instance:
```bash
hack/run-postgresql.sh
...
# tail the logs for issues
docker logs -f optimus-db
```

Then run the API in a container and link the database container instance.
```bash
hack/run-in-docker.sh
# create a new self-signed certificate to be able to start the API.
root@bcfd9ce0c4dc:/go/src/github.com/cloudflavor/optimus-api# openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem
Generating a RSA private key
.....................+++++
................+++++
writing new private key to 'key.pem'
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:xx
State or Province Name (full name) [Some-State]:xx
Locality Name (eg, city) []:xxxxxxx
Organization Name (eg, company) [Internet Widgits Pty Ltd]:xxxxxx
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:xxxxxx.xxxxx.xxx
Email Address []:xxxx@xxx.xxx

# point the env variables to the TLS cert path and start the server:
export OPTIMUS_LOG_LEVEL=7
export OPTIMUS_SSL_CERT_PATH="/go/src/github.com/cloudflavor/optimus-api/certificate.pem"
export OPTIMUS_SSL_CERT_KEY_PATH="/go/src/github.com/cloudflavor/optimus-api/key.pem"
go run cmd/api/main.go
```
Index will return the current available API versions:
```bash
 curl -kv https://127.0.0.1:8000/

* Connection state changed (MAX_CONCURRENT_STREAMS == 250)!
< HTTP/2 200
< content-type: application/json
< content-length: 25
< date: Mon, 03 Jun 2019 21:50:31 GMT
<
[{"version":"v1alpha1"}]
* Connection #0 to host 127.0.0.1 left intact

```


Running the tests:

```bash
# Inside the docker container created above
go get github.com/smartystreets/goconvey

make test
....
4 total assertions

--- PASS: TestMiddlewareFcuntionality (0.00s)
PASS
coverage: 90.0% of statements
ok  	github.com/cloudflavor/optimus-api/pkg/middleware	1.014s	coverage: 90.0% of statements
```
