FROM golang:latest
ADD . /go/src/RedisParser
WORKDIR /go/src/RedisParser
CMD ["go","test","-v","RedisParser/parser"]
