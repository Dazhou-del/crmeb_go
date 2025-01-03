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

func newSystemGroup(db *gorm.DB, opts ...gen.DOOption) systemGroup {
	_systemGroup := systemGroup{}

	_systemGroup.systemGroupDo.UseDB(db, opts...)
	_systemGroup.systemGroupDo.UseModel(&model.SystemGroup{})

	tableName := _systemGroup.systemGroupDo.TableName()
	_systemGroup.ALL = field.NewAsterisk(tableName)
	_systemGroup.ID = field.NewInt64(tableName, "id")
	_systemGroup.Name = field.NewString(tableName, "name")
	_systemGroup.Info = field.NewString(tableName, "info")
	_systemGroup.FormID = field.NewInt64(tableName, "form_id")
	_systemGroup.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemGroup.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemGroup.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemGroup.fillFieldMap()

	return _systemGroup
}

// systemGroup 组合数据表
type systemGroup struct {
	systemGroupDo systemGroupDo

	ALL       field.Asterisk
	ID        field.Int64  // 组合数据ID
	Name      field.String // 数据组名称
	Info      field.String // 简介
	FormID    field.Int64  // form 表单 id
	CreatedAt field.Int64  // 创建时间
	UpdatedAt field.Int64  // 修改时间
	DeletedAt field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (s systemGroup) Table(newTableName string) *systemGroup {
	s.systemGroupDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemGroup) As(alias string) *systemGroup {
	s.systemGroupDo.DO = *(s.systemGroupDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemGroup) updateTableName(table string) *systemGroup {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Info = field.NewString(table, "info")
	s.FormID = field.NewInt64(table, "form_id")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemGroup) WithContext(ctx context.Context) ISystemGroupDo {
	return s.systemGroupDo.WithContext(ctx)
}

func (s systemGroup) TableName() string { return s.systemGroupDo.TableName() }

func (s systemGroup) Alias() string { return s.systemGroupDo.Alias() }

func (s systemGroup) Columns(cols ...field.Expr) gen.Columns { return s.systemGroupDo.Columns(cols...) }

func (s *systemGroup) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemGroup) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["info"] = s.Info
	s.fieldMap["form_id"] = s.FormID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemGroup) clone(db *gorm.DB) systemGroup {
	s.systemGroupDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemGroup) replaceDB(db *gorm.DB) systemGroup {
	s.systemGroupDo.ReplaceDB(db)
	return s
}

type systemGroupDo struct{ gen.DO }

type ISystemGroupDo interface {
	gen.SubQuery
	Debug() ISystemGroupDo
	WithContext(ctx context.Context) ISystemGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemGroupDo
	WriteDB() ISystemGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemGroupDo
	Not(conds ...gen.Condition) ISystemGroupDo
	Or(conds ...gen.Condition) ISystemGroupDo
	Select(conds ...field.Expr) ISystemGroupDo
	Where(conds ...gen.Condition) ISystemGroupDo
	Order(conds ...field.Expr) ISystemGroupDo
	Distinct(cols ...field.Expr) ISystemGroupDo
	Omit(cols ...field.Expr) ISystemGroupDo
	Join(table schema.Tabler, on ...field.Expr) ISystemGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDo
	Group(cols ...field.Expr) ISystemGroupDo
	Having(conds ...gen.Condition) ISystemGroupDo
	Limit(limit int) ISystemGroupDo
	Offset(offset int) ISystemGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemGroupDo
	Unscoped() ISystemGroupDo
	Create(values ...*model.SystemGroup) error
	CreateInBatches(values []*model.SystemGroup, batchSize int) error
	Save(values ...*model.SystemGroup) error
	First() (*model.SystemGroup, error)
	Take() (*model.SystemGroup, error)
	Last() (*model.SystemGroup, error)
	Find() ([]*model.SystemGroup, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemGroup, err error)
	FindInBatches(result *[]*model.SystemGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemGroup) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemGroupDo
	Assign(attrs ...field.AssignExpr) ISystemGroupDo
	Joins(fields ...field.RelationField) ISystemGroupDo
	Preload(fields ...field.RelationField) ISystemGroupDo
	FirstOrInit() (*model.SystemGroup, error)
	FirstOrCreate() (*model.SystemGroup, error)
	FindByPage(offset int, limit int) (result []*model.SystemGroup, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemGroupDo) Debug() ISystemGroupDo {
	return s.withDO(s.DO.Debug())
}

func (s systemGroupDo) WithContext(ctx context.Context) ISystemGroupDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemGroupDo) ReadDB() ISystemGroupDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemGroupDo) WriteDB() ISystemGroupDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemGroupDo) Session(config *gorm.Session) ISystemGroupDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemGroupDo) Clauses(conds ...clause.Expression) ISystemGroupDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemGroupDo) Returning(value interface{}, columns ...string) ISystemGroupDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemGroupDo) Not(conds ...gen.Condition) ISystemGroupDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemGroupDo) Or(conds ...gen.Condition) ISystemGroupDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemGroupDo) Select(conds ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemGroupDo) Where(conds ...gen.Condition) ISystemGroupDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemGroupDo) Order(conds ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemGroupDo) Distinct(cols ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemGroupDo) Omit(cols ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemGroupDo) Join(table schema.Tabler, on ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemGroupDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemGroupDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemGroupDo) Group(cols ...field.Expr) ISystemGroupDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemGroupDo) Having(conds ...gen.Condition) ISystemGroupDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemGroupDo) Limit(limit int) ISystemGroupDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemGroupDo) Offset(offset int) ISystemGroupDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemGroupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemGroupDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemGroupDo) Unscoped() ISystemGroupDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemGroupDo) Create(values ...*model.SystemGroup) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemGroupDo) CreateInBatches(values []*model.SystemGroup, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemGroupDo) Save(values ...*model.SystemGroup) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemGroupDo) First() (*model.SystemGroup, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroup), nil
	}
}

func (s systemGroupDo) Take() (*model.SystemGroup, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroup), nil
	}
}

func (s systemGroupDo) Last() (*model.SystemGroup, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroup), nil
	}
}

func (s systemGroupDo) Find() ([]*model.SystemGroup, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemGroup), err
}

func (s systemGroupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemGroup, err error) {
	buf := make([]*model.SystemGroup, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemGroupDo) FindInBatches(result *[]*model.SystemGroup, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemGroupDo) Attrs(attrs ...field.AssignExpr) ISystemGroupDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemGroupDo) Assign(attrs ...field.AssignExpr) ISystemGroupDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemGroupDo) Joins(fields ...field.RelationField) ISystemGroupDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemGroupDo) Preload(fields ...field.RelationField) ISystemGroupDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemGroupDo) FirstOrInit() (*model.SystemGroup, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroup), nil
	}
}

func (s systemGroupDo) FirstOrCreate() (*model.SystemGroup, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemGroup), nil
	}
}

func (s systemGroupDo) FindByPage(offset int, limit int) (result []*model.SystemGroup, count int64, err error) {
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

func (s systemGroupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemGroupDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemGroupDo) Delete(models ...*model.SystemGroup) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemGroupDo) withDO(do gen.Dao) *systemGroupDo {
	s.DO = *do.(*gen.DO)
	return s
}
