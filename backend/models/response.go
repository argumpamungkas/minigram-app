package models

type ReponseInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type ReponsePostingAll struct {
	ReponseInfo
	Data []ResponsePosting `json:"data"`
}

type ReponsePostingById struct {
	ReponseInfo
	Data ResponsePosting `json:"data"`
}

type ResponsePosting struct {
	UserId   uint    `json:"user_id"`
	Username string  `json:"username"`
	Avatar   *string `json:"avatar"`
	Caption  *string `json:"caption"`
	Photo    string  `json:"photo"`
	Likes    int     `json:"likes"`
	Comments int     `json:"comments"`
}
