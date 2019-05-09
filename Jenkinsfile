pipeline {
    agent {
        label 'windows && docker'
    }
    stages {
        stage('Build') {
            steps {
                powershell '''
function docker-compose {
    docker-compose.exe @Args
    if ($LASTEXITCODE) {
        throw "failed to run docker-compose $Args with Exit Code $LASTEXITCODE"
    }
}

docker-compose up --build -d
try {
    docker-compose ps
    docker-compose exec -T hello cmd /c dir
    docker-compose exec -T etcd /app/etcd --version
    docker-compose exec -T -e ETCDCTL_API=3 etcd /app/etcdctl version
    docker-compose exec -T -e ETCDCTL_API=3 etcd /app/etcdctl endpoint health
    docker-compose exec -T -e ETCDCTL_API=3 etcd /app/etcdctl put foo bar
    docker-compose exec -T -e ETCDCTL_API=3 etcd /app/etcdctl get foo
    $port = (docker-compose port hello 8888) -replace '.+:(\\d+)','$1'
    (New-Object WebClient).DownloadString("http://localhost:$port")
} finally {
    docker-compose logs
    docker-compose down
}
'''
            }
        }
    }
}
