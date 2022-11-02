FROM golang:latest
RUN mkdir /app
ENV GO_ENV dev
ADD . /app
WORKDIR /app/cmd
RUN go build -o main .
EXPOSE 4002
CMD ["/app/cmd/main"]