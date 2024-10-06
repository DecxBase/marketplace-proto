.PHONY: gen
gen:
	@decx protoc compile

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: verify
verify:
	go mod verify
