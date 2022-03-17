package user

type User struct {
	UUID        string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
}
