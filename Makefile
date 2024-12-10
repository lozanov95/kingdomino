prod:
	git pull
	make run-container

dev:
	make -j 2 dev-fe dev-be

dev-fe:
	npm run --prefix gameclient dev

dev-be:
	cd backend && go run main.go -port 80

run-container:
	docker compose up -d --build

test:
	cd backend && go test -v --cover ./...

install-fe:
	cd gameclient && npm install

install-be:
	cd backend && go mod download && go mod verify

install:
	make -j 2 install-fe install-be