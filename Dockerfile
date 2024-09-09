from node:20

WORKDIR /usr/src/app

RUN apt update 

COPY package*.json ./

COPY . .

RUN npm install

EXPOSE 80
EXPOSE 443

CMD ["npm", "run", "dev"]



