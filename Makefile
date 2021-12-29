.PHONY: playground
playground:
	go1.18beta1 run cmd/molizen/molizen.go -source ./playground/user/user.go -destination ./playground/actor/user.go