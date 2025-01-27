
FROM alpine:latest

RUN mkdir /app

COPY ./build_app/brokerApp /app

CMD [ "/app/brokerApp" ]