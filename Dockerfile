# Dockerfile.postgres

# Use the official PostgreSQL image from Docker Hub
FROM postgres:13

# Environment variables
ENV POSTGRES_USER=youruser
ENV POSTGRES_PASSWORD=yourpassword
ENV POSTGRES_DB=database

# Expose the PostgreSQL port
EXPOSE 5432
