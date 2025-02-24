.PHONY: gen-demo-auth
gen-demo-auth:
	@cd demo/auth && go mod init auth && cwgo server -I ../../idl --type RPC --module auth --service auth --idl ../../idl/auth.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/Tinuvile/goShop/app/frontend -I ../../idl

.PHONY: gen-rpc-client
gen-rpc-client:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/Tinuvile/goShop/rpc_gen --I ../idl --idl ../idl/user.proto

.PHONY: gen-rpc-server
gen-rpc-server:
	@cd app/user && cwgo server --type RPC --service user --module github.com/Tinuvile/goShop/app/user --pass "-use github.com/Tinuvile/goShop/rpc_gen/kitex_gen" --I ../../idl --idl ../../idl/user.proto