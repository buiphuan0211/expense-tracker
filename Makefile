#!bin/bash

export DOMAIN_EXPENSE_TRACKER_ADMIN=localhost:8000
export ZOOKEEPER_URI=127.0.0.1:2181
export ZOOKEEPER_PREFIX_EXPENSE_TRACKER_COMMON=/expense_tracker/common
export ZOOKEEPER_PREFIX_EXPENSE_TRACKER_ADMIN=/expense_tracker/admin

make run-admin:
	go run cmd/admin/main.go

swagger-admin:
	swag init -d ./ -g cmd/admin/main.go \
    -o ./docs/admin