package service

type MeetingCreateRequest struct {
	Name     string `json:"name,omitempty"`
	CreateAt int64  `json:"create_at"`
	EndAt    int64  `json:"end_at"`
}
