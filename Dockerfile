ARG port

FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

#RUN go build main.go

EXPOSE $port

CMD [ "go", "run", "main.go"]
