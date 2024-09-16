from node:20

WORKDIR /usr/src/app

RUN apt update 

COPY package*.json ./

COPY . .

RUN npm install

RUN npm run build 



