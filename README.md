# Todo App

A full-stack to-do application built with Go (Gin) for the backend, React for the frontend, and MongoDB for the database. Uses Docker Compose to manage the stack in containers.

## 🚀 Features
- CRUD operations for to-dos
- Sorting by due date and title
- Search functionality
- Dockerized with Docker Compose for easy setup

## 🛠️ Technologies
- **Backend**: Go (Gin), MongoDB
- **Frontend**: React
- **Database**: MongoDB
- **Containerization**: Docker Compose

## 🏗️ Project Structure
```plaintext
.
├── backend
│   ├── db
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── models
│   └── routes
├── docker-compose.yml
├── frontend
│   ├── Dockerfile
│   ├── node_modules
│   ├── package-lock.json
│   ├── package.json
│   ├── public
│   ├── README.md
│   └── src
└── README.md
```

## ⚙️ Installation and Setup

### Prerequisites
- Docker & Docker Compose must be installed on your system.

### Setup
1. **Clone the repository**:
    ```bash
    git clone https://github.com/JialongHe/Todo-App
    cd todo-app
    ```

2. **Build and start the app**:
    ```bash
    docker-compose up --build
    ```

    This will:
    - Build the frontend and backend images
    - Start the MongoDB, backend, and frontend services
    - Expose the frontend at [http://localhost:3000](http://localhost:3000) and the backend at [http://localhost:8080](http://localhost:8080)

3. **Open the app** in your browser:
    - Frontend: [http://localhost:3000](http://localhost:3000)
    - Backend API: [http://localhost:8080](http://localhost:8080)

### Development Mode
- To reflect changes in the code, modify files in the `frontend/` or `backend/` folders, then rebuild the images:
    ```bash
    docker-compose up --build
    ```

## 🔧 Environment Variables

- **REACT_APP_API_URL**: Backend API URL (used by frontend). Adjust if you want to run locally.
- **MONGO_URI**: MongoDB URI (used by backend).