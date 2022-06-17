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
    <button @click="HalfDuplexEcho">HalfDuplexEcho</button> <br>
    单边消息
    双边数据流消息
  </div>
</template>

<script>

import { grpc } from "@improbable-eng/grpc-web";
import { EchoService, EchoServiceClient } from "./lib/echo_pb_service";
import {
  ClientStreamingEchoRequest,
  EchoRequest,
  Empty,
  ServerStreamingEchoRequest
} from "./lib/echo_pb";
const host = "http://192.168.1.68:18081";
console.log("EchoServiceClient:", EchoServiceClient);

export default {
  name: 'App',
  data() {
    return {
      client: null,
    }
  },
  created() {
    this.client = new EchoServiceClient(host)
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
    requestGrpcClient(tag, func, req) {
      console.log(`==========================\n${tag}`)
      grpc.client(func, {
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
    },
    FullDuplexEcho() {
      const req = new EchoRequest();
      req.setMessage("Hello" + (new Date()).toISOString());
      this.requestGrpcClient("FullDuplexEcho", EchoService.FullDuplexEcho, req);
    },
    HalfDuplexEcho() {
      const req = new EchoRequest();
      req.setMessage("Hello" + (new Date()).toISOString());
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
