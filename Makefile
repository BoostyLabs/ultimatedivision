# Variables
LATEST_COMMIT := $$(git rev-parse HEAD)
VERSION ?= latest
HOST_FOR_DOCKER_IMAGE ?= dcloud.hicrystal.com
ENVIRONMENT ?= ""

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build_dist: ## Build dist folder that needed for frontend.
	cd web/console && npm ci && npm run build


build_nft_signer: ## Build NFT Signer docker image.
	docker build -f ./deploy/nftsigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer:$(VERSION) . && docker build -f ./deploy/nftsigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer$(ENVIRONMENT):latest .

push_nft_signer: ## Push NFT Signer docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer:$(VERSION) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_signer$(ENVIRONMENT):latest


build_nft_drop: ## Build NFT Drop docker image.
	docker build -f ./deploy/nftdrop.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_drop:$(VERSION) . && docker build -f ./deploy/nftdrop.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_drop$(ENVIRONMENT):latest .

push_nft_drop: ## Push NFT Drop docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_drop:$(VERSION) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_nft_drop$(ENVIRONMENT):latest


build_currency_signer: ## Build currency signer docker image.
	docker build -f ./deploy/currencysigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer:$(VERSION) . && docker build -f ./deploy/currencysigner.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer$(ENVIRONMENT):latest .

push_currency_signer: ## Push currency signer docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer:$(VERSION) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_currency_signer$(ENVIRONMENT):latest


build_card_generator: ## Build Card Generator docker image.
	docker build -f ./deploy/cardgenerator.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator:$(VERSION) . && docker build -f ./deploy/cardgenerator.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator$(ENVIRONMENT):latest .

push_card_generator: ## Push Card Generator docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator:$(VERSION) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division_card_generator$(ENVIRONMENT):latest


build_app: ## Build Application docker image.
	make build_dist && docker build -f ./deploy/ultimatedivision.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division:$(VERSION) . && docker build -f ./deploy/ultimatedivision.Dockerfile -t $(HOST_FOR_DOCKER_IMAGE)/ultimate_division$(ENVIRONMENT):latest .

push_app: ## Push Application docker image.
	docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division:$(VERSION) && docker push $(HOST_FOR_DOCKER_IMAGE)/ultimate_division$(ENVIRONMENT):latest


build: ## Build all necessary docker images.
	make build_app build_nft_signer build_nft_drop build_currency_signer build_card_generator

push: ## Push all necessary docker images.
	make push_app push_nft_signer build_nft_drop push_currency_signer build_card_generator

docker: ## Build and push all necessary docker images.
	make build push