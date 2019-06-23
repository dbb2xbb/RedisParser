package parser

import (
	"fmt"
	"strconv"
)

// 批量回复信息
type BatchResp struct {
	Value interface{}
	Length int
}

// 多条批量回复信息
type MultiBatchResults struct {
	LineNum int
	resps []BatchResp
}

/**************************************************************/

// 多条批量回复
func MultiBatchResp(mstr string) MultiBatchResults {
	start, end := 0, 0
	mb := MultiBatchResults{}
	for i := 0; i < len(mstr)-4; i++ {
		if string(mstr[i:i+4]) == `\r\n`  {
			switch string(mstr[start:end+1][0]) {
			case "*":
				mb.LineNum = IntResp(mstr[start:end+1])
				i += 4
				start, end = i, i
			case "$":
				if IntResp(mstr[start:end+1]) != -1 {
					for i += 4; string(mstr[i:i+4]) != `\r\n`;i++{end = i}
					b := BatResp(mstr[start:end+1+4])
					mb.resps = append(mb.resps, *b)
				}
				i += 4
				start, end = i, i
			case ":":
				mb.resps = append(mb.resps, BatchResp{
					Value: IntResp(mstr[start: end+1]),
				})
				i += 4
				start, end = i, i
			case "+":
				mb.resps = append(mb.resps, BatchResp{
					Value:StatusResp(mstr[start:end+1]),
				})
				i += 4
				start, end = i, i
			case "-":
				mb.resps = append(mb.resps, BatchResp{
					Value:ErrResp(mstr[start:end+1]),
				})
				i += 4
				start, end = i, i
			}
		} else {
			end = i
		}
	}
	return mb
}

// 错误回复
func ErrResp(estr string) string {
	return fmt.Sprintf(`"%s"`,estr[1:len(estr)-4])
}

// 状态回复
func StatusResp(rstr string) string {
	return fmt.Sprintf(`"%s"`,rstr[1:len(rstr)-4])
}

// 整数回复
func IntResp(istr string) int {
	end := len(istr)
	if len(istr) > 4 && istr[len(istr)-4:] == `\r\n` {
		end = len(istr)-4
	}
	res, _ := strconv.Atoi(istr[1:end])
	return res
}

// 批量回复
func BatResp(bstr string) *BatchResp {
	b := &BatchResp{}
	start, end := 0,0
	for; bstr[end:end+4] != `\r\n`; { end ++ }
	Len := IntResp(bstr[start:end])
	if Len != -1 {
		end += 4
		start = end
		for ;string(bstr[end:end+4]) != `\r\n`;end++{ }
		Value := fmt.Sprintf(`"%s"`,bstr[start:end])
		b.Value = Value
		b.Length = Len
		return b
	}
	return nil
}

func RedisParser(RedisStr string) interface{} {
	var res interface{}
	switch string(RedisStr[0]) {
	// 多条批量回复
	case "*":
		res = MultiBatchResp(RedisStr)
	// 错误回复
	case "-":
		res = ErrResp(RedisStr)
	// 状态回复
	case "+":
		res = StatusResp(RedisStr)
	// 整数回复
	case ":":
		res = IntResp(RedisStr)
	// 批量回复
	case "$":
		res = BatResp(RedisStr)
	}
	return res
}



