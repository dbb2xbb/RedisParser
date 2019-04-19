#bin/bash

cd `dirname $0`
echo `pwd`

go test -v ./parser/redisParser_test.go ./parser/redisParse.go
