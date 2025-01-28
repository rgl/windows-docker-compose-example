# enable strict mode and fail the job when there is an unhandled exception.
Set-StrictMode -Version Latest
$FormatEnumerationLimit = -1
$ErrorActionPreference = 'Stop'
$ProgressPreference = 'SilentlyContinue'
trap {
    Write-Host "ERROR: $_"
    ($_.ScriptStackTrace -split '\r?\n') -replace '^(.*)$', 'ERROR: $1' | Write-Host
    ($_.Exception.ToString() -split '\r?\n') -replace '^(.*)$', 'ERROR EXCEPTION: $1' | Write-Host
    Exit 1
}

# wrap the docker command (to make sure this script aborts when it fails).
function docker {
    docker.exe @Args | Out-String -Stream -Width ([int]::MaxValue)
    if ($LASTEXITCODE) {
        throw "$(@('docker')+$Args | ConvertTo-Json -Compress) failed with exit code $LASTEXITCODE"
    }
}

docker compose up --build -d
try {
    docker compose ps
    docker compose exec -T etcd etcd --version
    docker compose exec -T etcd etcdctl version
    docker compose exec -T etcd etcdctl endpoint health
    docker compose exec -T etcd etcdctl put foo bar
    docker compose exec -T etcd etcdctl get foo
    $port = (docker compose port hello 8888) -replace '.+:(\d+)', '$1'
    $endpoint = "http://localhost:$port"
    Write-Host "Invoking the hello endpoint at $endpoint..."
    (New-Object System.Net.WebClient).DownloadString($endpoint)
}
finally {
    docker compose logs
    docker compose down --volumes
}
