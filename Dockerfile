FROM golang:1.15-buster as builder

WORKDIR /app

COPY go.mod go.sum ./
#COPY compress@v1.11.9 /go/pkg/mod/github.com/klauspost
#COPY compress@v1.12.2 /go/pkg/mod/github.com/klauspost
#COPY compress@v1.13.0 /go/pkg/mod/github.com/klauspost

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o courier-management .
#RUN apk --update add redis
#RUN apk add tcpdump
#RUN apk add busybox-extras 

FROM alpine:3.11.6

WORKDIR /app/
ENV APP_ENV ${APP_ENV}

COPY --from=builder /app/courier-management ./
COPY --from=builder /app/config* ./
COPY --from=builder /app/keys/ ./
COPY --from=builder /app/open-app.html ./
COPY --from=builder /app/docker-entrypoint.sh ./

RUN ls ./
USER root
#EXPOSE 8086
ENTRYPOINT ["sh", "/app/docker-entrypoint.sh"]
ENTRYPOINT ["sh", "./docker-entrypoint.sh"]
CMD ["./docker-entrypoint.sh"]
