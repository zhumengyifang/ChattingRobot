package model

type Chatting struct {
	ReqType    int        `json:"reqType"`
	Perception *Perception `json:"perception"`
	UserInfo   *UserInfo   `json:"userInfo"`
}

type InputText struct {
	Text string `json:"text"`
}

type Perception struct {
	InputText *InputText `json:"inputText"`
}

type UserInfo struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}