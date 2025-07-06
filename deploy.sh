#!/bin/bash

# Set default values
PORT=3000
COMPOSE_FILE="docker-compose.yaml"
ENVIRONMENT=$1

if [[ "$ENVIRONMENT" == "dev" ]]; then
    TAG="cin114:dev"
    API_TAG="cin114-api:dev"
    NAME="cin114-dev"
    HOSTNAME="devnextjs"
    PORT=3001
    COMPOSE_FILE="docker-compose.dev.yaml"
    export DEPLOY_ENV="dev"
    export IMAGE_TAG="dev"
    echo "Deploying to development environment, $PROJECT_DIR"
elif [[ "$ENVIRONMENT" == "prod" ]]; then
    TAG="cin114:prod"
    API_TAG="cin114-api:prod"
    NAME="cin114-prod"
    HOSTNAME="prodnextjs"
    PROJECT_DIR="$PROJECT_DIR/cin114"
    COMPOSE_FILE="docker-compose.yaml"
    export DEPLOY_ENV="prod"
    export IMAGE_TAG="prod"
    echo "Deploying to production environment, $PROJECT_DIR"
else
    echo "Usage: $0 [dev|prod]"
    exit 1
fi

# Pull the specific pre-built images from the registry
echo "Pulling latest pre-built images..."
docker pull ghcr.io/identityofsine/$TAG
docker pull ghcr.io/identityofsine/$API_TAG

if [ $? -ne 0 ]; then
    echo "Failed to pull the latest images."
    exit 1
fi

# Stop and remove existing containers using docker-compose
echo "Stopping existing containers..."
docker compose -f $COMPOSE_FILE down

# Clean up unused volumes
docker volume prune -f

# Start the services using docker-compose
echo "Starting services with docker-compose..."
docker compose -f $COMPOSE_FILE up -d

if [ $? -ne 0 ]; then
    echo "Failed to start services with docker-compose."
    exit 1
fi

echo "Deployment completed successfully!"