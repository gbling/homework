# Dockerfile
FROM golang:alpine AS builder
WORKDIR /go/src/project/
COPY ./module2/ /go/src/project/
RUN go build -o /go/src/project/server

FROM alpine:latest as prod
COPY --from=builder /go/src/project/server /bin/server
EXPOSE 8080
ENTRYPOINT ["/bin/server"]
