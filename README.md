# About

[![Build](https://github.com/rgl/windows-docker-compose-example/actions/workflows/build.yml/badge.svg)](https://github.com/rgl/windows-docker-compose-example/actions/workflows/build.yml)

This is an example on how to start a docker compose environment in a (remote) docker host.

**NB** This is similar to [rgl/ubuntu-docker-compose-example](https://github.com/rgl/ubuntu-docker-compose-example).

## Usage

```bash
# if required, set the DOCKER_HOST environment variable.
# this make the docker client use this dockerd.
# NB you must start dockerd with -H tcp://0.0.0.0:2375
#export DOCKER_HOST=tcp://localhost:2375

# create the environment defined in docker-compose.yml
# and leave it running in the background.
docker compose up --build -d

# show running containers.
docker compose ps

# execute command inside the containers.
docker compose exec -T etcd etcd --version
docker compose exec -T etcd etcdctl version
docker compose exec -T etcd etcdctl endpoint health
docker compose exec -T etcd etcdctl put foo bar
docker compose exec -T etcd etcdctl get foo

# get the allocated hello port and create an endpoint url based in
# the DOCKER_HOST environment variable host.
hello_endpoint="$(
python <<'EOF'
import os
import urllib.parse
import subprocess

p = subprocess.Popen(
    ['docker', 'compose', 'port', 'hello', '8888'],
    text=True,
    stdout=subprocess.PIPE,
    stderr=subprocess.STDOUT)
stdout, stderr = p.communicate()
if 'DOCKER_HOST' in os.environ:
    docker_host_ip_address = urllib.parse.urlparse(os.environ['DOCKER_HOST']).netloc.split(':')[0]
    hello_port = stdout.strip().split(':')[-1]
    hello_endpoint = f'http://{docker_host_ip_address}:{hello_port}'
else:
    hello_endpoint = f'http://{stdout.strip()}'
print(hello_endpoint)
EOF
)"

# invoke the hello endpoint.
wget -qO- $hello_endpoint

# show logs.
docker compose logs

# destroy the environment.
docker compose down
```
