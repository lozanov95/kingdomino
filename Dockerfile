FROM node as fe
WORKDIR /fe
COPY gameclient/package.json .
RUN npm install
COPY gameclient ./
RUN npm run build

FROM golang:1.20
WORKDIR /usr/src/kingdomino
COPY ./go.mod ./go.sum ./main.go ./
RUN go mod download && go mod verify
COPY ./backend ./backend/
RUN go build -v -o /usr/local/bin/kingdomino ./...
COPY --from=fe /fe/dist/ ./ui/
CMD ["kingdomino"]