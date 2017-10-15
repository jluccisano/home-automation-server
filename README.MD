# Home Automation Server

[![tests][tests]][tests-url]

```bash
go build .

./home-automation-server
```
### Docker

#### Build

```bash
sudo docker build -t jluccisano/home-automation-server .
```

#### Run

```bash
sudo docker run -dit --name home-automation-server \
    -p 8514:8514 \
    -v $PWD/conf/conf-prod.yaml:/go/src/app/conf/conf-prod.yaml \
    jluccisano/home-automation-server
```

#### Check

```bash
curl -H "Content-Type: application/json" \
     -X POST -d '{"id":"foo","password":"bar"}' \
     http://localhost:8514/authenticate
```

[tests]: http://img.shields.io/travis/jluccisano/home-automation-server.svg
[tests-url]: https://travis-ci.org/jluccisano/home-automation-server
