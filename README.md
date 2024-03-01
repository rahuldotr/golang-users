# Go REST API Project with Docker

This repository contains a Go REST API project (Project 3) that can be run using Docker. The project includes database migrations using edgedb.

## Prerequisites

- Docker installed on your machine
- Docker Compose installed on your machine

## Running the Project

To run the project, execute the following command:

```docker-compose up -d```

This command will build and start the Docker containers defined in the ```docker-compose.yml``` file in detached mode, allowing them to run in the background.

## Database Migrations

### Creating Migration

To create a migration, execute the following command:

```docker-compose exec edgedb /bin/bash -c "edgedb --tls-security=insecure -H localhost -P 5656 --password -u alokin -d u_data migration create"```

This command will create a new migration for the database.

### Applying Migration

To apply the migration, execute the following command:

```docker-compose exec edgedb /bin/bash -c "edgedb --tls-security=insecure -H localhost -P 5656 --password -u alokin -d u_data migrate"```

This command will apply the migration to the database.
