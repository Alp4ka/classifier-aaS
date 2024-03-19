FROM golang:1.19 AS builder

RUN echo "machine gitlab.com login bremk password glpat-zUnRSgzgnfhnkKYQzNks" > ~/.netrc

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./user ./cmd/app/main.go

FROM alpine:latest AS run

RUN apk add curl

COPY --from=builder /build/user /usr/local/bin/user

COPY docs /docs

COPY configs /configs

CMD ["user"]
