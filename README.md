# go-nginx

一个使用go实现的仿niginx的高性能web服务器

使用fasthttp时静态页面的性能：
![fasthttp.png](md/images/fasthttp.png)

使用net.http时静态页面的性能：
![img.png](md/images/net.http.png)

可以看出fasthttp的性能确实强悍，而nginx最大的特点就是性能强悍，同时由于我们对于路由匹配策略的需求不是很强，所以我们这里就采用fasthttp作为我们的http框架。
