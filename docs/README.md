# Go Backend Template Documentation Mirror

This directory mirrors the source structure of the project to explain the purpose, interface design, and usage instructions of each module.

## Directory Mapping

```
├── cmd/               --> [docs/cmd/]              Application Entrypoint
├── pkg/
│   ├── controllers/   --> [docs/pkg/controllers/]  HTTP Controllers
│   ├── db/            --> [docs/pkg/db/]           SQLC & Migration Schemas
│   ├── dto/           --> [docs/pkg/dto/]          Data Transfer Objects & Validation
│   ├── middlewares/   --> [docs/pkg/middlewares/]  Echo Middlewares
│   └── utils/         --> [docs/pkg/utils/]        Helpers & Setup Utils
└── tests/             --> [docs/tests/]            Testing Strategies & Mocks
```

## Running the Server

### Local Development
1. Start PostgreSQL and Redis instances.
2. Edit `.env` values (copied from `.env.example`).
3. Run the development server with hot-reloading:
   ```bash
   make dev
   ```
4. Run standard Go server:
   ```bash
   make run
   ```

### Docker Development
1. Build and run via Docker Compose:
   ```bash
   docker compose up --build
   ```

## Interactive API Docs
Once running, visit `http://localhost:8080/docs` to view the interactive Scalar API documentation UI.
