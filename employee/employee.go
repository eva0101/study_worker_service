package employee

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Position string `json:"position"`
}
