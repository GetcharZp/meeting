package service

type MeetingCreateRequest struct {
	Name     string `json:"name,omitempty"`
	CreateAt int64  `json:"create_at"`
	EndAt    int64  `json:"end_at"`
}

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}
