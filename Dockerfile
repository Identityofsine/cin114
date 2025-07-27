FROM node:20

# Set timezone to Eastern Time
ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Set build-time arguments
ARG BUILD_DATE
ARG BUILD_ID
ARG NEXT_PUBLIC_BRANCH
ARG NEXT_PUBLIC_API_BASE_URL

# Set as environment variables (available during build and runtime)
ENV BUILD_DATE=$BUILD_DATE
ENV BUILD_ID=$BUILD_ID
ENV NEXT_PUBLIC_BRANCH=$NEXT_PUBLIC_BRANCH
ENV NEXT_PUBLIC_API_BASE_URL=$NEXT_PUBLIC_API_BASE_URL

WORKDIR /usr/src/app

RUN apt update --allow-insecure-repositories 

COPY package*.json ./

COPY . .

RUN npm install

RUN npm run build 

# You can verify the variables are set by adding:
RUN bash -c 'echo -e "Build Date: $BUILD_DATE" && echo -e "Commit Hash: $BUILD_ID" && echo -e "Branch: $NEXT_PUBLIC_BRANCH"'

CMD ["npm", "start"]
