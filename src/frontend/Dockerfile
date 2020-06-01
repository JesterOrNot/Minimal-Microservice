FROM node:12.17.0-stretch-slim

WORKDIR /application

COPY . ./

RUN npm install && npm run build

CMD ["npm", "start"]
