FROM golang:1.19

WORKDIR /usr/src/app

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download && go mod verify

COPY ./backend .
COPY ./gameclient/dist ./ui

RUN go build -v -o /usr/local/bin/ ./...

CMD ["backend"]