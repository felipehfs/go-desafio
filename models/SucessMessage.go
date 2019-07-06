package models

type SuccessMessage struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Tokenjwt string    `json:"tokenjwt"`
	Expires  string    `json:"expires"`
	Tokenmsg string    `json:"tokenmsg"`
	Login    LoginInfo `json:"Login" `
}

type LoginInfo struct {
	ID         int    `json:"id"`
	UUIDuser   string `json:"uuiduser"`
	Avatarurl  string `json:"avatarurl"`
	Avatartype string `json:"avatartype"`
	Name       string `json:"name"`
	DataStart  string `json:"datastart"`
}
