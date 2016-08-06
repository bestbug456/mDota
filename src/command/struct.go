package command

type Userdata struct {
	Feature []float64 `json:"Feature,omitempty"`
	Name    string    `json:"Name,omitempty"`
}

type UsersFile struct {
	Users []Userdata `json:"users,omitempty"`
}
