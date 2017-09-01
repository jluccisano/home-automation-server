FROM resin/raspberrypi3-golang:1.8
RUN apt-get update \
    && apt-get install unzip
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
CMD ["go-wrapper", "run"] # ["home-automation-server"]