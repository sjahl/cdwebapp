FROM ubuntu:20.04

EXPOSE 8082

RUN mkdir /app

COPY  ./cdwebapp /app

WORKDIR /app

RUN chmod +x ./cdwebapp

ENTRYPOINT ["./cdwebapp"]