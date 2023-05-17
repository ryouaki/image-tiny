# image-tiny
A image optimize service for PNG(done), JPEG(doing), JPG(doing).

## For development
先确保提前安装Cargo工具。因为底层imagequant是Rust开发的。具体安装方法参考[安装Rust](https://www.rust-lang.org/zh-CN/tools/install)

也可以参考[Rust 及其包管理Cargo的安装使用](https://www.cnblogs.com/yucloud/p/rust_cargo.html) 这个比较细

```sh
  $ ./build.sh // 初始化开发环境。
  // 如果是m1芯片的mac 使用下面命令行初始化
  $ make init
  // 如果是其他架构的cpu。需要自己指定CPU架构
  $ ./build.sh xxxx
```

## For production
确保构建服务器有Cargo。修改Makefile文件。指定构建服务器CPU架构
```sh
  // ./Makefile
  # rustup target list
  all: init

  init:
    ./build.sh aarch64-apple-darwin 

  build:
    ./build.sh aarch64-apple-darwin  // <<<< 修改这里
    go build -ldflags "-s -w" -o ./bin/gocf ./main.go 
```

```sh
  $ make build // 打包
```

## 测试
该接口直接返回图片内容。前端需要自己处理一下
```sh
  curl --location 'localhost:8080/compress' --form 'file=@"/Users/user/Documents/WorkDocs/image-tiny/demo.png"'
```

