package dbops

import (
	"log"
	"time"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"/video_server/api/defs"
	"/video_server/api/utils"
)

func AddUserCredential(loginName string,pwd string) error {
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)") //准备要执行的sql
	if err != nil {
		return err
	}
	_,err = stmtIns.Exec(loginName,pwd)  //exec表示执行sql
	if err != nil {
		return err
	}
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error) {
	stmtOut,err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s",err)
		return "",err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "",err
	}
	stmtOut.Close()
	return pwd,nil
}

func DeleteUser(loginName string,pwd string) error {
	stmtDel,err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("deleteUser error : %s",err)
		return err
	}

	_,err = stmtDel.Exec(loginName,pwd) 
	if err != nil {
		return err
	}

	stmtDel.Close()
	return nil
}

func AddNewVideo(aid int,name string)(*defs.VideoInfo,error){
	vid,err := utils.NewUUID()
	if err != nil {
		return nil,err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns,err := dbConn.Prepare(`INSERT INTO video_info
	(id,author_id,name,display_ctime) VALUES(?,?,?,?)`)

	if err != nil {
		return nil,err
	}

	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil {
		return nil,err
	}

	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}

	defer stmtIns.Close()
	return res,nil
}

func GetVideoInfo(vid string) (&defs.VideoInfo,error)