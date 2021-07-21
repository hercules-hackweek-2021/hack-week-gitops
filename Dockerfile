FROM golang:1.16 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /opt/app
COPY . .
RUN go build -o bin/server

FROM scratch AS server
ENTRYPOINT ["/opt/app/server"]
COPY --from=builder /opt/app/bin/server /opt/app/server
