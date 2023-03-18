FROM golang:1.19-alpine
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

#RUN go mod tidy

RUN go build -o /remood-backend-exec

EXPOSE 8080
CMD [ "/remood-backend-exec" ]

