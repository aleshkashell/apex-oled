ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: run/main
run/main:
	go run ./cmd/apex-oled