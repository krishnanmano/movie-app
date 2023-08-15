package request

type BookingRequest struct {
	UserName string `json:"user_name"`
	EmailId  string `json:"email_id"`
	Seats    []uint `json:"seats"`
	ShowId   uint   `json:"show_id"`
}
