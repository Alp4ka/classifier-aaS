FROM golang:1.21.4 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./app github.com/Alp4ka/classifier-aaS/cmd/app

FROM alpine:latest AS run

RUN apk add curl

COPY --from=builder /build/app /usr/local/bin/app

COPY --from=builder /build/cmd/app/docs ./docs

ENV PG_DSN="postgres://db:db@pg:5432/classifier-aas?sslmode=disable"

CMD ["app"]
