FROM node:18.15-alpine 

WORKDIR /app

COPY package.json .

RUN npm i

COPY . .

CMD ["npm", "run", "dev" ]
