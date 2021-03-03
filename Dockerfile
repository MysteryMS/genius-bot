FROM golang:alpine AS builder
WORKDIR /src
COPY . /src
RUN cd ./src && go build -o genius

FROM alpine
WORKDIR /app
COPY --from=builder /src/src/genius /app

EXPOSE 80
ENTRYPOINT ./genius
