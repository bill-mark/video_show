package dbops

import (
	"log"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	stmtOut.QueryRow(loginName).Scan(&pwd)
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

	stmtDel.Exec(loginName,pwd) 
	stmtDel.Close()
	return nil
}