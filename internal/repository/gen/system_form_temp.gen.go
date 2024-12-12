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

func newSystemFormTemp(db *gorm.DB, opts ...gen.DOOption) systemFormTemp {
	_systemFormTemp := systemFormTemp{}

	_systemFormTemp.systemFormTempDo.UseDB(db, opts...)
	_systemFormTemp.systemFormTempDo.UseModel(&model.SystemFormTemp{})

	tableName := _systemFormTemp.systemFormTempDo.TableName()
	_systemFormTemp.ALL = field.NewAsterisk(tableName)
	_systemFormTemp.ID = field.NewInt64(tableName, "id")
	_systemFormTemp.Name = field.NewString(tableName, "name")
	_systemFormTemp.Info = field.NewString(tableName, "info")
	_systemFormTemp.Content = field.NewString(tableName, "content")
	_systemFormTemp.CreatedAt = field.NewInt64(tableName, "created_at")
	_systemFormTemp.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_systemFormTemp.DeletedAt = field.NewField(tableName, "deleted_at")

	_systemFormTemp.fillFieldMap()

	return _systemFormTemp
}

// systemFormTemp 表单模板
type systemFormTemp struct {
	systemFormTempDo systemFormTempDo

	ALL       field.Asterisk
	ID        field.Int64  // 表单模板id
	Name      field.String // 表单名称
	Info      field.String // 表单简介
	Content   field.String // 表单内容
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s systemFormTemp) Table(newTableName string) *systemFormTemp {
	s.systemFormTempDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s systemFormTemp) As(alias string) *systemFormTemp {
	s.systemFormTempDo.DO = *(s.systemFormTempDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *systemFormTemp) updateTableName(table string) *systemFormTemp {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Info = field.NewString(table, "info")
	s.Content = field.NewString(table, "content")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *systemFormTemp) WithContext(ctx context.Context) ISystemFormTempDo {
	return s.systemFormTempDo.WithContext(ctx)
}

func (s systemFormTemp) TableName() string { return s.systemFormTempDo.TableName() }

func (s systemFormTemp) Alias() string { return s.systemFormTempDo.Alias() }

func (s systemFormTemp) Columns(cols ...field.Expr) gen.Columns {
	return s.systemFormTempDo.Columns(cols...)
}

func (s *systemFormTemp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *systemFormTemp) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["info"] = s.Info
	s.fieldMap["content"] = s.Content
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s systemFormTemp) clone(db *gorm.DB) systemFormTemp {
	s.systemFormTempDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s systemFormTemp) replaceDB(db *gorm.DB) systemFormTemp {
	s.systemFormTempDo.ReplaceDB(db)
	return s
}

type systemFormTempDo struct{ gen.DO }

type ISystemFormTempDo interface {
	gen.SubQuery
	Debug() ISystemFormTempDo
	WithContext(ctx context.Context) ISystemFormTempDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISystemFormTempDo
	WriteDB() ISystemFormTempDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISystemFormTempDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISystemFormTempDo
	Not(conds ...gen.Condition) ISystemFormTempDo
	Or(conds ...gen.Condition) ISystemFormTempDo
	Select(conds ...field.Expr) ISystemFormTempDo
	Where(conds ...gen.Condition) ISystemFormTempDo
	Order(conds ...field.Expr) ISystemFormTempDo
	Distinct(cols ...field.Expr) ISystemFormTempDo
	Omit(cols ...field.Expr) ISystemFormTempDo
	Join(table schema.Tabler, on ...field.Expr) ISystemFormTempDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISystemFormTempDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISystemFormTempDo
	Group(cols ...field.Expr) ISystemFormTempDo
	Having(conds ...gen.Condition) ISystemFormTempDo
	Limit(limit int) ISystemFormTempDo
	Offset(offset int) ISystemFormTempDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemFormTempDo
	Unscoped() ISystemFormTempDo
	Create(values ...*model.SystemFormTemp) error
	CreateInBatches(values []*model.SystemFormTemp, batchSize int) error
	Save(values ...*model.SystemFormTemp) error
	First() (*model.SystemFormTemp, error)
	Take() (*model.SystemFormTemp, error)
	Last() (*model.SystemFormTemp, error)
	Find() ([]*model.SystemFormTemp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemFormTemp, err error)
	FindInBatches(result *[]*model.SystemFormTemp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SystemFormTemp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISystemFormTempDo
	Assign(attrs ...field.AssignExpr) ISystemFormTempDo
	Joins(fields ...field.RelationField) ISystemFormTempDo
	Preload(fields ...field.RelationField) ISystemFormTempDo
	FirstOrInit() (*model.SystemFormTemp, error)
	FirstOrCreate() (*model.SystemFormTemp, error)
	FindByPage(offset int, limit int) (result []*model.SystemFormTemp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISystemFormTempDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s systemFormTempDo) Debug() ISystemFormTempDo {
	return s.withDO(s.DO.Debug())
}

func (s systemFormTempDo) WithContext(ctx context.Context) ISystemFormTempDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s systemFormTempDo) ReadDB() ISystemFormTempDo {
	return s.Clauses(dbresolver.Read)
}

func (s systemFormTempDo) WriteDB() ISystemFormTempDo {
	return s.Clauses(dbresolver.Write)
}

func (s systemFormTempDo) Session(config *gorm.Session) ISystemFormTempDo {
	return s.withDO(s.DO.Session(config))
}

func (s systemFormTempDo) Clauses(conds ...clause.Expression) ISystemFormTempDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s systemFormTempDo) Returning(value interface{}, columns ...string) ISystemFormTempDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s systemFormTempDo) Not(conds ...gen.Condition) ISystemFormTempDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s systemFormTempDo) Or(conds ...gen.Condition) ISystemFormTempDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s systemFormTempDo) Select(conds ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s systemFormTempDo) Where(conds ...gen.Condition) ISystemFormTempDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s systemFormTempDo) Order(conds ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s systemFormTempDo) Distinct(cols ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s systemFormTempDo) Omit(cols ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s systemFormTempDo) Join(table schema.Tabler, on ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s systemFormTempDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s systemFormTempDo) RightJoin(table schema.Tabler, on ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s systemFormTempDo) Group(cols ...field.Expr) ISystemFormTempDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s systemFormTempDo) Having(conds ...gen.Condition) ISystemFormTempDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s systemFormTempDo) Limit(limit int) ISystemFormTempDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s systemFormTempDo) Offset(offset int) ISystemFormTempDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s systemFormTempDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISystemFormTempDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s systemFormTempDo) Unscoped() ISystemFormTempDo {
	return s.withDO(s.DO.Unscoped())
}

func (s systemFormTempDo) Create(values ...*model.SystemFormTemp) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s systemFormTempDo) CreateInBatches(values []*model.SystemFormTemp, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s systemFormTempDo) Save(values ...*model.SystemFormTemp) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s systemFormTempDo) First() (*model.SystemFormTemp, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemFormTemp), nil
	}
}

func (s systemFormTempDo) Take() (*model.SystemFormTemp, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemFormTemp), nil
	}
}

func (s systemFormTempDo) Last() (*model.SystemFormTemp, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemFormTemp), nil
	}
}

func (s systemFormTempDo) Find() ([]*model.SystemFormTemp, error) {
	result, err := s.DO.Find()
	return result.([]*model.SystemFormTemp), err
}

func (s systemFormTempDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SystemFormTemp, err error) {
	buf := make([]*model.SystemFormTemp, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s systemFormTempDo) FindInBatches(result *[]*model.SystemFormTemp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s systemFormTempDo) Attrs(attrs ...field.AssignExpr) ISystemFormTempDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s systemFormTempDo) Assign(attrs ...field.AssignExpr) ISystemFormTempDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s systemFormTempDo) Joins(fields ...field.RelationField) ISystemFormTempDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s systemFormTempDo) Preload(fields ...field.RelationField) ISystemFormTempDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s systemFormTempDo) FirstOrInit() (*model.SystemFormTemp, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemFormTemp), nil
	}
}

func (s systemFormTempDo) FirstOrCreate() (*model.SystemFormTemp, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SystemFormTemp), nil
	}
}

func (s systemFormTempDo) FindByPage(offset int, limit int) (result []*model.SystemFormTemp, count int64, err error) {
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

func (s systemFormTempDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s systemFormTempDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s systemFormTempDo) Delete(models ...*model.SystemFormTemp) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *systemFormTempDo) withDO(do gen.Dao) *systemFormTempDo {
	s.DO = *do.(*gen.DO)
	return s
}