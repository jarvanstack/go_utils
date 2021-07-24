package gormu

import (
	"fmt"
	"testing"
)

type TestUser struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

var db *DB

func init() {
	db, _ = NewDB("root:xxxxx@tcp(172.16.16.17:3306)/youke")
}
func TestCreateSysResult(t *testing.T) {
	err := db.DB.AutoMigrate(&TestUser{})
	if err != nil {
		fmt.Printf("err=%#v\n", err)
	}
}
func TestInsertResult(t *testing.T) {
	sql := `
	INSERT INTO youku.test_users 
	( id, a_user_id, b_user_id, a_result, b_result, round_num )
	VALUES
	(?,?,'?','?','?',?)
	`
	id := "1"

	s := db.DB.Exec(sql, id, 1, 1, 1, 1, 1).Error
	fmt.Printf("s=%#v\n", s)

}
func Test原始SQL(t *testing.T) {
	var user TestUser
	sql := `
	SELECT
		sys_user.user_id, 
		sys_user.username, 
		sys_user.email, 
		sys_user.password, 
		sys_user.nickname, 
		sys_user.avatar
	FROM
		sys_user
	WHERE
		sys_user.password = ? AND
		sys_user.username = ?
	limit 1
	`
	db.DB.Raw(sql, "admin", "admin").Scan(&user)
	fmt.Printf("user=%#v\n", user)
}
