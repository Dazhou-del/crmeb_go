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

func newEbSystemRoleMenu(db *gorm.DB, opts ...gen.DOOption) ebSystemRoleMenu {
	_ebSystemRoleMenu := ebSystemRoleMenu{}

	_ebSystemRoleMenu.ebSystemRoleMenuDo.UseDB(db, opts...)
	_ebSystemRoleMenu.ebSystemRoleMenuDo.UseModel(&model.EbSystemRoleMenu{})

	tableName := _ebSystemRoleMenu.ebSystemRoleMenuDo.TableName()
	_ebSystemRoleMenu.ALL = field.NewAsterisk(tableName)
	_ebSystemRoleMenu.Rid = field.NewInt32(tableName, "rid")
	_ebSystemRoleMenu.MenuID = field.NewInt32(tableName, "menu_id")

	_ebSystemRoleMenu.fillFieldMap()

	return _ebSystemRoleMenu
}

// ebSystemRoleMenu 角色菜单关联表
type ebSystemRoleMenu struct {
	ebSystemRoleMenuDo ebSystemRoleMenuDo

	ALL    field.Asterisk
	Rid    field.Int32 // 角色id
	MenuID field.Int32 // 权限id

	fieldMap map[string]field.Expr
}

func (e ebSystemRoleMenu) Table(newTableName string) *ebSystemRoleMenu {
	e.ebSystemRoleMenuDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebSystemRoleMenu) As(alias string) *ebSystemRoleMenu {
	e.ebSystemRoleMenuDo.DO = *(e.ebSystemRoleMenuDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebSystemRoleMenu) updateTableName(table string) *ebSystemRoleMenu {
	e.ALL = field.NewAsterisk(table)
	e.Rid = field.NewInt32(table, "rid")
	e.MenuID = field.NewInt32(table, "menu_id")

	e.fillFieldMap()

	return e
}

func (e *ebSystemRoleMenu) WithContext(ctx context.Context) IEbSystemRoleMenuDo {
	return e.ebSystemRoleMenuDo.WithContext(ctx)
}

func (e ebSystemRoleMenu) TableName() string { return e.ebSystemRoleMenuDo.TableName() }

func (e ebSystemRoleMenu) Alias() string { return e.ebSystemRoleMenuDo.Alias() }

func (e ebSystemRoleMenu) Columns(cols ...field.Expr) gen.Columns {
	return e.ebSystemRoleMenuDo.Columns(cols...)
}

func (e *ebSystemRoleMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebSystemRoleMenu) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 2)
	e.fieldMap["rid"] = e.Rid
	e.fieldMap["menu_id"] = e.MenuID
}

func (e ebSystemRoleMenu) clone(db *gorm.DB) ebSystemRoleMenu {
	e.ebSystemRoleMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebSystemRoleMenu) replaceDB(db *gorm.DB) ebSystemRoleMenu {
	e.ebSystemRoleMenuDo.ReplaceDB(db)
	return e
}

type ebSystemRoleMenuDo struct{ gen.DO }

type IEbSystemRoleMenuDo interface {
	gen.SubQuery
	Debug() IEbSystemRoleMenuDo
	WithContext(ctx context.Context) IEbSystemRoleMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbSystemRoleMenuDo
	WriteDB() IEbSystemRoleMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbSystemRoleMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbSystemRoleMenuDo
	Not(conds ...gen.Condition) IEbSystemRoleMenuDo
	Or(conds ...gen.Condition) IEbSystemRoleMenuDo
	Select(conds ...field.Expr) IEbSystemRoleMenuDo
	Where(conds ...gen.Condition) IEbSystemRoleMenuDo
	Order(conds ...field.Expr) IEbSystemRoleMenuDo
	Distinct(cols ...field.Expr) IEbSystemRoleMenuDo
	Omit(cols ...field.Expr) IEbSystemRoleMenuDo
	Join(table schema.Tabler, on ...field.Expr) IEbSystemRoleMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemRoleMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemRoleMenuDo
	Group(cols ...field.Expr) IEbSystemRoleMenuDo
	Having(conds ...gen.Condition) IEbSystemRoleMenuDo
	Limit(limit int) IEbSystemRoleMenuDo
	Offset(offset int) IEbSystemRoleMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemRoleMenuDo
	Unscoped() IEbSystemRoleMenuDo
	Create(values ...*model.EbSystemRoleMenu) error
	CreateInBatches(values []*model.EbSystemRoleMenu, batchSize int) error
	Save(values ...*model.EbSystemRoleMenu) error
	First() (*model.EbSystemRoleMenu, error)
	Take() (*model.EbSystemRoleMenu, error)
	Last() (*model.EbSystemRoleMenu, error)
	Find() ([]*model.EbSystemRoleMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemRoleMenu, err error)
	FindInBatches(result *[]*model.EbSystemRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbSystemRoleMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbSystemRoleMenuDo
	Assign(attrs ...field.AssignExpr) IEbSystemRoleMenuDo
	Joins(fields ...field.RelationField) IEbSystemRoleMenuDo
	Preload(fields ...field.RelationField) IEbSystemRoleMenuDo
	FirstOrInit() (*model.EbSystemRoleMenu, error)
	FirstOrCreate() (*model.EbSystemRoleMenu, error)
	FindByPage(offset int, limit int) (result []*model.EbSystemRoleMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbSystemRoleMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebSystemRoleMenuDo) Debug() IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Debug())
}

func (e ebSystemRoleMenuDo) WithContext(ctx context.Context) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebSystemRoleMenuDo) ReadDB() IEbSystemRoleMenuDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebSystemRoleMenuDo) WriteDB() IEbSystemRoleMenuDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebSystemRoleMenuDo) Session(config *gorm.Session) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebSystemRoleMenuDo) Clauses(conds ...clause.Expression) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebSystemRoleMenuDo) Returning(value interface{}, columns ...string) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebSystemRoleMenuDo) Not(conds ...gen.Condition) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebSystemRoleMenuDo) Or(conds ...gen.Condition) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebSystemRoleMenuDo) Select(conds ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebSystemRoleMenuDo) Where(conds ...gen.Condition) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebSystemRoleMenuDo) Order(conds ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebSystemRoleMenuDo) Distinct(cols ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebSystemRoleMenuDo) Omit(cols ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebSystemRoleMenuDo) Join(table schema.Tabler, on ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebSystemRoleMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebSystemRoleMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebSystemRoleMenuDo) Group(cols ...field.Expr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebSystemRoleMenuDo) Having(conds ...gen.Condition) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebSystemRoleMenuDo) Limit(limit int) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebSystemRoleMenuDo) Offset(offset int) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebSystemRoleMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebSystemRoleMenuDo) Unscoped() IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebSystemRoleMenuDo) Create(values ...*model.EbSystemRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebSystemRoleMenuDo) CreateInBatches(values []*model.EbSystemRoleMenu, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebSystemRoleMenuDo) Save(values ...*model.EbSystemRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebSystemRoleMenuDo) First() (*model.EbSystemRoleMenu, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemRoleMenu), nil
	}
}

func (e ebSystemRoleMenuDo) Take() (*model.EbSystemRoleMenu, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemRoleMenu), nil
	}
}

func (e ebSystemRoleMenuDo) Last() (*model.EbSystemRoleMenu, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemRoleMenu), nil
	}
}

func (e ebSystemRoleMenuDo) Find() ([]*model.EbSystemRoleMenu, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbSystemRoleMenu), err
}

func (e ebSystemRoleMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemRoleMenu, err error) {
	buf := make([]*model.EbSystemRoleMenu, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebSystemRoleMenuDo) FindInBatches(result *[]*model.EbSystemRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebSystemRoleMenuDo) Attrs(attrs ...field.AssignExpr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebSystemRoleMenuDo) Assign(attrs ...field.AssignExpr) IEbSystemRoleMenuDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebSystemRoleMenuDo) Joins(fields ...field.RelationField) IEbSystemRoleMenuDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebSystemRoleMenuDo) Preload(fields ...field.RelationField) IEbSystemRoleMenuDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebSystemRoleMenuDo) FirstOrInit() (*model.EbSystemRoleMenu, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemRoleMenu), nil
	}
}

func (e ebSystemRoleMenuDo) FirstOrCreate() (*model.EbSystemRoleMenu, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemRoleMenu), nil
	}
}

func (e ebSystemRoleMenuDo) FindByPage(offset int, limit int) (result []*model.EbSystemRoleMenu, count int64, err error) {
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

func (e ebSystemRoleMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebSystemRoleMenuDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebSystemRoleMenuDo) Delete(models ...*model.EbSystemRoleMenu) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebSystemRoleMenuDo) withDO(do gen.Dao) *ebSystemRoleMenuDo {
	e.DO = *do.(*gen.DO)
	return e
}
