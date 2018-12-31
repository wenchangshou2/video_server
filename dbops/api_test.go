package dbops

import (
	"fmt"
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
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)

}
func testAddUser(t *testing.T) {
	fmt.Println("test adduser")
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("error of adduser:%v", err)
	}
	fmt.Println("success11")
}
func testGetUser(t *testing.T) {
	user, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}
	fmt.Println("user", user)
}
func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting fail")
	}
}
