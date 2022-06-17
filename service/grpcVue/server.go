package grpcVue

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"runtime/debug"
	"test/service/grpcVue/echovue"
)

func getWrapServer(grpcServer *grpc.Server) *grpcweb.WrappedGrpcServer {
	options := []grpcweb.Option{
		grpcweb.WithAllowedRequestHeaders([]string{"*", "X-Nanoapp-Remoteaddr", "X-Nanoapp-Domain"}),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(*http.Request) bool { return true }),
	}

	if os.Getenv("NPROXY_DISABLE_WS_PING") == "" {
		// see also: https://github.com/improbable-eng/grpc-web/issues/713
		options = append(options, grpcweb.WithWebsocketPingInterval(time.Second*10))
	}
	return grpcweb.WrapServer(grpcServer, options...)
}

// func setupDomain(req *http.Request) error {
// 	origin := req.Header.Get("Origin")
// 	if strings.HasPrefix(origin, "http://") {
// 		origin = origin[7:]
// 	} else if strings.HasPrefix(origin, "https://") {
// 		origin = origin[8:]
// 	} else {
// 		return errors.New("invalid origin header")
// 	}

// 	domain := strings.Split(origin, ":")[0]
// 	if domain == "localhost" || domain == "127.0.0.1" {
// 		domain = "*"
// 	}
// 	req.Header.Set("x-nanoapp-domain", domain)
// 	return nil
// }

func Main() {

	// 退出时自动恢复 服务
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s\n\n%s\n", r, debug.Stack())
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 18081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	echoService := echovue.Server{}

	grpcServer := grpc.NewServer()

	echovue.RegisterEchoServiceServer(grpcServer, &echoService)
	wrappedGrpc := getWrapServer(grpcServer)

	log.Println("Register NProxy at", lis.Addr().String())

	http.Serve(lis, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		// isOption := req.Method == "OPTIONS"

		// if !isOption {
		// 	// 设置http client的TCP地址信息，以便可以反向查找到请求者的UID。
		// 	req.Header.Set("x-nanoapp-remoteaddr", req.RemoteAddr)
		// 	if setupDomain(req) != nil {
		// 		resp.WriteHeader(http.StatusBadRequest)
		// 		return
		// 	}

		// 	for _, plg := range d.routePlugins {
		// 		if !plg.BeforeForward(req, resp) {
		// 			return
		// 		}
		// 	}
		// }

		wrappedGrpc.ServeHTTP(resp, req)

		// if !isOption {
		// 	for _, plg := range d.routePlugins {
		// 		plg.AfterForward(req, resp)
		// 	}
		// }
	}))
}
