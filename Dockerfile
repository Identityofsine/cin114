FROM node:20

# Set build-time arguments
ARG BUILD_DATE
ARG BUILD_ID

# Set as environment variables (available during build and runtime)
ENV BUILD_DATE=$BUILD_DATE
ENV BUILD_ID=$BUILD_ID
ENV NEXT_PUBLIC_BRANCH=$NEXT_PUBLIC_BRANCH

WORKDIR /usr/src/app

RUN apt update 

COPY package*.json ./

COPY . .

RUN npm install

RUN npm run build 

# You can verify the variables are set by adding:
RUN bash -c 'echo -e "Build Date: $BUILD_DATE" && echo -e "Commit Hash: $BUILD_ID" && echo -e "Branch: $NEXT_PUBLIC_BRANCH"'

CMD ["npm", "start"]
