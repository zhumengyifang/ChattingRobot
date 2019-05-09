package model

type Chatting struct {
	ReqType    int        `json:"reqType"`
	Perception *Perception `json:"perception"`
	UserInfo   *UserInfo   `json:"userInfo"`
}
