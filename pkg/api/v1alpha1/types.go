package v1alpha1

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Namespace struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}
