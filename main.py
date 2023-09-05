import docker
client = docker.from_env()

containers = client.containers.list()

for container in containers:
    id = container.attrs.get('Id')
    running = container.attrs.get('State').get('Running')
    ip = container.attrs.get('NetworkSettings').get('IPAddress')
