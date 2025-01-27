
FROM alpine:latest

RUN mkdir /app

COPY ./build_app/bankApp /app

CMD [ "/app/bankApp" ]