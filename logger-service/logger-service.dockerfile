
FROM alpine:latest

RUN mkdir /app

COPY ./build_app/loggerServiceApp /app

CMD [ "/app/loggerServiceApp" ]