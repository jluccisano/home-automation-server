FROM resin/raspberrypi3-golang:1.8
RUN apt-get update \
    && apt-get install unzip

RUN mkdir public
WORKDIR /go/src/app/public
RUN wget https://github.com/jluccisano/home-automation-webapp/releases/download/v1.0/home-automation-webapp.zip
RUN unzip ./home-automation-webapp.zip
RUN rm ./home-automation-webapp.zip

WORKDIR /go/src/app

COPY . .
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
CMD ["go-wrapper", "run"] # ["home-automation-server"]