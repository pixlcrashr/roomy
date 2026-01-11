package migrations

import "embed"

// PostgreSQL contains the embedded PostgreSQL migration files.
//
//go:embed postgres/*.sql
var postgresFS embed.FS
