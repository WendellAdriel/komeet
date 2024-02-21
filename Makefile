# Well documented Makefiles
DEFAULT_GOAL := help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-40s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

configure: ## Setup the application for the first time
	cp config.sample.json config.json \
	&& cp secrets.sample.json secrets.json

build: ## Builds the application
	rm -rf dist \
	&& mkdir dist \
	&& cd src \
	&& go build -ldflags "-w -s" -o ../dist/komeet . \
	&& chmod +x ../dist/komeet \
	&& cd ../web \
	&& yarn build \
	&& mv dist ../dist/web

run: ## Runs the application
	./dist/komeet serve

create-user: ## Creates a new user for the application
	./dist/komeet create-user --name="$(NAME)" --email=$(EMAIL) --password=$(PASSWORD)