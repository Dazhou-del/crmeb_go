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

func newEbWechatReply(db *gorm.DB, opts ...gen.DOOption) ebWechatReply {
	_ebWechatReply := ebWechatReply{}

	_ebWechatReply.ebWechatReplyDo.UseDB(db, opts...)
	_ebWechatReply.ebWechatReplyDo.UseModel(&model.EbWechatReply{})

	tableName := _ebWechatReply.ebWechatReplyDo.TableName()
	_ebWechatReply.ALL = field.NewAsterisk(tableName)
	_ebWechatReply.ID = field.NewInt32(tableName, "id")
	_ebWechatReply.Keywords = field.NewString(tableName, "keywords")
	_ebWechatReply.Type = field.NewString(tableName, "type")
	_ebWechatReply.Data = field.NewString(tableName, "data")
	_ebWechatReply.Status = field.NewInt32(tableName, "status")
	_ebWechatReply.CreateTime = field.NewTime(tableName, "create_time")
	_ebWechatReply.UpdateTime = field.NewTime(tableName, "update_time")

	_ebWechatReply.fillFieldMap()

	return _ebWechatReply
}

// ebWechatReply 微信关键字回复表
type ebWechatReply struct {
	ebWechatReplyDo ebWechatReplyDo

	ALL        field.Asterisk
	ID         field.Int32  // 微信关键字回复id
	Keywords   field.String // 关键字
	Type       field.String // 回复类型
	Data       field.String // 回复数据
	Status     field.Int32  // 回复状态 0=不可用  1 =可用
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebWechatReply) Table(newTableName string) *ebWechatReply {
	e.ebWechatReplyDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebWechatReply) As(alias string) *ebWechatReply {
	e.ebWechatReplyDo.DO = *(e.ebWechatReplyDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebWechatReply) updateTableName(table string) *ebWechatReply {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Keywords = field.NewString(table, "keywords")
	e.Type = field.NewString(table, "type")
	e.Data = field.NewString(table, "data")
	e.Status = field.NewInt32(table, "status")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebWechatReply) WithContext(ctx context.Context) IEbWechatReplyDo {
	return e.ebWechatReplyDo.WithContext(ctx)
}

func (e ebWechatReply) TableName() string { return e.ebWechatReplyDo.TableName() }

func (e ebWechatReply) Alias() string { return e.ebWechatReplyDo.Alias() }

func (e ebWechatReply) Columns(cols ...field.Expr) gen.Columns {
	return e.ebWechatReplyDo.Columns(cols...)
}

func (e *ebWechatReply) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebWechatReply) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 7)
	e.fieldMap["id"] = e.ID
	e.fieldMap["keywords"] = e.Keywords
	e.fieldMap["type"] = e.Type
	e.fieldMap["data"] = e.Data
	e.fieldMap["status"] = e.Status
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebWechatReply) clone(db *gorm.DB) ebWechatReply {
	e.ebWechatReplyDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebWechatReply) replaceDB(db *gorm.DB) ebWechatReply {
	e.ebWechatReplyDo.ReplaceDB(db)
	return e
}

type ebWechatReplyDo struct{ gen.DO }

type IEbWechatReplyDo interface {
	gen.SubQuery
	Debug() IEbWechatReplyDo
	WithContext(ctx context.Context) IEbWechatReplyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbWechatReplyDo
	WriteDB() IEbWechatReplyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbWechatReplyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbWechatReplyDo
	Not(conds ...gen.Condition) IEbWechatReplyDo
	Or(conds ...gen.Condition) IEbWechatReplyDo
	Select(conds ...field.Expr) IEbWechatReplyDo
	Where(conds ...gen.Condition) IEbWechatReplyDo
	Order(conds ...field.Expr) IEbWechatReplyDo
	Distinct(cols ...field.Expr) IEbWechatReplyDo
	Omit(cols ...field.Expr) IEbWechatReplyDo
	Join(table schema.Tabler, on ...field.Expr) IEbWechatReplyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbWechatReplyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbWechatReplyDo
	Group(cols ...field.Expr) IEbWechatReplyDo
	Having(conds ...gen.Condition) IEbWechatReplyDo
	Limit(limit int) IEbWechatReplyDo
	Offset(offset int) IEbWechatReplyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbWechatReplyDo
	Unscoped() IEbWechatReplyDo
	Create(values ...*model.EbWechatReply) error
	CreateInBatches(values []*model.EbWechatReply, batchSize int) error
	Save(values ...*model.EbWechatReply) error
	First() (*model.EbWechatReply, error)
	Take() (*model.EbWechatReply, error)
	Last() (*model.EbWechatReply, error)
	Find() ([]*model.EbWechatReply, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbWechatReply, err error)
	FindInBatches(result *[]*model.EbWechatReply, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbWechatReply) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbWechatReplyDo
	Assign(attrs ...field.AssignExpr) IEbWechatReplyDo
	Joins(fields ...field.RelationField) IEbWechatReplyDo
	Preload(fields ...field.RelationField) IEbWechatReplyDo
	FirstOrInit() (*model.EbWechatReply, error)
	FirstOrCreate() (*model.EbWechatReply, error)
	FindByPage(offset int, limit int) (result []*model.EbWechatReply, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbWechatReplyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebWechatReplyDo) Debug() IEbWechatReplyDo {
	return e.withDO(e.DO.Debug())
}

func (e ebWechatReplyDo) WithContext(ctx context.Context) IEbWechatReplyDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebWechatReplyDo) ReadDB() IEbWechatReplyDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebWechatReplyDo) WriteDB() IEbWechatReplyDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebWechatReplyDo) Session(config *gorm.Session) IEbWechatReplyDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebWechatReplyDo) Clauses(conds ...clause.Expression) IEbWechatReplyDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebWechatReplyDo) Returning(value interface{}, columns ...string) IEbWechatReplyDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebWechatReplyDo) Not(conds ...gen.Condition) IEbWechatReplyDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebWechatReplyDo) Or(conds ...gen.Condition) IEbWechatReplyDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebWechatReplyDo) Select(conds ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebWechatReplyDo) Where(conds ...gen.Condition) IEbWechatReplyDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebWechatReplyDo) Order(conds ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebWechatReplyDo) Distinct(cols ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebWechatReplyDo) Omit(cols ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebWechatReplyDo) Join(table schema.Tabler, on ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebWechatReplyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebWechatReplyDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebWechatReplyDo) Group(cols ...field.Expr) IEbWechatReplyDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebWechatReplyDo) Having(conds ...gen.Condition) IEbWechatReplyDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebWechatReplyDo) Limit(limit int) IEbWechatReplyDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebWechatReplyDo) Offset(offset int) IEbWechatReplyDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebWechatReplyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbWechatReplyDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebWechatReplyDo) Unscoped() IEbWechatReplyDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebWechatReplyDo) Create(values ...*model.EbWechatReply) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebWechatReplyDo) CreateInBatches(values []*model.EbWechatReply, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebWechatReplyDo) Save(values ...*model.EbWechatReply) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebWechatReplyDo) First() (*model.EbWechatReply, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatReply), nil
	}
}

func (e ebWechatReplyDo) Take() (*model.EbWechatReply, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatReply), nil
	}
}

func (e ebWechatReplyDo) Last() (*model.EbWechatReply, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatReply), nil
	}
}

func (e ebWechatReplyDo) Find() ([]*model.EbWechatReply, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbWechatReply), err
}

func (e ebWechatReplyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbWechatReply, err error) {
	buf := make([]*model.EbWechatReply, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebWechatReplyDo) FindInBatches(result *[]*model.EbWechatReply, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebWechatReplyDo) Attrs(attrs ...field.AssignExpr) IEbWechatReplyDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebWechatReplyDo) Assign(attrs ...field.AssignExpr) IEbWechatReplyDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebWechatReplyDo) Joins(fields ...field.RelationField) IEbWechatReplyDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebWechatReplyDo) Preload(fields ...field.RelationField) IEbWechatReplyDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebWechatReplyDo) FirstOrInit() (*model.EbWechatReply, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatReply), nil
	}
}

func (e ebWechatReplyDo) FirstOrCreate() (*model.EbWechatReply, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbWechatReply), nil
	}
}

func (e ebWechatReplyDo) FindByPage(offset int, limit int) (result []*model.EbWechatReply, count int64, err error) {
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

func (e ebWechatReplyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebWechatReplyDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebWechatReplyDo) Delete(models ...*model.EbWechatReply) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebWechatReplyDo) withDO(do gen.Dao) *ebWechatReplyDo {
	e.DO = *do.(*gen.DO)
	return e
}