# Flows Microservices

Backend microservices for Flows application

---

## Tech Stacks

- Go
- Postgres
- NodeJS
- ExpressJS
- MongoDB

---

## Development

To run services, you can run manually from service directory or from root using [docker compose](https://docs.docker.com/compose/)

- start all service:

```bash
docker-compose up -d
```

- start a specific service (use --build tag to rebuild image):

```bash
docker-compose up -d <service>
```

- stop services:

```bash
docker-compose down
```

- stop specific service:

```bash
docker-compose down <service>
```
