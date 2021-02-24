FROM golang:1.15.8 AS buildstage

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go build

FROM golang:1.15.8

EXPOSE 8081

RUN mkdir /app

COPY  --from=buildstage ./app/cdwebapp /app

WORKDIR /app

ENTRYPOINT ["./cdwebapp"]