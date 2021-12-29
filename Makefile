# Variables
LATEST_COMMIT := $$(git rev-parse HEAD)
VERSION ?= latest
VER ?= latest

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build_nft_signer_prod: ## Build NFT Signer docker image for Production.
	cp ./deploy/configs/prod/config_nft_signer.json ./ && docker build -f ./deploy/NFTSigner.Dockerfile -t docker.hicrystal.com/ud_nft_signer_prod:$(VERSION) .

push_nft_signer_prod: ## Push NFT Signer docker image for Production.
	docker push docker.hicrystal.com/ud_nft_signer_prod:$(VERSION)

build_nft_signer_test: ## Build NFT Signer docker image for Testing.
	cp ./deploy/configs/test/config_nft_signer.json ./docker build -f ./deploy/NFTSigner.Dockerfile -t docker.hicrystal.com/ud_nft_signer_test:$(VERSION) .

push_nft_signer_test: ## Push NFT Signer docker image for Testing.
	docker push docker.hicrystal.com/ud_nft_signer_test:$(VERSION)

build_currency_signer_prod: ## Build Currency Signer docker image for Production.
	cp ./deploy/configs/prod/config_currency_signer.json ./ && docker build -f ./deploy/CurrencySigner.Dockerfile -t docker.hicrystal.com/ud_currency_signer_prod:$(VERSION) .

push_currency_signer_prod: ## Push Currency Signer docker image for Production.
	docker push docker.hicrystal.com/ud_currency_signer_prod:$(VERSION)

build_currency_signer_test: ## Build Currency Signer docker image for Testing.
	cp ./deploy/configs/test/config_currency_signer.json ./docker build -f ./deploy/CurrencySigner.Dockerfile -t docker.hicrystal.com/ud_currency_signer_test:$(VERSION) .

push_currency_signer_test: ## Push Currency Signer docker image for Testing.
	docker push docker.hicrystal.com/ud_currency_signer_test:$(VERSION)

build_app_prod: ## Build Application docker image for Production.
	cp ./deploy/configs/prod/config.json ./ && docker build -f ./deploy/AppProd.Dockerfile -t docker.hicrystal.com/ud_app_prod:$(VERSION) .

push_app_prod: ## Push Application docker image for Production.
	docker push docker.hicrystal.com/ud_app_prod:$(VERSION)

build_app_test: ## Build Application docker image for Testing.
	cp ./deploy/configs/test/config.json ./docker build -f ./deploy/AppTest.Dockerfile -t docker.hicrystal.com/ud_app_test:$(VERSION) .

push_app_test: ## Push Application docker image for Testing.
	docker push docker.hicrystal.com/ud_app_test:$(VERSION)

build_prod: ## Build all necessary docker images for Production.
	make build_app_prod build_nft_signer_prod build_currency_signer_prod VERSION=$(VER)

push_prod: ## Push all necessary docker images for Production.
	make push_app_prod push_nft_signer_prod push_currency_signer_prod VERSION=$(VER)

build_test: ## Build all necessary docker images for Testing.
	make build_app_test build_nft_signer_test build_currency_signer_test VERSION=$(VER)

push_test: ## Push all necessary docker images for Testing.
	make push_app_test push_nft_signer_test push_currency_signer_test VERSION=$(VER)

