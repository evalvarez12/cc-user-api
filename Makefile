#--------- RUN ON HOST -----------------
# Get docker images
reset-database:
	cat sql/0000-init.sql | PGPASSWORD=pass psql -h localhost -Umazing financy && \
	cat sql/0001-triggers.sql | PGPASSWORD=pass psql -h localhost -Umazing financy

deploy:
	rsync --delete -av . --exclude .git --exclude data ham:~/deploy/ham-api2 && \
	ssh -t ham "make -C deploy/ham-api2 images" && \
	ssh -t ham "make -C deploy/ham-api2 api"

deploy-database:
	rsync --delete -av . --exclude .git --exclude data ham:~/deploy/ham-api2 && \
	ssh -t ham "cat deploy/ham-api2/sql/0000-init.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Umazing financy" && \
	ssh -t ham "cat deploy/ham-api2/sql/0001-triggers.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Umazing financy"

deploy-hello:
	rsync --delete -av . --exclude .git --exclude data hellohamm:~/deploy/hamm-api && \
	ssh -t hellohamm "make -C deploy/hamm-api images" && \
	ssh -t hellohamm "make -C deploy/hamm-api api"

deploy-database-hello:
	rsync --delete -av . --exclude .git --exclude data hellohamm:~/deploy/hamm-api && \
	ssh -t hellohamm "cat deploy/hamm-api/sql/0000-init.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Umazing financy" && \
	ssh -t hellohamm "cat deploy/hamm-api/sql/0001-triggers.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Umazing financy"

images:
	docker pull postgres:9.4.5 && \
	docker build -t mazing:server .

# Create DB container
database:
	(docker stop postgres || exit 0) && \
	(docker rm postgres || exit 0) && \
	docker run \
		-d \
		-v $$PWD/data:/var/lib/postgresql/data \
		-p 127.0.0.1:15432:5432 \
		--name postgres postgres:9.4.5 && \
	sleep 10 && \
	psql -h127.0.0.1 -p15432 -Upostgres -c "CREATE ROLE mazing PASSWORD 'pass' NOSUPERUSER NOCREATEDB NOCREATEROLE INHERIT LOGIN" &&\
	psql -h127.0.0.1 -p15432 -Upostgres -c "CREATE DATABASE financy" &&\
	cat sql/0000-init.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Umazing financy &&\
	cat sql/0001-triggers.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Umazing financy

database-shell:
	docker exec -it postgres psql -Umazing financy

api:
	(docker stop server || exit 0) && \
	(docker rm server || exit 0) && \
	docker run \
		-d \
		-p 0.0.0.0:8082:8082 \
		--name server \
		--link postgres \
		mazing:server
