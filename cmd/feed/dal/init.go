package dal

import "github.com/1037group/dousheng/cmd/feed/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
