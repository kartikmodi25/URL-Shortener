postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=252900 -d postgres:16.2-alpine
createdb:
	docker exec -it postgres16 createdb --username=postgres --owner=postgres url_data
run:
	go run main.go

.PHONY: run postgres createdb