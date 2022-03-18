.PHONY: scenario
scenario:
	go run playground/scenarios/scenario$(WHAT)/main.go

# update actor in scenarios
.PHONY: scenario-update
scenario-update:
	go run cmd/molizen/molizen.go -source ./playground/scenarios/scenario$(WHAT)/user/user.go -destination ./playground/scenarios/scenario$(WHAT)/actor/user.go

.PHONY: lint
lint:
	golangci-lint run --timeout 30m ./...

.PHONY: format
format:
	golangci-lint run --fix ./...
