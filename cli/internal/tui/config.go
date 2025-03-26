package tui

// ProjectConfig represents the configuration for a new project
type ProjectConfig struct {
	ProjectName    string
	Database       string
	ORM            string
	Auth           string
	Docker         string
	Features       []string
	GraphQLLibrary string
}

// Database options
var DatabaseOptions = []string{
	"PostgreSQL",
	"MongoDB",
	"MySQL",
	"SQLite",
}

// ORM options based on selected database
func GetORMOptions(db string) []string {
	switch db {
	case "PostgreSQL", "MySQL", "SQLite":
		return []string{
			"GORM",
			"Ent",
			"SQLC",
			"Raw SQL",
		}
	case "MongoDB":
		return []string{
			"Official Go Driver",
			"mgm",
			"Raw Driver",
		}
	default:
		return []string{"GORM"}
	}
}

// Authentication options
var AuthOptions = []string{
	"JWT",
	"OAuth",
	"None",
}

// Docker options
var DockerOptions = []string{
	"None",
	"Development only",
	"Full (development + production)",
}

// Feature options
var FeatureOptions = []string{
	"WebSocket Subscriptions",
	"Redis Caching",
	"Background Jobs",
	"Metrics & Monitoring",
	"File Upload Support",
}

// GraphQL library options
var GraphQLOptions = []string{
	"gqlgen",
	"graphql-go",
}