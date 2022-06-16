FROM node:16 as build
WORKDIR /app
RUN npm i --location=global npm
COPY package*.json ./
RUN npm i
COPY public ./public
COPY src ./src
RUN ls -lah \
    && npm run build
RUN ls -lah

FROM nginx
RUN mkdir /usr/share/nginx/html/dist/
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
