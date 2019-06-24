demo:
	docker build -t redis-parser:1.0 .
pkg:
	docker save -o redisParser.tar.gz redis-parser
clean:
	rm -rf redisParser.tar.gz
	docker rmi redis-parser:1.0
	
