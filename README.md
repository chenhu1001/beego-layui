# beego-layui
采用beego和layui开发的脚手架

## 使用 GoReleaser 发布你的应用
初始化
```
goreleaser init
```

执行自动发布流程
Mac
```
brew install goreleaser
goreleaser --rm-dist
goreleaser --snapshot --skip-publish --rm-dist
```
Linux
```
./goreleaser --rm-dist
./goreleaser --snapshot --skip-publish --rm-dist
```
Windows
```
goreleaser.exe --rm-dist
goreleaser.exe --snapshot --skip-publish --rm-dist
```

## 常见问题
### 1、go:linkname must refer to declared function or variable
出现此问题一般是在go 1.18版本，解决办法是单独执行：go get -u golang.org/x/sys