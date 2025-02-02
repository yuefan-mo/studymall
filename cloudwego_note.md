# CloudWeGo 学习
## 参考网站：
Hertz文档网站：`https://www.cloudwego.io/zh/docs/hertz/`
视频案例代码网站：`https://github.com/cloudwego/biz-demo`

## hello_world 开发环境配置
使用`go env `检查`GOPROXY`的值，修改为`go env -w GOPROXY=https://goproxy.cn,direct`

- 使用 `go mod init github.com/yuefan-mo/studymall`(自己的仓库地址)为项目添加`go mod`管理项目依赖 

- 引入hertz: `go get -u github.com/cloudwego/hertz`

- 使用`go mod tidy`整理代码：（用于整理和清理 Go 模块依赖）
```
其主要功能包括：

- 移除未使用的依赖：它会自动删除 `go.mod` 和 `go.sum` 文件中列出的那些在代码中没有实际使用的模块和依赖。

- 添加缺失的依赖：它会根据代码中实际使用到的模块，确保 `go.mod` 文件中包含正确的依赖，并将缺失的模块添加进去。

- 更新 `go.sum` 文件：确保 `go.sum` 文件中包含了正确的校验和，避免不必要的或冗余的条目。
```

最后打印`hello world`的代码不是视频中的`ctx`，而是`c`如下：
```go
	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello world"))
	})
```

## 脚手架
### IDL（接口描述语言）
每个线程用不同的计算机语言，如同时用java和go语言
