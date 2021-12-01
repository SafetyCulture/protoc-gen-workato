package template

func (t *WorkatoTemplate) generateEnumPicklists() {
	for _, enum := range t.enums {
		pickListDef := &PicklistDefinition{
			Name:   enumPicklistName(enum),
			Values: []PicklistValue{},
		}

		removeUnspecifiedValue(enum)

		for _, value := range enum.Values {
			desc := value.Description
			if desc == "" {
				desc = value.Name
			}
			pickListDef.Values = append(pickListDef.Values, PicklistValue{value.Name, desc})
		}

		t.Picklists = append(t.Picklists, pickListDef)
	}
}
