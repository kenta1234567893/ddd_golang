.PHONY: up

up: ## Docker 起動
	docker compose up -d

.PHONY: down
down: ## Docker 終了
	@cd ${PWD}/docker && docker compose down

.PHONY: reup
reup: ## Docker 再ビルド
	@cd ${PWD}/docker && \
	docker compose down && \
	docker compose up -d --build

.PHONY: start
start: ## Docker　開始
	@cd ${PWD}/docker && docker compose start

.PHONY: stop
stop: ## Docker 停止
	@cd ${PWD}/docker && docker compose stop

.PHONY: restart
restart: ## Docker 再起動
	@cd ${PWD}/docker && docker compose restart


 
# makeコマンド実行時に引数を指定しなかった場合はhelpを実行する。
.DEFAULT_GOAL := help
 
help:  ## 本Makefileの使い方を表示する。
	@echo ""
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST)  \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'
	@echo ""