package converter

import (
	"github.com/google/uuid"
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func APIKeyToAPI(m *model.APIKey) *gen.APIKey {
	if m == nil {
		return nil
	}
	// Extract prefix from value (first 8 chars)
	keyPrefix := m.Value
	if len(keyPrefix) > 8 {
		keyPrefix = keyPrefix[:8]
	}
	k := &gen.APIKey{
		ID:          m.ID,
		Name:        m.Name,
		KeyPrefix:   keyPrefix,
		Permissions: []string{}, // Permissions would be derived from user
		CreatedAt:   m.CreatedAt,
	}
	if m.LastUsedAt != nil {
		k.LastUsedAt.SetTo(*m.LastUsedAt)
	}
	if m.ExpiresAt != nil {
		k.ExpiresAt.SetTo(*m.ExpiresAt)
	}
	return k
}

func APIKeysToAPI(models []*model.APIKey) []gen.APIKey {
	result := make([]gen.APIKey, len(models))
	for i, m := range models {
		result[i] = *APIKeyToAPI(m)
	}
	return result
}

func CreateAPIKeyRequestToModel(req *gen.CreateAPIKeyRequest, userID uuid.UUID, keyValue string) *model.APIKey {
	if req == nil {
		return nil
	}
	m := &model.APIKey{
		UserID: userID,
		Name:   req.Name,
		Value:  keyValue,
	}
	if req.ExpiresAt.IsSet() {
		exp := req.ExpiresAt.Value
		m.ExpiresAt = &exp
	}
	return m
}
