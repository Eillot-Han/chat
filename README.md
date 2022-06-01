# chat


### 开启redis服务
```
> $ cd到本机redis
> $ redis-server.exe redis.windows.conf
```
### 开启后端接口
```
> $ cd chet-golang
> $ go run main.go
```
### 开启websocket连接
```
> $ cd chet-websocket
> $ go run main.go
```
### 打开前端页面
```
> $ cd chet-react
> $ npm i 安装依赖
> $ npm run start
```

账号：789386699 密码：123
账号：1223344 密码：123
账号：155594166 密码：123

题目：用go实现一个即时通讯聊天程序： - 注册/登录/下线/注销 - 一对一单聊 - 多人群聊 - 单聊/群聊中的文件发送与接收 - 聊天记录的保存与查找

后端接口全部实现
由于时间不够，前端实现了登录/下线/一对一单聊/多人群聊/所有历史纪录显示/根据内容查找聊天记录/保存聊天记录/增加群聊/删除群聊/删除好友
