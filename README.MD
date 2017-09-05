# Home Automation Server

[![tests][tests]][tests-url]

```bash
go build .

./home-automation-server
```

sudo docker build -t jluccisano/home-automation-server .
sudo docker run -dit --name home-automation-server \
    -p 8514:8514 \
    -v $PWD/conf/conf-prod.yaml:/go/src/app/conf/conf-prod.yaml \
    jluccisano/home-automation-server
sudo docker exec -i -t e6bd4cc550cf /bin/bash


travis encrypt DOCKER_USER=username --add
travis encrypt DOCKER_PASS=password --add

uname -m
armv7l


[tests]: http://img.shields.io/travis/jluccisano/home-automation-server.svg
[tests-url]: https://travis-ci.org/jluccisano/home-automation-server