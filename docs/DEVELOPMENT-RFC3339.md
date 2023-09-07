# 开发者手册-RFC3339

本手册为 Stash-RFC3339 改版的开发者手册，主要记录 RFC3339 改版的特殊流程，本手册没有记录的内容均以原版手册 DEVELOPMENT.md 为准

## Windows 上使用 WSL 编译 Linux 版的步骤

### 1. 安装 Ubuntu

1. 启动 Windows 的 WSL 功能
2. wsl --install Ubuntu（安装 Ubuntu，会要求创建新角色和密码，跟着走就行，初始化完成后自动进入 Ubuntu 控制台）
3. sudo -i（直接获取 root 权限，省得之后麻烦）

### 2. 安装较为简单的包（此步骤开始尽可能让所有的网络请求走代理，直到编译完成为止，可以使用 TUN 或虚拟网卡）

1. apt-get update（升级安装的包，刚安装好的 Ubuntu 是老版的需要升级）
2. apt-get install git gcc ffmpeg -y（安装 go 和 nodejs 以外的需要用的包，这两个包 Ubuntu 库里的版本太低了要装新的）
3. apt install make（安装 make 功能）

### 3. 安装 Go

1. arch（查看系统架构）
2. 去官方安装手册那里下载对应架构的文件并放到 Ubuntu 的文件目录里，然后命令行进入这个目录
3. rm -rf /usr/local/go（移除可能已经安装的旧的 Go）
4. tar -C /usr/local -xzf 包文件名（安装）
5. export PATH=$PATH:/usr/local/go/bin（设置 PATH）
6. go version（检查 Go 版本，检查到则说明安装成功）

### 4. 安装 Nodejs 并开启所需服务
1. apt-get install -y ca-certificates curl gnupg（安装了一些需要的包）
2. mkdir -p /etc/apt/keyrings（准备安装位置）
3. curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg（似乎是下载什么东西）
4. NODE_MAJOR=20（设置要安装的版本）
5. echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | sudo tee /etc/apt/sources.list.d/nodesource.list（安装？）
6. apt-get update（这一步是有用的，前面的操作会导进一些老的包需要更新）
7. apt-get install nodejs -y（安装）
8. node -v（检查 nodejs 版本，检查到则说明安装成功）
9. corepack enable（开启 Nodejs 的核心包）
10. corepack prepare yarn@stable --activate（启动核心包的 yarn 服务）

### 5. 编译

1. 进入源码所在目录，需要和最外层的 README 在一个层级
2. make release（编译）
3. 编译完成后的文件将在当前目录，文件名为"stash"，没有后缀名

### 6. 使用

1. 将编译后的文件放到主机里
2. ./stash（Linux 的运行文件指令）
3. nohup ./stash-linux >logs/20230828.log 2>&1 &（不占用 shell 窗口、不随窗口关闭停止、输出log、错误信息视为log输出）

## Windows 上使用 WSL 编译 Docker Image 的步骤

### 0. 首先完成前面的【使用 WSL 编译 Linux 版的步骤】部分

### 1. 安装 Docker

1.  snap install docker（安装 Docker）

### 2. 编译

1. 进入源码所在目录，需要和最外层的 README 在一个层级
2. make docker-build（编译，编译后的 Image 在 Docker 内部）
3. docker image list（查看 Docker 内部的 Image 列表，找到其中那个名字带 Stash 的）
4. docker save stash/build -o ./stash-image（将 Image 取出来。stash/build 这部分对应上一步找到的那个 Image 的名字；./stash-image 是保存文件，表示存在当前目录的 stash-image 文件）
5. chmod 777 stash-image（修改导出的 Image 文件为所有用户都有所有权限，防止因为 Windows 用户对 WSL 文件权限不足导致无法复制和移动这个文件）

### 3. 使用

1. 把映像移动到主机存储目录里
2. docker load -i stash-image（将映像导入 Docker，stash-image 是映像文件的文件名）
3. 按照 docker 的标准方式创建容器之后运行，或者使用 docker-compose 运行，这部分在冷备份里有。

## Docker 版到 Linux 版的数据迁移

由于跨版本设置是无法迁移的，不仅是语言设置、显示设置

### 1. 准备迁移的数据

1. 进入 Stash 的设置界面，在任务里找到备份，点击备份。这个备份是对数据库进行备份，备份的数据库就在数据库的同级目录。
2. 将 config.yml、icon.png、stash-go.sqlite 之外的内容复制出来。config 因为跨平台无法通用；icon 都一样不需要复制；sqlite 文件因为 SQLite 的运行时有缓存的特性不能用原来的要用备份的。【注意】如果此步骤出现了类似权限的问题，使用 chmod 777 -R ./* 将目录下的所有文件设为所有用户都具备完全操作权限

### 2. 新 Stash 服务的准备
1. 开启新的 Stash，正常进行初始化。
2. 初始化完毕后关闭新的 Stash，数据迁移期间不能开机否则数据可能错乱。

### 3. 数据转移
1. 将从旧 Stash 复制出来的数据覆盖到新 Stash 的目录里，位置要对应。【注意】如果此步骤出现了类似权限的问题，使用 chmod 777 -R ./* 将目录下的所有文件设为所有用户都具备完全操作权限
2. 启动新的 Stash，此时可以迁移的数据已经完成了迁移但是很可能无法正常显示。
3. 将旧的 Stash 的设置在新 Stash 里重新设置一遍，尤其是收藏库的目录一定要相同。
4. 网页清空缓存重载界面，此时应当正常显示。