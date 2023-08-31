FROM debian:11

RUN apt-get update && apt-get install -y golang&& apt-get install -y curl \
    && apt-get install -y golang ca-certificates

WORKDIR /app

COPY . .

RUN go build -o app

EXPOSE 8089

CMD ["./app"]