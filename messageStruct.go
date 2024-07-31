package messageservice

type MessageDB struct {
	Id             int     `json:"id" db:"id"`
	Message        string  `json:"message" db:"message" binding:"message"`
	Status         string  `json:"status" db:"status" binding:"status"`
	Processed_time *string `json:"processed_time,omitempty" db:"processed_time"`
	Date_create    string  `json:"date_create" db:"date_create"`
}
type NewMessage struct {
	Id      int    `json:"id"`
	Key     string `json:"key"`
	Payload string `json:"message"`
}
