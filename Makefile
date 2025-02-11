.PHONY: gen-demo-auth
gen-demo-auth:
	@cd demo/auth && go mod init auth && cwgo server -I ../../idl --type RPC --module auth --service auth --idl ../../idl/auth.proto

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/Tinuvile/goShop/app/frontend -I ../../idl

