package template

import "github.com/SafetyCulture/protoc-gen-workato/template/schema"

func (t *WorkatoTemplate) generateEnumPicklists() {
	for _, enum := range t.enums {
		pickListDef := &schema.PicklistDefinition{
			Name:   enumPicklistName(enum),
			Values: []schema.PicklistValue{},
		}

		removeUnspecifiedValue(enum)

		for _, value := range enum.Values {
			if !t.checkVisibility(value.Option("google.api.value_visibility")) {
				continue
			}

			desc := value.Description
			if desc == "" {
				desc = value.Name
			}
			pickListDef.Values = append(pickListDef.Values, schema.PicklistValue{
				Key:   value.Name,
				Value: escapeStringValue(desc),
			})
		}

		t.Picklists = append(t.Picklists, pickListDef)
	}
}
