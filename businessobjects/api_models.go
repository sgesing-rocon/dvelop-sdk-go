package businessobjects

type EntityType struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	State      string     `json:"state"`
	Key        Key        `json:"key"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Indexed  bool   `json:"indexed"`
	Type     string `json:"type"`
	State    string `json:"state"`
}

type Key struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	State string `json:"state"`
}
