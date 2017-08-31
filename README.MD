```bash
go build .

./home-automation-server
```

sudo docker build -t home-server .
sudo docker run -it --name home-server \
    -p 8514:8514 \
    -v $PWD/conf/conf-prod.yaml:/go/src/app/conf/conf-prod.yaml \
    home-server
sudo docker exec -i -t e6bd4cc550cf /bin/bash
