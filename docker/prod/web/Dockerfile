FROM node:18-alpine

WORKDIR /web
COPY package.json .
RUN yarn install

COPY . .

CMD [ "yarn", "dev" ]