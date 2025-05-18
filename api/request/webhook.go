package request

type HookRequest struct {
	Pusher
	Compare    string `json:"compare"`
	Repository string `json:"repository"`
	Before     string `json:"before"`
	After      string `json:"after"`
}

type Pusher struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Date     string `json:"date"`
}
