FROM golang:1.19.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine

COPY --from=builder /app/server .

CMD [ "./server" ]
