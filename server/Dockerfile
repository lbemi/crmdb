FROM golang:1.19-alpine as builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o lbemi-server ./cmd && chmod +x /server/lbemi-server


FROM busybox as runner
WORKDIR /app
COPY --from=builder /app/lbemi-server /app
COPY --from=builder /app/config.yaml /app
ENTRYPOINT ["./lbemi-server"]