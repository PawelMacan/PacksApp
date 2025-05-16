# 📦 Pack Optimizer App

This is a recruitment assignment project that provides an HTTP API and UI for calculating the most optimal combination of product packs to fulfill customer orders. It consists of a **Go backend (microservice)** and a **Next.js frontend**, containerized with Docker.

---

## 📘 Problem Statement

One of our product lines ships in various **pack sizes**:

- 250 items  
- 500 items  
- 1000 items  
- 2000 items  
- 5000 items  

### Customer order constraints:

1. Only **whole packs** may be used (no partial packs).  
2. Use the **minimum number of items** to fulfill the order.  
3. If multiple combinations have the same number of items, prefer the one with **fewer packs**.

> Example: for an order of `12001` items, the optimal result would be:  
> `2 × 5000`, `1 × 2000`, `1 × 250` → total 12250 → overage = 249

---

## 📂 Project Structure

```
project-root/
├── backend/        # Go backend (HTTP API)
├── frontend/       # Next.js frontend (React + Tailwind)
├── docker-compose.yml
```

---

## 🚀 Getting Started

### 🔧 Prerequisites

- Go 1.21+ or 1.24+
- Node.js 20+
- Docker & Docker Compose

---

## 🧪 Running Locally (without Docker)

### 🔹 Backend

```bash
cd backend
go run ./cmd/server
```

> API will be available at `http://localhost:8080`

### 🔹 Frontend

**Important**: Before running the frontend, you need to set up the environment variables.

1. Create a `.env.local` file in the `frontend` directory with the following content:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

2. Then start the frontend server:

```bash
cd frontend
npm install
npm run dev
```

> UI will be available at `http://localhost:3000`

**Note**: The `.env.local` file is required for the frontend to connect to the backend API when running locally. Without this file, you'll see connection errors.

---

## 🐳 Running with Docker

```bash
docker-compose up --build
```

This runs:
- `backend` on port `8080`
- `frontend` on port `3000`

> Access the app at [http://localhost:3000](http://localhost:3000)

---

## 🧪 Testing the Backend

Unit tests are available for the pack calculation logic.

```bash
cd backend
go test ./internal/service/...
```

---

## 🧠 Project Architecture

### Go Backend
- Frameworkless HTTP API
- Clean architecture with SOLID principles
- JSON config for pack sizes (`packs.json`)
- Easily extendable with new pack sizes

### Next.js Frontend
- React + Tailwind CSS
- Single-page interface for order input and result display
- Environment-based API URL configuration

---

## 🔍 Sample API Request

**Endpoint:** `POST /calculate`

**Request body:**

```json
{
  "amount": 12001
}
```

**Response:**

```json
{
  "packs": {
    "5000": 2,
    "2000": 1,
    "250": 1
  },
  "total_items": 12250,
  "requested_amount": 12001,
  "overage": 249,
  "total_packs": 4
}
```

---

## 📄 Future Improvements

- Persist user request history
- Support for authentication (e.g., for admin API)
- Deploy to cloud (e.g., Vercel, Render, Railway)
- Add unit/integration tests for the frontend

---

## 🧑‍💻 Author

Developed as part of a technical recruitment task.  
Stack: **Go**, **Next.js (React)**, **Tailwind CSS**, **Docker**.
