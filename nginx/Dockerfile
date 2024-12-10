FROM node:23 AS build
WORKDIR /fe
COPY gameclient/package.json .
RUN npm install
COPY gameclient/ .
RUN npm run build

FROM nginx:1.27.3
COPY --from=build /fe/dist /usr/share/nginx/html