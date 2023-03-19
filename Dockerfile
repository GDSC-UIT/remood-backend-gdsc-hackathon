FROM golang:1.19-alpine
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o ./remood-backend-exec


EXPOSE 8080
ENTRYPOINT ["./remood-backend-exec"]

