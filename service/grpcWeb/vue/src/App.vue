<template>
  <HelloWorld msg="Welcome to Your Vue.js App"/>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

import { grpc } from "@improbable-eng/grpc-web";
import { EchoService, EchoServiceClient } from "./lib/echo_pb_service";
import { EchoRequest } from "./lib/echo_pb";
const host = "http://192.168.1.68:18080";

export default {
  name: 'App',
  components: {
    HelloWorld
  },
  created: () => {
    const req = new EchoRequest();
    const grpcRequest = grpc.unary(EchoService.Echo, {
      request: req,
      host: host,
      onMessage: res => {
        console.log("all ok. got onMessage: ", JSON.stringify(res));
      },
      onEnd: res => {
        // const { status, statusMessage, headers, message, trailers } = res;
        // if (status === grpc.Code.OK && message) {
        //   console.log("all ok. got onEnd: ", message.toObject());
        // }
        console.log("all ok. got onEnd: ", JSON.stringify(res));
      }
    });
    grpcRequest.close();
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
