PROJECT_NAME="HealthHub"
LOCAL_CONFIG=config/local.yml


.PHONY: tag
tag:
	./scripts/tag.sh

.PHONY: run
run:
	go run cmd/sso/main.go --config "${LOCAL_CONFIG}"
