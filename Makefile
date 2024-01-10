# Variables
ENVIRONMENT ?=
IMAGE_TAG ?= latest
REGISTRY_HOST ?= ghcr.io/boostylabs
LATEST_COMMIT := $$(git rev-parse --short HEAD)

APP_IMAGE = ultimate_division
NFT_SIGNER_IMAGE = ultimate_division_nft_signer
CURRENCY_SIGNER_IMAGE = ultimate_division_currency_signer

IMAGE_APP_BACKUP = $(REGISTRY_HOST)/$(APP_IMAGE)$(ENVIRONMENT):$(LATEST_COMMIT)
IMAGE_APP_LATEST = $(REGISTRY_HOST)/$(APP_IMAGE)$(ENVIRONMENT):$(IMAGE_TAG)

IMAGE_NFT_SIGNER_BACKUP = $(REGISTRY_HOST)/$(NFT_SIGNER_IMAGE)$(ENVIRONMENT):$(LATEST_COMMIT)
IMAGE_NFT_SIGNER_LATEST = $(REGISTRY_HOST)/$(NFT_SIGNER_IMAGE)$(ENVIRONMENT):$(IMAGE_TAG)

IMAGE_CURRENCY_SIGNER_BACKUP = $(REGISTRY_HOST)/$(CURRENCY_SIGNER_IMAGE)$(ENVIRONMENT):$(LATEST_COMMIT)
IMAGE_CURRENCY_SIGNER_LATEST = $(REGISTRY_HOST)/$(CURRENCY_SIGNER_IMAGE)$(ENVIRONMENT):$(IMAGE_TAG)

help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
%:
	@:

build_dist: ## Build dist folder that needed for frontend.
	cd web/console && npm ci && npm run build

build_app: ## Build app docker image.
	DOCKER_BUILDKIT=1 docker build -f ./deploy/ultimatedivision.Dockerfile -t $(IMAGE_APP_BACKUP) . && DOCKER_BUILDKIT=1 docker build -f ./deploy/ultimatedivision.Dockerfile -t $(IMAGE_APP_LATEST) .

push_app: ## Push app docker image.
	docker push $(IMAGE_APP_BACKUP) && docker push $(IMAGE_APP_LATEST)

build_nft_signer: ## Build nft_signer docker image.
	DOCKER_BUILDKIT=1 docker build -f ./deploy/nftsigner.Dockerfile -t $(IMAGE_NFT_SIGNER_BACKUP) . && DOCKER_BUILDKIT=1 docker build -f ./deploy/nftsigner.Dockerfile -t $(IMAGE_NFT_SIGNER_LATEST) .

push_nft_signer: ## Push nft_signer docker image.
	docker push $(IMAGE_NFT_SIGNER_BACKUP) && docker push $(IMAGE_NFT_SIGNER_LATEST)

build_currency_signer: ## Build currency_signer docker image.
	DOCKER_BUILDKIT=1 docker build -f ./deploy/currencysigner.Dockerfile -t $(IMAGE_CURRENCY_SIGNER_BACKUP) . && DOCKER_BUILDKIT=1 docker build -f ./deploy/currencysigner.Dockerfile -t $(IMAGE_CURRENCY_SIGNER_LATEST) .

push_currency_signer: ## Push currency_signer docker image.
	docker push $(IMAGE_CURRENCY_SIGNER_BACKUP) && docker push $(IMAGE_CURRENCY_SIGNER_LATEST)

docker: ## Build and push all docker images.
	make build_app push_app \
		 build_nft_signer push_nft_signer \
		 build_currency_signer push_currency_signer

run_local: ## Build and run app locally.
	make build_dist && cd deploy/local && docker-compose up

stop_local: ## Stop app locally.
	cd deploy/local && docker-compose down
