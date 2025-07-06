          #!/bin/bash
          set -e  # Exit immediately if a command exits with a non-zero status
          set -x  # Print each command before executing it
          # Set environment variables for deployment
          export NEXT_PUBLIC_BRANCH="dev"
          export NODE_ENV="development"
          export BUILD_DATE="$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
          export BUILD_ID="$(git rev-parse --short HEAD 2>/dev/null || echo 'unknown')"
          # Navigate to the homebrew directory
          cd ~/homebrew || exit 1
          # Run the deployment script with the 'dev' argument
          ./deploy.sh dev
