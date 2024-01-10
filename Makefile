dev_server:
	nodemon --exec go run main.go --signal SIGTERM

pre_commit:
	pre-commit run --all-files

docs:
	swag init

PHONY: start_dev pre_commit docs
