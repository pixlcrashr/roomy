package converter

import (
	"time"

	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func BlockingToAPI(m *model.Blocking) *gen.Blocking {
	if m == nil {
		return nil
	}
	b := &gen.Blocking{
		ID:           m.ID,
		EntityType:   gen.BlockingEntityType(m.EntityType),
		EntityId:     m.EntityID,
		BlockingType: gen.BlockingBlockingType(m.BlockingType),
		StartTime:    m.StartTime,
		EndTime:      m.EndTime,
		IsRecurring:  m.IsRecurring,
		Source:       gen.BlockingSourceOwn, // Default, may be overridden
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
	if m.Name != nil {
		b.Name.SetTo(*m.Name)
	}
	if m.Reason != nil {
		b.Reason.SetTo(*m.Reason)
	}
	if m.RecurrenceRule != nil {
		b.RecurrenceRule.SetTo(*m.RecurrenceRule)
	}
	if m.RecurrenceDuration != nil {
		b.RecurrenceDuration.SetTo(time.Duration(*m.RecurrenceDuration).String())
	}
	if m.RecurrenceEnd != nil {
		b.RecurrenceEnd.SetTo(*m.RecurrenceEnd)
	}
	return b
}

func BlockingsToAPI(models []*model.Blocking) []gen.Blocking {
	result := make([]gen.Blocking, len(models))
	for i, m := range models {
		result[i] = *BlockingToAPI(m)
	}
	return result
}

func CreateBlockingRequestToModel(req *gen.CreateBlockingRequest, entityType string, entityID uuid.UUID) *model.Blocking {
	if req == nil {
		return nil
	}
	m := &model.Blocking{
		EntityType:   entityType,
		EntityID:     entityID,
		BlockingType: model.BlockingType(req.BlockingType),
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		IsRecurring:  req.IsRecurring.Or(false),
	}
	if req.Name.IsSet() {
		name := req.Name.Value
		m.Name = &name
	}
	if req.Reason.IsSet() {
		reason := req.Reason.Value
		m.Reason = &reason
	}
	if req.RecurrenceRule.IsSet() {
		rule := req.RecurrenceRule.Value
		m.RecurrenceRule = &rule
	}
	if req.RecurrenceDuration.IsSet() {
		dur, err := time.ParseDuration(req.RecurrenceDuration.Value)
		if err == nil {
			durNanos := int64(dur)
			m.RecurrenceDuration = &durNanos
		}
	}
	if req.RecurrenceEnd.IsSet() {
		end := req.RecurrenceEnd.Value
		m.RecurrenceEnd = &end
	}
	return m
}
