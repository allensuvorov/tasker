# Tasker – Distributed HTTP Task Scheduler

Tasker is a lightweight, distributed job scheduler designed to handle HTTP tasks asynchronously. Built with Go, it leverages a microservices architecture to efficiently manage task queuing, execution, and status tracking.

---

## 🧩 Architecture Overview

Tasker comprises two decoupled microservices:

1. **API Server**  
   Handles incoming task submissions via HTTP, persists them to a PostgreSQL database, and provides endpoints to query task statuses.

2. **Agent Worker**  
   Continuously polls the database for new tasks, executes them concurrently, and updates their statuses upon completion.


---

## 📦 Features

- Asynchronous HTTP task execution
- Concurrent task processing with Go routines
- RESTful API for task management
- PostgreSQL integration for reliable persistence
- Modular microservices architecture
- Environment-based configuration

---

## ⚙️ Configuration

Set the PostgreSQL connection string using the `DATABASE_DSN` environment variable:

```bash
export DATABASE_DSN="postgres://user:password@localhost:5432/tasker_db?sslmode=disable"
