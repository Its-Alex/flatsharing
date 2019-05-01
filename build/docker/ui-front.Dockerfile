FROM node:12-alpine as builder

COPY internal/ui-front /app

WORKDIR /app

RUN npm install && \
    npm run build

FROM nginx:1.16-alpine

COPY --from=builder /app/build* /usr/share/nginx/html/
COPY --from=builder /app/nginx.conf /etc/nginx/conf.d/default.conf
