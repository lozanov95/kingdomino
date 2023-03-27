pretty:
	npx prettier --write .

build-fe:
	npm --prefix ./gameclient run build

run-dev:
	make -j 2 run-dev-fe run-dev-be

run-dev-fe:
	npm run --prefix gameclient dev

run-dev-be:
	go run main.go -port 8080

image:
	docker image build -t dominoapp .

container:
	docker container rm dominocontainer -f
	docker container run -dp 8080:80 --rm --name dominocontainer dominoapp

run-prod:	build-fe 	image 	container