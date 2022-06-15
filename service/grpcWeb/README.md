
使用 [Bloom RPC](https://github.com/bloomrpc/bloomrpc/releases) 进行测试
在只有server 的情况下，可以使用BloomRPC这套工具来模拟Client 对gRPC server 发送请求，功能就类似在Restful 中使用的Postman。

前端请求和 `Bloom RPC` 测试结果一样，根本就没有到达服务端，服务端无论是否开启结果都一样。

从[官方博客](https://grpc.io/blog/state-of-grpc-web/)了解到，存在两个不同的库可以进行 `grpc-web` 的工作：
一个是 [grpc/grpc-web](https://github.com/grpc/grpc-web) 另一个是 [improbable-eng/grpc-web](https://github.com/improbable-eng/grpc-web)

两个代理库之间功能没有区别，区别在于如何部署。Envoy 将适合某些场景，而进程内 Go 代理有其自身的优势。
前者是官方 gRPC 客户端，后者支持 Fetch API。

前者支持一元和服务器端流式传输，但仅在与该grpcwebtext模式一起使用时。该模式仅完全支持一元请求 grpcweb。使用 Google Closure library 14 base 以 JavaScript 实现。它在 npm 上以grpc-web15的形式提供。它最初附带一个作为 NGINX 扩展16实现的代理，但后来加倍使用 Envoy 代理 HTTP 过滤器17，自 v1.4.0 以来的所有版本都可用。
后者支持一元和服务器端流式传输，并具有根据浏览器功能在 XHR 和 Fetch 之间自动选择的实现。用 TypeScript 实现，在 npm 上作为@improbable-eng/grpc-web10可用。还有一个可用的 Go 代理，既可以作为可以导入现有 Go gRPC 服务器的包11，也可以作为独立代理，用于将任意 gRPC 服务器公开给 gRPC-Web 前端12。
这两种模式指定了在请求和响应中对 protobuf 有效负载进行编码的不同方式。

在浏览器端没有直接的gRPC支持。
在web浏览器端，目前有3个选择，分别是：
方案一: grpc/grpc-web
作者:   gRPC官方
简介:   通过envoy进行反向代理，对gRPC服务和web http/1.1进行互相翻译。必须有一个envoy或nginx代理。客户端通过protoc生成js或ts代码
特点:   官方支持。直接生成commonJS或TS客户端代码。需要独立proxy（envoy或nginx）
地址:   https://github.com/grpc/grpc-web
示例:   https://github.com/grpc/grpc-web/blob/master/net/grpc/gateway/examples/helloworld/
备注:   默认使用Envoy作为网关代理，对Envoy的支持已经内置在gRPC-Web 库里

方案二: improbable-eng/grpc-web
作者:	improbable-eng
简介:	嵌入代码在gRPC服务器端，直接把gRPC协议翻译为gRPC-web协议。同时生成浏览器端TypeScript代码，供浏览器直接调用服务。
特点:	直接生成TSt客户端代码，可在浏览器直接使用。proxy代理内嵌在服务器端。
地址:   https://github.com/improbable-eng/grpc-web
示例:   https://github.com/improbable-eng/grpc-web/blob/master/client/grpc-web


方案三: grpc-ecosystem/grpc-gateway
作者:	grpc-ecosystem
简介:	通过protoc插件的方式，对protobuf定义里的annotation进行处理，自动生成反向代理服务器代码以及RESTful API 的swagger 描述信息文件。
特点:	直接生成RESRful API代理服务器、REST API文档。需要手写浏览器端RESTful client代码。


improbable-eng/grpc-web简介
improbable-eng/grpc-web 库基于 Golang 和 TypeScript:

grpcweb - 这是一个可以把现有 grpc.Server 包装成 gRPC-web http.Handler 的 GoLang 包。同时支持HTTP2和HTTP/1.1。
grpcwebproxy - 基于GoLang的独立反向代理服务，为经典 gRPC 服务（比如：java或c++）提供代理，把他们的服务通过 gRPC-web 暴露给现代浏览器。
ts-protoc-gen - protoc (protocol buffers 编译器) 的 TypeScript 代码生成插件，生成 TypeScript 强类型消息类和方法定义代码。
@improbable-eng/grpc-web - 供浏览器和 NodeJS 使用的 TypeScript gRPC-Web 客户端库。

为什么使用 improbable-eng/grpc-web?
对于 gRPC-Web, 要创建定义良好、易于解释的浏览器前端代码和微服务之间的API，是一件及其简单的事情。前端开发有以下意义重大的好处：

不再需要到处寻找API文档 —— .proto 就是典型的API规范格式。
不在需要手写 JSON 调用对象 —— 所有的请求和响应都是强类型、代码自动生成的，我们在IDE可以方便的使用代码自动提示，提高编程效率。
不再需要处理HTTP的method、header、body以及底层网络 —— 所有事情都有 grpc.invoke 处理了。
不用需要反复猜测错误代码的含义 —— gRPC 状态码是典型的在 API 中表示问题的方法。
不需要为了避免并发链接而采用低效的一次性网络连接 —— gRPC-Web 是基于 HTTP2的，支持一个连接上多路传输多路数据流。
没有从服务器读取流式数据的问题 —— gRPC-Web 支持一对一远程调用和一对多数据流请求。
处理新二进制时，更少的数据处理错误 —— 前端、后端同时兼容的请求及响应.
总而言之，gRPC-Web 把前端代码和微服务之间的交互从手工编写HTTP请求代码转换成已经定义好的用户业务函数。



grpc-ecosystem/grpc-gateway简介
这个 grpc-gateway 是Google protocol buffers 编译器protoc的一个插件。它读取 protobuf 服务定义文件，然后生成 一个反向代理服务器，用来“把一个RESTful API调用翻译成gRPC”。这个服务器根据你的定义文件中的 google.api.http annotation来生成的。这可以让你在服务器里同时提供给RPC和RESTful风格的API服务。
grpc-ecosystem/go-grpc-middleware简介
gRPC的调用需要考虑安全认证问题。可以使用 grpc-ecosystem/go-grpc-middleware 实现。
在 gRPC 服务器中可以使用中间件 参考：用GoLang开发gRPC service中间件 在用户应用逻辑调用之前，或在客户端调用服务器之前执行相关操作。这是实现用户认证auth, 日志logging, 消息message, 数据验证validation, 重试或性能监控的完美方式。


两种实现方式，均通过 子目录 web(evnoy) 和 vue(improbable-eng) 中的 docker-composer.yaml 方式测试启动
