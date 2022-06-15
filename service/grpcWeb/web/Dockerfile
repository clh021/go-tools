FROM node:16 as build
WORKDIR /app
COPY package*.json ./
RUN yarn install
RUN mkdir lib
COPY lib/echo_grpc_web_pb.js ./lib/
COPY lib/echo_pb.js ./lib/
COPY client.js ./
COPY echoapp.js ./
COPY package.sh ./
RUN ls -lah \
    && ./package.sh

FROM nginx
RUN mkdir /usr/share/nginx/html/dist/
COPY --from=build /app/dist/main.js /usr/share/nginx/html/dist/
COPY echotest.html /usr/share/nginx/html/index.html
EXPOSE 80
