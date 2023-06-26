#build stage
FROM golang:1.20-alpine AS builder

WORKDIR /linkshorter
COPY . .
RUN go build -o app ./cmd/main.go

#run stage
FROM alpine
WORKDIR /linkshorter
COPY --from=builder /linkshorter/app .
COPY /config.env .

EXPOSE 8000 9000
CMD ./app