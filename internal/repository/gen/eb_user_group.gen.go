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

func newEbUserGroup(db *gorm.DB, opts ...gen.DOOption) ebUserGroup {
	_ebUserGroup := ebUserGroup{}

	_ebUserGroup.ebUserGroupDo.UseDB(db, opts...)
	_ebUserGroup.ebUserGroupDo.UseModel(&model.EbUserGroup{})

	tableName := _ebUserGroup.ebUserGroupDo.TableName()
	_ebUserGroup.ALL = field.NewAsterisk(tableName)
	_ebUserGroup.ID = field.NewInt32(tableName, "id")
	_ebUserGroup.GroupName = field.NewString(tableName, "group_name")

	_ebUserGroup.fillFieldMap()

	return _ebUserGroup
}

// ebUserGroup 用户分组表
type ebUserGroup struct {
	ebUserGroupDo ebUserGroupDo

	ALL       field.Asterisk
	ID        field.Int32
	GroupName field.String // 用户分组名称

	fieldMap map[string]field.Expr
}

func (e ebUserGroup) Table(newTableName string) *ebUserGroup {
	e.ebUserGroupDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebUserGroup) As(alias string) *ebUserGroup {
	e.ebUserGroupDo.DO = *(e.ebUserGroupDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebUserGroup) updateTableName(table string) *ebUserGroup {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.GroupName = field.NewString(table, "group_name")

	e.fillFieldMap()

	return e
}

func (e *ebUserGroup) WithContext(ctx context.Context) IEbUserGroupDo {
	return e.ebUserGroupDo.WithContext(ctx)
}

func (e ebUserGroup) TableName() string { return e.ebUserGroupDo.TableName() }

func (e ebUserGroup) Alias() string { return e.ebUserGroupDo.Alias() }

func (e ebUserGroup) Columns(cols ...field.Expr) gen.Columns { return e.ebUserGroupDo.Columns(cols...) }

func (e *ebUserGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebUserGroup) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 2)
	e.fieldMap["id"] = e.ID
	e.fieldMap["group_name"] = e.GroupName
}

func (e ebUserGroup) clone(db *gorm.DB) ebUserGroup {
	e.ebUserGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebUserGroup) replaceDB(db *gorm.DB) ebUserGroup {
	e.ebUserGroupDo.ReplaceDB(db)
	return e
}

type ebUserGroupDo struct{ gen.DO }

type IEbUserGroupDo interface {
	gen.SubQuery
	Debug() IEbUserGroupDo
	WithContext(ctx context.Context) IEbUserGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbUserGroupDo
	WriteDB() IEbUserGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbUserGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbUserGroupDo
	Not(conds ...gen.Condition) IEbUserGroupDo
	Or(conds ...gen.Condition) IEbUserGroupDo
	Select(conds ...field.Expr) IEbUserGroupDo
	Where(conds ...gen.Condition) IEbUserGroupDo
	Order(conds ...field.Expr) IEbUserGroupDo
	Distinct(cols ...field.Expr) IEbUserGroupDo
	Omit(cols ...field.Expr) IEbUserGroupDo
	Join(table schema.Tabler, on ...field.Expr) IEbUserGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbUserGroupDo
	Group(cols ...field.Expr) IEbUserGroupDo
	Having(conds ...gen.Condition) IEbUserGroupDo
	Limit(limit int) IEbUserGroupDo
	Offset(offset int) IEbUserGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserGroupDo
	Unscoped() IEbUserGroupDo
	Create(values ...*model.EbUserGroup) error
	CreateInBatches(values []*model.EbUserGroup, batchSize int) error
	Save(values ...*model.EbUserGroup) error
	First() (*model.EbUserGroup, error)
	Take() (*model.EbUserGroup, error)
	Last() (*model.EbUserGroup, error)
	Find() ([]*model.EbUserGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserGroup, err error)
	FindInBatches(result *[]*model.EbUserGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbUserGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbUserGroupDo
	Assign(attrs ...field.AssignExpr) IEbUserGroupDo
	Joins(fields ...field.RelationField) IEbUserGroupDo
	Preload(fields ...field.RelationField) IEbUserGroupDo
	FirstOrInit() (*model.EbUserGroup, error)
	FirstOrCreate() (*model.EbUserGroup, error)
	FindByPage(offset int, limit int) (result []*model.EbUserGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbUserGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebUserGroupDo) Debug() IEbUserGroupDo {
	return e.withDO(e.DO.Debug())
}

func (e ebUserGroupDo) WithContext(ctx context.Context) IEbUserGroupDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebUserGroupDo) ReadDB() IEbUserGroupDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebUserGroupDo) WriteDB() IEbUserGroupDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebUserGroupDo) Session(config *gorm.Session) IEbUserGroupDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebUserGroupDo) Clauses(conds ...clause.Expression) IEbUserGroupDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebUserGroupDo) Returning(value interface{}, columns ...string) IEbUserGroupDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebUserGroupDo) Not(conds ...gen.Condition) IEbUserGroupDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebUserGroupDo) Or(conds ...gen.Condition) IEbUserGroupDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebUserGroupDo) Select(conds ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebUserGroupDo) Where(conds ...gen.Condition) IEbUserGroupDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebUserGroupDo) Order(conds ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebUserGroupDo) Distinct(cols ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebUserGroupDo) Omit(cols ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebUserGroupDo) Join(table schema.Tabler, on ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebUserGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebUserGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebUserGroupDo) Group(cols ...field.Expr) IEbUserGroupDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebUserGroupDo) Having(conds ...gen.Condition) IEbUserGroupDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebUserGroupDo) Limit(limit int) IEbUserGroupDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebUserGroupDo) Offset(offset int) IEbUserGroupDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebUserGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserGroupDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebUserGroupDo) Unscoped() IEbUserGroupDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebUserGroupDo) Create(values ...*model.EbUserGroup) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebUserGroupDo) CreateInBatches(values []*model.EbUserGroup, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebUserGroupDo) Save(values ...*model.EbUserGroup) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebUserGroupDo) First() (*model.EbUserGroup, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserGroup), nil
	}
}

func (e ebUserGroupDo) Take() (*model.EbUserGroup, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserGroup), nil
	}
}

func (e ebUserGroupDo) Last() (*model.EbUserGroup, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserGroup), nil
	}
}

func (e ebUserGroupDo) Find() ([]*model.EbUserGroup, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbUserGroup), err
}

func (e ebUserGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserGroup, err error) {
	buf := make([]*model.EbUserGroup, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebUserGroupDo) FindInBatches(result *[]*model.EbUserGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebUserGroupDo) Attrs(attrs ...field.AssignExpr) IEbUserGroupDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebUserGroupDo) Assign(attrs ...field.AssignExpr) IEbUserGroupDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebUserGroupDo) Joins(fields ...field.RelationField) IEbUserGroupDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebUserGroupDo) Preload(fields ...field.RelationField) IEbUserGroupDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebUserGroupDo) FirstOrInit() (*model.EbUserGroup, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserGroup), nil
	}
}

func (e ebUserGroupDo) FirstOrCreate() (*model.EbUserGroup, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserGroup), nil
	}
}

func (e ebUserGroupDo) FindByPage(offset int, limit int) (result []*model.EbUserGroup, count int64, err error) {
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

func (e ebUserGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebUserGroupDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebUserGroupDo) Delete(models ...*model.EbUserGroup) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebUserGroupDo) withDO(do gen.Dao) *ebUserGroupDo {
	e.DO = *do.(*gen.DO)
	return e
}