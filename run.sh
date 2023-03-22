#!/bin/bash
# Build the app
go build -o rayanesh-backend ./cmd/Backend
# Set the config variables
export REDIS_URL=127.0.0.1:6379
export REDIS_DB=1
export REDIS_USERNAME=
export REDIS_PASSWORD=
export DATABASE_DSN="host=127.0.0.1 user=postgres password=1234 dbname=rayanesh port=5432"
export LISTEN_ADDRESS=127.0.0.1:35892
# Run the app
./rayanesh-backend