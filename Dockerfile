FROM golang:1.17-alpine

WORKDIR /app
COPY . .
RUN go build

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/cdwebapp .
EXPOSE 8081

ENTRYPOINT ["./cdwebapp"]
