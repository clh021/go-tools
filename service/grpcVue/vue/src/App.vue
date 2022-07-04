<template>
  <div>
    消息内容:<input type="text"> <br>
    <button @click="Echo">Echo</button> <br>
    <button @click="EchoAbort">EchoAbort</button> <br>
    <button @click="NoOp">NoOp</button> <br>
    <button @click="ServerStreamingEcho">ServerStreamingEcho</button> <br>
    <button @click="ServerStreamingEchoAbort">ServerStreamingEchoAbort</button> <br>
    <button @click="ClientStreamingEcho">ClientStreamingEcho</button> <br>
    <button @click="FullDuplexEcho">FullDuplexEcho</button> <br>
    <button @click="HalfDuplexEcho">HalfDuplexEcho</button>
  </div>
</template>

<script>

import { grpc } from "@improbable-eng/grpc-web";
import { EchoService, EchoServiceClient } from "./lib/echov_pb_service";
import {
  ClientStreamingEchoRequest,
  EchoRequest,
  Empty,
  ServerStreamingEchoRequest
} from "./lib/echov_pb";
const locat = document.location;
const apiPort = location.port == '8081' ? 18080 : 18081;
const host = `${locat.protocol}//${locat.hostname}:${apiPort}`;
console.log("EchoServiceClient:", EchoServiceClient);

// 'myTransport' is configured to send Browser cookies along with cross-origin requests.
// const myTransport = grpc.CrossBrowserHttpTransport({ withCredentials: true });
// Specify the default transport before any requests are made.
// grpc.setDefaultTransport(myTransport);
export default {
  name: 'App',
  data() {
    return {
      client: null,
    }
  },
  created() {
    // this.client = new EchoServiceClient(host)
  },
  methods:{
    requestGrpcUnary(tag, func, req) {
      console.log(`==========================\n${tag}`)
      grpc.unary(func, {
        request: req,
        host: host,
        onMessage: res => {
          console.log("all ok. got onMessage: ", JSON.stringify(res));
        },
        onEnd: res => {
          // console.log("all ok. got res: ", JSON.stringify(res));
          // const { status, statusMessage, headers, message, trailers } = res;
          // console.log("status: ", status);
          // console.log("statusMessage: ", statusMessage);
          // console.log("headers: ", headers);
          // console.log("trailers: ", trailers);
          const { status, message } = res;
          if (status === grpc.Code.OK && message) {
            console.log("all ok. got onEnd: ", message.toObject());
          }
        }
      });
    },
    requestGrpcInvoke(tag, func, req) {
      console.log(`==========================\n${tag}`)
      grpc.invoke(func, {
        request: req,
        host: host,
        onMessage: res => {
          console.log("all ok. got onMessage: ", JSON.stringify(res));
        },
        onEnd: res => {
          // console.log("all ok. got res: ", JSON.stringify(res));
          // const { status, statusMessage, headers, message, trailers } = res;
          // console.log("status: ", status);
          // console.log("statusMessage: ", statusMessage);
          // console.log("headers: ", headers);
          // console.log("trailers: ", trailers);
          const { status, message } = res;
          if (status === grpc.Code.OK && message) {
            console.log("all ok. got onEnd: ", message.toObject());
          }
        }
      });
    },
    initClient() {

    },
    requestGrpcClient(tag, func, req) {
      console.log(`==========================\n${tag}`)
      const client = grpc.client(func, {
        debug: true,
        host: host,
        // transport: myTransport,
        transport: grpc.WebsocketTransport(),
      });
      client.onHeaders((headers) => {
        console.log("onHeaders", headers);
      });
      client.onMessage((message) => {
        console.log("onMessage", message, message.messageCount);
      });
      client.onEnd((status, statusMessage, trailers) => {
        console.log("onEnd", status, statusMessage, trailers);
      });
      client.start();
      for (let i = 0; i < 5; i++) {
        setTimeout(() => {
          req.setMessage(`${i} Hello` + (new Date()).toISOString());
          client.send(req);
        }, i*1000)
      }
      setTimeout(() => {
        client.finishSend();
      }, 6000)
    },
    Echo() {
      const req = new EchoRequest();
      req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcUnary("Echo", EchoService.Echo, req);
    },
    EchoAbort() {
      const req = new EchoRequest();
      req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcUnary("EchoAbort", EchoService.EchoAbort, req);
    },
    NoOp() {
      const req = new Empty();
      this.requestGrpcUnary("NoOp", EchoService.NoOp, req);
    },
    ServerStreamingEcho() {
      const req = new ServerStreamingEchoRequest();
      this.requestGrpcInvoke("ServerStreamingEcho", EchoService.ServerStreamingEcho, req);
    },
    ServerStreamingEchoAbort() {
      const req = new ServerStreamingEchoRequest();
      req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcInvoke("ServerStreamingEchoAbort", EchoService.ServerStreamingEchoAbort, req);
    },
    ClientStreamingEcho() {
      const req = new ClientStreamingEchoRequest();
      req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcClient("ClientStreamingEcho", EchoService.ClientStreamingEcho, req);
      // 一次获取客户端，多次发送
      // 然后服务端应该可以不停的打印出每次发送过去的消息
      // 服务端还返回收到消息的总次数
    },
    FullDuplexEcho() {
      const req = new EchoRequest();
      // req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcClient("FullDuplexEcho", EchoService.FullDuplexEcho, req);
    },
    HalfDuplexEcho() {
      const req = new EchoRequest();
      // req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcClient("HalfDuplexEcho", EchoService.HalfDuplexEcho, req);
    }
  },
  mounted() {
    console.log("mounted")
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
