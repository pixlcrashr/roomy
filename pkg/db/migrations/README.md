# Database Migrations

This directory contains PostgreSQL database migrations managed by [golang-migrate](https://github.com/golang-migrate/migrate).

## Directory Structure

```
migrations/
├── embed.go                          # Go embed directive for SQL files
├── migrate.go                        # Migration helper functions
├── README.md                         # This file
└── postgresql/                       # PostgreSQL-specific migrations
    ├── 000001_initial_schema.up.sql
    ├── 000001_initial_schema.down.sql
    └── ...
```

## Migration File Naming Convention

Migration files **MUST** follow this exact naming pattern:

```
{version}_{description}.{direction}.sql
```

Where:
- **`version`**: 6-digit zero-padded sequential number (e.g., `000001`, `000002`)
- **`description`**: Snake_case description of the migration (e.g., `initial_schema`, `add_users_table`)
- **`direction`**: Either `up` (apply) or `down` (rollback)

### Examples

```
000001_initial_schema.up.sql      # Creates initial tables
000001_initial_schema.down.sql    # Drops all initial tables
000002_add_audit_log.up.sql       # Adds audit_log table
000002_add_audit_log.down.sql     # Drops audit_log table
```

## Creating a New Migration

### Step 1: Determine the Next Version Number

List existing migrations and increment the highest version by 1:

```bash
ls src/backend/pkg/database/migrations/postgresql/
```

### Step 2: Create Migration Files

Create **both** up and down migration files:

```bash
# Replace NNNNNN with the next version number and description with your migration name
touch src/pkg/db/migrations/postgresql/NNNNNN_description.up.sql
touch src/pkg/db/migrations/postgresql/NNNNNN_description.down.sql
```

### Step 3: Write the Up Migration

The up migration should contain SQL statements to apply the schema change:

```sql
-- Example: 000002_add_audit_log.up.sql
CREATE TABLE audit_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    table_name VARCHAR(255) NOT NULL,
    record_id UUID NOT NULL,
    action VARCHAR(50) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_audit_log_table_record ON audit_log(table_name, record_id);
```

### Step 4: Write the Down Migration

The down migration **MUST** reverse everything done in the up migration:

```sql
-- Example: 000002_add_audit_log.down.sql
DROP INDEX IF EXISTS idx_audit_log_table_record;
DROP TABLE IF EXISTS audit_log;
```

## Critical Rules for AI Agents

### DO:

1. **Always create both up AND down migrations** - Never create only one direction
2. **Use sequential version numbers** - Check existing migrations first
3. **Make down migrations fully reversible** - They must undo everything in the up migration
4. **Use `IF EXISTS` / `IF NOT EXISTS`** - Prevents errors on re-runs
5. **Drop in reverse order** - Due to foreign key constraints, drop dependent objects first
6. **Test both directions** - Run up, then down, then up again to verify

### DON'T:

1. **Never modify existing migrations** - Create a new migration instead
2. **Never skip version numbers** - Must be sequential (000001, 000002, 000003...)
3. **Never use timestamps as versions** - Use 6-digit sequential numbers
4. **Never leave down migrations empty** - They must properly rollback
5. **Never delete migration files** - This breaks version tracking

### SQL Best Practices:

```sql
-- Use explicit schema for clarity
CREATE TABLE public.my_table (...);

-- Always specify NOT NULL where appropriate
column_name VARCHAR(255) NOT NULL,

-- Use TIMESTAMPTZ for timestamps (timezone-aware)
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

-- Use UUID for primary keys
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

-- Name constraints explicitly for easier debugging
CONSTRAINT fk_orders_customer FOREIGN KEY (customer_id) REFERENCES customers(id),

-- Name indexes explicitly
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
```

## Running Migrations

### Apply All Pending Migrations

```bash
go run main.go migrate up
```

### Rollback Last Migration

```bash
go run main.go migrate down
```

### Rollback All Migrations

```bash
go run main.go migrate down --all
```

### Check Current Version

```bash
go run main.go migrate version
```

### Apply/Rollback N Steps

```bash
# Apply 2 migrations
go run main.go migrate steps 2

# Rollback 1 migration
go run main.go migrate steps -1
```

### Force Version (Fix Dirty State)

If a migration fails partway through, the database may be in a "dirty" state:

```bash
# Force version to N (use after manually fixing the database)
go run main.go migrate force N

# Clear version entirely
go run main.go migrate force -1
```

## Troubleshooting

### Dirty Database State

If you see "dirty database version N", it means a migration failed partway through:

1. Check what was partially applied
2. Manually fix the database state
3. Force the version: `go run main.go migrate force N`

### Migration Not Found

Ensure:
1. Files are in `postgresql/` subdirectory
2. File names match the exact pattern
3. Both `.up.sql` and `.down.sql` exist

### Embedded Files Not Updating

After adding new migration files, rebuild the binary:

```bash
go build -o roomy.exe main.go
```

## Schema Documentation

The initial schema (`000001_initial_schema`) creates these tables:

| Table | Description |
|-------|-------------|
| `accounts` | Chart of accounts for bookkeeping |
| `fiscal_years` | Fiscal year periods |
| `receipts` | Stored receipt documents |
| `transactions` | Financial transactions |
| `events` | Event type definitions |
| `event_schemas` | Versioned JSON schemas for events |
| `transaction_templates` | Templates for auto-generating transactions |
| `event_instances` | Instances of events with validated data |
| `event_instance_transactions` | Links events to generated transactions |

Custom types:
- `transaction_direction` - ENUM: 'debit', 'credit'
