package api

type Task struct {
	ID int
	Title string
	Done bool
}

type APIResponse struct {
    Status  int         `json:"status"`
    Data    interface{} `json:"data"`
    Message string      `json:"message"`
}

type TaskStore struct {
	Tasks []Task
	Counter int
}