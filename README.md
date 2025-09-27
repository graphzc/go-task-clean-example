# Go Clean Architecture Template

A production-ready Go REST API template following Clean Architecture principles with dependency injection, comprehensive tooling, and best practices.

## 🏗️ Architecture

This template implements Clean Architecture (also known as Hexagonal Architecture or Ports and Adapters Architecture) which provides:

- **Separation of Concerns**: Clear boundaries between different layers
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Testability**: Easy to unit test business logic in isolation
- **Maintainability**: Changes in one layer don't affect others
- **Framework Independence**: Business logic is independent of external frameworks

### Directory Structure

```text
.
├── cmd/                          # Application entry points
│   └── api/                      # REST API application
│       ├── main.go               # Application entry point
│       ├── di/                   # Dependency injection setup
│       └── server/               # Server configuration
├── internal/                     # Private application code
│   ├── config/                   # Configuration management
│   ├── domain/                   # Business entities and enums
│   │   ├── entities/             # Domain entities
│   │   └── enums/                # Domain enumerations
│   ├── dto/                      # Data Transfer Objects
│   ├── handlers/                 # HTTP request handlers (Controllers)
│   ├── infrastructure/           # External dependencies (databases, auth, etc.)
│   ├── repositories/             # Data access layer
│   ├── router/                   # HTTP routing
│   ├── services/                 # Business logic layer
│   └── utils/                    # Utility functions
├── docker-compose.yml            # Development environment setup
├── Dockerfile                    # Container configuration
├── Makefile                      # Build and development tasks
├── go.mod                        # Go module definition
└── go.sum                        # Go module checksums
```

### Layer Responsibilities

- **Entities**: Core business objects with enterprise-wide business rules
- **Services**: Application business rules and use cases
- **Repositories**: Data access abstraction layer
- **Handlers**: HTTP request/response handling (presentation layer)
- **Infrastructure**: External systems integration (databases, auth, etc.)

## 🚀 Features

- ✅ **Clean Architecture** with proper layer separation
- ✅ **Dependency Injection** using Google Wire
- ✅ **HTTP Server** with Echo framework
- ✅ **Database Integration** with SQLX and PostgreSQL
- ✅ **Request Validation** with go-playground/validator
- ✅ **Error Handling** with structured error responses
- ✅ **CORS Support** with configurable origins
- ✅ **Environment Configuration** with validation
- ✅ **Logging** with structured logging (zerolog)
- ✅ **Mock Generation** for testing
- ✅ **Docker Support** with multi-stage builds
- ✅ **Development Tools** (Makefile, hot reload ready)

## 🛠️ Tech Stack

- **Go 1.25+** - Programming language
- **Echo v4** - HTTP web framework
- **SQLX** - SQL extensions for database operations
- **PostgreSQL** - Primary database
- **Google Wire** - Dependency injection
- **Zerolog** - Structured logging
- **Testify** - Testing toolkit
- **Mockery** - Mock generation
- **Docker & Docker Compose** - Containerization

## 🏃 Quick Start

### Prerequisites

- Go 1.25 or higher
- Docker and Docker Compose
- Make (optional, for convenience commands)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/graphzc/go-task-clean-example.git
   cd go-clean-template
   ```

2. **Install Go dependencies:**

   ```bash
   go mod tidy
   ```

3. **Install development tools:**

   ```bash
   # Install Wire for dependency injection
   go install github.com/google/wire/cmd/wire@latest
   
   # Install Mockery for mock generation
   go install github.com/vektra/mockery/v2@latest
   
   # Install wiresetgen (if using custom wire sets)
   go install github.com/your-org/wiresetgen@latest
   ```

4. **Set up environment variables:**

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

5. **Start the database:**

   ```bash
   docker-compose up -d postgres
   ```

6. **Generate code and run the application:**

   ```bash
   make generate
   make start
   ```

The API will be available at `http://localhost:8080`

## ⚙️ Configuration

The application uses environment variables for configuration. Create a `.env` file based on the example:

```env
# Server Configuration
PORT=8080
LOG_FORMAT=json

# CORS Configuration
CORS_ALLOW_ORIGINS=http://localhost:3000,http://localhost:8080

# Database Configuration
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_NAME=yourapp
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=yourpassword
DATABASE_SSL_MODE=disable

# Optional: Google Cloud Storage
GOOGLE_APP_CREDENTIALS=path/to/credentials.json
UPLOAD_SLIP_BUCKET=your-bucket-name
```

## 🔨 Development

### Available Make Commands

```bash
# Generate all code (DI, mocks) and format
make generate

# Start the application
make start

# Run tests with coverage
make test

# Format code
make fmt

# Clean and regenerate mocks
make mock-clean mock-generate

# Generate dependency injection code
make di-generate
```

### Adding New Features

1. **Create Domain Entity** (if needed):

   ```go
   // internal/domain/entities/user.go
   type User struct {
       ID    int64  `json:"id"`
       Email string `json:"email"`
       Name  string `json:"name"`
   }
   ```

2. **Define Repository Interface**:

   ```go
   // internal/repositories/user/base.go
   type Repository interface {
       Create(ctx context.Context, user *entities.User) error
       GetByID(ctx context.Context, id int64) (*entities.User, error)
   }
   ```

3. **Implement Repository**:

   ```go
   // internal/repositories/user/base.go
   type repository struct {
       db *sqlx.DB
   }

   func NewRepository(db *sqlx.DB) Repository {
       return &repository{db: db}
   }
   ```

4. **Create Service**:

   ```go
   // internal/services/user/base.go
   type Service interface {
       CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*entities.User, error)
   }
   ```

5. **Add Handler**:

   ```go
   // internal/handlers/user/base.go
   type Handler interface {
       CreateUser(c echo.Context) error
   }
   ```

6. **Update Dependency Injection**:
   Add your new dependencies to the wire sets in `cmd/api/di/wire.go`

7. **Register Routes**:
   Add routes in `internal/router/api_router.go`

8. **Generate Code**:

   ```bash
   make generate
   ```

### Testing

Run tests with coverage:

```bash
make test
```

The template includes:

- Unit tests for services and repositories
- Integration tests for handlers
- Mock generation for all interfaces
- Test utilities and helpers

### Database Migrations

While not included in the base template, you can easily add database migrations using tools like:

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [pressly/goose](https://github.com/pressly/goose)

## 🐳 Docker

### Development with Docker Compose

```bash
# Start all services (app + database)
docker-compose up

# Start only the database
docker-compose up postgres

# Build and run the application
docker-compose up --build app
```

### Production Docker Build

```bash
# Build the Docker image
docker build -t go-clean-template .

# Run the container
docker run -p 8080:8080 --env-file .env go-clean-template
```

## 📚 API Documentation

### Health Check

```http
GET /health
```

### Example Endpoints (customize based on your entities)

```http
GET    /api/v1/foos          # List all foos
POST   /api/v1/foos          # Create a new foo
GET    /api/v1/foos/:id      # Get foo by ID
PUT    /api/v1/foos/:id      # Update foo
DELETE /api/v1/foos/:id      # Delete foo
```

## 🧪 Testing Strategy

The template supports multiple testing approaches:

- **Unit Tests**: Test individual functions and methods in isolation
- **Integration Tests**: Test the interaction between layers
- **Repository Tests**: Test database operations (use testcontainers for real DB tests)
- **Handler Tests**: Test HTTP endpoints with mocked dependencies

## 🔧 Customization

### Adding More Database Providers

1. Create new database implementation in `internal/infrastructure/database/`
2. Update configuration to support multiple providers
3. Use dependency injection to switch between implementations

### Integration with External Services

1. Create interfaces in appropriate service layers
2. Implement concrete types in `internal/infrastructure/`
3. Use dependency injection for testability

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Clean Architecture concepts by Robert C. Martin
- Echo framework team for the excellent HTTP framework
- Google Wire team for dependency injection
- The Go community for excellent tooling and libraries

---

## 🆘 Need Help?

- Check the [Issues](https://github.com/graphzc/go-task-clean-example/issues) for common problems
- Create a new issue if you find bugs or have feature requests
- Review the code comments and documentation in the source files

Happy coding! 🚀
