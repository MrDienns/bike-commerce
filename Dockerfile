FROM golang:latest
EXPOSE 8080
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main", "serve"]