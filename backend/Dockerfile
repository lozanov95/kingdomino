FROM golang:1.23 AS be
WORKDIR /usr/src/kingdomino
COPY backend/go.mod backend/go.sum ./
RUN go mod download && go mod verify
COPY backend/main.go ./
COPY backend/game game/
RUN CGO_ENABLED=0 go build -v -o . ./...

FROM ubuntu:latest AS user
RUN useradd -u 10001 user

FROM scratch
WORKDIR /app
COPY --from=user /etc/passwd /etc/passwd
COPY --from=be /usr/src/kingdomino/kingdomino kingdomino
USER user
CMD ["./kingdomino"]