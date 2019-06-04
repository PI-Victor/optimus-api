#!/bin/bash
echo "docker logs -f optimus-db"
docker run -p 5432:5432 --name optimus-db -e POSTGRES_USER=tac -e POSTGRES_DB=optimus-db -e POSTGRES_PASSWORD=test -d postgres:11.1
