## 设计原则

1. 尽可能小的体积(功能扩充时，体积是重用的)
2. 尽可能简单的使用方式

## 目标功能清单

> 对比 shell 脚本方式，不能比 shell 体积更大，不能额外增加环境依赖
 
1. MySQL备份 更适合通过脚本或者客户端工具搞定
2. MySQL恢复 更适合通过脚本或者客户端工具搞定
3. MySQL查询 更适合通过脚本或者客户端工具搞定
4. MySQL管理 更适合通过脚本或者客户端工具搞定
5. SSH 管理
6. 计划任务管理  更适合通过脚本或者系统服务搞定
7. 计划任务日志  更适合通过脚本或者系统服务搞定
8. 文件修改
9. 文件在线管理 FileBrowse
10. 把 caddy 融合进来，工具增强，特别是 blog,ssl 服务
11. 前端打开限制条件(电脑设备，操作系统，电池，用户名，)
    ```
    Continent, Country and city
    Date of last visit
    Referring website or search engine (including search term)
    Time spent on the website
    Browser and operating system
    IP Address
    Language
    Browser
    OS
    Screen size
    ```
12. 完全基于规则的数据库操作接口
13. Spf13 工具用上
14. 邮件通知接口，短信通知接口，微信通知接口，短链接通知接口(二次验证消息)
15. NSQ 	NSQ 是无中心设计、节点自动注册和发现的开源消息系统。可作为内部通讯框架的基础，易于配置和发布。
16. go-tcp-proxy 	go-tcp-proxy是一个简单的tcp代理， 可以用于tcp端口转发， 也可以用做http代理使用
17. myLG是一个开源的网络工具集，它包含了很多不同类型的网络诊断工具, 功能包括ping，trace， bgp， dns lookup， 端口扫描， 局域网网络发现，提供web界面, tcpdump等
18. cow 使用二级代理；支持多种协议：sock5、http
19. Tyk 是一个开源的、轻量级的、快速可伸缩的 API 网关，支持配额和速度限制，支持认证和数据分析，支持多用户多组织，提供全 RESTful API。
20. KodeRunr (读作 code runner) 是款我在闲暇时间用Go语言编写的应用。顾名思义，你可以用它在网页上、命令行里写程序，贴代码，与此同时无需在本地安装任何编程语言。支持Ruby, Python, GO, Swift, C, Elixir等
21. scheduler scheduler专门进行任务的调度分发任务工作，各个任务的具体任务执行分配到各个项目中， 从而达到对任务的统一配置和管理。 该工具提供了web管理界面
22. 本地通知
23. 模拟键盘鼠标操作
24. 截取屏幕，识别屏幕文字
25. 浏览器环境检测
26. 检查客户是否开启了开发者工具栏，如果开启了就尝试关闭或者停止服务。

## 开发技术相关

见 service/README.md