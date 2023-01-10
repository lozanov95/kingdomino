npm --prefix ./gameclient run build
docker image build -t dominoapp .
docker container rm dominocontainer -f
docker container run -dp 8080:80 --rm --name dominocontainer dominoapp
