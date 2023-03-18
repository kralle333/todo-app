
local-postgres:
	docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=admin postgres

local-migrate-up:
	migrate -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' -path db/migrations up
