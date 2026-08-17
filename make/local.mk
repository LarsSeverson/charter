LOCAL_COMPOSE := docker compose --project-directory . -f deploy/local/compose.yaml

local-up:
	$(LOCAL_COMPOSE) up --build

local-down:
	$(LOCAL_COMPOSE) down

local-down-volumes:
	$(LOCAL_COMPOSE) down --volumes

.PHONY: local-up local-down local-down-volumes
