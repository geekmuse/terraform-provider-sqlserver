INSTALL_TO := /usr/local/bin

.DEFAULT_TARGET: build

.PHONY: build
build:
	go build -o terraform-provider-sqlserver

.PHONY: tf_init
tf_init:
	terraform init

.PHONY: test
test: build tf_init