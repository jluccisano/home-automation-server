FROM resin/raspberrypi3-golang:1.8
WORKDIR /go/src/app
RUN mkdir public
RUN wget -P ./public https://github.com/jluccisano/home-automation-webapp/releases/download/v1.0/home-automation-webapp.zip
COPY . .
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
CMD ["go-wrapper", "run"] # ["home-automation-server"]