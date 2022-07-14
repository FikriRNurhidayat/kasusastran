MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=${DATABASE_SSL_MODE}

develop: format
	docker-compose down &> /dev/null
	docker-compose up -d &> /dev/null
	docker-compose logs -f kasusastran-httpd kasusastran-workerd

start: build
	./out/bin/kasusastran

install: build
	cp out/bin/kasusastran "${HOME}/.local/bin/kasusastran"

uninstall:
	rm -rf "${HOME}/.local/bin/kasusastran"

build:
	mkdir -p out/bin &> /dev/null
	go build -o out/bin/kasusastran-serve cmd/serve/main.go
	go build -o out/bin/kasusastran-work cmd/work/main.go

init:
	go mod tidy
	go mod vendor

mocks: format
	go install github.com/vektra/mockery/v2@latest 1>/dev/null
	rm -rf mocks
	mockery --all --keeptree --dir app
	mockery --all --output mocks/package --srcpkg google.golang.org/grpc/grpclog

apis:
	buf generate
	$(MAKE) format

query:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest 1> /dev/null
	rm -rf ./app/domain/query/*
	sqlc generate
	$(MAKE) mocks 

migration:
	$(eval timestamp := $(shell date +%s))
	touch db/migrations/$(timestamp)_${name}.up.sql
	touch db/migrations/$(timestamp)_${name}.down.sql

rollbackdb:
	echo "y" | docker run -v ${PWD}/db/migrations:/migrations \
		--rm -i --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} down "${step}"

migratedb:
	docker run -v ${PWD}/db/migrations:/migrations \
		--rm -it --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} up

createdb:
	createdb ${DATABASE_NAME}

dropdb:
	dropdb ${DATABASE_NAME}

seeddb:
	cp ./db/seeds.sql ./db/seeds.sql.bak
	envsubst < ./db/seeds.sql.bak > ./db/seeds.sql
	sed -i 's/COPY/\\copy/g' ./db/seeds.sql
	psql -a -f ./db/seeds.sql ${DATABASE_URL}
	psql -a -f ./db/reset.sql ${DATABASE_URL} 1> /dev/null
	mv ./db/seeds.sql.bak ./db/seeds.sql

cleandb:
	psql -a -f ./db/clean.sql ${DATABASE_URL}

format:
	buf format -w
	go fmt ./...

test: format
	go install gotest.tools/gotestsum@latest 1> /dev/null
	go install github.com/boumenot/gocover-cobertura@latest 1> /dev/null
	go install github.com/ggere/gototal-cobertura@latest 1> /dev/null
	gotestsum --format testname --junitfile junit.xml -- -coverprofile=coverage.lcov.info -covermode count ./... 
	gocover-cobertura < coverage.lcov.info > coverage.xml
	gototal-cobertura < coverage.xml

setupdb: createdb migratedb seeddb
