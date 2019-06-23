#!/bin/bash
cd `dirname $0`
echo `pwd`

case $1 in
c|container)
	docker run -it --rm redis-parser:$2
	;;

b|bare)
	go test -v ./parser/redisParser_test.go ./parser/redisParse.go
	;;
*)
	echo "Usage: ./test.sh [c/container|b/bare]"
	exit 1
	;;
esac
exit 0


