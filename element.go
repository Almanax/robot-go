package robot

type Properties map[string][]string

type Element struct {
	ID     string     `json:"id"`
	Source string     `json:"source,omitempty"`
	Target string     `json:"target,omitempty"`
	Data   Properties `json:"data"`
}

type Elements []Element

func (l Elements) WithType(t string) Elements {
	var found Elements
	for _, e := range l {
		if e.HasType(t) {
			found = append(found, e)
		}
	}
	return found
}

func (e Element) IsNode() bool {
	return e.Source == ""
}

func (e Element) IsEdge() bool {
	return !e.IsNode()
}

func (e Element) Types() []string {
	return e.Data[RDF_TYPE]
}

func (e Element) Labels() []string {
	return e.Data[RDFS_LABEL]
}

func (e Element) HasType(t string) bool {
	return e.IsNode() && StringInSlice(t, e.Types())
}

func NewNode(id string, data Properties) Element {
	return Element{ID: id, Data: data}
}

func NewEdge(subj, pred, obj string) Element {
	id := subj + pred + obj
	data := Properties{RDF_PROPERTY: []string{pred}}
	return Element{ID: id, Data: data, Source: subj, Target: obj}
}
