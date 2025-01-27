
FROM alpine:latest

RUN mkdir /app

COPY ./build_app/listenerApp /app

CMD [ "/app/listenerApp" ]