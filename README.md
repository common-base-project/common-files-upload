# 公共文件管理系统后台
## 功能
1. 支持文件上传到服务器
2. 支持文件上传到 ipfs
3. 支持文件上传并入库
4. 支持文件下载
5. 支持文件在线预览


## 通用文件上传
    设置环境变量  export ENV_SERVER_MODE=dev

## 打包
    make docker-all VERSION="staging_v0.0.1" ENV_SERVER_MODE="staging"
    make docker-all VERSION="staging_v0.0.1" ENV_SERVER_MODE="dev"
    make docker-all VERSION="prod_v0.0.2" ENV_SERVER_MODE="prod"

    docker buildx build --platform linux/amd64 --no-cache -f Dockerfile-prod -t ccr.ccs.tencentyun.com/game-center/common-files:staging_v0.0.1 .
    docker buildx build --platform linux/amd64 -f Dockerfile-prod -t ccr.ccs.tencentyun.com/game-center/common-files:staging_v0.0.2 .
    docker build -f Dockerfile-prod -t registry.cn-hongkong.aliyuncs.com/game-center/common-files:staging_v0.0.1 .

    docker push ccr.ccs.tencentyun.com/game-center/common-files:staging_v0.0.2

## sipue
    docker login harbor.test.sipue.cn --username=gongzhigang-sipue -p mtgxxxxx

    docker buildx build --platform linux/amd64 -f Dockerfile-prod -t harbor.test.sipue.cn/sipue/common-files:staging_v1.0.2 .
    docker push harbor.test.sipue.cn/sipue/common-files:staging_v1.0.2

## 生成`swagger`文档
```
Download Swag for Go by using:
    go get -u github.com/swaggo/swag/cmd/swag
Run the Swag in your Go project root folder which contains main.go file, Swag will parse comments and generate required files(docs folder and docs/doc.go).
    swag init
    
# 基于Makefile
make swagger

# OR 使用swag命令
swag init --generalInfo ./cmd/gin-admin/main.go --output ./internal/app/swagger
swag init -g ./src/cmd/main.go

```

## 基于 docker 容器开发
```text
为什么要在开发中使用Docker?
    // https://studygolang.com/articles/04818
    1 一致的开发环境 使用Docker，可以保证整个研发团队使用一致的开发环境。
    2 开发环境与最终的生产环境保持一致 这减少了部署出错的可能性。
    3 简化了编译和构建的复杂性 对于一些动辄数小时的编译和构建工作，可以用Docker来简化。
    4 在开发时只需Docker 无需在自己的开发主机上搭建各种编程语言环境。
    5 可以使用同一编程语言的多个版本 可以使用同一编程语言（如python, python, ruby, ruby, java, node）等的多个版本，无需解决多版本冲突的问题。
    6 部署很简单 应用程序在容器中运行，和在生产环境中部署运行是一样的。只需打包你的代码并部署到带有同样镜像的服务器上，或者是把代码连同原镜像建立一个新Docker镜像再直接运行新镜像。
    7 使用自己喜欢的开发IDE 仍然可以继续使用自己喜欢的开发IDE，无需运行VirtualBox虚拟机或SSH。

怎样在开发中使用Docker?
    采用Docker开发与普通开发不同之处有两点：
    1 确保所有的依赖都放入了工作目录
    2 修改构建和运行命令，使之在Docker容器中可以运行

案例
    把Golang开发的代码复制到GOPATH指定的位置，且系统必须安装Go环境。假设程序就是hello.go，且使用了websocket。那么过程如下：
    安装依赖:
        github.com/gorilla/websocket
    建立应用程序
        docker run --rm -v "$GOPATH":/gopath -v "$(pwd)":/app -w /app google/golang sh -c 'go build -o hello'
    注意要把本地GOPATH挂载到容器。

    运行应用程序
        docker run --rm -v "$(pwd)":/app -w /app google/golang sh -c './hello'
    注意，要保持容器的一直可用，比如RocksDB实例的一直可用，这样就不会使得容器每次运行都删除了RocksDB实例，可以执行：

        docker run --name goapp -v "$GOPATH":/gopath -v "$(pwd)":/app -w /app google/golang sh -c 'go build -o hello && ./hello' || docker start -ia goapp


本项目本地开发步骤：
前提（可选）：
    需要安装 air 工具: https://github.com/cosmtrek/air
一 直接下载源代码到本地用 IDE 这本地调试开发
二 基于 docker 环境开发
    1 安装 docker
    2 下载开发镜像 'registry.aibee.cn/qa/unicorn-files:dev_v1' 或者基于源代码编译docker镜像
        docker build -f dev.Dockerfile -t registry.aibee.cn/qa/unicorn-files:dev_v1 .
        docker push registry.aibee.cn/qa/unicorn-files:dev_v1
    
    3 推荐 vscode 基于 docker 开发
        a vscode 需要安装 "Remote - Containers" 工具
        b 这 vscode 编辑器选择快捷键 Cmd + shift + p 输入 "Remote-Containers: Attach to Running Container……" 然后选择 unicorn-files
        c vscode 打开文件夹，打开 "/opt/app" 目录即可开发
        d 可以这 Container 里直接使用自己的git， 也可以直接调试等
        
     
```
