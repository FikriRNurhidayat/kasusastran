MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=${DATABASE_SSL_MODE}

develop:
	go run main.go

start: build
	./out/bin/kasusastran

install: build
	cp out/bin/kasusastran "${HOME}/.local/bin/kasusastran"

uninstall:
	rm -rf "${HOME}/.local/bin/kasusastran"

build:
	mkdir -p out/bin &> /dev/null
	go build -o out/bin/kasusastran main.go 

init:
	go mod vendor
	go install

mocks: format
	rm -rf mocks
	docker run \
		-it \
		--rm \
		-u 1000:1000 \
		-v ${PWD}:/service \
		-w /service \
		vektra/mockery --all --keeptree --dir app

query:
	rm -rf ./app/domain/query/*
	docker run \
		--rm \
		-u 1000:1000 \
		-v ${PWD}:/opt/app \
		-w /opt/app \
		kjconroy/sqlc generate
	$(MAKE) mocks 

migration:
	$(eval timestamp := $(shell date +%s))
	touch db/migrations/$(timestamp)_${name}.up.sql
	touch db/migrations/$(timestamp)_${name}.down.sql

rollback:
	echo "y" | docker run -v ${PWD}/db/migrations:/migrations \
		--rm -i --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} down 

migrate:
	docker run -v ${PWD}/db/migrations:/migrations \
		--rm -it --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} up

create:
	createdb ${DATABASE_NAME}

drop:
	dropdb ${DATABASE_NAME}

seed: clean
	cp ./db/seeds.sql ./db/seeds.sql.bak
	envsubst < ./db/seeds.sql.bak > ./db/seeds.sql
	sed -i 's/COPY/\\copy/g' ./db/seeds.sql
	psql -a -f ./db/seeds.sql ${DATABASE_URL}
	psql -a -f ./db/reset.sql ${DATABASE_URL} 1> /dev/null
	mv ./db/seeds.sql.bak ./db/seeds.sql

clean:
	psql -a -f ./db/clean.sql ${DATABASE_URL}
	$(MAKE) migrate 

format:
	go fmt ./...

test: format
	go install gotest.tools/gotestsum@latest 1> /dev/null
	go install github.com/boumenot/gocover-cobertura@latest 1> /dev/null
	go install github.com/ggere/gototal-cobertura@latest 1> /dev/null
	gotestsum --format testname --junitfile junit.xml -- -coverprofile=coverage.lcov.info -covermode count ./... 
	gocover-cobertura < coverage.lcov.info > coverage.xml
	gototal-cobertura < coverage.xml


setup: create migrate seed
