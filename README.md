# Docker & Kubernetes Learning Documentation

**Internship: SingleStore** | **Video:** [Docker & Kubernetes Tutorial](https://www.youtube.com/watch?v=kTp5xUtcalw) | **Progress:** Up to 3:20:00

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
- [Kubernetes](#kubernetes)
- [Workloads](#workloads)
- [Hands-On Practice](#hands-on-practice)

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

## Kubernetes

> **In simple words:** Kubernetes (K8s) is like an orchestra conductor for containers. It automatically manages hundreds of containers across multiple servers - starting them, restarting if they crash, scaling up/down based on load, and routing traffic.

### Architecture
```
Container → Pod → Node → Cluster
```

**Think of it:** Container = musician, Pod = small band, Node = stage, Cluster = entire concert venue

### Master Node (Control Plane)

> **What it does:** The "brain" of Kubernetes - makes all decisions about where to run containers, monitors health, and maintains desired state.

- **kube-apiserver:** REST API, only entry point
- **etcd:** Key-value store, cluster state (single source of truth)
- **kube-controller-manager:** Controller of controllers
- **cloud-controller-manager:** Cloud provider integration
- **kube-scheduler:** Assigns pods to nodes
- **Addons:** DNS, Dashboard, monitoring

### Worker Node

> **What it does:** The "workers" that actually run your containers. Master tells them what to do, they execute.

- **kubelet:** Manages pod lifecycle
- **kube-proxy:** Network proxy
- **container runtime:** Runs containers (Docker/containerd)

### Core Concepts

**Context:** Access parameters (cluster, user, namespace)
```bash
minikube start
kubectl config get-contexts
kubectl config use-context <name>
```

**Namespaces:** Group and isolate resources
```bash
kubectl get namespaces
kubectl create namespace dev
kubectl apply -f app.yaml -n dev
kubectl delete namespace dev  # Deletes all resources!
```

**Node Pools:** Group of nodes with same size/config

---

## Pods

> **In simple words:** A pod is the smallest unit in Kubernetes - a wrapper around one or more containers that need to work closely together. Like a shared apartment for containers.

### What?
Smallest K8s unit - encapsulates 1+ containers, shares network/storage

**Scale by adding pods, not containers per pod**

**Key rule:** One application instance = One pod. Need more capacity? Create more pods, don't stuff more containers in one pod.

### Lifecycle States
`Pending` → `Running` → `Succeeded`/`Failed`/`Unknown`/`CrashLoopBackOff`

### Commands
```bash
kubectl get pods -o wide
kubectl describe pod <name>      # Check Events section
kubectl apply -f pod.yaml
kubectl delete pod <name>
kubectl exec -it <pod> -- /bin/bash
kubectl logs -f <pod>
```

### Example Pod
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app
  labels:
    app: frontend
spec:
  containers:
  - name: nginx
    image: nginx:1.21
    ports:
    - containerPort: 80
```

### Init Containers

> **Simple explanation:** Helper containers that run BEFORE your main app starts. Like a checklist - "Is database ready? Are config files loaded? Good, now start the main app."

Run **before** main container - for setup tasks

```yaml
spec:
  containers:
  - name: myapp
    image: busybox
  initContainers:
  - name: init-db
    image: busybox
    command: ['sh', '-c', 'until nslookup db; do sleep 2; done']
```

---

## Labels & Selectors

> **In simple words:** Labels are like sticky notes you put on things (pods, nodes) with info like "environment=production" or "app=frontend". Selectors are filters that find things based on those sticky notes.

### Labels
Key-value pairs for identification
```yaml
metadata:
  labels:
    app: frontend
    environment: prod
    tier: web
```

### Selectors
Filter objects by labels

> **Example use:** "Show me all pods where app=frontend AND environment=production"

```bash
kubectl get pods -l app=frontend
kubectl get pods -l app=frontend,environment=prod
```

### Use Case: Node Selection

> **Real scenario:** You want your database to run only on nodes with fast SSD storage, not slow hard drives.

```bash
# Label node
kubectl label nodes node-1 disktype=ssd

# Pod uses nodeSelector
spec:
  nodeSelector:
    disktype: ssd
```

**Verify endpoints:**
```bash
kubectl get pods -o wide
kubectl get endpoints <service>  # Compare IPs
```

---

## Kubernetes Networking

> **In simple words:** How containers talk to each other. Containers in the same pod are like roommates (share address, use room numbers). Pods are like different apartments (need full addresses). External access needs a doorman (Service).

### Communication Model
1. **Same pod:** `localhost:<port>` (shared IP)
2. **Different pods:** Pod IP addresses
3. **External access:** Service (LoadBalancer/NodePort)

### Multi-Container Pod Example
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: two-containers
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
    - containerPort: 80
  - name: busybox
    image: busybox
    command: ['sh', '-c', 'sleep 3600']
```

**Test communication:**
```bash
kubectl create -f two-containers.yaml
kubectl exec -it two-containers --container busybox -- /bin/sh
wget -qO- localhost:80  # Access nginx from busybox
```

**Port Forwarding:**

> **What it does:** Temporarily expose a pod to your local machine for testing - like a temporary tunnel from your laptop to the pod.

```bash
kubectl port-forward pod/my-pod 8080:80
curl localhost:8080
```

---

## Workloads

> **In simple words:** Workloads are different ways to run your applications in Kubernetes. Think of them as different "templates" for running containers based on what your app needs - some apps need to run forever, some need to run once, some need persistent data, etc.

### Deployments

> **Simple explanation:** The most common way to run apps in Kubernetes. Kubernetes automatically keeps the right number of copies running, handles updates smoothly, and restarts pods if they crash. Like having an auto-pilot that maintains your app.

**What it does:**
- Manages stateless applications
- Ensures desired number of replicas are always running
- Handles rolling updates and rollbacks
- Automatically replaces failed pods

**When to use:** Web applications, APIs, microservices - anything that doesn't need persistent storage or specific identity

**Example Deployment:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.21
        ports:
        - containerPort: 80
```

**Commands:**
```bash
# Create deployment
kubectl apply -f deployment.yaml

# View deployments
kubectl get deployments

# Scale deployment
kubectl scale deployment nginx-deployment --replicas=5

# Update image (rolling update)
kubectl set image deployment/nginx-deployment nginx=nginx:1.22

# Check rollout status
kubectl rollout status deployment/nginx-deployment

# Rollback to previous version
kubectl rollout undo deployment/nginx-deployment

# View rollout history
kubectl rollout history deployment/nginx-deployment
```

**How it works:** Deployment → Creates ReplicaSet → Creates Pods

### ReplicaSets

> **Simple explanation:** Ensures a specific number of identical pods are always running. If a pod dies, ReplicaSet immediately creates a new one. Think of it as a "copy machine" that maintains exact copies.

**What it does:**
- Maintains specified number of pod replicas
- Automatically replaces failed pods
- Uses labels to track which pods it manages

**Important:** Don't create ReplicaSets directly! Use Deployments instead. Deployments automatically create and manage ReplicaSets for you.

**When Kubernetes uses it:** Behind every Deployment is a ReplicaSet doing the actual work

**Example:**
```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: frontend
spec:
  replicas: 3
  selector:
    matchLabels:
      tier: frontend
  template:
    metadata:
      labels:
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: gcr.io/google-samples/gb-frontend:v5
```

```bash
kubectl get replicasets
kubectl describe rs frontend
```

### StatefulSets

> **Simple explanation:** For apps that need a stable identity and persistent storage - like databases. Each pod gets a fixed name (app-0, app-1, app-2) that doesn't change even if the pod restarts. Like hotel rooms with fixed numbers.

**What it does:**
- Provides stable, unique network identifiers for each pod
- Ensures ordered, graceful deployment and scaling
- Maintains persistent storage per pod
- Pods have predictable names: `<statefulset-name>-0`, `<statefulset-name>-1`, etc.

**When to use:** Databases (MySQL, MongoDB), message queues (Kafka), any app that needs:
- Stable persistent storage
- Stable network identity
- Ordered deployment/scaling

**Key differences from Deployment:**
- Pods are NOT interchangeable
- Each pod has persistent identifier
- Pods are created/deleted in order (0, 1, 2, ...)
- Each pod can have its own PersistentVolume

**Example:**
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
spec:
  serviceName: "mysql"
  replicas: 3
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql:5.7
        ports:
        - containerPort: 3306
        volumeMounts:
        - name: mysql-storage
          mountPath: /var/lib/mysql
  volumeClaimTemplates:
  - metadata:
      name: mysql-storage
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 1Gi
```

**Pod naming:** If StatefulSet name is `mysql`, pods will be: `mysql-0`, `mysql-1`, `mysql-2`

```bash
kubectl get statefulsets
kubectl get pods  # See ordered naming
kubectl scale statefulset mysql --replicas=5
```

### DaemonSets

> **Simple explanation:** Runs one copy of your pod on EVERY node in the cluster. Like installing a security camera on every floor of a building. Used for monitoring, logging, or anything that needs to run on all machines.

**What it does:**
- Ensures a copy of pod runs on all (or some) nodes
- Automatically adds pod to new nodes when they join
- Removes pod when node is removed
- Like a background daemon/service on every machine

**When to use:**
- Log collectors (Fluentd, Logstash)
- Monitoring agents (Prometheus Node Exporter)
- Network plugins
- Storage daemons
- Anything that needs to run on every node

**Example:**
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd
spec:
  selector:
    matchLabels:
      name: fluentd
  template:
    metadata:
      labels:
        name: fluentd
    spec:
      containers:
      - name: fluentd
        image: fluentd:latest
```

**No replicas field!** Kubernetes automatically calculates how many pods based on number of nodes.

```bash
kubectl get daemonsets
kubectl describe ds fluentd
```

### Jobs

> **Simple explanation:** For tasks that run once and then stop - like a batch job or script. Creates pods that do their work and exit. Think of it as a "to-do item" that gets checked off when complete.

**What it does:**
- Runs a task to completion
- Creates one or more pods
- Retries if pod fails
- Tracks successful completions
- Stops when task is done

**When to use:**
- Batch processing
- Data migrations
- Backups
- One-time setup tasks
- Any task that should run to completion

**Example:**
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: data-import
spec:
  completions: 1  # How many successful completions needed
  parallelism: 1  # How many pods to run at once
  backoffLimit: 3  # Retry attempts
  template:
    spec:
      containers:
      - name: import
        image: busybox
        command: ["echo", "Importing data..."]
      restartPolicy: Never  # Important!
```

**Commands:**
```bash
kubectl apply -f job.yaml
kubectl get jobs
kubectl describe job data-import
kubectl logs job/data-import  # See output
kubectl delete job data-import  # Clean up
```

**Parallel jobs:** Can run multiple pods simultaneously by setting `parallelism: 3`

### CronJobs

> **Simple explanation:** Like Jobs, but runs on a schedule - hourly, daily, etc. Think of it as setting up recurring reminders or scheduled tasks, like cron in Linux.

**What it does:**
- Creates Jobs on a schedule
- Uses cron syntax for timing
- Perfect for recurring tasks
- Each run creates a new Job

**When to use:**
- Daily backups
- Periodic reports
- Data cleanup tasks
- Scheduled API calls
- Any recurring task

**Example:**
```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: daily-backup
spec:
  schedule: "0 2 * * *"  # Every day at 2 AM
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: backup
            image: backup-tool
            command: ["/bin/sh", "-c", "backup-database.sh"]
          restartPolicy: OnFailure
```

**Cron schedule format:**
```
# ┌───────────── minute (0 - 59)
# │ ┌───────────── hour (0 - 23)
# │ │ ┌───────────── day of month (1 - 31)
# │ │ │ ┌───────────── month (1 - 12)
# │ │ │ │ ┌───────────── day of week (0 - 6) (Sunday to Saturday)
# │ │ │ │ │
# * * * * *
```

**Common schedules:**
```
"0 2 * * *"        # Every day at 2 AM
"*/15 * * * *"     # Every 15 minutes
"0 0 * * 0"        # Every Sunday at midnight
"0 */6 * * *"      # Every 6 hours
"30 3 * * 1-5"     # 3:30 AM on weekdays
```

```bash
kubectl get cronjobs
kubectl describe cronjob daily-backup
kubectl get jobs --watch  # Watch jobs being created
```

---

## Services

> **In simple words:** Services are like phone numbers for your pods. Pods come and go (changing IP addresses), but a Service gives you one stable address to reach them. It also load-balances traffic across multiple pods.

### Why Services?

**Problem:** Pods are ephemeral - they die and get recreated with new IP addresses. How do you reliably connect to them?

**Solution:** Services provide:
- Stable IP address and DNS name
- Load balancing across pods
- Service discovery

### Service Types

#### 1. ClusterIP (Default)

> **Simple explanation:** Internal-only access. Like an office phone extension - only people inside the office (cluster) can call it. Perfect for backend services that don't need external access.

**Use case:** Internal communication between microservices

**Example:**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-service
spec:
  type: ClusterIP
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 80        # Service port
    targetPort: 8080  # Pod port
```

**How it works:** Gets internal IP like `10.0.171.239` - only accessible within cluster

```bash
kubectl get services
# backend-service   ClusterIP   10.0.171.239   <none>   80/TCP
```

#### 2. NodePort

> **Simple explanation:** Opens a specific port (30000-32767) on EVERY node in the cluster. External users can access via `<NodeIP>:<NodePort>`. Like opening the same door on every building in a complex.

**Use case:** Simple external access for development/testing (not recommended for production)

**Example:**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: web-service
spec:
  type: NodePort
  selector:
    app: web
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
    nodePort: 30007  # Optional, Kubernetes will assign if omitted
```

**Access:** `http://<any-node-ip>:30007`

**Traffic flow:** External Client → Node IP:30007 → NodePort Service → ClusterIP → Pod

**Limitation:** No load balancing between nodes - client connects to one node

#### 3. LoadBalancer

> **Simple explanation:** Creates a cloud load balancer (AWS ELB, GCP Load Balancer, etc.) with a public IP. Traffic is distributed evenly across all pods. Like a receptionist who evenly distributes visitors. **Only works on cloud providers!**

**Use case:** Production external access with load balancing

**Example:**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: web-loadbalancer
spec:
  type: LoadBalancer
  selector:
    app: web
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```

**How it works:**
1. Kubernetes requests load balancer from cloud provider
2. Cloud creates load balancer with public IP
3. Traffic: External Client → Load Balancer → Nodes → Pods

```bash
kubectl get service web-loadbalancer
# NAME               TYPE           EXTERNAL-IP
# web-loadbalancer   LoadBalancer   34.123.45.67
```

**Access:** `http://34.123.45.67` (the external IP)

**Traffic flow:** External Client → Cloud Load Balancer → NodePort → ClusterIP → Pod

#### 4. ExternalName

> **Simple explanation:** Maps a Kubernetes service name to an external DNS name. Like creating a contact shortcut - instead of typing the full external address, use a short internal name.

**Use case:** Connecting to external services (external databases, APIs)

**Example:**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: external-database
spec:
  type: ExternalName
  externalName: my-database.example.com
```

**Usage:** Pods can connect to `external-database` instead of `my-database.example.com`

### Service vs Selector

```yaml
selector:
  app: my-app
```

Service finds pods with matching labels and sends traffic to them. Change the label on a pod → Service stops sending traffic to it.

### Service Discovery

Pods can reach services using:
- **DNS name:** `service-name` (same namespace) or `service-name.namespace.svc.cluster.local`
- **Environment variables:** Automatically injected

```bash
# Inside a pod
curl http://backend-service
curl http://backend-service.default.svc.cluster.local
```

### Commands

```bash
# Create service
kubectl apply -f service.yaml

# List services
kubectl get services
kubectl get svc  # Short form

# Describe service (shows endpoints)
kubectl describe service backend-service

# See which pods are behind service
kubectl get endpoints backend-service
```

---

## ConfigMaps & Secrets

### ConfigMaps

> **In simple words:** A place to store non-sensitive configuration data (like config files, environment variables) separately from your code. Change config without rebuilding containers!

**What it stores:**
- Configuration files
- Environment variables
- Command-line arguments
- Any non-sensitive data

**Why use it:**
- Decouple configuration from container image
- Same image works in dev/test/prod with different configs
- Change config without rebuilding image

#### Creating ConfigMaps

**Method 1: From literal values**
```bash
kubectl create configmap app-config \
  --from-literal=DATABASE_URL=mysql://db:3306 \
  --from-literal=LOG_LEVEL=debug
```

**Method 2: From file**
```bash
# Create config file first
echo "max_connections=100" > mysql.conf
echo "port=3306" >> mysql.conf

kubectl create configmap mysql-config --from-file=mysql.conf
```

**Method 3: YAML file**
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  DATABASE_URL: "mysql://db:3306"
  LOG_LEVEL: "debug"
  config.json: |
    {
      "port": 8080,
      "timeout": 30
    }
```

#### Using ConfigMaps

**Option 1: All keys as environment variables**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp
spec:
  containers:
  - name: app
    image: myapp:latest
    envFrom:
    - configMapRef:
        name: app-config  # All keys become env variables
```

**Option 2: Specific keys as environment variables**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp
spec:
  containers:
  - name: app
    image: myapp:latest
    env:
    - name: DB_URL  # Custom env variable name
      valueFrom:
        configMapKeyRef:
          name: app-config
          key: DATABASE_URL
    - name: LOG
      valueFrom:
        configMapKeyRef:
          name: app-config
          key: LOG_LEVEL
```

**Option 3: Mount as volume (files)**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp
spec:
  containers:
  - name: app
    image: myapp:latest
    volumeMounts:
    - name: config-volume
      mountPath: /etc/config  # Files appear here
  volumes:
  - name: config-volume
    configMap:
      name: app-config
```

**Result:** `/etc/config/DATABASE_URL` and `/etc/config/LOG_LEVEL` files created

### Secrets

> **In simple words:** Like ConfigMaps, but for sensitive data (passwords, API keys, tokens). Values are base64-encoded (not encrypted, just obfuscated). For real security, use external secret managers.

**What it stores:**
- Passwords
- API keys
- Tokens
- SSH keys
- TLS certificates

**Key difference from ConfigMap:** Values are base64-encoded and hidden when you describe them

#### Creating Secrets

**Method 1: From literal values**
```bash
kubectl create secret generic db-secret \
  --from-literal=username=admin \
  --from-literal=password=super-secret-123
```

**Method 2: YAML file**
```bash
# Encode values first
echo -n 'admin' | base64  # YWRtaW4=
echo -n 'super-secret-123' | base64  # c3VwZXItc2VjcmV0LTEyMw==
```

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: db-secret
type: Opaque
data:
  username: YWRtaW4=  # base64 encoded
  password: c3VwZXItc2VjcmV0LTEyMw==  # base64 encoded
```

**Method 3: From files**
```bash
echo -n 'admin' > ./username.txt
echo -n 'super-secret-123' > ./password.txt

kubectl create secret generic db-secret \
  --from-file=./username.txt \
  --from-file=./password.txt
```

#### Using Secrets

**Option 1: All keys as environment variables**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp
spec:
  containers:
  - name: app
    image: myapp:latest
    envFrom:
    - secretRef:
        name: db-secret
```

**Option 2: Specific keys as environment variables**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp
spec:
  containers:
  - name: app
    image: myapp:latest
    env:
    - name: DB_USER
      valueFrom:
        secretKeyRef:
          name: db-secret
          key: username
    - name: DB_PASS
      valueFrom:
        secretKeyRef:
          name: db-secret
          key: password
```

**Option 3: Mount as volume**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp
spec:
  containers:
  - name: app
    image: myapp:latest
    volumeMounts:
    - name: secret-volume
      mountPath: /etc/secrets
      readOnly: true
  volumes:
  - name: secret-volume
    secret:
      secretName: db-secret
```

### ConfigMap vs Secret

| Feature | ConfigMap | Secret |
|---------|-----------|--------|
| Purpose | Non-sensitive config | Sensitive data |
| Encoding | Plain text | Base64 encoded |
| Visibility | Visible in describe | Hidden in describe |
| Use case | Config files, env vars | Passwords, keys, tokens |

### Important Notes

⚠️ **Secrets are NOT encrypted**, just base64-encoded! For production:
- Use external secret managers (HashiCorp Vault, AWS Secrets Manager)
- Enable encryption at rest in etcd
- Limit RBAC access to secrets

📝 **Updating:** Changes to ConfigMaps/Secrets:
- Mounted as volumes: Updated automatically (with delay)
- Environment variables: Require pod restart

```bash
# View ConfigMaps/Secrets
kubectl get configmaps
kubectl get secrets

# View contents
kubectl describe configmap app-config  # Shows all data
kubectl describe secret db-secret      # Hides values

# Get actual secret values
kubectl get secret db-secret -o yaml

# Decode secret
echo 'YWRtaW4=' | base64 --decode  # admin

# Delete
kubectl delete configmap app-config
kubectl delete secret db-secret
```

---

## Hands-On Practice

### Template for Each Exercise

**Exercise: [Title]**  
**Date:** YYYY-MM-DD  
**Objective:** Brief description

**Commands:**
```bash
# Commands used
```

**Outcome:** Results  
**Learnings:** Key takeaways  
**Issues:** Problems and solutions

---

## Resources

- [Docker Docs](https://docs.docker.com/) | [Kubernetes Docs](https://kubernetes.io/docs/)
- [Tutorial Video](https://www.youtube.com/watch?v=kTp5xUtcalw&t=7723s)
- [Docker Cheat Sheet](https://docs.docker.com/get-started/docker_cheatsheet.pdf) | [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)

---

## Progress

- [x] Microservices & Cloud Native
- [x] Containers & Docker  
- [x] Docker Compose
- [x] Kubernetes Architecture
- [x] Pods, Labels, Networking
- [x] Workloads (Deployments, ReplicaSets, StatefulSets, DaemonSets, Jobs, CronJobs)
- [x] Services (ClusterIP, NodePort, LoadBalancer, ExternalName)
- [x] ConfigMaps & Secrets
- [ ] Ingress & Ingress Controllers
- [ ] Persistent Volumes & Storage
- [ ] RBAC & Security
- [ ] Helm & Package Management
- [ ] Advanced Topics

**Last Updated:** [Date]  
**Current Progress:** Completed Workloads & Services sections (up to ~4:30:00)