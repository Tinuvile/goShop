//package main
//
//import (
//	"fmt"
//	"log"
//
//	"github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth/authservice" // 客户端代码
//	"github.com/cloudwego/kitex/client"
//	consul "github.com/kitex-contrib/registry-consul"
//)
//
//func main() {
//	// 1. 创建Consul解析器
//	resolver, err := consul.NewConsulResolver("127.0.0.1:8500")
//	if err != nil {
//		log.Fatal("创建解析器失败:", err)
//	}
//
//	// 2. 创建客户端（服务名称需要与服务端保持一致）
//	authClient, err := authservice.NewClient(
//		"auth", // 服务名称
//		client.WithResolver(resolver),
//	)
//	if err != nil {
//		log.Fatal("创建客户端失败:", err)
//	}
//
//	fmt.Printf("%v", authClient)
//
//	//// 3. 测试Token颁发功能
//	//deliverResp, err := authClient.DeliverTokenByRPC(context.Background(), &auth.DeliverTokenReq{
//	//	UserId: 1001,
//	//})
//	//if err != nil {
//	//	log.Fatal("颁发Token失败:", err)
//	//}
//	//log.Printf("颁发的Token: %s", deliverResp.Token)
//	//
//	//// 4. 测试Token验证功能
//	//verifyResp, err := authClient.VerifyTokenByRPC(context.Background(), &auth.VerifyTokenReq{
//	//	Token: deliverResp.Token,
//	//})
//	//if err != nil {
//	//	log.Fatal("验证Token失败:", err)
//	//}
//	//log.Printf("验证结果: %t", verifyResp.Res)
//}

package main

import (
	"fmt"
	"github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth"
)

func main() {
	resp := &auth.DeliveryResp{Token: "test"}
	fmt.Println(resp.Token) // 应该能直接输出
}
