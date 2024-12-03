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

func newEbSystemConfig(db *gorm.DB, opts ...gen.DOOption) ebSystemConfig {
	_ebSystemConfig := ebSystemConfig{}

	_ebSystemConfig.ebSystemConfigDo.UseDB(db, opts...)
	_ebSystemConfig.ebSystemConfigDo.UseModel(&model.EbSystemConfig{})

	tableName := _ebSystemConfig.ebSystemConfigDo.TableName()
	_ebSystemConfig.ALL = field.NewAsterisk(tableName)
	_ebSystemConfig.ID = field.NewInt32(tableName, "id")
	_ebSystemConfig.Name = field.NewString(tableName, "name")
	_ebSystemConfig.Title = field.NewString(tableName, "title")
	_ebSystemConfig.FormID = field.NewInt32(tableName, "form_id")
	_ebSystemConfig.Value = field.NewString(tableName, "value")
	_ebSystemConfig.Status = field.NewBool(tableName, "status")
	_ebSystemConfig.CreateTime = field.NewTime(tableName, "create_time")
	_ebSystemConfig.UpdateTime = field.NewTime(tableName, "update_time")

	_ebSystemConfig.fillFieldMap()

	return _ebSystemConfig
}

// ebSystemConfig 配置表
type ebSystemConfig struct {
	ebSystemConfigDo ebSystemConfigDo

	ALL        field.Asterisk
	ID         field.Int32  // 配置id
	Name       field.String // 字段名称
	Title      field.String // 字段提示文字
	FormID     field.Int32  // 表单id
	Value      field.String // 值
	Status     field.Bool   // 是否隐藏
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebSystemConfig) Table(newTableName string) *ebSystemConfig {
	e.ebSystemConfigDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebSystemConfig) As(alias string) *ebSystemConfig {
	e.ebSystemConfigDo.DO = *(e.ebSystemConfigDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebSystemConfig) updateTableName(table string) *ebSystemConfig {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Name = field.NewString(table, "name")
	e.Title = field.NewString(table, "title")
	e.FormID = field.NewInt32(table, "form_id")
	e.Value = field.NewString(table, "value")
	e.Status = field.NewBool(table, "status")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebSystemConfig) WithContext(ctx context.Context) IEbSystemConfigDo {
	return e.ebSystemConfigDo.WithContext(ctx)
}

func (e ebSystemConfig) TableName() string { return e.ebSystemConfigDo.TableName() }

func (e ebSystemConfig) Alias() string { return e.ebSystemConfigDo.Alias() }

func (e ebSystemConfig) Columns(cols ...field.Expr) gen.Columns {
	return e.ebSystemConfigDo.Columns(cols...)
}

func (e *ebSystemConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebSystemConfig) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 8)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
	e.fieldMap["title"] = e.Title
	e.fieldMap["form_id"] = e.FormID
	e.fieldMap["value"] = e.Value
	e.fieldMap["status"] = e.Status
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebSystemConfig) clone(db *gorm.DB) ebSystemConfig {
	e.ebSystemConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebSystemConfig) replaceDB(db *gorm.DB) ebSystemConfig {
	e.ebSystemConfigDo.ReplaceDB(db)
	return e
}

type ebSystemConfigDo struct{ gen.DO }

type IEbSystemConfigDo interface {
	gen.SubQuery
	Debug() IEbSystemConfigDo
	WithContext(ctx context.Context) IEbSystemConfigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbSystemConfigDo
	WriteDB() IEbSystemConfigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbSystemConfigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbSystemConfigDo
	Not(conds ...gen.Condition) IEbSystemConfigDo
	Or(conds ...gen.Condition) IEbSystemConfigDo
	Select(conds ...field.Expr) IEbSystemConfigDo
	Where(conds ...gen.Condition) IEbSystemConfigDo
	Order(conds ...field.Expr) IEbSystemConfigDo
	Distinct(cols ...field.Expr) IEbSystemConfigDo
	Omit(cols ...field.Expr) IEbSystemConfigDo
	Join(table schema.Tabler, on ...field.Expr) IEbSystemConfigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemConfigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemConfigDo
	Group(cols ...field.Expr) IEbSystemConfigDo
	Having(conds ...gen.Condition) IEbSystemConfigDo
	Limit(limit int) IEbSystemConfigDo
	Offset(offset int) IEbSystemConfigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemConfigDo
	Unscoped() IEbSystemConfigDo
	Create(values ...*model.EbSystemConfig) error
	CreateInBatches(values []*model.EbSystemConfig, batchSize int) error
	Save(values ...*model.EbSystemConfig) error
	First() (*model.EbSystemConfig, error)
	Take() (*model.EbSystemConfig, error)
	Last() (*model.EbSystemConfig, error)
	Find() ([]*model.EbSystemConfig, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemConfig, err error)
	FindInBatches(result *[]*model.EbSystemConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbSystemConfig) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbSystemConfigDo
	Assign(attrs ...field.AssignExpr) IEbSystemConfigDo
	Joins(fields ...field.RelationField) IEbSystemConfigDo
	Preload(fields ...field.RelationField) IEbSystemConfigDo
	FirstOrInit() (*model.EbSystemConfig, error)
	FirstOrCreate() (*model.EbSystemConfig, error)
	FindByPage(offset int, limit int) (result []*model.EbSystemConfig, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbSystemConfigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebSystemConfigDo) Debug() IEbSystemConfigDo {
	return e.withDO(e.DO.Debug())
}

func (e ebSystemConfigDo) WithContext(ctx context.Context) IEbSystemConfigDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebSystemConfigDo) ReadDB() IEbSystemConfigDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebSystemConfigDo) WriteDB() IEbSystemConfigDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebSystemConfigDo) Session(config *gorm.Session) IEbSystemConfigDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebSystemConfigDo) Clauses(conds ...clause.Expression) IEbSystemConfigDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebSystemConfigDo) Returning(value interface{}, columns ...string) IEbSystemConfigDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebSystemConfigDo) Not(conds ...gen.Condition) IEbSystemConfigDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebSystemConfigDo) Or(conds ...gen.Condition) IEbSystemConfigDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebSystemConfigDo) Select(conds ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebSystemConfigDo) Where(conds ...gen.Condition) IEbSystemConfigDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebSystemConfigDo) Order(conds ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebSystemConfigDo) Distinct(cols ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebSystemConfigDo) Omit(cols ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebSystemConfigDo) Join(table schema.Tabler, on ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebSystemConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebSystemConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebSystemConfigDo) Group(cols ...field.Expr) IEbSystemConfigDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebSystemConfigDo) Having(conds ...gen.Condition) IEbSystemConfigDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebSystemConfigDo) Limit(limit int) IEbSystemConfigDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebSystemConfigDo) Offset(offset int) IEbSystemConfigDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebSystemConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemConfigDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebSystemConfigDo) Unscoped() IEbSystemConfigDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebSystemConfigDo) Create(values ...*model.EbSystemConfig) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebSystemConfigDo) CreateInBatches(values []*model.EbSystemConfig, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebSystemConfigDo) Save(values ...*model.EbSystemConfig) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebSystemConfigDo) First() (*model.EbSystemConfig, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemConfig), nil
	}
}

func (e ebSystemConfigDo) Take() (*model.EbSystemConfig, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemConfig), nil
	}
}

func (e ebSystemConfigDo) Last() (*model.EbSystemConfig, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemConfig), nil
	}
}

func (e ebSystemConfigDo) Find() ([]*model.EbSystemConfig, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbSystemConfig), err
}

func (e ebSystemConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemConfig, err error) {
	buf := make([]*model.EbSystemConfig, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebSystemConfigDo) FindInBatches(result *[]*model.EbSystemConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebSystemConfigDo) Attrs(attrs ...field.AssignExpr) IEbSystemConfigDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebSystemConfigDo) Assign(attrs ...field.AssignExpr) IEbSystemConfigDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebSystemConfigDo) Joins(fields ...field.RelationField) IEbSystemConfigDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebSystemConfigDo) Preload(fields ...field.RelationField) IEbSystemConfigDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebSystemConfigDo) FirstOrInit() (*model.EbSystemConfig, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemConfig), nil
	}
}

func (e ebSystemConfigDo) FirstOrCreate() (*model.EbSystemConfig, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemConfig), nil
	}
}

func (e ebSystemConfigDo) FindByPage(offset int, limit int) (result []*model.EbSystemConfig, count int64, err error) {
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

func (e ebSystemConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebSystemConfigDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebSystemConfigDo) Delete(models ...*model.EbSystemConfig) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebSystemConfigDo) withDO(do gen.Dao) *ebSystemConfigDo {
	e.DO = *do.(*gen.DO)
	return e
}
