#!/bin/bash

# Start PostgreSQL container
docker run -d --name mydb --network=mynetwork \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=yahia2002 \
  -e POSTGRES_DB=toolsproject \
  -p 5432:5432 \
  -v mydb_data:/var/lib/postgresql/data postgresimage:0.6

echo "Waiting for PostgreSQL container to start..."
sleep 5

# Start Backend container
docker run -d --name backendcon --network mynetwork \
  -p 8080:8080 backendgo:1.4

# Start Frontend container
docker run -d --name frontend --network mynetwork \
  -p 4200:80 frontend:0.6

echo "Containers started successfully!"
