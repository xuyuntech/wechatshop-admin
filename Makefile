
TAG=${shell git describe --tag --long}
PREFIX=

db:
	docker rm -f wechatshop-db || true
	docker run -d --name wechatshop-db -v `pwd`/db_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=wechatshop -p 3306:3306 mysql
