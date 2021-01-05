FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY dist .

#EXPOSE 1323

ENTRYPOINT ["/app/kinsta"]