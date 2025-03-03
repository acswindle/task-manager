FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY database ./database
COPY internal ./internal

RUN go build -o /task-manager

CMD [ "/task-manager" ]
