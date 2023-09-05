from io import BytesIO
from docker import APIClient
dockerfile = '''
# Shared Volume
FROM busybox:buildroot-2014.02
VOLUME /data
CMD ["/bin/sh"]
'''
f = BytesIO(dockerfile.encode('utf-8'))
cli = APIClient(base_url='tcp://127.0.0.1:2375')
response = [line for line in cli.build(
    fileobj=f, rm=True, tag='yourname/volume'
)]
