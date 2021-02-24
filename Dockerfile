FROM golang:1.15.8

EXPOSE 8081

RUN mkdir /app

COPY  ./cdwebapp /app

WORKDIR /app

RUN chmod +x ./cdwebapp

ENTRYPOINT ["./cdwebapp"]