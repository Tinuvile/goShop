package main

import (
	"github.com/Tinuvile/goShop/demo/auth/biz/dal"
	"github.com/joho/godotenv"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
	"time"

	"github.com/Tinuvile/goShop/demo/auth/conf"
	"github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 初始化服务并启动服务器
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	opts := kitexInit() // 初始化配置

	dal.Init()

	svr := authservice.NewServer(new(AuthServiceImpl), opts...) // 创建服务实例

	err = svr.Run() // 启动服务
	if err != nil {
		klog.Error(err.Error())
	} // 错误日志
}

// 配置Kitex服务参数，包括地址、服务注册、日志等
func kitexInit() (opts []server.Option) {
	// address 地址
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info 服务发现
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		//ServiceName: conf.GetConf().Kitex.Service,
		ServiceName: "auth",
	}))

	// Consul服务注册
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
