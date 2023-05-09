SHELL := /bin/bash

all: help

help :
	@echo "Help information, please run specific target:"
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
		IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf " - %-20s %s\n" $$help_command $$help_info ; \
	done


install: ## Setting up APP
	@docker compose down -v
	@docker compose build
	@docker compose up -d kafka && printf "init Kafka ..." && sleep 10 && echo " continue" && sleep 1
	@docker compose up -d

install-local: ## Setting up API in dev mode
	@cp deployments/docker-compose.dev.override.yml docker-compose.override.yml
	@touch .env.local
	@make install

destroy: ## Uninstall project environment
	@read -p $$'\e  \033[0;36mDestroy old environment?\033[0m \033[0;33mWARNING: It will remove all images, local configurations and database!\033[0m [\033[0;32mno\033[0m]: ' -r destroy; \
	destroy=$${destroy:-"no"}; \
	if [[ $$destroy =~ ^[yY][eE][sS]|[yY]$$ ]]; then \
		printf "\033[0;36mStop and remove docker containers, images, volumes ...\033[0m\n"; \
		docker compose down -v --remove-orphans --rmi local; \
		printf "\033[0;32mApplication environment destroyed!\033[0m\n"; \
	fi;