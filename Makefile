prod:
	git pull
	make run-container

pretty:
	npx prettier --write .

dev:
	make -j 2 dev-fe dev-be

dev-fe:
	npm run --prefix gameclient dev

dev-be:
	go run main.go -port 8080

run-container:
	docker image build . -t domino
	docker container rm -f c-domino
	docker container run --restart always -dp 8080:80 --name c-domino domino

test:
	go test -v ./...