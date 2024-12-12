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

func newSystemRoleMenu(db *gorm.DB, opts ...gen.DOOption) systemRoleMenu {
	_systemRoleMenu := systemRoleMenu{}

	_systemRoleMenu.systemRoleMenuDo.UseDB(db, opts...)
	_systemRoleMenu.systemRoleMenuDo.UseModel(&model.SystemRoleMenu{})

	tableName := _systemRoleMenu.systemRoleMenuDo.TableName()
	_systemRoleMenu.ALL = field.NewAsterisk(tableName)
	_systemRoleMenu.Rid = field.NewInt64(tableName, "rid")
	_systemRoleMenu.MenuID = field.NewInt64(tableName, "menu_id")
	_systemRoleMenu.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemRoleMenu.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemRoleMenu.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemRoleMenu.fillFieldMap()

	return _systemRoleMenu
}

// systemRoleMenu 角色菜单关联表
type systemRoleMenu struct {
	systemRoleMenuDo systemRoleMenuDo

	ALL       field.Asterisk
	Rid       field.Int64 // 角色id
	MenuID    field.Int64 // 权限id
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s systemRoleMenu) Table(newTableName string) *systemRoleMenu {
	s.systemRoleMenuDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemRoleMenu) As(alias string) *systemRoleMenu {
	s.systemRoleMenuDo.DO = *(s.systemRoleMenuDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemRoleMenu) updateTableName(table string) *systemRoleMenu {
	s.ALL = field.NewAsterisk(table)
	s.Rid = field.NewInt64(table, "rid")
	s.MenuID = field.NewInt64(table, "menu_id")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemRoleMenu) WithContext(ctx context.Context) ISystemRoleMenuDo {
	return s.systemRoleMenuDo.WithContext(ctx)
}

func (s systemRoleMenu) TableName() string { return s.systemRoleMenuDo.TableName() }

func (s systemRoleMenu) Alias() string { return s.systemRoleMenuDo.Alias() }

func (s systemRoleMenu) Columns(cols ...field.Expr) gen.Columns {
	return s.systemRoleMenuDo.Columns(cols...)
}

func (s *systemRoleMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemRoleMenu) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["rid"] = s.Rid
	s.fieldMap["menu_id"] = s.MenuID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemRoleMenu) clone(db *gorm.DB) systemRoleMenu {
	s.systemRoleMenuDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemRoleMenu) replaceDB(db *gorm.DB) systemRoleMenu {
	s.systemRoleMenuDo.ReplaceDB(db)
	return s
}

type systemRoleMenuDo struct{ gen.DO }

type ISystemRoleMenuDo interface {
	gen.SubQuery
	Debug() ISystemRoleMenuDo
	WithContext(ctx context.Context) ISystemRoleMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemRoleMenuDo
	WriteDB() ISystemRoleMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemRoleMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemRoleMenuDo
	Not(conds ...gen.Condition) ISystemRoleMenuDo
	Or(conds ...gen.Condition) ISystemRoleMenuDo
	Select(conds ...field.Expr) ISystemRoleMenuDo
	Where(conds ...gen.Condition) ISystemRoleMenuDo
	Order(conds ...field.Expr) ISystemRoleMenuDo
	Distinct(cols ...field.Expr) ISystemRoleMenuDo
	Omit(cols ...field.Expr) ISystemRoleMenuDo
	Join(table schema.Tabler, on ...field.Expr) ISystemRoleMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemRoleMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemRoleMenuDo
	Group(cols ...field.Expr) ISystemRoleMenuDo
	Having(conds ...gen.Condition) ISystemRoleMenuDo
	Limit(limit int) ISystemRoleMenuDo
	Offset(offset int) ISystemRoleMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemRoleMenuDo
	Unscoped() ISystemRoleMenuDo
	Create(values ...*model.SystemRoleMenu) error
	CreateInBatches(values []*model.SystemRoleMenu, batchSize int) error
	Save(values ...*model.SystemRoleMenu) error
	First() (*model.SystemRoleMenu, error)
	Take() (*model.SystemRoleMenu, error)
	Last() (*model.SystemRoleMenu, error)
	Find() ([]*model.SystemRoleMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemRoleMenu, err error)
	FindInBatches(result *[]*model.SystemRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemRoleMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemRoleMenuDo
	Assign(attrs ...field.AssignExpr) ISystemRoleMenuDo
	Joins(fields ...field.RelationField) ISystemRoleMenuDo
	Preload(fields ...field.RelationField) ISystemRoleMenuDo
	FirstOrInit() (*model.SystemRoleMenu, error)
	FirstOrCreate() (*model.SystemRoleMenu, error)
	FindByPage(offset int, limit int) (result []*model.SystemRoleMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemRoleMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemRoleMenuDo) Debug() ISystemRoleMenuDo {
	return s.withDO(s.DO.Debug())
}

func (s systemRoleMenuDo) WithContext(ctx context.Context) ISystemRoleMenuDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemRoleMenuDo) ReadDB() ISystemRoleMenuDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemRoleMenuDo) WriteDB() ISystemRoleMenuDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemRoleMenuDo) Session(config *gorm.Session) ISystemRoleMenuDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemRoleMenuDo) Clauses(conds ...clause.Expression) ISystemRoleMenuDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemRoleMenuDo) Returning(value interface{}, columns ...string) ISystemRoleMenuDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemRoleMenuDo) Not(conds ...gen.Condition) ISystemRoleMenuDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemRoleMenuDo) Or(conds ...gen.Condition) ISystemRoleMenuDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemRoleMenuDo) Select(conds ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemRoleMenuDo) Where(conds ...gen.Condition) ISystemRoleMenuDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemRoleMenuDo) Order(conds ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemRoleMenuDo) Distinct(cols ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemRoleMenuDo) Omit(cols ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemRoleMenuDo) Join(table schema.Tabler, on ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemRoleMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemRoleMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemRoleMenuDo) Group(cols ...field.Expr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemRoleMenuDo) Having(conds ...gen.Condition) ISystemRoleMenuDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemRoleMenuDo) Limit(limit int) ISystemRoleMenuDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemRoleMenuDo) Offset(offset int) ISystemRoleMenuDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemRoleMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemRoleMenuDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemRoleMenuDo) Unscoped() ISystemRoleMenuDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemRoleMenuDo) Create(values ...*model.SystemRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemRoleMenuDo) CreateInBatches(values []*model.SystemRoleMenu, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemRoleMenuDo) Save(values ...*model.SystemRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemRoleMenuDo) First() (*model.SystemRoleMenu, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemRoleMenu), nil
	}
}

func (s systemRoleMenuDo) Take() (*model.SystemRoleMenu, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemRoleMenu), nil
	}
}

func (s systemRoleMenuDo) Last() (*model.SystemRoleMenu, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemRoleMenu), nil
	}
}

func (s systemRoleMenuDo) Find() ([]*model.SystemRoleMenu, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemRoleMenu), err
}

func (s systemRoleMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemRoleMenu, err error) {
	buf := make([]*model.SystemRoleMenu, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemRoleMenuDo) FindInBatches(result *[]*model.SystemRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemRoleMenuDo) Attrs(attrs ...field.AssignExpr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemRoleMenuDo) Assign(attrs ...field.AssignExpr) ISystemRoleMenuDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemRoleMenuDo) Joins(fields ...field.RelationField) ISystemRoleMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemRoleMenuDo) Preload(fields ...field.RelationField) ISystemRoleMenuDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemRoleMenuDo) FirstOrInit() (*model.SystemRoleMenu, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemRoleMenu), nil
	}
}

func (s systemRoleMenuDo) FirstOrCreate() (*model.SystemRoleMenu, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemRoleMenu), nil
	}
}

func (s systemRoleMenuDo) FindByPage(offset int, limit int) (result []*model.SystemRoleMenu, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s systemRoleMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemRoleMenuDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemRoleMenuDo) Delete(models ...*model.SystemRoleMenu) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemRoleMenuDo) withDO(do gen.Dao) *systemRoleMenuDo {
	s.DO = *do.(*gen.DO)
	return s
}