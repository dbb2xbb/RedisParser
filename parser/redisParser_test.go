package parser

import (
	"fmt"
	"testing"
)

var tests = []struct{
	TestStr string
	result string
}{
	{`*3\r\n$3\r\nSET\r\n$-1\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n`, 	""},
	{`*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n`, 	""},
	{`+OK\r\n`, 	""},
	{`:1000\r\n`, ""},
	{`-ERR unknown command 'foobar'\r\n`,""},
	{`-WRONGTYPE Operation against a key holding the wrong kind of value\r\n`,""},
	{`*3\r\n$-1\r\n$2\r\nHI\r\n`,""},
	{`$5\r\nHello\r\n`,""},
	{`$-1\r\n`,""},
}

func TestRedisParser(t *testing.T) {
	for _, test := range tests {
		fmt.Println(RedisParser(test.TestStr))
	}
}



