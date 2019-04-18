package dao

import "github.com/xormplus/xorm"

type Base struct {
	Orm *xorm.Engine
}
