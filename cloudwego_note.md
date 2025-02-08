# CloudWeGo 学习
## 参考网站：
Hertz文档网站：`https://www.cloudwego.io/zh/docs/hertz/`
视频案例代码网站：`https://github.com/cloudwego/biz-demo`

## 将Vscode中Go的工具全部添加下载
- 在vscode界面打开命令面板（ctrl+shift+p）
- 输入 `Go: Install/Update Tools`，勾选全部工具，点击确认

## git ssh连接超时解决方法：
将ssh端口22切换为端口443，即在C盘的.ssh文件下新建config文件,内容如下：
```
Host github.com
    Hostname ssh.github.com
    Port 443
    User git
    IdentityFile ~/.ssh/id_rsa  # 替换为你的私钥路径
```


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
每个线程用不同的计算机语言，如同时用java和go语言,也**可以将一个go程序封装成一个接口被另一个go程序调用**
两种常见的IDL：`thrift`和`protobuf`

#### 相关网站：
**cwgo(代码生成工具)：**
网址：`https://github.com/cloudwego/cwgo` （网站下有cwgo的说明文档可以查看）
安装命令：`GOPROXY=https://goproxy.cn/,direct go install github.com/cloudwego/cwgo@latest`  
如果提前设置好`GOPROXY`,直接使用  `go install github.com/cloudwego/cwgo@latest`


**thriftgo(thrift依赖于此工具)：**
网址：`https://github.com/cloudwego/thriftgo`
安装命令：`go install github.com/cloudwego/thriftgo@latest`

**protobuf:**
网址：`https://protobuf.com.cn/programming-guides/proto3/`
下载路径：`https://github.com/protocolbuffers/protobuf/releases` （注意下载windows对应的win64.zip）


#### thrift运行代码：
- 先检测环境路径中有没有cwgo的路径，将模块添加到/bin路径中
`cwgo --help` 检查
- **视频中没有，但需要：** 在初始文件夹目录下运行指令`go work init`(如果有多个 Go 模块在不同的目录下，并且希望它们能够互相引用和协作，你可以使用 `go work init` 来创建一个工作区，确保它们之间的依赖关系能够正确处理。)
- 创建demo/demo_thrift文件夹，输入指令：`cwgo server --type RPC --module github.com/yuefan-mo/studymall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift`
- `go work use .`：将会在 go.work 文件中记录 /project 目录，使得工作区能够管理 /moduleA 和 /moduleB 中的模块依赖关系。

 #### proto运行代码：
 - 需要将 proto中的protoc.exe文件和include文件夹添加到环境变量中
 - 输入指令：`cwgo server -I ../../idl --type RPC --module github.com/yuefan-mo/studymall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto`

## 服务注册与服务发现
Kitex官网：``
consul安装注册：`https://duoke360.com/tutorial/hertz/h149
`
