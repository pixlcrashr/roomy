
When modifying the database models, the following steps has to be followed to ensure that migrations will be created successfully:

1. Modify the database GORM models in `pkg/db/models`.
2. Create up- and down-migrations for the model changes for PostgreSQL in `pkg/db/migrations/postgres/`. Read `pkg/db/migrations/README.md` for a detailed description on how the structure of migration files should be.
3. Use `go run main.go -c config.dev.yaml migrate up` to apply the migrations to the development database.
