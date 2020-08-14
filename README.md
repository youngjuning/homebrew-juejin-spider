# 掘金专栏小爬虫

这个脚本很简单，爬取掘金专栏详情页并把内容转成Markdown之后保存在本地。这是在写 [极速入门Go并爬取掘金专栏 | 🏆 技术专题第二期](https://juejin.im/post/6860522117423857678) 时写的小玩意。

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
$ juejin -post 6860522117423857678 -root ~/Desktop
```

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

### 2、将可执行文件打包成 tar.gz 的格式

```shell
$ tar zcvf juejin_0.0.1.tar.gz juejin
```

![](https://i.loli.net/2020/08/14/u4ZWeM1UlIqALzH.png)

上传到git，供配方软连接到这个脚本文件。

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
