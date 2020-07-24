FROM golang:1.14 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM ubuntu:16.04

RUN apt-get -y update && apt-get --yes install cron libssl-dev

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY docker/conf/stock-cron /etc/cron.d/root

CMD ["cron", "-f"]