package models

type Swipe struct {
	UserId       string `json:"user_id"`
	TargetUserId string `json:"target_user_id"`
	Swipe        bool   `json:"swipe"`
}
