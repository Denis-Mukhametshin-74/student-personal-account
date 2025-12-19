package dto

type ScheduleItem struct {
	Time    string `json:"time"`
	Subject string `json:"subject"`
	Room    string `json:"room"`
	Teacher string `json:"teacher"`
}

type ScheduleDay struct {
	Day     string         `json:"day"`
	Lessons []ScheduleItem `json:"lessons,omitempty"`
}

type ScheduleResponse struct {
	Group string        `json:"group"`
	Days  []ScheduleDay `json:"days,omitempty"`
}
