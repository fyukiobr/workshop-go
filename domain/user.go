package domain

type User struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name,omitempty"`
	Age       int                    `json:"age,omitempty"`
	Phones    []string               `json:"phones,omitempty"`
	Relatives map[string]interface{} `json:"-"`
}

func (u *User) String() string {
	return "Name: " + u.Name
}
