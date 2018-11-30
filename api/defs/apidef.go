package defs

type UserCredential struct {
	Username string 'json:"user_name"'
	Pwd string 'json:"pwd"'
}

//response
type SigedUp struct {
	Success bool 'json:"success"'
	SessionId string 'json:"session_id"'
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

type SimpleSession struct{
	Username string
	TTL int64
}