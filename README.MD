```bash
go build .

./home-automation-server
```

sudo docker build -t home-server .
sudo docker run -dit -p 8514:8514 -v /conf.yaml:/go/src/app/conf.yaml --name home-server home-server
sudo docker exec -i -t e6bd4cc550cf /bin/bash
