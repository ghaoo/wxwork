## 配置rboot机器人
[文档](https://www.kancloud.cn/ghaoo/rboot/1476890)

## ssh + nginx 内网穿透

### 在可以访问外网的服务器安装 Nginx
1. 创建 nginx 配置文件
    `88` 为外网访问地址
    `8888` 用来转发 88 端口的内容
    ```conf
    upstream _proxy {
            server 127.0.0.1:8888;
    }
    
    server {
            listen 88; 
            server_name localhost;
            location / {
                    proxy_pass http://_proxy/;
            }
    }
    ```
   
2. 加载配置文件

### 在内网上建立 ssh 反向隧道/远程转发

即把服务器请求转发到本地
```bash
ssh -fCNR [B机器IP或省略]:[B机器端口]:[A机器的IP]:[A机器的sshd端口] [登录B机器的用户名@B机器的IP] -p [B机器的sshd端口]
```
内网机器的服务端口是8080，反向代理到外网机器的8999端口，在内网机器上操作：
```bash
ssh -fCNR 8999:localhost:8080 -o ServerAliveInterval=60 root@123.123.123.123 -p 22
```
