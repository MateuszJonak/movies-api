#!make
#include .env
#export $(shell sed 's/=.*//' .env)

# autodocumenting feature
# insert all documentation after ## tag
.DEFAULT_GOAL := help

.PHONY: help
help: ## display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' `echo $(MAKEFILE_LIST) | awk '{gsub(/\.env/,"")}1'` | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2, $$3}'
# ./autodocumenting

# settings
jwt_keys_dir=middlewares

# jwt default passphrase
jwt_pass=moviesapipasword

jwt-keys: ## generate new jwt-keys for new projects
	ssh-keygen -t rsa -b 4096 -f $(jwt_keys_dir)/private.key \
	&& openssl rsa -in $(jwt_keys_dir)/private.key -pubout -outform PEM -out $(jwt_keys_dir)/public.key.pub