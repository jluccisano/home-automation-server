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

Generate private and public key:

```bash
mkdir -p $PWD/resources
```

Private key:
```bash
openssl genrsa -out $PWD/resources/rs256-4096-private.rsa 4096
```

Public key:
```bash
openssl rsa -in $PWD/resources/rs256-4096-private.rsa -pubout -outform PEM -out $PWD/resources/rs256-4096-public.pem
```

```bash
mv $PWD/conf/conf.yml.tpl $PWD/conf/conf-prod.yml
```

Adapt conf:

```yml
url: your_syno_root_url
user: your_user_acount
passwd: your_pass
publicKey: $PWD/resources/rs256-4096-public.pem
privateKey: $PWD/resources/rs256-4096-private.rsa
```

```bash
sudo docker run -dit --name home-automation-server \
    -p 8514:8514 \
    -v $PWD/conf/conf-prod.yaml:/go/src/app/conf/conf-prod.yaml \
    -v $PWD/resources/:/go/src/app/resources \
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
