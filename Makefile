generate:
	go run github.com/99designs/gqlgen generate

develop:
	go run server.go

localDb_up:
	cd local_db && docker-compose up -d

localDb_down:
	cd local_db && docker-compose down
