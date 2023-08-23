## 开发环境搭建

### MySQL安装配置 

1. 安装mysql-server
2. 打开/etc/mysql/dibian.cnf 找到默认用户名密码
3. mysql -u debian-sys-maint -p
4. alter user root@localhost identified with mysql_native_password by '这里是密码';
5. 在Ubuntu software软件中心搜索并安装mysql-workbench-community

	解决无法链接问题sudo snap connect mysql-workbench-community:password-manager-service :password-manager-service

### GO环境配置

1. https://go.dev/dl/ 下载对应版本的压缩包
2. 解压到/usr/local/go
3. sudo vim ~/.bashrc
4. export GOROOT=/usr/local/go
5. export GOPROXY=https://goproxy.cn
6. export PATH=$PATH:$GOROOT/bin

### Vue环境安装

1. https://nodejs.org/zh-cn/download/ 下载node
2. 解压到/usr/local目录
3. sudo ln /usr/local/node-v16.18.0-linux-x64/bin/node  /usr/local/bin/node
4. sudo ln -s /usr/local/node-v16.18.0-linux-x64/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm

### 初始化（首次编译运行前需要执行）

1. 在mysql创建数据库goadmin, SQL: create database goadmin;
2. cd ./go-admin修改数据库相关配置 ./config/setting.yaml
3. 安装后端依赖项
go get go-admin/common/middleware

go get go-admin/common/file_store

go get github.com/go-admin-team/go-admin-core/sdk/pkg/casbin@v1.4.1-0.20220809101213-21187928f7d9

go get github.com/go-admin-team/go-admin-core/sdk/pkg/ws@v1.4.1-0.20220809101213-21187928f7d9

go get github.com/go-admin-team/go-admin-core/sdk/pkg/utils@v1.4.1-0.20220809101213-21187928f7d9

go get go-admin/cmd

go build

./go-admin migrate

4. 安装前端依赖项 cd ./go-admin-ui

npm install (报错时添加--force)

### 允许许远程访问的修改
	VUE_APP_BASE_API = 'http://localhost:8000'
	改为VUE_APP_BASE_API = 'http://10.40.29.152:8000'
	修改数据库表sys_config中的系统logo : localhost改为10.40.29.152


### 启动

1. 启动后端 ./go-admin server

2. 启动前端 npm run dev



### 客户端编译
1. sudo apt install curl 
   sudo apt-get install libcurl4-openssl-dev

curl -X POST  -H "Content-Type:application/json" http://10.40.29.152:8000/api/v1/model-compile-apply/update_noauth/16 --data '{"id":16,"applyTime":"2022-11-14T15:28:20+08:00","applicant":"刘方盛1","status":"failed","name":"啊啊啊","type":"啊啊","onnx":"e2b186d17a4ad4fed6a5ef62f36ab340.png","onnx_local":"住宿2.png","md5":"啊","nspCnt":"2","compileType":"INT8","channelOrder":"NHWC","calibDataset":"","postCode":"","createdAt":"2022-11-14T15:28:42+08:00","updatedAt":"2022-11-14T15:36:05+08:00","createBy":1,"updateBy":1}'
