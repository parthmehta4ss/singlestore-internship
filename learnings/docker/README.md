# Docker Learning Documentation

**Internship: SingleStore** | **Video:** [Docker & Kubernetes Tutorial](https://www.youtube.com/watch?v=kTp5xUtcalw)

---

## Table of Contents
- [Microservices Architecture](#microservices-architecture)
- [Cloud Native](#cloud-native)
- [Containers](#containers)
- [Docker](#docker)
- [Docker Volumes](#docker-volumes)
- [YAML](#yaml)
- [Docker Compose](#docker-compose)
- [Container Registries](#container-registries)

---

## Microservices Architecture

> **In simple words:** Instead of building one giant application where everything is connected (monolithic), break it into smaller independent services that talk to each other via APIs.

### Monolithic vs Microservices

**Monolithic:** All components (backend, frontend, DB) tightly coupled - difficult to scale individually

**Microservices:** Independent services communicating via APIs
- Deploy/upgrade/rollback independently
- Scale specific services
- Different tech stacks per service
- Failure isolation

**Why it matters:** Like LEGO blocks - you can replace or upgrade one piece without rebuilding everything.

---

## Cloud Native

> **In simple words:** Building apps specifically designed to run in the cloud - using modern tools like containers, making them easy to scale, update, and manage automatically.

### Core Components
Containers + Service Meshes + Microservices + Immutable Infrastructure + Declarative APIs

### Trail Map
1. **Containerization** → Package and deploy apps
2. **CI/CD** → Automate deployment
3. **Orchestration** → Kubernetes with Helm
4. **Observability** → Monitoring (Fluentd/Jaeger)
5. **Service Meshes** → Advanced cluster functionality
6. **Networking & Policies** → Security implementation

---

## Containers

> **In simple words:** A container is like a backpack with everything your app needs (code, tools, libraries) to run anywhere - your laptop, server, or cloud. No more "but it works on my machine!" problems.

### What & Why?
**Container** = Code + Runtime + System tools + Libraries in one deployable unit

**Benefits:** Portable, fast, isolated, consistent across environments

### Containers vs VMs
| Feature | Containers | VMs |
|---------|-----------|-----|
| Size | MBs | GBs |
| Startup | Seconds | Minutes |
| Isolation | Process-level | Full OS |

**Think of it:** Container = lightweight suitcase, VM = moving entire house

### Image Layers
- Lower layers: Read-only (shared)
- Top layer: Read-write (container-specific)
- Only changed layers downloaded on updates

**Analogy:** Like a cake with layers - if you modify the frosting (top layer), you don't need to rebake the entire cake

---

## Docker

> **In simple words:** Docker is the tool that creates and manages containers. Think of it as a factory that packages your app into containers and runs them.

### Essential Commands
```bash
# Images
docker pull <image>
docker images
docker rmi <image>

# Containers
docker run -d -p 8080:80 <image>    # Run detached, map ports
docker ps                           # List running
docker ps -a                        # List all
docker stop/start <container>
docker logs <container>
docker exec -it <container> bash    # Interactive shell

# Build
docker build -t <name>:<tag> .
```

### Dockerfile Example

> **What is Dockerfile?** A recipe/instruction manual telling Docker how to build your container - what base to use, what files to copy, what commands to run.

```dockerfile
FROM node:16-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["node", "app.js"]
```

---

## Docker Volumes

> **In simple words:** Volumes are external storage folders where containers can save data. When container dies, the volume stays - like saving your game progress to the cloud instead of just your console.

### Why?
> **Containers are Ephemeral** - Data is lost when container is deleted

**Solution:** Store data outside container in volumes

### Commands
```bash
docker volume create my-data
docker volume ls
docker run -v my-data:/app/data my-image
docker volume rm my-data
```

### In Docker Compose
```yaml
services:
  db:
    volumes:
      - db-data:/var/lib/postgresql/data
volumes:
  db-data:
```

---

## YAML

> **In simple words:** YAML is a simple, human-readable way to write configuration files. Instead of complex code, you write "key: value" pairs that anyone can understand.

**YAML Ain't Markup Language** - Human-readable config format

### Syntax
```yaml
# Key-value
name: John
age: 30

# Nested (2 spaces)
person:
  name: John
  address:
    city: NYC

# Lists
fruits:
  - apple
  - banana

# Comments with #
```

**Rules:** 2 spaces (no tabs), space after colon, strings can be unquoted

---

## Docker Compose

> **In simple words:** Instead of running multiple `docker run` commands for each container (frontend, backend, database), write one YAML file and start everything with one command. Like a playlist for your containers.

### What?
Run multiple containers with single YAML file

### Commands
```bash
docker compose up -d        # Start detached
docker compose down         # Stop and remove
docker compose ps           # List services
docker compose logs -f      # Follow logs
docker compose exec <svc> <cmd>
```

### Example: Full-Stack App
```yaml
version: '3.8'
services:
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    depends_on:
      - backend
    environment:
      - API_URL=http://backend:5000

  backend:
    build: ./backend
    ports:
      - "5000:5000"
    depends_on:
      - db
    environment:
      - DATABASE_URL=postgres://user:pass@db:5432/mydb
    restart: unless-stopped

  db:
    image: postgres:14
    volumes:
      - db-data:/var/lib/postgresql/data
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
    expose:
      - "5432"  # Internal only

volumes:
  db-data:
```

### Key Concepts
**Networking:** Services communicate using service names as hostnames  
**Dependencies:** `depends_on` ensures start order  
**Restart Policies:** `always`, `no`, `on-failure`, `unless-stopped`  
**Environment Variables:** Inline, .env file, or command line

---

## Container Registries

> **In simple words:** Like GitHub for code, but for container images. A central place to store and share your containers - either publicly (Docker Hub) or privately (company's own registry).

**Registry** = Central repository for container images (Docker Hub, AWS ECR, GCR, ACR)

### Benefits
- Public or private
- Cloud registries near your app = faster pulls
- Local caching of layers

```bash
docker pull nginx
docker tag my-image:latest registry.com/my-image:v1
docker push registry.com/my-image:v1
```

---

## Resources

- [Docker Documentation](https://docs.docker.com/)
- [Docker Cheat Sheet](https://docs.docker.com/get-started/docker_cheatsheet.pdf)
- [Tutorial Video](https://www.youtube.com/watch?v=kTp5xUtcalw)

---

## Progress

- [x] Microservices & Cloud Native
- [x] Containers & Docker  
- [x] Docker Volumes
- [x] YAML
- [x] Docker Compose
- [x] Container Registries

## Hello


**Last Updated:** [Date]