package home

import (
	"github.com/go-home-admin/home/bootstrap/utils"
	"gorm.io/gorm"
)

// @Table("users")
type UsersTable struct {
	gorm.Model
	Id       uint32 `json:"id"`
	UserName string `json:"user_name"`
}

func (receiver *UsersTable) TableName() string {
	return "users"
}

func (receiver *UsersTable) DB() *gorm.DB {
	return nil
}

func (receiver *UsersTable) WhereId(Id uint32) *UsersTable {
	return receiver
}

func (receiver *UsersTable) OrWhereId(Id uint32) *UsersTable {
	return receiver
}

func (receiver *UsersTable) WhereUserName(val string) *UsersTable {
	return receiver
}

func (receiver *UsersTable) Find() *UsersTable {
	return receiver
}

func (receiver *UsersTable) First() *UsersTable {
	return receiver
}

func (receiver *UsersTable) And(fuc func(table *UsersTable)) *UsersTable {
	return receiver
}

func (receiver *UsersTable) Or(fuc func(table *UsersTable)) *UsersTable {
	return receiver
}

func test() {
	u := &UsersTable{}
	data := u.WhereUserName("test").And(func(table *UsersTable) {
		table.WhereId(1).OrWhereId(2)
	}).Or(func(table *UsersTable) {
		table.WhereId(2).Or(func(table *UsersTable) {
			table.WhereId(1)
		})
	}).Find()

	// select * form users where user_name = ? and (id = ? or id = ?) or (id = ? or (id = ?))
	utils.Dump(data)

	data = u.WhereId(1).First()

}
