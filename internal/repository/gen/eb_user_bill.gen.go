// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"crmeb_go/internal/model"
)

func newEbUserBill(db *gorm.DB, opts ...gen.DOOption) ebUserBill {
	_ebUserBill := ebUserBill{}

	_ebUserBill.ebUserBillDo.UseDB(db, opts...)
	_ebUserBill.ebUserBillDo.UseModel(&model.EbUserBill{})

	tableName := _ebUserBill.ebUserBillDo.TableName()
	_ebUserBill.ALL = field.NewAsterisk(tableName)
	_ebUserBill.ID = field.NewInt32(tableName, "id")
	_ebUserBill.UID = field.NewInt32(tableName, "uid")
	_ebUserBill.LinkID = field.NewString(tableName, "link_id")
	_ebUserBill.Pm = field.NewInt32(tableName, "pm")
	_ebUserBill.Title = field.NewString(tableName, "title")
	_ebUserBill.Category = field.NewString(tableName, "category")
	_ebUserBill.Type = field.NewString(tableName, "type")
	_ebUserBill.Number = field.NewFloat64(tableName, "number")
	_ebUserBill.Balance = field.NewFloat64(tableName, "balance")
	_ebUserBill.Mark = field.NewString(tableName, "mark")
	_ebUserBill.Status = field.NewBool(tableName, "status")
	_ebUserBill.CreateTime = field.NewTime(tableName, "create_time")
	_ebUserBill.UpdateTime = field.NewTime(tableName, "update_time")

	_ebUserBill.fillFieldMap()

	return _ebUserBill
}

// ebUserBill 用户账单表
type ebUserBill struct {
	ebUserBillDo ebUserBillDo

	ALL        field.Asterisk
	ID         field.Int32   // 用户账单id
	UID        field.Int32   // 用户uid
	LinkID     field.String  // 关联id
	Pm         field.Int32   // 0 = 支出 1 = 获得
	Title      field.String  // 账单标题
	Category   field.String  // 明细种类
	Type       field.String  // 明细类型
	Number     field.Float64 // 明细数字
	Balance    field.Float64 // 剩余
	Mark       field.String  // 备注
	Status     field.Bool    // 0 = 带确定 1 = 有效 -1 = 无效
	CreateTime field.Time    // 添加时间
	UpdateTime field.Time    // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebUserBill) Table(newTableName string) *ebUserBill {
	e.ebUserBillDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebUserBill) As(alias string) *ebUserBill {
	e.ebUserBillDo.DO = *(e.ebUserBillDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebUserBill) updateTableName(table string) *ebUserBill {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.UID = field.NewInt32(table, "uid")
	e.LinkID = field.NewString(table, "link_id")
	e.Pm = field.NewInt32(table, "pm")
	e.Title = field.NewString(table, "title")
	e.Category = field.NewString(table, "category")
	e.Type = field.NewString(table, "type")
	e.Number = field.NewFloat64(table, "number")
	e.Balance = field.NewFloat64(table, "balance")
	e.Mark = field.NewString(table, "mark")
	e.Status = field.NewBool(table, "status")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebUserBill) WithContext(ctx context.Context) IEbUserBillDo {
	return e.ebUserBillDo.WithContext(ctx)
}

func (e ebUserBill) TableName() string { return e.ebUserBillDo.TableName() }

func (e ebUserBill) Alias() string { return e.ebUserBillDo.Alias() }

func (e ebUserBill) Columns(cols ...field.Expr) gen.Columns { return e.ebUserBillDo.Columns(cols...) }

func (e *ebUserBill) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebUserBill) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 13)
	e.fieldMap["id"] = e.ID
	e.fieldMap["uid"] = e.UID
	e.fieldMap["link_id"] = e.LinkID
	e.fieldMap["pm"] = e.Pm
	e.fieldMap["title"] = e.Title
	e.fieldMap["category"] = e.Category
	e.fieldMap["type"] = e.Type
	e.fieldMap["number"] = e.Number
	e.fieldMap["balance"] = e.Balance
	e.fieldMap["mark"] = e.Mark
	e.fieldMap["status"] = e.Status
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebUserBill) clone(db *gorm.DB) ebUserBill {
	e.ebUserBillDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebUserBill) replaceDB(db *gorm.DB) ebUserBill {
	e.ebUserBillDo.ReplaceDB(db)
	return e
}

type ebUserBillDo struct{ gen.DO }

type IEbUserBillDo interface {
	gen.SubQuery
	Debug() IEbUserBillDo
	WithContext(ctx context.Context) IEbUserBillDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbUserBillDo
	WriteDB() IEbUserBillDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbUserBillDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbUserBillDo
	Not(conds ...gen.Condition) IEbUserBillDo
	Or(conds ...gen.Condition) IEbUserBillDo
	Select(conds ...field.Expr) IEbUserBillDo
	Where(conds ...gen.Condition) IEbUserBillDo
	Order(conds ...field.Expr) IEbUserBillDo
	Distinct(cols ...field.Expr) IEbUserBillDo
	Omit(cols ...field.Expr) IEbUserBillDo
	Join(table schema.Tabler, on ...field.Expr) IEbUserBillDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserBillDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbUserBillDo
	Group(cols ...field.Expr) IEbUserBillDo
	Having(conds ...gen.Condition) IEbUserBillDo
	Limit(limit int) IEbUserBillDo
	Offset(offset int) IEbUserBillDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserBillDo
	Unscoped() IEbUserBillDo
	Create(values ...*model.EbUserBill) error
	CreateInBatches(values []*model.EbUserBill, batchSize int) error
	Save(values ...*model.EbUserBill) error
	First() (*model.EbUserBill, error)
	Take() (*model.EbUserBill, error)
	Last() (*model.EbUserBill, error)
	Find() ([]*model.EbUserBill, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserBill, err error)
	FindInBatches(result *[]*model.EbUserBill, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbUserBill) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbUserBillDo
	Assign(attrs ...field.AssignExpr) IEbUserBillDo
	Joins(fields ...field.RelationField) IEbUserBillDo
	Preload(fields ...field.RelationField) IEbUserBillDo
	FirstOrInit() (*model.EbUserBill, error)
	FirstOrCreate() (*model.EbUserBill, error)
	FindByPage(offset int, limit int) (result []*model.EbUserBill, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbUserBillDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebUserBillDo) Debug() IEbUserBillDo {
	return e.withDO(e.DO.Debug())
}

func (e ebUserBillDo) WithContext(ctx context.Context) IEbUserBillDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebUserBillDo) ReadDB() IEbUserBillDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebUserBillDo) WriteDB() IEbUserBillDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebUserBillDo) Session(config *gorm.Session) IEbUserBillDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebUserBillDo) Clauses(conds ...clause.Expression) IEbUserBillDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebUserBillDo) Returning(value interface{}, columns ...string) IEbUserBillDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebUserBillDo) Not(conds ...gen.Condition) IEbUserBillDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebUserBillDo) Or(conds ...gen.Condition) IEbUserBillDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebUserBillDo) Select(conds ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebUserBillDo) Where(conds ...gen.Condition) IEbUserBillDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebUserBillDo) Order(conds ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebUserBillDo) Distinct(cols ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebUserBillDo) Omit(cols ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebUserBillDo) Join(table schema.Tabler, on ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebUserBillDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebUserBillDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebUserBillDo) Group(cols ...field.Expr) IEbUserBillDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebUserBillDo) Having(conds ...gen.Condition) IEbUserBillDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebUserBillDo) Limit(limit int) IEbUserBillDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebUserBillDo) Offset(offset int) IEbUserBillDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebUserBillDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserBillDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebUserBillDo) Unscoped() IEbUserBillDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebUserBillDo) Create(values ...*model.EbUserBill) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebUserBillDo) CreateInBatches(values []*model.EbUserBill, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebUserBillDo) Save(values ...*model.EbUserBill) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebUserBillDo) First() (*model.EbUserBill, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserBill), nil
	}
}

func (e ebUserBillDo) Take() (*model.EbUserBill, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserBill), nil
	}
}

func (e ebUserBillDo) Last() (*model.EbUserBill, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserBill), nil
	}
}

func (e ebUserBillDo) Find() ([]*model.EbUserBill, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbUserBill), err
}

func (e ebUserBillDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserBill, err error) {
	buf := make([]*model.EbUserBill, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebUserBillDo) FindInBatches(result *[]*model.EbUserBill, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebUserBillDo) Attrs(attrs ...field.AssignExpr) IEbUserBillDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebUserBillDo) Assign(attrs ...field.AssignExpr) IEbUserBillDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebUserBillDo) Joins(fields ...field.RelationField) IEbUserBillDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebUserBillDo) Preload(fields ...field.RelationField) IEbUserBillDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebUserBillDo) FirstOrInit() (*model.EbUserBill, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserBill), nil
	}
}

func (e ebUserBillDo) FirstOrCreate() (*model.EbUserBill, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserBill), nil
	}
}

func (e ebUserBillDo) FindByPage(offset int, limit int) (result []*model.EbUserBill, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e ebUserBillDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebUserBillDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebUserBillDo) Delete(models ...*model.EbUserBill) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebUserBillDo) withDO(do gen.Dao) *ebUserBillDo {
	e.DO = *do.(*gen.DO)
	return e
}
