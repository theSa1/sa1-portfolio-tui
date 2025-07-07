FROM mirror.gcr.io/golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ssh .

FROM mirror.gcr.io/alpine:3.21.2
WORKDIR /root/
ENV PORT=2222
EXPOSE 2222
VOLUME /root/.ssh
COPY --from=builder /app/ssh .
CMD ./ssh