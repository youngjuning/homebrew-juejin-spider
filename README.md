# 掘金专栏小爬虫

这个脚本很简单，爬取掘金专栏详情页并把内容转成Markdown之后保存在本地。

## 安装

```sh
// 在执行这个命令的时候，brew会自动去更新自己的formula仓库，会耗时几分钟。。。
$ brew tap youngjuning/juejin-spider https://github.com/youngjuning/homebrew-juejin-spider.git
// 下载、安装 hello 脚本
$ brew install youngjuning/juejin-spider/juejin
```

## 升级

```sh
$ brew reinstall juejin
```

## 使用

```sh
$ juejin -h
Usage of juejin:
  -post string
    	文章编号 (default "6859538537830858759")
  -root string
    	文件保存的根目录 (default "/Users/yangjunning/juejin")
```

## 打包并使用Homebrew发布脚本

### 1、打包成可执行文件

```sh
$ go build juejin.go
```

会在当前目录下生成一个叫 juejin 的可执行文件，`./juejin` 是可以执行的，也可以使用`go build -o=/usr/local/bin juejin.go` 或 `go build -o=$GOPATH/bin/ juejin.go` 放到已经注册的系统路径中。

> 黑客是要有追求的，不可能做个玩具出来。而且Go本身就是运行起来不靠任何依赖和环境，我不能要求使用工具的人还得装个go的环境。我第一个想到的就是将我的脚本发到Homebrew，谢天谢地，[使用HomeBrew发布脚本](https://www.jianshu.com/p/e88831aac62a) 详细地讲解了这个过程。

### 2、将可执行文件打包成 tar.gz 的格式

```shell
$ tar zcvf juejin_0.0.1.tar.gz juejin
```

![](https://i.loli.net/2020/08/14/u4ZWeM1UlIqALzH.png)

### 3、使用 `brew create <git-url>` 创建药方

```sh
$ brew create https://github.com/youngjuning/homebrew-juejin-spider/raw/master/juejin_0.0.1.tar.gz
```

> 注意，这里因为我没有更改brew的镜像源，所以创建成功后文件会出现在`/usr/local/Homebrew/Library/Taps/homebrew/homebrew-core/Formula/juejin.rb` 我们可以手动将这个 hello.rb 文件剪切到自己的Formula文件夹下`/usr/local/Homebrew/Library/Taps/youngjuning/homebrew-juejin-spider/Formula/juejin.rb`。

我们需要对安装方式做一下调整:

```ruby
def install
    bin.install "juejin"
end
```

做完这些操作后，保存，提交到git上。

### 4、安装脚本

```shell
$ brew install youngjuning/juejin-spider/juejin
```

执行 `juejin -h` 检查是否成功：

```shell
$ juejin -h
Usage of juejin:
  -post string
    	文章编号 (default "6859538537830858759")
  -root string
    	文件保存的根目录 (default "/Users/yangjunning/juejin")
```
