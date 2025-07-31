FROM golang:1.24.5-nanoserver-ltsc2022 AS builder
WORKDIR /Users/ContainerUser/hello
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -ldflags="-s"

FROM mcr.microsoft.com/windows/nanoserver:ltsc2022
EXPOSE 8888
WORKDIR /app
COPY --from=builder /Users/ContainerUser/hello/hello.exe ./
ENTRYPOINT ["/app/hello.exe"]
