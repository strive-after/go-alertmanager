一 项目简介
1.module定义了从alert 那里接受信息的结构体以及以md格式发送至叮叮的结构体

2.notifier从alert接受信息

3.transformer针对接受结构体我们那一些字段出来拼接成一个字符串整体来发送给叮叮

4.叮叮发送格式https://ding-doc.dingtalk.com/doc#/serverapi3/iydd5h

二 项目部署

git clone  https://github.com/strive-after/go-alertmanager.git

可执行程序是linux系统版本

1.物理机部署

因为这个程序需要读取环境变量token

设置环境变量token=xxxx

这个值是你叮叮机器人的token

./webhook 就可以启动

2.docker部署参照k8s部署方式不过build镜像的时候需要修改dockerfile 里面加环境变量token=xxx

3.k8s部署

需要打包项目目录中有dockerfile 

docker build -t xxxx  .

deploy yaml跟svcyaml  直接apply就可以吧镜像跟token替换成自己的

三

部署完成之后alert需要修改webhook的对接地址才可以发送消息

程序默认监听端口8080  url为/webhook

如果你物理机部署那么 ip:8080/webhook 这个就是alert的连接webhook的地址