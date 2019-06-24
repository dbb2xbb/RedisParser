### Redis协议解析工具，可支持如下几种redis信息解析:
#### 支持功能：
- 状态回复
- 错误回复
- 整数回复
- 批量回复
- 多条批量回复

#### 测试方法：
   本程序支持两种测试方法：
   1. 裸机运行。宿主机配置Golang开发环境，使用`go test`命令测试。
   2. 生成docker镜像，将程序容器化再运行`go test`。
   
   这里推荐第二种方式，可以有效地将测试环境与宿主机隔离，
   避免污染宿主机运行环境。
   
   具体步骤如下：
   
   一. 裸机运行：
   1. 新增`redisParser_test.go`中的`tests`结构体的测试字符串
   2. 执行`test.sh bare` 或 `test.sh b`
   
   二. 容器运行：
   1. 新增`redisParser_test.go`中的`tests`结构体的测试字符串
   2. 在Makefile所在目录中执行`make`
   3. 执行 `test.sh container` 或 `test.sh c`
   
   *若要删除此镜像，则执行`make clean`即可
