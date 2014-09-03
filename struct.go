package main

type Feedback struct {
	FeedbackForm
	Time		string
}

type FeedbackForm struct {
	Id			int64  `form:"Id"`
	Name		string `form:"Name"`
	Url			string `form:"Url"`
	Info		string `form:"Info"`
	Operator	string `form:"Operator"`
	Solution	string `form:"solution"`
}

type CrmUser struct {
	Id			int64
	EmpId		string
	UserForm
	time		string
	priority	string
}

type UserForm struct {
	Username	string `form:"Username"`
	Password	string `form:"Password"`
}

type HostInfo struct {
	Host		string `json:"host"`
	Info		string `json:"info"`
	Option		string `json:"option"`
}

type HostsList struct {
	Param		[]string `json:"hosts"`
}
