#!/bin/bash

docker run --name optimus-openapi-gen --rm -ti -v ${PWD}:/local -w /local openapitools/openapi-generator-cli $1
