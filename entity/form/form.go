package form

type FormCreateBook struct {
	Title   string `json:"title,omitempty"`
	Author  string `json:"author,omitempty"`
	Edition int    `json:"edition,omitempty"`
}
