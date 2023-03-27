FROM golang:1.20

WORKDIR /usr/src/kingdomino

COPY ./go.mod ./go.sum ./main.go ./
RUN go mod download && go mod verify

COPY ./backend ./backend/
COPY ./gameclient/dist ./ui

RUN go build -v -o /usr/local/bin/ ./...

CMD ["kingdomino"]