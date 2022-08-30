FROM node:18-alpine AS builder

WORKDIR /web
COPY package.json .
RUN yarn install
COPY . .
RUN yarn build

FROM nginx:1-alpine
COPY --from=builder /web/dist /dist
COPY nginx.conf /etc/nginx/nginx.conf