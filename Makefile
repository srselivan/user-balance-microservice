run:
	go run ./cmd/.
docker:
	docker run -p 3306:3306 --rm --name mysql_dev -e MYSQL_ROOT_PASSWORD=root -d mysql
	docker exec -it mysql_dev bash
stop:
	docker stop mysql_dev