package converter

import (
	"github.com/pixlcrashr/roomy/pkg/api/ogen/gen"
	"github.com/pixlcrashr/roomy/pkg/db/model"
)

func QRTemplateToAPI(m *model.QRTemplate) *gen.QRTemplate {
	if m == nil {
		return nil
	}
	return &gen.QRTemplate{
		ID:           m.ID,
		Name:         m.Name,
		HtmlTemplate: m.HTMLTemplate,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func QRTemplatesToAPI(models []*model.QRTemplate) []gen.QRTemplate {
	result := make([]gen.QRTemplate, len(models))
	for i, m := range models {
		result[i] = *QRTemplateToAPI(m)
	}
	return result
}

func CreateQRTemplateRequestToModel(req *gen.CreateQRTemplateRequest) *model.QRTemplate {
	if req == nil {
		return nil
	}
	return &model.QRTemplate{
		Name:         req.Name,
		HTMLTemplate: req.HtmlTemplate,
	}
}

func UpdateQRTemplateRequestToModel(req *gen.UpdateQRTemplateRequest, existing *model.QRTemplate) {
	if req == nil || existing == nil {
		return
	}
	if req.Name.IsSet() {
		existing.Name = req.Name.Value
	}
	if req.HtmlTemplate.IsSet() {
		existing.HTMLTemplate = req.HtmlTemplate.Value
	}
}
