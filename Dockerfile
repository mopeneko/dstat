FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app main.go

FROM gcr.io/distroless/base

COPY --from=builder /app/app /

CMD ["/app"]
