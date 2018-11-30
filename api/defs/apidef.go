package defs

type UserCredential struct {
	Username string 'json:"user_name"'
	Pwd string 'json:"pwd"'
}

//date model
type VideoInfo struct {
	 ID string
	 AuthorId int
	 DisplayCtime string
}

type Comments struct {
	Id string
	VideoId string
	AuthorId int
	Content string
}