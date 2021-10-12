## 背景

`GSQL` 最开始因为 [@scially](https://github.com/scially) 在内部业务系统开发中涉及大量Postgis的SQL操作，平时都是Java+SpringBoot就开始干了，咔咔咔先是一上午的环境配置，再是咔咔咔一下午的调试，一天就过去了。所有就需要一个很简单的一个程序，只需要配置SQL和MapperPath，然后直接调用REST服务就行，基于这个想法，搞了这个项目。


这个仓库的目标是：

1. 把SQL转换成REST服务，支持POST、GET等情况，支持跨域。  

## 安装

这个项目使用 [Go](https://golang.org/) 。请确保你本地安装了它们。在工程根目录下，执行

```bash
$ go build main/mian.go -o gsql
```

## 使用说明

1、具体配置参考examples下的conf文件。   
2、数据库配置参考beego的ORM框架

## 示例

在工程根目录下，执行
```bash
$ ./gsql -p 9000 -f ./examples/postgresql.conf
```

## 相关仓库

- [resquel](https://github.com/formio/resquel) — Easily convert your SQL database into a REST API。


## 反馈
1. 有问题提出issue就行，不过不一定有能力修改  
2. 害，代码水平就这样，这么多年，改不过来了。


## 使用许可

[MIT](LICENSE) © scially