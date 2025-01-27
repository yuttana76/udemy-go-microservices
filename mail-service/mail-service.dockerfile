
FROM alpine:latest

RUN mkdir /app

COPY ./build_app/mailerApp /app
COPY templates /templates

CMD [ "/app/mailerApp" ]