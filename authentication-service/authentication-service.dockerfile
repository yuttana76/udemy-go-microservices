FROM alpine:latest

RUN mkdir /app

COPY ./build_app/authApp /app

CMD [ "/app/authApp"]