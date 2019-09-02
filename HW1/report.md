# 实验报告：搭建云桌面私有云服务

> 16305204 郑佳豪

## 云桌面

相信中山大学的同学对**云桌面**并不陌生，每当期中考和期末考时，计算机和软件工程的学生都得在某云桌面上，进行上机考试。腾讯云主机、阿里云主机、Vultr 以及搬瓦工，都在提供与云桌面类似的服务（VPS）。

云桌面，是当前**移动办公**的高效解决方案。在使用云桌面技术后，用户无需再购买电脑主机，主机所包含的 CPU、内存、硬盘、网卡等组件全部在后端的服务器中虚拟出来，用户在安装客户端后，通过特有的通信协议访问后端服务器上的虚拟机主机来实现交互式操作，达到与电脑一致的体验效果。

## 搭建服务

现在，就让我们使用 VMware Workstation 搭建私有的云桌面服务。

### 下载

- 系统镜像

  我们在中科大镜像源下载 [CentOS7](http://mirrors.ustc.edu.cn/centos/7.6.1810/isos/x86_64/) ，这里我选择的是 CentOS-7-x86_64-Minimal-1810.iso 镜像文件。

- VMware Workstation

### 创建虚拟机

当我们将系统镜像文件下载完毕后，需要在 VMware Workstation 创建虚拟机，并安装 CentOS 。按照提示，我们应该很快完成安装步骤，随即进入系统的 CLI 界面。在安装过程中，我们设定 Root Password 为 `jiahonzheng123456` ，同时创建了用户名为 `centos` 且密码为空的用户。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_1.jpg)

### 配置网络

在创建完虚拟机后，我们需要对虚拟机和主机的网络进行相关的配置。

#### 开启网络服务

由于我们在创建虚拟机时，选择的是 NAT 网卡模式，进入系统后，理论上来说，系统是可以访问外网的，但在执行完 `ping www.baidu.com` 后，我们发现系统返回以下的错误。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_3.jpg)

这是因为 CentOS7 并未默认开启网络服务。为了开启网络服务，我们需要修改 `/etc/sysconfig/network-scripts/` 目录下的网卡配置 `ONBOOT` 条目。

```bash
ls /etc/sysconfig/network-scripts/
```

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_4.jpg)

执行 `ls` 命令，结果如上图，其中的 `ifcfg-ens33` 即为网卡配置文件，我们使用 `vim` 修改其中的 `ONBOOT` 条目为 `yes` 。

```bash
vi /etc/sysconfig/network-scripts/ifcfg-ens33
```

```bash
# 原本是 ONBOOT=no ，将其修改为以下形式
ONBOOT=yes
```

修改完 `ONBOOT` 的值为 `yes` 后，我们重启网络服务，使更改生效。

```bash
service network restart
```

经过上述的设置，现在系统已可正常访问外网。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_5.jpg)

#### 设置虚拟网卡

为实现云桌面服务，主机需要创建两张虚拟网卡，一张为 NAT 模式，一张为仅主机模式。NAT 模式的虚拟网卡使子网内的所有虚拟机可共享主机网络，即可访问外网；仅主机模式的虚拟网卡则作为局域网的 DHCP 服务器。

在 VMware 的 Edit -> Virtual Network Editor 窗口中，我们可以看到 VMware 已自动创建上述两张虚拟网卡，其中 VMnet1 为仅主机模式，VMnet8 为 NAT 模式。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_2.jpg)

仅仅设置了主机的虚拟网卡是不够的，我们还需要为虚拟机添加新的虚拟网卡（由于我们使用 NAT 模式创建了虚拟机，故虚拟机已经有一张连接 VMnet8 的 NAT 虚拟网卡），我们需要将新添加的虚拟网卡设置为仅主机模式，让其连接 VMnet1 。

我们可以在 VMware 的 VM -> Settings 中为虚拟机创建和配置虚拟网卡。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_6.jpg)

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_7.jpg)

### SSH 登录

在配置完虚拟机的网络后，我们可以尝试在主机上使用 `SSH` 连接到虚拟机。首先，我们需要执行 `ip addr` 命令，获取虚拟机的 IP 地址，如下图所示，当前虚拟机的 IP 地址为 `192.168.11.128` 。随后，我们在主机的 `Bash` 上，执行 `ssh root@192.168.11.128` 命令，输入密码后，即可进入虚拟机的远程 CLI 中。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_8.jpg)

### 远程桌面登录

到这里，不知大家有没有发现，我们创建的虚拟机是没有图形化桌面的，只有黑白色的 CLI 窗口，为了完成图形化的远程桌面控制，我们需要为其安装 GUI 界面，这里我们选择的是 `GNOME` 的桌面图形界面。

执行以下命令，安装 `GNOME Desktop` 。

```bash
yum groupinstall "GNOME Desktop"
```

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_9.jpg)

在安装完毕后，我们需要设置 GUI 界面开启自启动，并且重启虚拟机使其生效。

```bash
ln -sf /lib/systemd/system/runlevel5.target /etc/systemd/system/default.target
shutdown -r now
```

重启完成后，点击 `centos` 的用户头像（由于我们设置其密码为空，故点击即进入系统桌面），即可进入系统桌面。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_10.jpg)

在安装完图形化界面后，我们需要进行远程桌面的相关配置。我们知道，Windows 自带的远程桌面协议是 `RDP` 协议，这种协议是 Linux 原生不支持的，因此我们需要使用第三方软件，`xrdp` 是一个很好的选择。

```bash
# 进入 root 模式：执行 su 命令，输入密码
yum install -y epel-release
# 这一步很关键！
yum -y install
https://archive.fedoraproject.org/pub/archive/epel/7/x86_64/Packages/x/xorgxrdp-0.2.9-1.el7.x86_64.rpm
# 安装 xrdp 和 tigervnc-server
yum -y install xrdp tigervnc-server
```

安装完 `xrdp` 和 `tigervnc-server` 后，我们启动 `xrdp` 服务，并添加至开机自启动项。

```bash
# 启动 xrdp 服务
systemctl start xrdp
# 添加 xrdp 至开机自启
systemctl enable xrdp
```

随后，我们执行 `netstat -antup | grep xrdp` 命令查看 xrdp 服务的运行状态。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_13.jpg)

从上图可知，`xrdp` 服务监听了 `3389` 端口，这端口号其实也是 Windows 下的远程桌面监听端口号。

我们需要配置 `firewall` 防火墙以及 `SELinux` 安全上下文。

```bash
# 放通 3389 端口上的所有 TCP 请求
firewall-cmd --permanent --add-port=3389/tcp
# 重载防火墙，使规则生效
firewall-cmd --reload
# 设定 xrdp 和 xrdp-sesman 的目标安全环境
chcon --type=bin_t /usr/sbin/xrdp
chcon --type=bin_t /usr/sbin/xrdp-sesman
```

在执行完上述指令后，我们即可在 Windows 的远程桌面连接工具下，享受虚拟机所提供的远程桌面服务。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_12.jpg)

用户名输入 `centos` ，密码为空，点击 `OK` 后，即可进入虚拟机远程桌面，如下图。

![](https://jiahonzheng-blog.oss-cn-shenzhen.aliyuncs.com/%E6%90%AD%E5%BB%BA%E4%BA%91%E6%A1%8C%E9%9D%A2%E7%A7%81%E6%9C%89%E4%BA%91%E6%9C%8D%E5%8A%A1_11.jpg)

### 克隆虚拟机

经过上述流程，我们已实现了一个可对外提供远程桌面服务的虚拟机实例。我们知道，云服务厂商肯定不会只有一个虚拟机实例，那么如何拥有大量的虚拟机实例呢？我们可以通过 VMware 提供的虚拟机复制功能，批量生产这些符合要求的虚拟机实例。

在 VMware Workstation 中，当我们进行虚拟机复制时，它会生成一个当前虚拟机的快照，随后使用该快照生成一个独立的虚拟机实例，且支持**链接复制**，使得多个虚拟机可共享同一份快照文件，进而节约了存储空间，降低了服务成本。

我们可在 VMware 的 VM -> Manage -> Clone 进行虚拟机的批量生产。

## 参考资料

- [How to interpret this yum dependency error?](https://serverfault.com/questions/979672/how-to-interpret-this-yum-dependency-error)
- [CentOS 7 如何安装 xrdp ？](https://www.jb51.net/os/RedHat/531360.html)
