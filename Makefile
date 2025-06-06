export MYSQL_URL='mysql://root:superSecretPassword@tcp(localhost:3306)/fastcampus'

migrate-create:
	@migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@migrate -database $(MYSQL_URL) -path scripts/migrations up

migrate-down:
	@migrate -database $(MYSQL_URL) -path scripts/migrations down

# Run syntax "make migrate-force VERSION=2"
VERSION ?= 1
migrate-force:
	@migrate -database $(MYSQL_URL) -path scripts/migrations force $(VERSION)