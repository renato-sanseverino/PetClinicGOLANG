#build stage
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk update && apk add openssl
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=petclinicgolang Version=0.0.1
EXPOSE 3000
