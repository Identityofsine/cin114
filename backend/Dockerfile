# Specifies a parent image
FROM golang:1.23.1-bookworm

# Set timezone to Eastern Time
ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Build arguments for dynamic configuration
ARG BUILD_DATE
ARG BUILD_ID
ARG GIT_COMMIT
ARG GIT_BRANCH
ARG VERSION=0.0.1
ARG STRIPE_SECRET_KEY
ARG STRIPE_WEBHOOK_SECRET
ARG STRIPE_REDIRECT_URL
ARG SMTP_SERVER
ARG SMTP_PORT
ARG SMTP_USERNAME
ARG SMTP_PASSWORD
ARG SMTP_FROM_EMAIL
ARG SMTP_FROM_NAME

# Set as environment variables (available during build and runtime)
ENV BUILD_DATE=$BUILD_DATE
ENV BUILD_ID=$BUILD_ID
ENV GIT_COMMIT=$GIT_COMMIT
ENV GIT_BRANCH=$GIT_BRANCH
ENV VERSION=$VERSION
ENV STRIPE_SECRET_KEY=$STRIPE_SECRET_KEY
ENV STRIPE_WEBHOOK_SECRET=$STRIPE_WEBHOOK_SECRET
ENV STRIPE_REDIRECT_URL=$STRIPE_REDIRECT_URL
ENV SMTP_SERVER=$SMTP_SERVER
ENV SMTP_PORT=$SMTP_PORT
ENV SMTP_USERNAME=$SMTP_USERNAME
ENV SMTP_PASSWORD=$SMTP_PASSWORD
ENV SMTP_FROM_EMAIL=$SMTP_FROM_EMAIL
ENV SMTP_FROM_NAME=$SMTP_FROM_NAME

# Creates an app directory to hold your app's source code
WORKDIR /usr/src/app

RUN apt-get update --allow-insecure-repositories

# Installs Node.js and npm
RUN apt install -y npm 
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

# Generate stripe.yaml with build-time Stripe configuration
RUN rm -f config/stripe.yaml && \
    echo "# Stripe Configuration (Generated at build time)" > config/stripe.yaml && \
    echo "stripe_secret_key: \"${STRIPE_SECRET_KEY:-}\"" >> config/stripe.yaml && \
    echo "webhook_secret: \"${STRIPE_WEBHOOK_SECRET:-}\"" >> config/stripe.yaml && \
    echo "stripe_redirect_url: \"${STRIPE_REDIRECT_URL:-}\"" >> config/stripe.yaml

# Generate mail.yaml with build-time SMTP configuration
RUN rm -f config/mail.yaml && \
    echo "# Mail Configuration (Generated at build time)" > config/mail.yaml && \
    echo "smtp_server: \"${SMTP_SERVER:-smtp.gmail.com}\"" >> config/mail.yaml && \
    echo "smtp_port: ${SMTP_PORT:-587}" >> config/mail.yaml && \
    echo "smtp_username: \"${SMTP_USERNAME:-}\"" >> config/mail.yaml && \
    echo "smtp_password: \"${SMTP_PASSWORD:-}\"" >> config/mail.yaml && \
    echo "from_email: \"${SMTP_FROM_EMAIL:-}\"" >> config/mail.yaml && \
    echo "from_name: \"${SMTP_FROM_NAME:-Cin114 Tickets}\"" >> config/mail.yaml

# Debug: Show the generated files
RUN cat config/server.yaml && echo "---" && cat config/stripe.yaml && echo "---" && cat config/mail.yaml

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o ./app ./cmd/server/main.go  

# Tells Docker which network port your container listens on
EXPOSE 3001 

# Specifies the executable command that runs when the container starts
CMD ["./app"]
