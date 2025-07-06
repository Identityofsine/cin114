#!/bin/bash

# Quick Development Workflow
# This script builds, pushes, and deploys in one command for rapid iteration

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

show_help() {
    echo -e "${BLUE}🚀 Quick Development Workflow${NC}"
    echo ""
    echo "This script builds, pushes, and deploys in one command for rapid iteration."
    echo ""
    echo "Usage: $0 [environment] [options]"
    echo ""
    echo "Environments:"
    echo "  dev     - Build, push, and deploy to development"
    echo "  prod    - Build, push, and deploy to production"
    echo ""
    echo "Options:"
    echo "  --frontend-only    - Only build/deploy frontend"
    echo "  --api-only        - Only build/deploy API"
    echo "  --no-cache        - Build without using cache"
    echo "  --build-only      - Only build and push, don't deploy"
    echo ""
    echo "Examples:"
    echo "  $0 dev                    # Full dev workflow"
    echo "  $0 prod --frontend-only   # Deploy only frontend to prod"
    echo "  $0 dev --no-cache         # Dev workflow without cache"
    echo "  $0 dev --build-only       # Build and push dev images only"
    echo ""
    echo "What this script does:"
    echo "  1. 🔨 Builds Docker images locally"
    echo "  2. 📦 Pushes images to GitHub Container Registry"
    echo "  3. 🚀 Deploys to your server"
    echo ""
    echo "Perfect for rapid development iteration! 🎯"
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

main() {
    local environment=""
    local build_options=()
    local deploy_only=false
    
    # Parse arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            dev|prod)
                environment=$1
                shift
                ;;
            --frontend-only|--api-only|--no-cache)
                build_options+=("$1")
                shift
                ;;
            --build-only)
                deploy_only=true
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
    
    # Validate environment
    if [[ "$environment" != "dev" && "$environment" != "prod" ]]; then
        log_error "Environment must be 'dev' or 'prod'"
        show_help
        exit 1
    fi
    
    echo -e "${BLUE}🚀 Starting Quick Development Workflow for $environment${NC}"
    echo ""
    
    # Step 1: Build and Push
    log_info "Step 1: Building and pushing images..."
    ./build-and-push.sh "$environment" "${build_options[@]}"
    
    if [ "$deploy_only" = false ]; then
        echo ""
        # Step 2: Deploy
        log_info "Step 2: Deploying to server..."
        ./deploy-new.sh "$environment"
        
        echo ""
        log_success "🎉 Complete workflow finished!"
        
        if [ "$environment" = "dev" ]; then
            echo -e "${GREEN}🌐 Development site: https://dev.cin114.net${NC}"
            echo -e "${GREEN}🔌 Development API: https://api.dev.cin114.net${NC}"
        else
            echo -e "${GREEN}🌐 Production site: https://cin114.net${NC}"
            echo -e "${GREEN}🔌 Production API: https://api.cin114.net${NC}"
        fi
    else
        echo ""
        log_success "🎉 Build and push completed!"
        log_info "Run './deploy-new.sh $environment' when you're ready to deploy"
    fi
}

# Run main function with all arguments
main "$@" 