# Todo App

A full-stack to-do application built with Go (Gin) for the backend, React for the frontend, and MongoDB for the database. It uses Docker Compose to run the entire stack in containers.

## 🚀 Features

- **Create, Read, Update, Delete (CRUD) operations** for to-dos
- **Sorting** by due date, title and order
- **Search functionality**
- **Responsive design** for mobile and desktop
- **Dockerized app** using Docker Compose for easy setup

## 🛠️ Technologies Used

- **Backend**: Go (Gin), MongoDB
- **Frontend**: React
- **Database**: MongoDB
- **Docker**: Docker Compose for containerization

## 🏗️ Project Structure

```plaintext
todo-app/
├── backend/              # Go backend (Gin)
│   ├── main.go
│   └── go.mod
├── frontend/             # React frontend
│   ├── public/
│   └── src/
├── docker-compose.yml    # Docker Compose config
├── .dockerignore
├── .gitignore
└── README.md

