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

### 编译完成后的文件将在当前目录，文件名为"stash"，没有后缀名。此文件直接在 Linux 系统中使用 ./stash 即可运行