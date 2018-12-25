#!/bin/bash
echo "docker logs -f optimus-db"
docker run -p 5432:5432 --name optimus-db -e POSTGRES_PASSWORD=test -e POSTGRES_USERNAME=tac -e POSTGRES_DB=optimus-db -d postgres:11.1
