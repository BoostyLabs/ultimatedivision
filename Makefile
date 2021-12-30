# Variables
LATEST_COMMIT := $$(git rev-parse HEAD)
VERSION ?= latest
VER ?= latest

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build_dist: ## Build dist folder that needed for frontend.
	cd web/console && npm i && npm run build

build_nft_signer_prod: ## Build NFT Signer docker image for Production.
	docker build -f ./deploy/NFTSigner.Dockerfile -t docker.hicrystal.com/ud_nft_signer_prod:$(VERSION) . && docker build -f ./deploy/NFTSigner.Dockerfile -t docker.hicrystal.com/ud_nft_signer_prod:latest .

push_nft_signer_prod: ## Push NFT Signer docker image for Production.
	docker push docker.hicrystal.com/ud_nft_signer_prod:$(VERSION) && docker push docker.hicrystal.com/ud_nft_signer_prod:latest

build_nft_signer_test: ## Build NFT Signer docker image for Testing.
	docker build -f ./deploy/NFTSigner.Dockerfile -t docker.hicrystal.com/ud_nft_signer_test:$(VERSION) . && docker build -f ./deploy/NFTSigner.Dockerfile -t docker.hicrystal.com/ud_nft_signer_test:latest .

push_nft_signer_test: ## Push NFT Signer docker image for Testing.
	docker push docker.hicrystal.com/ud_nft_signer_test:$(VERSION) && docker push docker.hicrystal.com/ud_nft_signer_test:latest

build_currency_signer_prod: ## Build Currency Signer docker image for Production.
	docker build -f ./deploy/CurrencySigner.Dockerfile -t docker.hicrystal.com/ud_currency_signer_prod:$(VERSION) . && docker build -f ./deploy/CurrencySigner.Dockerfile -t docker.hicrystal.com/ud_currency_signer_prod:latest .

push_currency_signer_prod: ## Push Currency Signer docker image for Production.
	docker push docker.hicrystal.com/ud_currency_signer_prod:$(VERSION) && docker push docker.hicrystal.com/ud_currency_signer_prod:latest

build_currency_signer_test: ## Build Currency Signer docker image for Testing.
	docker build -f ./deploy/CurrencySigner.Dockerfile -t docker.hicrystal.com/ud_currency_signer_test:$(VERSION) . && docker build -f ./deploy/CurrencySigner.Dockerfile -t docker.hicrystal.com/ud_currency_signer_test:latest .

push_currency_signer_test: ## Push Currency Signer docker image for Testing.
	docker push docker.hicrystal.com/ud_currency_signer_test:$(VERSION) && docker push docker.hicrystal.com/ud_currency_signer_test:latest

build_app_prod: ## Build Application docker image for Production.
	make build_dist && docker build -f ./deploy/AppProd.Dockerfile -t docker.hicrystal.com/ud_app_prod:$(VERSION) . && make build_dist && docker build -f ./deploy/AppProd.Dockerfile -t docker.hicrystal.com/ud_app_prod:latest .

push_app_prod: ## Push Application docker image for Production.
	docker push docker.hicrystal.com/ud_app_prod:$(VERSION) && docker push docker.hicrystal.com/ud_app_prod:latest

build_app_test: ## Build Application docker image for Testing.
	make build_dist && docker build -f ./deploy/AppTest.Dockerfile -t docker.hicrystal.com/ud_app_test:$(VERSION) . && make build_dist && docker build -f ./deploy/AppTest.Dockerfile -t docker.hicrystal.com/ud_app_test:latest .

push_app_test: ## Push Application docker image for Testing.
	docker push docker.hicrystal.com/ud_app_test:$(VERSION) && docker push docker.hicrystal.com/ud_app_test:latest

build_prod: ## Build all necessary docker images for Production.
	make build_app_prod build_nft_signer_prod build_currency_signer_prod

push_prod: ## Push all necessary docker images for Production.
	make push_app_prod push_nft_signer_prod push_currency_signer_prod

build_test: ## Build all necessary docker images for Testing.
	make build_app_test build_nft_signer_test build_currency_signer_test

push_test: ## Push all necessary docker images for Testing.
	make push_app_test push_nft_signer_test push_currency_signer_test

docker_prod: ## Build and push all necessary docker images for Production.
	make build_prod push_prod

docker_test: ## Build and push all necessary docker images for Testing.
	make build_test push_test