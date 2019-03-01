package main

/*
  个人认为soft delete的最佳实践是用deleted_at默认值设置为"0001-01-01 00:00:00",然后要删除是设置为当前时间，优点:
1.非null值，可以和其他列建立联合唯一索引，保证其他列的唯一性
2.非0, 1值，不会导致无法"重复删除", 比如已经设置了("test", 1)，下个("test", 1)无法被插入
3.在xorm中会自动加上and (deleted_at is null or delete_at = "0001-01-01"),导致无法使用到索引
*/

import (
	"time"

	logger "github.com/caitinggui/seelog"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *xorm.Engine

// CondDeleted returns the conditions whether a record is soft deleted.
//func (engine *Engine) CondDeleted(colName string) builder.Cond {
//if engine.dialect.DBType() == core.MSSQL {
//return builder.IsNull{colName}

//}
//return builder.IsNull{colName}.Or(builder.Eq{colName: zeroTime1})

//}

type Category struct {
	Id        uint64    `xorm:"pk autoincr"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	Name      string    `xorm:"varchar(32) notnull unique(ux_name_deleted)"`
	DeletedAt time.Time `xorm:"deleted default('0001-01-01 00:00:00') notnull unique(ux_name_deleted)"`
}

func (self *Category) Insert() error {
	res, err := DB.Insert(self)
	logger.Info("插入结果: ", res, err)
	return err
}

func (self *Category) Delete() error {
	res, err := DB.Delete(self)
	logger.Info("删除结果: ", res, err)
	return err
}

func GetCategoryByName(name string) (cate Category, err error) {
	DB.Where("name = ?", name).Get(&cate)
	logger.Info("查询结果: ", cate, err)
	return
}

func main() {
	defer logger.Flush()
	DB, _ = xorm.NewEngine("sqlite3", "./test.db")
	DB.ShowSQL(true)
	cate := Category{}
	DB.CreateTables(&cate)
	DB.CreateIndexes(&cate)
	DB.CreateUniques(&cate)
	cate1 := Category{Name: "test1"}
	cate1.Insert()
	cate1.Delete()
	cate2 := Category{Name: "test1"}
	cate2.Insert()
	GetCategoryByName("test1")
}
