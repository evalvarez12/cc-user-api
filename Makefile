#--------- RUN ON HOST -----------------
# Get docker images
reset-database:
	cat sql/0000-init.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Ucc cc_users

images:
	docker pull postgres:9.4.5 && \
	docker build -t cc:user_api .

# Create DB container
database:
	(docker stop postgres || exit 0) && \
	(docker rm postgres || exit 0) && \
	docker run \
		-d \
		-p 127.0.0.1:15432:5432 \
		--name postgres postgres:9.4.5 && \
	sleep 10 && \
	docker exec postgres psql -h127.0.0.1 -p5432 -Upostgres -c "CREATE ROLE cc PASSWORD 'pass' NOSUPERUSER NOCREATEDB NOCREATEROLE INHERIT LOGIN" &&\
	docker exec postgres psql -h127.0.0.1 -p5432 -Upostgres -c "CREATE DATABASE cc_users"
	cat sql/0000-init.sql | PGPASSWORD=pass psql -h127.0.0.1 -p15432 -Ucc cc_users

database-shell:
	docker exec -it postgres psql -Ucc cc_users

api:
	(docker stop user_api || exit 0) && \
	(docker rm user_api || exit 0) && \
	docker run \
		-d \
		-p 0.0.0.0:8082:8082 \
		--name user_api\
		--link postgres \
		cc:user_api

update-api:
	docker cp . user_api:/go/src/github.com/arbolista-dev/cc-user-api
