# iot-message-console
IOT MESSAGE FOR IOT EDGE SERVER

目的
* 更方便的查看日志
* 避免损伤嵌入式设备的存储器


访问地址

```
http://iotlog.doublechaintech.com/

```
nginx 配置

````
upstream websocket {
    server localhost:8181;
}

server {
#	listen 80 default_server;
#	listen [::]:80 default_server;

	root /var/www/html;

	# Add index.php to the list if you are using PHP
	index index.html index.htm index.nginx-debian.html;

	server_name iotlog.doublechaintech.com;

	location / {
		# First attempt to serve request as file, then
		# as directory, then fall back to displaying a 404.
	proxy_pass http://websocket;
    	proxy_http_version 1.1;
    	proxy_set_header Upgrade $http_upgrade;
    	proxy_set_header Connection "Upgrade";


	}

	location /ws {
	proxy_pass http://websocket;
    	proxy_http_version 1.1;
    	proxy_set_header Upgrade $http_upgrade;
    	proxy_set_header Connection "Upgrade";
	proxy_set_header Origin "";

	}


}


````
