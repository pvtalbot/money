FROM node:lts-alpine

WORKDIR /web
COPY package.json .
RUN yarn install

COPY . .

CMD [ "yarn", "dev" ]