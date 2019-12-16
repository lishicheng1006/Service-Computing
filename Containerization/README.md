# 服务计算：MySQL 与容器化

本次课程作业围绕着 MySQL 和容器化而展开的。

## Hello Docker

### 编写 Dockerfile

我们需要编写 Dockerfile ，才可以构建 Docker 镜像，以下是一个简单的镜像构建文件。

```dockerfile
FROM ubuntu
ENTRYPOINT ["top", "-b"]
CMD ["-c"]
```

在该文件中，我们指明镜像基于 `ubuntu:latest` 镜像，镜像启动后运行 `top -b -c` 命令。

执行 `docker build . -t hello` 命令，我们即可构建一个简单的 Docker 镜像。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_1.png)

### 运行容器

在构建完容器后，我们可执行以下命令运行容器。

```bash
docker run -it -rm hello -H
```

其中，`-it` 表示**可交互式的 TTY 界面**，`-rm` 表示容器运行完毕后删除此容器。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_2.png)

## MySQL Docker

我们可执行以下命令运行一个 MySQL 容器，在命令中我们通过 `-e` 设置容器的环境变量参数，设定了 MySQL 数据库密码。

```dockerfile
docker run -p 3336:3306 --name hellomysql -e MYSQL_ROOT_PASSWORD=Abcd1234 -d mysql:5.7
```

上述命令执行结果如下。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_3.png)

我们可执行 `docker ps -a | grep hellomysql` 命令查看容器是否正常运行，该命令执行结果如下图所示，输出结果中的 `Up 3 minutes` 表明容器已正常运行。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_4.png)

### 持久化

数据库的一个很重要的特性是**持久化**，而上述执行的命令是不具备持久化能力的，因为一旦容器停止运行，所存储的数据都丢失了。

因此，我们需要创建 Volumes ，并与镜像链接，实现数据持久化。首先，我们需要删除原先创建的镜像，具体命令如下。

```bash
docker rm -fv $(docker ps -a | grep hellomysql |awk '{print $1}')
```

随后，我们创建名为 `hellovolume` 的卷。

```bash
docker volume create hellovolume
```

最后，我们将 `hellovolume` 卷与容器链接，并启动容器。

```
docker run -p 3336:3306 --name hellomysql -e MYSQL_ROOT_PASSWORD=Abcd1234 -v hellovolume:/var/lib/mysql -d mysql:5.7
```

上述命令的执行结果如下图所示。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_5.png)

我们现在创建另外一个容器，作为 MySQL 数据库的客户端，实现与 `hellomysql` 的交互，执行命令如下：

```bash
docker run --name myclient --link hellomysql:mysql -it mysql:5.7 bash
```

在该命令中，我们使用 `--link` 参数，使得 `myclient` 和 `hellomysql` 容器链接在一起，命令执行结果如下图所示。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_6.png)

## 使用 Docker Compose

[Docker Compose](https://docs.docker.com/compose/) 是一个多容器应用自动化部署工具，可大大提高容器部署效率，其通过解析 `docker-compose.yml` 文件中的配置，实现容器的自动化部署。

现在，我们想运行多个容器：MySQL 和 Adminer ，其中[Adminer]( https://www.adminer.org/ ) 是一个数据库管理工具，提供了一个 Web 页面，用于管理数据库。我们编写以下的 `docker-compose.yml` 文件。

```yaml
version: "3"
services:
  mysql:
    image: mysql:5.7
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: Abcd1234
  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080
```

在上述文件中，我们声明了两个 `services` ，分别是 `mysql` 和 `adminer` 。在 `mysql` 中，我们指明了所使用的镜像、密码配置和重启策略；在 `adminer` 中，我们指明了所使用的镜像、重启策略和端口映射。

通过执行 `docker-compose up -d` 命令，我们即可启动 `mysql` 和 `adminer` 两个容器。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_7.png)

在成功运行上述命令后，我们即可在浏览器中访问 `:8080` 端口，进入 Adminer 管理页面。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_8.png)

输出 MySQL 账户名和密码后，我们即可进入到数据库管理页面。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_9.png)

最终，我们通过执行以下命令来终止上述两个容器的运行。

```bash
docker-compose down
```

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E5%AE%B9%E5%99%A8%E5%8C%96%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%AE%B9%E5%99%A8%E6%9C%8D%E5%8A%A1_10.png)