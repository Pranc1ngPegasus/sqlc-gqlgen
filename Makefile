.PHONY: format
format:
	$(call format)

.PHONY: generate.gqlgen
generate.gqlgen:
	@go run github.com/99designs/gqlgen generate
	$(call format)

.PHONY: generate.sqlc
generate.sqlc:
	@go run github.com/kyleconroy/sqlc/cmd/sqlc generate -f sqlc.yaml
	$(call format)

.PHONY: generate.wire
generate.wire:
	@go run github.com/google/wire/cmd/wire ./...
	$(call format)

.PHONY: migrate.plan
migrate.plan:
	@go run github.com/k0kubun/sqldef/cmd/psqldef -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} ${DATABASE_NAME} --dry-run < schema.sql

.PHONY: migrate.apply
migrate.apply:
	@go run github.com/k0kubun/sqldef/cmd/psqldef -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} ${DATABASE_NAME} < schema.sql

define format
	@go mod tidy
endef
