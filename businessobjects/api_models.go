package businessobjects

type CustomModelList struct {
	Items []CustomModel `json:"value"`
}

type CustomModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

type CreateCustomModelParams struct {
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

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
