FROM mcr.microsoft.com/windows/nanoserver:ltsc2022 AS builder
WORKDIR /Users/ContainerUser/etcd
ADD https://github.com/etcd-io/etcd/releases/download/v3.6.5/etcd-v3.6.5-windows-amd64.zip ../etcd.zip
RUN tar xf ../etcd.zip --strip-components=1 *.exe

FROM mcr.microsoft.com/windows/nanoserver:ltsc2022
EXPOSE 2379
WORKDIR /app
COPY --from=builder /Users/ContainerUser/etcd ./
ENTRYPOINT ["/app/etcd.exe", "--advertise-client-urls=http://0.0.0.0:2379", "--listen-client-urls=http://0.0.0.0:2379"]
