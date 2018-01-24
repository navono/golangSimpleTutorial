# GolangSimpleTutorial
![](https://img.shields.io/badge/language-Golang-orange.svg)
[![](https://img.shields.io/badge/license-MIT-000000.svg)](https://github.com/navono/golangSimpleTutorial.git/blob/master/LICENSE)

# 环境配置
`VS Code`下安装`Go`插件后，提示需要安装一些工具，但是在墙内会经常导致失败。可以[在此](https://pan.baidu.com/s/1mjLPKfe)下载已经编译好的`Windows`版本。将这些二进制文件放置到`GOPATH\bin`目录下即可。

# 运行
> go run main.go


# 解释说明
## goroutine
`goroutine`是什么？是`OS线程`？它是如何工作的？我们可以随意创建任意数量吗？
> `goroutine`是`Go`特有的，它们不是`OS线程`，或者说连线程都不是。
  <br>它们由`Go`的`runtime`统一管理，一种高层级的抽象的`coroutine`。`coroutine`是并发的`subroutine`（functions, closures or methods in Go），同时是不可抢占的（也就是不能被中断），相反，`coroutine`是可重入（reentry）和挂起的（suspension）。
  <br>`goroutine`不支持可重入和挂起，这些统一由`Go`的`runtime`管理。`runtime`观察`goroutine`的行为，在`goroutine`阻塞时，对其挂起，非阻塞时，恢复其运行。
  <br>由此看出，`goroutine`是一种特殊的`coroutine`。

`goroutine`是隐式并发的。但是这并不是说并发是`coroutine`的一个属性。还是需要有个宿主（Host）来同时管理多个`coroutine`，然后对其进行调度。

`Go`的并发采用的是`fork-join`模型。

# 删除临时文件
> go clean -i