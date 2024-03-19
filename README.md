# Project Name: seucy

## Folder Structure

The project is organized into several directories, each serving a specific purpose. Here's an overview of the folder structure:

### Backend Service (`backend-service`)

- `cmd/backend`: Main application entry point.
- `internal/controllers`: Controllers handling business logic.
- `internal/middleware`: Middleware for request processing.
- `internal/models`: Data models for the backend.
- `internal/routes`: Routing configuration.
- `internal/services`: Business logic services.
- `internal/utils`: Utility functions.
- `internal/app.go`: Main application configuration.
- `go.mod`, `go.sum`: Go module files.

### Client View (`client-view`)

- `cmd/client-view`: Main application entry point.
- `internal`: Internal application logic.
- `public`: Static assets like HTML and CSS.
  - `public/index.html`: Main HTML file.
  - `public/styles.css`: Stylesheet.
- `go.mod`, `go.sum`: Go module files.

### Engine API (`engine-api`)

- `cmd/engine-api`: Main application entry point.
- `internal/controllers`: Controllers handling API logic.
- `internal/models`: Data models for the API.
- `internal/services`: Business logic services for the recommendation engine.
- `internal/utils`: Utility functions.
- `go.mod`, `go.sum`: Go module files.

### Scripts

- Additional scripts for project-related tasks.

### Tests

- Unit tests and test-related files.

### `.gitignore`

- Specifies intentionally untracked files to ignore when using Git.

### `README.md`

- This file, providing an overview of the project structure and usage instructions.

### `docker-compose.yml`

- Docker Compose configuration for containerized development.

### `Makefile`

- Makefile with common development tasks.

### `go.mod`, `go.sum`

- Go module files for managing dependencies.

## How to Use

### Backend Service (`backend-service`):

- Handle user authentication, manage movie and content data.
- Implement business logic in controllers, services, and configure routes in routes.
- Adjust the `app.go` file for main application configuration.

### Client View (`client-view`):

- Develop the frontend using HTML and CSS in the `public` directory.
- Configure client-side logic in the `cmd/client-view/main.go` file.

### Engine API (`engine-api`):

- Focus on the recommendation engine logic in the `services` and `controllers` directories.
- Define API routes and models in the `routes` and `models` directories.

### Scripts:

- Use the `scripts` directory for additional project-related tasks.

### Tests:

- Write unit tests in the `tests` directory to ensure code quality.

### Docker Compose:

- Utilize the `docker-compose.yml` file for containerized development and deployment.

### Makefile:

- Leverage the `Makefile` for common development tasks.

## Contributing

Feel free to contribute to this project by opening issues, submitting pull requests, or suggesting improvements.
