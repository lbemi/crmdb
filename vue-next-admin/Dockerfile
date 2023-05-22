FROM node:20.1.0-alpine3.17 as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN npm config set registry https://registry.npmmirror.com && npm install && npm run build


FROM  nginx:1.24.0-alpine as runner
COPY --from=builder /app/dist /usr/share/nginx/html
