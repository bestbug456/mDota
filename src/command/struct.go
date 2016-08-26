package command

type Userdata struct {
	Feature []float64 `json:"Feature,omitempty"`
	Name    string    `json:"Name,omitempty"`
}

type UsersFile struct {
	Users []Userdata `json:"users,omitempty"`
}

type Testdata struct {
	Feature []float64 `json:"i,omitempty"`
	Name    []string  `json:"o,omitempty"`
}

type TestFile struct {
	Users []Testdata `json:"training_data,omitempty"`
}
