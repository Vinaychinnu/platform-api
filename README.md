# platform-api

platform-api is a minimal backend service written in Go that models
platform-level metadata such as projects, environments, and releases.

The goal of this service is to demonstrate clean backend design,
clear separation of concerns, and production-aware service structure.

## What this service is

- A REST API written in Go
- A single source of truth for deployment metadata
- Structured using clean architectural boundaries

## What this service is NOT

- A Kubernetes controller
- A CI/CD system
- An authentication or authorization service
- A microservices demo or framework showcase

## Project structure

    cmd/api/
        main.go          Entry point for the HTTP service

    internal/
        transport/http/
            router.go    HTTP routing and handlers

    pkg/
        config/
            config.go    Environment-based configuration loading

## Running locally

Start the service:

    go run cmd/api/main.go

The server will start on port 8080 by default.

You can change the port using an environment variable:

    HTTP_PORT=9090 go run cmd/api/main.go

## Health check

Verify the service is running:

    curl http://localhost:8080/health

A successful response returns HTTP 200 with an empty body.
