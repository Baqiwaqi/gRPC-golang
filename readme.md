# GRPC Golang

## Idea

`Heb een idee voor je als ik op vakantie ben:`

`Kun je de quote tool die je aan het maken bent maken met de volgende features?`

- Golang gRPC server
- Firebase database (backend db)
- React gRPC client-
- Realtime updates
- Run in docker

## Todo
<!-- create Todo's here -->
- [ ] Create Dockerfile for generating gRPC proto files
- [ ] Create docker-compose.yml

## Commands

This commmand will show usefull information about the conatiner.

```bash
docker inspect "container name"
```

## Server

Golang server

### Run Sever

```bash
go run main.go
```

Open [http://localhost:3000](http://localhost:50052) with your browser to see the result.

### Docker Server

```bash
docker build -t baqiwaqi/grpoc-golang:1.0 .  
docker run -p  5001:8080 ff4523b66553   
```

## CLient

Next.js web app

### Run Client

``` bash
cd src
cd client
npm run dev
# or
yarn dev
```

### Docker Client

```bash
cd src
cd client
docker build -t nextjs-docker:1.0 .  
docker run -p 3000:3000 e13f3c029bbe  
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.
