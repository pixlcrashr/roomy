package query

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

type AuditLogQuery interface {
	// SELECT * FROM @@table
	// {{where}}
	//   {{if userID != nil}} user_id = @userID {{end}}
	//   {{if action != nil}} AND action = @action {{end}}
	//   {{if entityType != nil}} AND entity_type = @entityType {{end}}
	//   {{if entityID != nil}} AND entity_id = @entityID {{end}}
	//   {{if startDate != nil}} AND created_at >= @startDate {{end}}
	//   {{if endDate != nil}} AND created_at <= @endDate {{end}}
	// {{end}}
	// ORDER BY created_at DESC
	// LIMIT @limit OFFSET @offset
	List(
		ctx context.Context,
		limit int,
		offset int,
		userID *uuid.UUID,
		action *string,
		entityType *string,
		entityID *uuid.UUID,
		startDate *time.Time,
		endDate *time.Time,
	) ([]*model.AuditLogEntry, error)

	// SELECT COUNT(*) FROM @@table
	// {{where}}
	//   {{if userID != nil}} user_id = @userID {{end}}
	//   {{if action != nil}} AND action = @action {{end}}
	//   {{if entityType != nil}} AND entity_type = @entityType {{end}}
	//   {{if entityID != nil}} AND entity_id = @entityID {{end}}
	//   {{if startDate != nil}} AND created_at >= @startDate {{end}}
	//   {{if endDate != nil}} AND created_at <= @endDate {{end}}
	// {{end}}
	CountAll(
		ctx context.Context,
		userID *uuid.UUID,
		action *string,
		entityType *string,
		entityID *uuid.UUID,
		startDate *time.Time,
		endDate *time.Time,
	) (int64, error)

	// INSERT INTO @@table (id, user_id, action, entity_type, entity_id, old_value, new_value, ip_address, user_agent, created_at)
	// VALUES (@id, @userID, @action, @entityType, @entityID, @oldValue, @newValue, @ipAddress, @userAgent, NOW())
	Insert(
		ctx context.Context,
		id uuid.UUID,
		userID *uuid.UUID,
		action string,
		entityType string,
		entityID uuid.UUID,
		oldValue *string,
		newValue *string,
		ipAddress *string,
		userAgent *string,
	) error
}
