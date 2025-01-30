.PHONY: gen-demo-auth
gen-demo-auth:
	@cd demo/auth && go mod init auth && cwgo server -I ../../idl --type RPC --module auth --service auth --idl ../../idl/auth.proto
