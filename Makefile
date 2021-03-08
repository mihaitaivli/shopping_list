generate:
	go run github.com/99designs/gqlgen generate

develop:
	go run server.go

localDb:
	cd local_db && docker-compose up -d