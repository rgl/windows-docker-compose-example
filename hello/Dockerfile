FROM golang:1.12-windowsservercore-1809 as builder
WORKDIR /Users/ContainerUser/hello
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go build -ldflags="-s"

FROM mcr.microsoft.com/windows/nanoserver:1809
EXPOSE 8888
WORKDIR /app
COPY --from=builder /Users/ContainerUser/hello/hello.exe ./
ENTRYPOINT ["/app/hello.exe"]
