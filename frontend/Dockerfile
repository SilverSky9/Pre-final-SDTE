# Build
FROM node:16.14.0-alpine3.14 as Build
WORKDIR /app
COPY package.json /app/
RUN ["npm", "install"]
COPY . /app/
RUN npm run build
# Deploy
FROM nginx:1.21.6-alpine
COPY --from=Build /app/build/ /usr/share/nginx/html/
EXPOSE 80
