FROM golang:1.18-alpine as build

RUN apk --no-cache add git

WORKDIR /app
COPY . .
RUN go build

FROM alpine:3.15.0
RUN apk --no-cache add postgresql14-client=14.5-r0
WORKDIR /app
COPY --from=build /app/cdwebapp .
EXPOSE 8081

ENTRYPOINT ["./cdwebapp"]
