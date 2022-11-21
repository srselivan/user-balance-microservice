local:
	go run ./cmd/.
docker:
	docker run -p 3306:3306 --rm --name dev -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=users -d mysql
stop:
	docker stop dev

build:
	docker compose build
run:
	docker compose up