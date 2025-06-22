## 📚 Localization / Russian Documentation

🇷🇺 Prefer reading in Russian? See [README.ru.md](./README.ru.md) for full documentation in Russian.

# Go Project Template

[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Powered-blue.svg)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

**This is not just another project – it’s a fully functional boilerplate for quickly bootstrapping your Go web
applications.**

It’s designed to save you from repetitive environment setup tasks. Instead of wasting time configuring Docker,
hot-reload, and debugging tools, you can focus directly on writing your business logic.  
Use this repository as a solid foundation for your next Go project!

## 🚀 How to Use This Template

There are two recommended ways to get started:

**1. (Recommended) Use the GitHub Template:**

Click the green **"Use this template"** button at the top of the repository page.  
GitHub will automatically create a new repository under your account with all files copied — but without any commit
history.

**2. Manual Setup:**

If you prefer to set things up locally:

```bash
# 1. Clone the repository with a new name
git clone https://github.com/scarymovie/go-template.git my-new-project

# 2. Navigate to the project directory
cd my-new-project

# 3. Remove template git history and initialize your own
rm -rf .git
git init
git add .
git commit -m "Initial commit from template"

# 4. Don’t forget to update the module name in go.mod
# and adjust service names in docker-compose files.
```

## 🚀 Getting Started

### Prerequisites

Make sure you have the following tools installed:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)  
  *(usually included with Docker Desktop)*

### Installation and Running the Project

1. **Start the development environment using Docker Compose:**

   This command will build a Docker image for your Go application and start the containers for both the app and the
   PostgreSQL database.

   ```bash
   docker-compose -f docker/docker-compose.development.yml up --build

   - `-f docker/docker-compose.development.yml`:  specifies the path to the compose file.
   - `up`: starts the services.
   - `--build`: forces Docker to rebuild the images before starting.
   ```

Once everything is running, you'll see logs from air and your application in the terminal.
The server will be accessible at http://localhost:8080.

## 💡 Usage

### Hot Reload with Air

Thanks to **Air**, any changes to `.go` or `.yaml` files inside the `app/` directory will automatically trigger
a rebuild and restart of your application inside the container.  
Just save your file and watch the logs in the terminal — no manual restarts needed.

### Remote Debugging with Delve

The environment is preconfigured for remote debugging.  
The `go` container runs the **Delve** debugger and listens on port `2345`.

#### Setting Up in GoLand

1. Go to **Run > Edit Configurations > + > Go Remote**.
2. Set **Host** to `localhost`.
3. Set **Port** to `2345`.
4. Make sure your Docker containers are running (`docker-compose ... up`).
5. Set a breakpoint in your code, for example in `app/cmd/app/main.go`.
6. Start the debug session.

The debugger will attach to the Go process running inside the container and pause execution at your breakpoint.  
To reconnect after disconnecting, you may need to make a code change (to trigger hot reload) or restart the container manually.

### Connecting to PostgreSQL

- **Host:** `localhost` (when connecting from the host machine, since the port is exposed),  
  or `template-postgres-development` (when connecting from another container in the `internal` Docker network).
- **User:** `db_user`
- **Password:** `db_password`
- **Database:** `db_database`


## 🛠 TODO / Planned Improvements

- [ ] Add CI/CD pipeline for production deployments
- [ ] Add unit & integration tests
- [ ] Improve error handling with middleware
- [ ] Add DI building