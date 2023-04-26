package service

import "time"

type MeetingListRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

type MeetingListReply struct {
	Identity string    `json:"identity"`
	Name     string    `json:"name,omitempty"`
	BeginAt  time.Time `json:"begin_at"`
	EndAt    time.Time `json:"end_at"`
}

type MeetingCreateRequest struct {
	Name    string `json:"name,omitempty"`
	BeginAt int64  `json:"begin_at"`
	EndAt   int64  `json:"end_at"`
}

type UserLoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type MeetingEditRequest struct {
	Identity string `json:"identity"`
	*MeetingCreateRequest
}
