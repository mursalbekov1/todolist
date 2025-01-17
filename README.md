# Golang To Do application

Welcome to TodoList, a simple task management application written in Go. TodoList allows you to manage your tasks efficiently, providing features to add, update, delete, and retrieve tasks.
## Features

- Facilitates proxying of HTTP requests and returning responses.
- Docker support for containerization.
- Task Management: Allows users to create, update, and delete tasks.
- Completion Tracking: Supports marking tasks as completed or active. -
- Date Management: Records the creation and completion dates of tasks.
- Flexible Filtering: Enables filtering tasks based on completion status and creation dates.
- RESTful API: Provides a RESTful API for easy integration with other applications.
- Simple Interface: User-friendly interface for managing tasks efficiently.
- Configurable Environment: Supports configuration via environment variables or configuration files.

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/mursalbekov1/todolist.git
```

### Run and build the service:

```bash
make run
```

### Deployment link:

```bash
https://todolist-jlm4.onrender.com/v1/healthCheck
```

### Request

```json
    {
      "title":"my task"
    }

```

### Response

```json
    {
      "id": 7,
      "title": "my task",
      "activeAt": "13 Jul 24 21:23 UTC",
      "completed": false
    }
```