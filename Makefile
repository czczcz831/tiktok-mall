.PHONY: auth
auth:
	@cd app/auth && cwgo server -I ../../idl --module github.com/czczcz831/tiktok-mall/app/auth --service auth --idl ../../idl/auth.proto

.PHONY: gen
gen: ## gen client code of {svc}. example: make gen svc=product
	@scripts/gen.sh ${svc}
	
.PHONY: run
run: ## run {svc} server. example: make run svc=product
	@scripts/run.sh ${svc}

.PHONY: tidy
tidy: ## run `go mod tidy` for all go module
	@scripts/tidy.sh