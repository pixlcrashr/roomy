# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Roomy** is a room/place reservation and management system designed for organizations to manage bookable spaces (desks, meeting rooms, seating areas, etc.) with self-service user reservations and fine-grained permission controls.

## Core Domain Model

The system is built around a hierarchical structure:

- **Building** (highest level) - Contains areas, has name, description, location, opening hours
- **Area** - Groups multiple places within a building, can have room plans and serve as templates for place configuration
- **Place** (lowest level) - Individual bookable units (desk, chair, meeting room, etc.)
  - Can be configured with: location, name, seat capacity, description, bookability settings
  - May have equipment/amenities (projector, whiteboard, etc.)
  - Can be disabled or have time-based restrictions

**Reservation** - Links users to places for specific time slots
  - Requires check-in (via QR code or manual) unless configured otherwise
  - Subject to decay/timeout if not checked in
  - Can be self-service or manually assigned by organization members

**User & Permission System**:
- OAuth 2.0 authentication via GitLab
- Fine-grained permission/authorization using permission groups
- Two default groups: "system" (all permissions, immutable) and "default" (baseline for new users, editable)
- Organization members have different roles (viewer, editor, admin)

## Key Features

### User-Facing
- Self-service reservation with calendar view
- Public room availability list (anonymous viewing)
- QR code check-in for reservations
- Email notifications (reservation confirmations, cancellations, reminders)
- Room plan visualization for areas
- Search and filtering by equipment, capacity, availability
- Recurring reservations
- Favorite places
- Personal reservation dashboard

### Organization Member Features
- Manual reservation assignment
- Place/area/building management (CRUD operations)
- Room plan editor
- Configurable constraints per place/area:
  - Maximum reservation duration
  - Maximum reservations per time period (day/week/month/year)
  - Required check-in settings and timeout
  - User whitelists
  - Booking windows (how far in advance)
  - Minimum booking duration
- Blocking/maintenance scheduling (one-time or recurring)
- QR code generation and printing (configurable HTML templates)
- Usage statistics and metrics (anonymous)
- Reservation data export (CSV)
- Audit logs for changes
- User and group management
- API key creation for integrations

### System Capabilities
- API for external integrations (protected by API keys)
- API keys inherit permissions from the creating user
- Privacy compliance (user details hidden from other users, visible to organization members)

## Development Notes

### Project Status
This is an early-stage project currently in the planning phase. The `planning/feature-list.md` contains 88 user stories defining the complete feature set.

### Technology Stack

**Backend:**
- Go with GORM (database ORM)
- GORM Gen for type-safe query generation
- PostgreSQL database

**Requirements to align with:**
- OAuth 2.0 GitLab integration
- QR code generation and scanning
- Calendar/scheduling functionality
- Permission/authorization system
- API development (OpenAPI 3.0 spec in `openapi.yaml`)
- Room plan image upload and editing

---

## Database Layer & GORM Gen

The project uses GORM Gen for type-safe, generated database queries. **Never write queries using models from `pkg/db/models` directly.** Instead, define query interfaces in `pkg/db/query` and use the generated implementations from `pkg/db/gen`.

### Directory Structure

```
pkg/db/
├── models/       # GORM model definitions (structs with gorm tags)
├── query/        # Query interface definitions (YOU WRITE THESE)
├── gen/          # Auto-generated query implementations (DO NOT EDIT)
└── migrations/   # Database migrations
```

### Generating Queries

Run from repository root:
```bash
go generate ./...
```

This reads interfaces from `pkg/db/query` and generates implementations in `pkg/db/gen`.

### Defining Query Interfaces

Query interfaces use SQL annotations in comments. GORM Gen parses these and generates type-safe Go code.

#### Template Placeholders

| Placeholder | Description |
|-------------|-------------|
| `@@table` | Escaped & quoted table name of the model |
| `@@<name>` | Escaped & quoted table/column from parameter |
| `@<name>` | SQL query parameter (safe interpolation) |

#### Return Types

| Type | Description |
|------|-------------|
| `gen.T` | Returns the model struct |
| `gen.M` | Returns a map |
| `gen.RowsAffected` | Returns affected row count (int64) |
| `error` | Returns error if any |

#### Basic Query Example

```go
// pkg/db/query/building.go
package query

import "gorm.io/gen"

type BuildingQuerier interface {
    // SELECT * FROM @@table WHERE id = @id
    GetByID(id string) (gen.T, error)

    // SELECT * FROM @@table WHERE name LIKE @name
    SearchByName(name string) ([]gen.T, error)

    // INSERT INTO @@table (name, description, location) VALUES (@name, @description, @location)
    Create(name, description, location string) (gen.RowsAffected, error)
}
```

### Dynamic SQL with Conditionals

Use template expressions for conditional query building:

#### If/Else Conditions

```go
type PlaceQuerier interface {
    // SELECT * FROM @@table
    // {{where}}
    //   {{if areaId != ""}}
    //     area_id = @areaId
    //   {{end}}
    //   {{if minCapacity > 0}}
    //     AND capacity >= @minCapacity
    //   {{end}}
    //   {{if isBookable != nil}}
    //     AND is_bookable = @isBookable
    //   {{end}}
    // {{end}}
    Search(areaId string, minCapacity int, isBookable *bool) ([]gen.T, error)
}
```

#### Where Clause Builder

The `{{where}}` block intelligently:
- Adds `WHERE` only if conditions exist
- Trims unnecessary `AND`/`OR` at the start

#### Set Clause Builder (for UPDATE)

```go
type PlaceQuerier interface {
    // UPDATE @@table
    // {{set}}
    //   {{if name != ""}} name = @name, {{end}}
    //   {{if description != ""}} description = @description, {{end}}
    //   updated_at = NOW()
    // {{end}}
    // WHERE id = @id
    Update(id, name, description string) (gen.RowsAffected, error)
}
```

The `{{set}}` block auto-removes trailing commas.

#### For Loop (IN clauses)

```go
type ReservationQuerier interface {
    // SELECT * FROM @@table WHERE status IN (
    //   {{for _, status := range statuses}}
    //     @status,
    //   {{end}}
    // )
    GetByStatuses(statuses []string) ([]gen.T, error)
}
```

### Time-Based Queries

```go
type BlockingQuerier interface {
    // SELECT * FROM @@table WHERE
    // {{where}}
    //   entity_type = @entityType AND entity_id = @entityId
    //   {{if !startDate.IsZero()}}
    //     AND start_time >= @startDate
    //   {{end}}
    //   {{if !endDate.IsZero()}}
    //     AND end_time <= @endDate
    //   {{end}}
    // {{end}}
    GetByEntityAndDateRange(entityType, entityId string, startDate, endDate time.Time) ([]gen.T, error)
}
```

### Using Generated Queries

```go
// In your service/handler code
import "your-module/pkg/db/gen"

func (s *BuildingService) GetBuilding(ctx context.Context, id string) (*models.Building, error) {
    // Use the generated query interface, NOT direct model access
    q := gen.Use(s.db)
    building, err := q.Building.GetByID(id)
    if err != nil {
        return nil, err
    }
    return &building, nil
}
```

### Key Rules

1. **Never query models directly** - Always use interfaces from `pkg/db/query`
2. **Run `go generate ./...`** after modifying query interfaces
3. **Don't edit `pkg/db/gen`** - It's auto-generated and will be overwritten
4. **Use template expressions** for dynamic queries instead of string concatenation
5. **All SQL parameters use `@param`** - This ensures proper escaping and prevents SQL injection

---

## License

GNU General Public License v3.0 (GPL-3.0)
