#!/bin/bash

# Local Build and Push Script
# This script builds Docker images locally and pushes them to GitHub Container Registry
# Use this for faster development iteration without waiting for GitHub Actions

set -e

# Configuration
REGISTRY="ghcr.io"
NAMESPACE="identityofsine"
FRONTEND_IMAGE="$REGISTRY/$NAMESPACE/cin114"
API_IMAGE="$REGISTRY/$NAMESPACE/cin114-api"
export BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
#short git commit hash
export GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "local")
export GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")
export BUILD_ID="${GIT_BRANCH}-${GIT_COMMIT:-local}"
export VERSION="0.0.1"

# Load environment variables from .env file if it exists
if [ -f ".env-$1" ]; then
    source .env-$1
    echo "Loaded environment variables from .env-$1 file"
elif [ -f "backend/.env-$1" ]; then
    source backend/.env-$1
    echo "Loaded environment variables from backend/.env-$1 file"
else
    echo "No .env-$1 file found, using environment variables or defaults"
fi

# Set default redirect URL if not set
export STRIPE_REDIRECT_URL="${STRIPE_REDIRECT_URL:-https://dev.cin114.net/thank-you}"

# Secret validation function
validate_secrets() {
    local env=$1
    local missing_secrets=()
    
    # Check for required secrets based on environment
    if [[ "$env" == "prod" ]]; then
        if [[ -z "${STRIPE_SECRET_KEY:-}" ]]; then
            missing_secrets+=("STRIPE_SECRET_KEY")
        fi
        if [[ -z "${STRIPE_WEBHOOK_SECRET:-}" ]]; then
            missing_secrets+=("STRIPE_WEBHOOK_SECRET")
        fi
        if [[ -z "${STRIPE_REDIRECT_URL:-}" ]]; then
            missing_secrets+=("STRIPE_REDIRECT_URL")
        fi
        if [[ -z "${SMTP_USERNAME:-}" ]]; then
            missing_secrets+=("SMTP_USERNAME")
        fi
        if [[ -z "${SMTP_PASSWORD:-}" ]]; then
            missing_secrets+=("SMTP_PASSWORD")
        fi
        if [[ -z "${SMTP_FROM_EMAIL:-}" ]]; then
            missing_secrets+=("SMTP_FROM_EMAIL")
        fi
    fi
    
    if [[ ${#missing_secrets[@]} -gt 0 ]]; then
        log_warning "Missing required secrets for $env environment: ${missing_secrets[*]}"
        log_warning "These secrets should be set as environment variables before building"
        return 1
    fi
    
    return 0
}

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

show_help() {
    echo -e "${BLUE}🛠️  Local Build and Push Script${NC}"
    echo ""
    echo "Usage: $0 [environment] [options]"
    echo ""
    echo "Environments:"
    echo "  dev     - Build and push development images"
    echo "  prod    - Build and push production images"
    echo ""
    echo "Options:"
    echo "  --frontend-only    - Only build and push frontend"
    echo "  --api-only        - Only build and push API"
    echo "  --no-cache        - Build without using cache"
    echo "  --no-push         - Build but don't push to registry"
    echo ""
    echo "Examples:"
    echo "  $0 dev                    # Build and push both dev images"
    echo "  $0 prod --frontend-only   # Build and push only prod frontend"
    echo "  $0 dev --no-cache         # Build dev images without cache"
    echo "  $0 dev --no-push          # Build dev images but don't push"
    echo ""
    echo "Prerequisites:"
    echo "  - Docker installed and running"
    echo "  - Authenticated with GitHub Container Registry"
    echo "  - Source code available in ../cin114/"
}

log_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

log_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

log_error() {
    echo -e "${RED}❌ $1${NC}"
}

check_prerequisites() {
    log_info "Checking prerequisites..."
    
    # Check if Docker is running
    if ! docker info >/dev/null 2>&1; then
        log_error "Docker is not running. Please start Docker and try again."
        exit 1
    fi
    
    # Check if we're in the right directory (should have backend/ and src/ directories)
    if [ ! -d "backend" ] || [ ! -d "src" ]; then
        log_error "This script should be run from the cin114 root directory."
        log_error "Expected to find backend/ and src/ directories here."
        exit 1
    fi
    
    # Check if we can access the registry
    if ! docker pull hello-world >/dev/null 2>&1; then
        log_warning "Docker pull test failed. You might need to authenticate with the registry."
    fi
    
    log_success "Prerequisites check passed"
}

authenticate_registry() {
    log_info "Checking registry authentication..."
    
    # Try to login to GitHub Container Registry
    if ! docker login $REGISTRY -u $(whoami) 2>/dev/null; then
        log_warning "Not authenticated with GitHub Container Registry."
        echo ""
        echo "To authenticate, you need a GitHub Personal Access Token with 'write:packages' permission."
        echo "Run: docker login $REGISTRY -u YOUR_GITHUB_USERNAME"
        echo "Then enter your Personal Access Token as the password."
        echo ""
        read -p "Do you want to authenticate now? (y/N): " -r
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            docker login $REGISTRY
        else
            log_error "Authentication required to push images."
            exit 1
        fi
    fi
    
    log_success "Registry authentication verified"
}

build_frontend() {
    local env=$1
    local tag=$2
    local cache_option=$3
    
    log_info "Building frontend image for $env environment..."
    
    # Debug: Print build arguments being passed
    echo "BUILD_DATE: $BUILD_DATE"
    echo "BUILD_ID: $BUILD_ID"
    echo "GIT_COMMIT: $GIT_COMMIT"
    echo "GIT_BRANCH: $GIT_BRANCH"
    echo "VERSION: $VERSION"
    echo "NEXT_PUBLIC_BRANCH: $env"

    # Build the frontend image
    docker build \
        $cache_option \
        -t $FRONTEND_IMAGE:$tag \
        --build-arg BUILD_DATE="$BUILD_DATE" \
        --build-arg BUILD_ID="$BUILD_ID" \
        --build-arg GIT_COMMIT="$GIT_COMMIT" \
        --build-arg GIT_BRANCH="$GIT_BRANCH" \
        --build-arg VERSION="$VERSION" \
        --build-arg NEXT_PUBLIC_BRANCH="$env" \
        -f ../cin114/Dockerfile \
        ../cin114/
    
    log_success "Frontend image built: $FRONTEND_IMAGE:$tag"
}

build_api() {
    local env=$1
    local tag=$2
    local cache_option=$3
    
    log_info "Building API image for $env environment..."
    
    # Debug: Print build arguments being passed
    echo "BUILD_DATE: $BUILD_DATE"
    echo "BUILD_ID: $BUILD_ID"
    echo "GIT_COMMIT: $GIT_COMMIT"
    echo "GIT_BRANCH: $GIT_BRANCH"
    echo "VERSION: $VERSION"
    echo "STRIPE_SECRET_KEY: ${STRIPE_SECRET_KEY:-not set}"
    echo "STRIPE_WEBHOOK_SECRET: ${STRIPE_WEBHOOK_SECRET:-not set}"
    echo "STRIPE_REDIRECT_URL: $STRIPE_REDIRECT_URL"
    echo "SMTP_SERVER: ${SMTP_SERVER:-smtp.gmail.com}"
    echo "SMTP_PORT: ${SMTP_PORT:-587}"
    echo "SMTP_USERNAME: ${SMTP_USERNAME:-not set}"
    echo "SMTP_PASSWORD: ${SMTP_PASSWORD:-not set}"
    echo "SMTP_FROM_EMAIL: ${SMTP_FROM_EMAIL:-not set}"
    echo "SMTP_FROM_NAME: ${SMTP_FROM_NAME:-Cin114 Tickets}"

    # Build the API image
    docker build \
        $cache_option \
        -t $API_IMAGE:$tag \
        --build-arg BUILD_DATE="$BUILD_DATE" \
        --build-arg BUILD_ID="$BUILD_ID" \
        --build-arg GIT_COMMIT="$GIT_COMMIT" \
        --build-arg GIT_BRANCH="$GIT_BRANCH" \
        --build-arg VERSION="$VERSION" \
        --build-arg STRIPE_SECRET_KEY="${STRIPE_SECRET_KEY:-}" \
        --build-arg STRIPE_WEBHOOK_SECRET="${STRIPE_WEBHOOK_SECRET:-}" \
        --build-arg STRIPE_REDIRECT_URL="$STRIPE_REDIRECT_URL" \
        --build-arg SMTP_SERVER="${SMTP_SERVER:-smtp.gmail.com}" \
        --build-arg SMTP_PORT="${SMTP_PORT:-587}" \
        --build-arg SMTP_USERNAME="${SMTP_USERNAME:-}" \
        --build-arg SMTP_PASSWORD="${SMTP_PASSWORD:-}" \
        --build-arg SMTP_FROM_EMAIL="${SMTP_FROM_EMAIL:-}" \
        --build-arg SMTP_FROM_NAME="${SMTP_FROM_NAME:-Cin114 Tickets}" \
        -f ../cin114/backend/Dockerfile \
        ../cin114/backend/
    
    log_success "API image built: $API_IMAGE:$tag"
}

push_image() {
    local image=$1
    local tag=$2
    
    log_info "Pushing $image:$tag..."
    docker push $image:$tag
    log_success "Pushed $image:$tag"
}

main() {
    local environment=""
    local build_frontend=true
    local build_api=true
    local use_cache=true
    local push_images=true
    
    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            dev|prod)
                environment=$1
                shift
                ;;
            --frontend-only)
                build_api=false
                shift
                ;;
            --api-only)
                build_frontend=false
                shift
                ;;
            --no-cache)
                use_cache=false
                shift
                ;;
            --no-push)
                push_images=false
                shift
                ;;
            --help|-h)
                show_help
                exit 0
                ;;
            *)
                log_error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done

    log_info "Cleaning up dev-data directory..."

    log_warning "This will remove all data in the dev-data directory!"
    sudo rm -rf /home/kevin/software/cin114/backend/dev-data/


    
    # Validate environment
    if [[ "$environment" != "dev" && "$environment" != "prod" ]]; then
        log_error "Environment must be 'dev' or 'prod'"
        show_help
        exit 1
    fi
    
    # Set cache option
    local cache_option=""
    if [ "$use_cache" = false ]; then
        cache_option="--no-cache"
        log_warning "Building without cache (this will take longer)"
    fi
    
    log_info "Starting build process for $environment environment"
    
    # Check prerequisites
    check_prerequisites
    
    # Validate secrets for production builds
    if [[ "$environment" == "prod" ]]; then
        if ! validate_secrets "$environment"; then
            log_warning "Continuing with build, but secrets should be configured for production"
        fi
    fi
    
    # Authenticate if we're going to push
    if [ "$push_images" = true ]; then
        authenticate_registry
    fi
    
    # Build images
    if [ "$build_frontend" = true ]; then
        build_frontend "$environment" "$environment" "$cache_option"
    fi
    
    if [ "$build_api" = true ]; then
        build_api "$environment" "$environment" "$cache_option"
    fi
    
    # Push images
    if [ "$push_images" = true ]; then
        if [ "$build_frontend" = true ]; then
            push_image "$FRONTEND_IMAGE" "$environment"
        fi
        
        if [ "$build_api" = true ]; then
            push_image "$API_IMAGE" "$environment"
        fi
        
        log_success "All images pushed successfully!"
        echo ""
        log_info "You can now deploy using: ./deploy-new.sh $environment"
    else
        log_success "All images built successfully!"
        echo ""
        log_info "Images are available locally. Use --no-push flag removed to push them."
    fi
}

# Run main function with all arguments
main "$@" 
