# gopather

GOPATH 自动设置工具, 配合zsh使用

## GOPATH设置多个目录的坏处

多个项目依赖同一个库的不同版本，并且不同版本存在冲突的情况。

## 本工具使用方法

1. 安装

```shell
go get github.com/fengcc/gopather
```

2. 在.zshrc中添加以下内容，并执行`source .zshrc`。GOPATHER为你的多个GOPATH目录

```shell
export GOPATHER=default:gopath2:gopath3

# 每次目录改变的时候都会被调用
chpwd() {
    export GOPATH=$(gopather $(pwd))
}
```

完成。

## 效果

当cd进入GOPATHER中指定的目录或其子目录时，GOPATH就会被自动设置为对应的目录。否则，GOPATH就会被设置成GOPATHER的第一个目录，再也不用为多个GOPATH烦恼了
