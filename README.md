# 🚀 GoGraphQLify

<p align="center">
  <img src="https://gqlgen.com/images/gqlgen-logo.jpg" alt="GoGraphQLify Logo" width="300" />
  <br>
  <em>The magical GraphQL starter kit for Go that writes the boilerplate for you!</em>
</p>

<div align="center">
  
![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)
![GraphQL](https://img.shields.io/badge/GraphQL-E10098?style=for-the-badge&logo=graphql)
![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)
[![Twitter Follow](https://img.shields.io/twitter/follow/gographqlify?style=for-the-badge&logo=twitter&color=1DA1F2)](https://twitter.com/gographqlify)

</div>

## ✨ Create GraphQL APIs in Go with Zero Boilerplate Pain

```bash
$ go-graphqlify create my-awesome-api

🧙‍♂️ Welcome to GoGraphQLify! Let's build something awesome together!

? Choose your database:
  ● PostgreSQL
  ○ MongoDB
  ○ MySQL
  ○ SQLite

? Select an ORM:
  ● GORM
  ○ Ent
  ○ SQLC

? Authentication method:
  ● JWT
  ○ OAuth
  ○ None

? Docker configuration:
  ○ None
  ○ Development only
  ● Full (development + production)

? Enable additional features:
  ✓ WebSocket Subscriptions
  ✓ Redis Caching
  ✓ Background Jobs

🚧 Crafting your GraphQL API... 
✅ Done! Your project is ready at ./my-awesome-api
```

## 🌟 Why GoGraphQLify?

Ever spent hours setting up a new GraphQL project in Go? Wiring up databases, authentication, middleware, error handling, subscriptions, and more?

**No more!** GoGraphQLify gives you a production-ready GraphQL API in seconds, not days.

### 🔥 Features

- **Interactive CLI** - Like `create-next-app` but for Go GraphQL APIs
- **Database Options** - PostgreSQL, MongoDB, MySQL, SQLite with migrations
- **ORM Flexibility** - Choose from GORM, Ent, or SQLC
- **Authentication Built-in** - JWT or OAuth2 ready to go
- **GraphQL Supercharged** - Dataloaders, subscriptions, error handling
- **Developer Experience** - Hot reload, GraphQL playground, detailed docs
- **Production Ready** - Docker, CI/CD, monitoring, structured logging
- **Modular Architecture** - Clean separation of concerns using Go best practices

## 🏃‍♂️ Quick Start

```bash
# Install the CLI tool
go install github.com/go-graphqlify/cli/cmd/go-graphqlify@latest

# Create a new project
go-graphqlify create my-api

# Move into project directory
cd my-api

# Run the development server
make run
```

Then visit http://localhost:8080/playground to start exploring your GraphQL API!

## 🏗️ What's Generated?

```
my-api/
├── cmd/                  # Entry points
│   └── server/           # API server
├── config/               # Configuration management
├── db/                   # Database migrations and models
├── docker/               # Docker configuration
├── graph/                # GraphQL specific code
│   ├── generated/        # Auto-generated GraphQL code
│   ├── resolvers/        # Your GraphQL resolvers
│   ├── dataloaders/      # N+1 query protection
│   └── schema.graphqls   # GraphQL schema
├── internal/             # Business logic
│   ├── auth/             # Authentication logic
│   ├── services/         # Business services
│   └── utils/            # Internal utilities
├── pkg/                  # Reusable packages
├── Dockerfile            # Production Docker build
├── docker-compose.yml    # Local development environment
├── Makefile              # Common commands
└── README.md             # Project-specific docs
```

## 💻 Development Commands

```bash
# Start the development server with hot reload
make run

# Run tests
make test

# Generate GraphQL code from schema
make generate

# Run database migrations
make migrate

# Build for production
make build
```

## 🧩 Add Features After Setup

Already have a project and want to add new features? No problem!

```bash
cd my-api
go-graphqlify add auth --type jwt
go-graphqlify add cache --type redis
go-graphqlify add subscription
```

## 📝 License

MIT License - see [LICENSE](./LICENSE) for details.

---

<p align="center">
  Made with ❤️ by the GoGraphQLify team
</p>
