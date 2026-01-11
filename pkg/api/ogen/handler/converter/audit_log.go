package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func AuditLogEntryToAPI(m *model.AuditLogEntry) *gen.AuditLogEntry {
	if m == nil {
		return nil
	}
	e := &gen.AuditLogEntry{
		ID:         m.ID,
		EntityType: gen.AuditLogEntryEntityType(m.EntityType),
		EntityId:   m.EntityID,
		Action:     gen.AuditLogEntryAction(m.Action),
		UserId:     m.UserID,
		Timestamp:  m.Timestamp,
	}
	// Changes is a JSON object, we'll leave it empty for now
	return e
}

func AuditLogEntriesToAPI(models []*model.AuditLogEntry) []gen.AuditLogEntry {
	result := make([]gen.AuditLogEntry, len(models))
	for i, m := range models {
		result[i] = *AuditLogEntryToAPI(m)
	}
	return result
}
