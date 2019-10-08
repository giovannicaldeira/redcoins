FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD ./redcoins /app
WORKDIR /app
EXPOSE 8080
CMD ["/app/redcoins"]