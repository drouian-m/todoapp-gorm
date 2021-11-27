# todoapp-gorm

[Gorm](https://gorm.io/) experiments

## Frontend

The frontend application is based on [shprink/web-components-todo](https://github.com/shprink/web-components-todo).

## Backend

Gorm experiments

TODO: refactor backend

## Run postgres instance

```sh
docker run --ulimit memlock=-1:-1 -it --rm=true --memory-swappiness=0 \
    --name postgres-gorm -e POSTGRES_USER=gorm \
    -e POSTGRES_PASSWORD=gorm -e POSTGRES_DB=gorm \
    -p 5432:5432 postgres:13.1
```

## Run application

Run the API server.
```sh
npm run api:run
```

Run the Sentil frontend.
```sh
npm run frontend:run
```