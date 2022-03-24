package demo

import (
	"database/sql"
	"github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	"gorm.io/gorm"
)

type TableName struct {
	Id       int64         `gorm:"primaryKey;column:id"`        //
	Date     database.Time `gorm:"column:date"`                 // 时间
	Guard    string        `gorm:"column:guard"`                // guard
	Ip       string        `gorm:"column:ip"`                   // ip
	Uri      string        `gorm:"primaryKey;column:uri"`       // 请求地址
	Params   string        `gorm:"column:params"`               // 参数
	UserId   int64         `gorm:"column:user_id"`              // uid
	User     string        `gorm:"column:user"`                 // uid
	HttpType string        `gorm:"primaryKey;column:http_type"` // http_type
}

func (receiver *TableName) TableName() string {
	return "action_logs"
}

type OrmTableName struct {
	db *gorm.DB
}

func NewOrmTableName() *OrmTableName {
	orm := &OrmTableName{}
	orm.db = providers.NewMysqlProvider().GetBean("mysql").(*gorm.DB)
	return orm
}

func (orm *OrmTableName) Connect() *gorm.DB {
	return orm.db
}

// Create insert the value into database
func (orm *OrmTableName) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmTableName) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmTableName) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

// First find first record that match given conditions, order by primary key
func (orm *OrmTableName) First(conds ...interface{}) (*TableName, *gorm.DB) {
	dest := &TableName{}
	return dest, orm.db.First(dest, conds...)
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmTableName) Take(conds ...interface{}) (*TableName, *gorm.DB) {
	dest := &TableName{}
	return dest, orm.db.Take(dest, conds...)
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmTableName) Last(conds ...interface{}) (*TableName, *gorm.DB) {
	dest := &TableName{}
	return dest, orm.db.Last(dest, conds...)
}

func (orm *OrmTableName) Find(conds ...interface{}) (*TableName, *gorm.DB) {
	dest := &TableName{}
	return dest, orm.db.Find(dest, conds...)
}

// FindInBatches find records in batches
func (orm *OrmTableName) FindInBatches(batchSize int, fc func(tx *gorm.DB, batch int) error) (*TableName, *gorm.DB) {
	dest := &TableName{}
	return dest, orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmTableName) FirstOrInit(conds ...interface{}) (*TableName, *gorm.DB) {
	dest := &TableName{}
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmTableName) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmTableName) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmTableName) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmTableName) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmTableName) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (orm *OrmTableName) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.Delete(value, conds...)
}

func (orm *OrmTableName) Count(count *int64) *gorm.DB {
	return orm.db.Count(count)
}

func (orm *OrmTableName) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmTableName) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmTableName) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (orm *OrmTableName) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

func (orm *OrmTableName) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmTableName) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmTableName) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmTableName) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmTableName) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmTableName) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmTableName) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmTableName) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmTableName) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmTableName) Insert(row *TableName) *gorm.DB {
	return orm.db.Create(row)
}

func (orm *OrmTableName) Inserts(rows []*TableName) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmTableName) Limit(limit int) *OrmTableName {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmTableName) Offset(offset int) *OrmTableName {
	orm.db.Offset(offset)
	return orm
}

//
//func (orm *OrmTableName) GetPaginate(page uint32, limit uint32) (OrmOrmTableNameList, int64) {
//	var total int64
//	var offset int
//	orm.Count(&total)
//	if total > 0 {
//		orm.Limit(int(limit))
//		if page > 1 {
//			offset = int((page - 1) * limit)
//		}
//		orm.Offset(offset)
//		return orm.Find(), total
//	}
//	// 查询不到数据时候
//	orm.list = make([]*OrmAdminGroup, 0)
//	return orm.Find(), total
//}

func (orm *OrmTableName) Get() []TableName {
	return nil
}
