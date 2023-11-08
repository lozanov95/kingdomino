FROM node as fe
WORKDIR /fe
COPY gameclient/package.json .
RUN npm install
COPY gameclient ./
RUN npm run build

FROM golang:1.20 as be
WORKDIR /usr/src/kingdomino
COPY ./go.mod ./go.sum ./main.go ./
RUN go mod download && go mod verify
COPY ./backend ./backend/
RUN CGO_ENABLED=0 go build -v -o . ./...

FROM scratch
WORKDIR /app
COPY --from=be /usr/src/kingdomino/kingdomino kingdomino
COPY --from=fe /fe/dist/ ./ui/
CMD ["./kingdomino"]