.PHONY: dbtest start mysql.start mysql.stop

dbtest:
	go test -v ./infrastructure/database/...

start:
	go build -o dbsample
	./dbsample

mysql.start:
	docker run --rm -d \
		-e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} \
		-e MYSQL_USER=${DB_USER} \
		-e MYSQL_PASSWORD=${DB_PASSWORD} \
		-e MYSQL_DATABASE=${DB_NAME} \
		-p ${DB_PORT}:3306 \
		-h ${DB_HOST} \
		--name ${DB_HOST} mysql:5.7
	mysql -h ${DB_HOST} --port ${DB_PORT} \
		-u${DB_USER} -p${DB_PASSWORD} < mysql/init.sql

mysql.stop:
	docker stop ${DB_HOST}
