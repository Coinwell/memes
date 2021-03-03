# build stage
FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build

# final stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates openssl

COPY --from=0 /app/n2n2-meme /app/
# COPY --from=0 /app/.env /app/

EXPOSE 5000
ENTRYPOINT ["/app/n2n2-meme"]