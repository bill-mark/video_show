package dbops

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TeseUserWorkFlow(t *testing.T){
    t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T){
	err := AddUserCredential("avenssi","123")
	if err != nil{
		t.Errorf("Error of adduser:%v",err)
	}
}

func testGetUser(t *testing.T){
	pwd,err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil{
		t.Errorf("Error of getuser")
	}
}

func testDeleteUser(t *testing.T){
	err := DeleteUser("avenssi","123")
	if err != nil{
		t.Errorf("Error of deletuser:%v",err)
	}
}

func testRegetUser(t *testing.T){
	pwd,err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("error of regetuser:%v", err)
	}

	if pwd != "" {
		t.Errorf("deleting user test failed")
	}
}