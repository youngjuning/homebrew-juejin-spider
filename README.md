![](https://i.loli.net/2020/08/13/92pykLqY7NFWi3f.png)

<!--more-->

## Go印象

2018年的某一天，我在和公司后端架构师聊天时说我想学学后端语言，除了Java有啥推荐，他告诉我他在学Go。然后跟我讲了一些诸如分布式、协程、大数据、爬虫......巴拉巴拉的我也听不太懂的概念。然后我说我还是学NodeJs吧。

![](https://i.loli.net/2020/08/13/cVomW7L9YOTw2uA.png)

之所以斗胆再战Go语言，完全是梁静茹（[上次征文]()）给了我勇气。如果你是后端大佬，实在嫌弃一小时入门的Goer，那点个赞出门左转也是可以的。

本文的主题是重点是极速、爬虫、掘金专栏，目的是使用Go写一个小工具把掘金专栏文章爬取下来，慢慢看。

## Let's Go

### 优势

- 语法简单，易上手（仅有25个关键保留字）
- 性能高、编译快，开发效率不比Python和Ruby低
- 部署方便，编译包小，几乎无依赖（二进制文件包可直接运行）这一点和Deno很像
- 原生支持并发（goroutine）
- 官方统一规范（gofmt、golint）又看到了Deno的影子
- 丰富的标准库，再次看到了Deno的影子

### 趋势

权威的趋势，优弧大佬已经讲的很清楚了，我这里补充一下GitHub的star趋势和行业薪资:

<div style="display:flex;justify-content:space-between">
	<img src="https://i.loli.net/2020/08/13/ZYijG4BOe2rxXMT.png" style="zoom:35%;" />
  <img src="https://i.loli.net/2020/08/13/FjxKWmUR354H8Ea.png" style="zoom:35%;" />
</div>

## Go简介

Go是Google开发的一种静态强类型、编译型、并发型并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。

## Go语言特点

1. Go是一种新的语言，是一种支持并发、带垃圾回收、可快速编译的静态语言。
2. Go为并发执行与通信提供了基本的支持，是天生的高性能服务开发语言。
3. Go结合了解释性语言的游刃有余，动态类型语言的开发效率，以及静态类型的安全性。
4. Go只需要用几秒的时间就可以编译一个大型的Go程序，部署也非常容易。
5. Go具有Python/Ruby的开发效率，同时又是C语言的运行性能（还是有一定差距的）。
6. Go简单易上手（只有25个关键保留字）
7. Go有自己的开发规范，还提供工具支持。

## Go安装配置

> 作者还有一篇[程序员的Mac开发环境【持续更新】](https://juejin.im/post/6844904083489308685)，记录了我的Mac上的开发环境，读者大大可以顺便给个Star吗？

```sh
$ brew install go
```

![](https://i.loli.net/2020/08/13/co5mg3QNti1UPur.png)

> 小技巧1：`ctrl+c` 可以跳过 `Updating Homebrew...`，要不卡到你怀疑人生。

> 小技巧2：如果你有时间等，可以带上 `-verbose` 参数，这样下载的时候会告诉你更新进度。

> 小技巧3：Homebrew 是同步的 GitHub 仓库，如果实在卡，请自行换 Homebrew 代理源

安装成功后，查看go版本：

```sh
$ go version
go version go1.14.7 darwin/amd64
```

配置环境变量：

```sh
$ open /usr/local/Cellar/go/
```

> 然后看一下自己的libexec在什么地方然后记录下整体的地址，我的地址是 `/usr/local/Cellar/go/1.14.7/libexec`

需要将这部分写入到 `nano ~/.bash_profile`：

```sh
#GO
export GOROOT=/usr/local/Cellar/go/1.14.7/libexec
export GOPATH=~/.go
export PATH=${PATH}:$GOPATH/bin
```

好习惯，` source ~/.bash_profile`

## Go 常用命令

1、`go build`：用于编译我们指定的源码文件或代码包以及它们的依赖包

2、`go clean`：用来移除当前源码包里面编译生成的文件

3、`go doc`：打印附于Go语言程序实体上的文档。我们可以通过把程序实体的标识符作为该命令的参数来达到查看其文档的目的。

4、`go fmt`：帮助格式化你的代码文件，你只需要执行 go fmt xxx.go 你的代码将会被修改为标准格式

5、`go get`：根据要求和实际情况从互联网上下载或更新指定的代码败一级依赖包，并对他们进行编译和安装

6、`go install`：用于编译并安装指定的代码包及它们的依赖包

7、`go run`：可以遍历源码并运行命令源码文件

## Go标准库

1、sync：提供了基本的同步原语。在多个goroutine访问共享资源的时候，需要使用sync中提供的锁机制。

2、os：提供了对操作系统功能的非平台相关访问接口。接口为Unix风格。提供的功能包括文件操作、进程管理、信号和用户账号等。

3、time：时间相关的处理

4、fmt：实现格式化的输入输出操作。

5、io：实现了一系列非平台相关的IO相关接口和实现，比如提供了对os中系统相关的IO功能的封装。我们在进行流式读写（比如读写文件）时，会用到该包。

6、http：提供web服务

7、string：处理字符串的一些函数集合，包括合并、查找、分割、比较、后缀检查、索引、大小写处理等等。

## VsCode插件推荐

- Go：Rich Go language support for Visual Studio Code

## Hello World

创建`helloworld.go`写入下面内容:

```go
package main  // 代码包声明语句。
import "fmt" // 系统包用来输出的

func main() {
  // 打印函数调用语句。用于打印输出信息。
  fmt.Println(sayHello("掘金"))
}

func sayHello(juejin string) string {
  return "Hello "+juejin
}
```

然后执行`go run helloworld.go`，好了你已经入门了。

## 七牛云镜像代理

打开你的终端并执行，Go 1.13 及以上可行，其他版本请阅读[Goproxy中国](https://goproxy.cn/) 查看文档

```shell
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

## 爬取掘金专栏

### 爬虫是什么

百度百科和维基百科对网络爬虫的定义：简单来说爬虫就是抓取目标网站内容的工具，一般是根据定义的行为自动进行抓取, 对网页或数据的分析与过滤；抓取的网页URL进行爬行策略。

简单来说就是把目标网页下载下来，然后通过解析、过滤、去重等一系列操作获得自己想要的数据并以相应的格式保存下来。大致流程如下图：

![](https://i.loli.net/2020/08/13/f1ODv42eWHjo3yq.png)

### 分析页面结构

**专栏详情链接**

```html
<a data-v-76f2cee1="" href="(.*?)" target="_blank" rel="" st:name="title" class="title">
```

**专栏标题**

```html
<h1 data-v-23a9d5ed="" class="article-title">(.*?)<\/h1>
```

**专栏正文**

```html
<div class="markdown-body">(.*?)<\/div>
```

### colloy极速上手

gocolly是用go实现的网络爬虫框架，目前在github上具有11K+星，名列go版爬虫程序榜首。gocolly快速优雅，在单核上每秒可以发起1K以上请求；以回调函数的形式提供了一组接口，可以实现任意类型的爬虫；依赖goquery库可以像jquery一样选择web元素。

gocolly的官方网站是http://go-colly.org/，提供了详细的文档和示例代码。安装colly:

```sh
$ go get -u github.com/gocolly/colly/...
```

在 `go.mod` 中管理依赖：

```go
module juejin.im/junning

go 1.14

require (
  github.com/gocolly/colly/v2 latest
)
```

在代码中导入包：

```go
import "github.com/gocolly/colly"
```

## 踩坑

1、**go: cannot find main module; see 'go help modules'**

因为项目根目录下面没有`go.mod`文件，创建这个文件，这个文件里面用来管理module的。具体可以参考：https://blog.csdn.net/benben_2015/article/details/82227338

## 参考资料

- [漫画 Go 语言 纯手绘版 | 掘金小册](https://juejin.im/book/6844733833401597966/section)
- [Go 爬虫之 colly 从入门到不放弃指南]()
- [go语言中文网](https://studygolang.com/)
- [MacOS 使用Homebrew 安装go](https://juejin.im/post/6844904147125272583#heading-4)
- [Golang 常用标准库](https://juejin.im/post/6844904048508796935)
- [go语言常用命令介绍](https://juejin.im/post/6844903844472684558)
- [爬虫基本原理](https://juejin.im/post/6844903908997857288)
- [Goproxy 中国](https://goproxy.cn/)：中国最可靠的 Go 模块代理。

## 超级赛亚人镇楼

<img src="https://i.loli.net/2020/08/12/FAT4jKHG2Es7rfN.gif" style="zoom:65%;" />

> 点赞等于学会，收藏等于精通，点赞加收藏是真爱！

[🏆 技术专题第二期 | 我与 Go 的那些事......](https://juejin.im/post/6859784103621820429)
