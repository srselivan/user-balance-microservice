local:
	go run ./cmd/.
docker:
	docker run -p 5432:5432 --rm --name pg -e POSTGRES_PASSWORD=root -e POSTGRES_DB=users -d postgres
exec:
	docker exec -it pg bash
stop:
	docker stop pg

build:
	docker compose build
run:
	docker compose up