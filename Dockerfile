FROM ubuntu:20.04

EXPOSE 8081

WORKDIR /app

COPY  ./cdwebapp .

RUN chmod +x ./cdwebapp

ENTRYPOINT ["./cdwebapp"]
