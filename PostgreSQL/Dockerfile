FROM postgres:latest

ENV POSTGRES_DB toolsproject
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD yahia2002

COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432
