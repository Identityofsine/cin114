# Specifies a parent image
FROM golang:1.23.1-bookworm

# Build arguments for dynamic configuration
ARG BUILD_DATE
ARG BUILD_ID
ARG GIT_COMMIT
ARG GIT_BRANCH
ARG VERSION=0.0.1
ARG STRIPE_SECRET_KEY
ARG STRIPE_WEBHOOK_SECRET
ARG STRIPE_REDIRECT_URL

# Set as environment variables (available during build and runtime)
ENV BUILD_DATE=$BUILD_DATE
ENV BUILD_ID=$BUILD_ID
ENV GIT_COMMIT=$GIT_COMMIT
ENV GIT_BRANCH=$GIT_BRANCH
ENV VERSION=$VERSION
ENV STRIPE_SECRET_KEY=$STRIPE_SECRET_KEY
ENV STRIPE_WEBHOOK_SECRET=$STRIPE_WEBHOOK_SECRET
ENV STRIPE_REDIRECT_URL=$STRIPE_REDIRECT_URL

# Creates an app directory to hold your app's source code
WORKDIR /usr/src/app

RUN apt-get update

# Installs Node.js and npm
RUN apt-get install -y npm 
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
RUN . ~/.nvm/nvm.sh && nvm install 20.0.0

# Installs nodemon
RUN npm install -g nodemon

# Copies everything from your root directory into /app
COPY . .

# Debug: Print build arguments to verify they're set
RUN echo "Build Date: $BUILD_DATE" && \
    echo "Build ID: $BUILD_ID" && \
    echo "Git Commit: $GIT_COMMIT" && \
    echo "Git Branch: $GIT_BRANCH" && \
    echo "Version: $VERSION" 

# Generate server.yaml with build-time information (remove existing file first)
RUN rm -f config/server.yaml && \
    echo "version: \"${VERSION}\"" > config/server.yaml && \
    echo "server_name: \"cin114-backend\"" >> config/server.yaml && \
    echo "iteration: 1" >> config/server.yaml && \
    echo "commit: \"${GIT_COMMIT:-unknown}\"" >> config/server.yaml && \
    echo "branch: \"${GIT_BRANCH:-unknown}\"" >> config/server.yaml && \
    echo "build_date: \"${BUILD_DATE:-unknown}\"" >> config/server.yaml && \
    echo "build_id: \"${BUILD_ID:-unknown}\"" >> config/server.yaml

# Debug: Show the generated server.yaml file
RUN cat config/server.yaml

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o ./app ./cmd/server/main.go  

# Tells Docker which network port your container listens on
EXPOSE 3001 

# Specifies the executable command that runs when the container starts
CMD ["./app"]
