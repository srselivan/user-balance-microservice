run:
	go run ./cmd/.
docker:
	docker run -p 3306:3306 --rm --name dev -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=recordings -d mysql
stop:
	docker stop dev