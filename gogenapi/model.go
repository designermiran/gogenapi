package gogenapi

const (
	AssociationNone      = 0
	AssociationBelongsTo = 1
	AssociationHasMany   = 2
	AssociationHasOne    = 3
)

type Model struct {
	Name   string
	Fields []*Field
}

func (m *Model) AllPreloadAssocs() []string {
	result := []string{}

	for _, field := range m.Fields {
		result = append(result, field.PreloadAssocs()...)
	}

	return result
}

type Models []*Model // implements Sort interface

func (m Models) Len() int {
	return len(m)
}

func (m Models) Less(i, j int) bool {
	return m[i].Name < m[j].Name
}

func (m Models) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type Field struct {
	Name        string
	JSONName    string
	Type        string
	Tag         string
	Association *Association
}

func (f *Field) PreloadAssocs() []string {
	if f.Association == nil || f.Association.Type == AssociationNone {
		return []string{}
	}

	result := []string{
		f.Name,
	}

	for _, field := range f.Association.Model.Fields {
		if field.Association == nil || field.Association.Type == AssociationNone {
			continue
		}

		result = append(result, f.Name+"."+field.Name)
	}

	return result
}

func (f *Field) IsAssociation() bool {
	return f.Association != nil && f.Association.Type != AssociationNone
}

func (f *Field) IsBelongsTo() bool {
	return f.Association != nil && f.Association.Type == AssociationBelongsTo
}

type Association struct {
	Type  int
	Model *Model
}
