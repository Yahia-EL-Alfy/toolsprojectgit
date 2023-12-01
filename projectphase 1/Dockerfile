# Stage 1: Build Angular App
FROM node:latest as node
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# Stage 2: Setup Nginx
FROM nginx:latest
COPY --from=node /app/dist/tools-project /usr/share/nginx/html
EXPOSE 80