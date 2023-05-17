FROM node:18.15-alpine 

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

EXPOSE 7677

# CMD ["npm", "run", "dev --host" ]
