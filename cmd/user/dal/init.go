package dal

import "github.com/1037group/dousheng/cmd/user/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
