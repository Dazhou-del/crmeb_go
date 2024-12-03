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

func newEbUserTag(db *gorm.DB, opts ...gen.DOOption) ebUserTag {
	_ebUserTag := ebUserTag{}

	_ebUserTag.ebUserTagDo.UseDB(db, opts...)
	_ebUserTag.ebUserTagDo.UseModel(&model.EbUserTag{})

	tableName := _ebUserTag.ebUserTagDo.TableName()
	_ebUserTag.ALL = field.NewAsterisk(tableName)
	_ebUserTag.ID = field.NewInt32(tableName, "id")
	_ebUserTag.Name = field.NewString(tableName, "name")

	_ebUserTag.fillFieldMap()

	return _ebUserTag
}

// ebUserTag 标签管理
type ebUserTag struct {
	ebUserTagDo ebUserTagDo

	ALL  field.Asterisk
	ID   field.Int32
	Name field.String // 标签名称

	fieldMap map[string]field.Expr
}

func (e ebUserTag) Table(newTableName string) *ebUserTag {
	e.ebUserTagDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebUserTag) As(alias string) *ebUserTag {
	e.ebUserTagDo.DO = *(e.ebUserTagDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebUserTag) updateTableName(table string) *ebUserTag {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Name = field.NewString(table, "name")

	e.fillFieldMap()

	return e
}

func (e *ebUserTag) WithContext(ctx context.Context) IEbUserTagDo {
	return e.ebUserTagDo.WithContext(ctx)
}

func (e ebUserTag) TableName() string { return e.ebUserTagDo.TableName() }

func (e ebUserTag) Alias() string { return e.ebUserTagDo.Alias() }

func (e ebUserTag) Columns(cols ...field.Expr) gen.Columns { return e.ebUserTagDo.Columns(cols...) }

func (e *ebUserTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebUserTag) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 2)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
}

func (e ebUserTag) clone(db *gorm.DB) ebUserTag {
	e.ebUserTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebUserTag) replaceDB(db *gorm.DB) ebUserTag {
	e.ebUserTagDo.ReplaceDB(db)
	return e
}

type ebUserTagDo struct{ gen.DO }

type IEbUserTagDo interface {
	gen.SubQuery
	Debug() IEbUserTagDo
	WithContext(ctx context.Context) IEbUserTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbUserTagDo
	WriteDB() IEbUserTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbUserTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbUserTagDo
	Not(conds ...gen.Condition) IEbUserTagDo
	Or(conds ...gen.Condition) IEbUserTagDo
	Select(conds ...field.Expr) IEbUserTagDo
	Where(conds ...gen.Condition) IEbUserTagDo
	Order(conds ...field.Expr) IEbUserTagDo
	Distinct(cols ...field.Expr) IEbUserTagDo
	Omit(cols ...field.Expr) IEbUserTagDo
	Join(table schema.Tabler, on ...field.Expr) IEbUserTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbUserTagDo
	Group(cols ...field.Expr) IEbUserTagDo
	Having(conds ...gen.Condition) IEbUserTagDo
	Limit(limit int) IEbUserTagDo
	Offset(offset int) IEbUserTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserTagDo
	Unscoped() IEbUserTagDo
	Create(values ...*model.EbUserTag) error
	CreateInBatches(values []*model.EbUserTag, batchSize int) error
	Save(values ...*model.EbUserTag) error
	First() (*model.EbUserTag, error)
	Take() (*model.EbUserTag, error)
	Last() (*model.EbUserTag, error)
	Find() ([]*model.EbUserTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserTag, err error)
	FindInBatches(result *[]*model.EbUserTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbUserTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbUserTagDo
	Assign(attrs ...field.AssignExpr) IEbUserTagDo
	Joins(fields ...field.RelationField) IEbUserTagDo
	Preload(fields ...field.RelationField) IEbUserTagDo
	FirstOrInit() (*model.EbUserTag, error)
	FirstOrCreate() (*model.EbUserTag, error)
	FindByPage(offset int, limit int) (result []*model.EbUserTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbUserTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebUserTagDo) Debug() IEbUserTagDo {
	return e.withDO(e.DO.Debug())
}

func (e ebUserTagDo) WithContext(ctx context.Context) IEbUserTagDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebUserTagDo) ReadDB() IEbUserTagDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebUserTagDo) WriteDB() IEbUserTagDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebUserTagDo) Session(config *gorm.Session) IEbUserTagDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebUserTagDo) Clauses(conds ...clause.Expression) IEbUserTagDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebUserTagDo) Returning(value interface{}, columns ...string) IEbUserTagDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebUserTagDo) Not(conds ...gen.Condition) IEbUserTagDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebUserTagDo) Or(conds ...gen.Condition) IEbUserTagDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebUserTagDo) Select(conds ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebUserTagDo) Where(conds ...gen.Condition) IEbUserTagDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebUserTagDo) Order(conds ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebUserTagDo) Distinct(cols ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebUserTagDo) Omit(cols ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebUserTagDo) Join(table schema.Tabler, on ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebUserTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebUserTagDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebUserTagDo) Group(cols ...field.Expr) IEbUserTagDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebUserTagDo) Having(conds ...gen.Condition) IEbUserTagDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebUserTagDo) Limit(limit int) IEbUserTagDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebUserTagDo) Offset(offset int) IEbUserTagDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebUserTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserTagDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebUserTagDo) Unscoped() IEbUserTagDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebUserTagDo) Create(values ...*model.EbUserTag) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebUserTagDo) CreateInBatches(values []*model.EbUserTag, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebUserTagDo) Save(values ...*model.EbUserTag) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebUserTagDo) First() (*model.EbUserTag, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserTag), nil
	}
}

func (e ebUserTagDo) Take() (*model.EbUserTag, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserTag), nil
	}
}

func (e ebUserTagDo) Last() (*model.EbUserTag, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserTag), nil
	}
}

func (e ebUserTagDo) Find() ([]*model.EbUserTag, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbUserTag), err
}

func (e ebUserTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserTag, err error) {
	buf := make([]*model.EbUserTag, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebUserTagDo) FindInBatches(result *[]*model.EbUserTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebUserTagDo) Attrs(attrs ...field.AssignExpr) IEbUserTagDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebUserTagDo) Assign(attrs ...field.AssignExpr) IEbUserTagDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebUserTagDo) Joins(fields ...field.RelationField) IEbUserTagDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebUserTagDo) Preload(fields ...field.RelationField) IEbUserTagDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebUserTagDo) FirstOrInit() (*model.EbUserTag, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserTag), nil
	}
}

func (e ebUserTagDo) FirstOrCreate() (*model.EbUserTag, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserTag), nil
	}
}

func (e ebUserTagDo) FindByPage(offset int, limit int) (result []*model.EbUserTag, count int64, err error) {
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

func (e ebUserTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebUserTagDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebUserTagDo) Delete(models ...*model.EbUserTag) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebUserTagDo) withDO(do gen.Dao) *ebUserTagDo {
	e.DO = *do.(*gen.DO)
	return e
}