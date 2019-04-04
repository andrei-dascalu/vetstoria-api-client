#!/bin/bash

export VETSTORIA_PORT=3100

export VETSTORIA_API_URL="https://virtserver.swaggerhub.com/vetstoria/Booking-API/1.0"
export VETSTORIA_API_SECRET="abc"
export VETSTORIA_API_KEY="123"
export VETSTORIA_API_USERAGENT=""

export VETSTORIA_REDIS_HOST="localhost"
export VETSTORIA_REDIS_PORT="6379"

export VETSTORIA_LOGGER_FORMAT="json"
export VETSTORIA_LOGGER_LEVEL="info"

export VETSTORIA_RABBIT_HOST="localhost"
export VETSTORIA_RABBIT_PORT="5672"

go run cmd/main.go
