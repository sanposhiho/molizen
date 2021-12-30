.PHONY: scenario
scenario:
	go1.18beta1 run playground/scenarios/scenario$(WHAT)/main.go

# update actor in scenarios
.PHONY: scenario-update
scenario-update:
	go1.18beta1 run cmd/molizen/molizen.go -source ./playground/scenarios/scenario$(WHAT)/user/user.go -destination ./playground/scenarios/scenario$(WHAT)/actor/user.go

