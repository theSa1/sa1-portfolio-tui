FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sa1_dev .

ENV PORT=2222
EXPOSE 2222
VOLUME /app/.ssh

CMD ["./sa1_dev"]