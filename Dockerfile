FROM golang:alphine AS builder
WORKDIR /src
COPY ./src /src
RUN cd ./src  && go build -o genius

FROM alphine
WORKDIR /app
COPY --from=builder /src/genius /app

EXPOSE 80
ENTRYPOINT ./genius