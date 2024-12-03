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

func newEbCategory(db *gorm.DB, opts ...gen.DOOption) ebCategory {
	_ebCategory := ebCategory{}

	_ebCategory.ebCategoryDo.UseDB(db, opts...)
	_ebCategory.ebCategoryDo.UseModel(&model.EbCategory{})

	tableName := _ebCategory.ebCategoryDo.TableName()
	_ebCategory.ALL = field.NewAsterisk(tableName)
	_ebCategory.ID = field.NewInt32(tableName, "id")
	_ebCategory.Pid = field.NewInt32(tableName, "pid")
	_ebCategory.Path = field.NewString(tableName, "path")
	_ebCategory.Name = field.NewString(tableName, "name")
	_ebCategory.Type = field.NewInt32(tableName, "type")
	_ebCategory.URL = field.NewString(tableName, "url")
	_ebCategory.Extra = field.NewString(tableName, "extra")
	_ebCategory.Status = field.NewBool(tableName, "status")
	_ebCategory.Sort = field.NewInt32(tableName, "sort")
	_ebCategory.CreateTime = field.NewTime(tableName, "create_time")
	_ebCategory.UpdateTime = field.NewTime(tableName, "update_time")

	_ebCategory.fillFieldMap()

	return _ebCategory
}

// ebCategory 分类表
type ebCategory struct {
	ebCategoryDo ebCategoryDo

	ALL        field.Asterisk
	ID         field.Int32
	Pid        field.Int32  // 父级ID
	Path       field.String // 路径
	Name       field.String // 分类名称
	Type       field.Int32  // 类型，1 产品分类，2 附件分类，3 文章分类， 4 设置分类， 5 菜单分类，6 配置分类， 7 秒杀配置
	URL        field.String // 地址
	Extra      field.String // 扩展字段 Jsos格式
	Status     field.Bool   // 状态, 1正常，0失效
	Sort       field.Int32  // 排序
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebCategory) Table(newTableName string) *ebCategory {
	e.ebCategoryDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebCategory) As(alias string) *ebCategory {
	e.ebCategoryDo.DO = *(e.ebCategoryDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebCategory) updateTableName(table string) *ebCategory {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Pid = field.NewInt32(table, "pid")
	e.Path = field.NewString(table, "path")
	e.Name = field.NewString(table, "name")
	e.Type = field.NewInt32(table, "type")
	e.URL = field.NewString(table, "url")
	e.Extra = field.NewString(table, "extra")
	e.Status = field.NewBool(table, "status")
	e.Sort = field.NewInt32(table, "sort")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebCategory) WithContext(ctx context.Context) IEbCategoryDo {
	return e.ebCategoryDo.WithContext(ctx)
}

func (e ebCategory) TableName() string { return e.ebCategoryDo.TableName() }

func (e ebCategory) Alias() string { return e.ebCategoryDo.Alias() }

func (e ebCategory) Columns(cols ...field.Expr) gen.Columns { return e.ebCategoryDo.Columns(cols...) }

func (e *ebCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebCategory) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 11)
	e.fieldMap["id"] = e.ID
	e.fieldMap["pid"] = e.Pid
	e.fieldMap["path"] = e.Path
	e.fieldMap["name"] = e.Name
	e.fieldMap["type"] = e.Type
	e.fieldMap["url"] = e.URL
	e.fieldMap["extra"] = e.Extra
	e.fieldMap["status"] = e.Status
	e.fieldMap["sort"] = e.Sort
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebCategory) clone(db *gorm.DB) ebCategory {
	e.ebCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebCategory) replaceDB(db *gorm.DB) ebCategory {
	e.ebCategoryDo.ReplaceDB(db)
	return e
}

type ebCategoryDo struct{ gen.DO }

type IEbCategoryDo interface {
	gen.SubQuery
	Debug() IEbCategoryDo
	WithContext(ctx context.Context) IEbCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbCategoryDo
	WriteDB() IEbCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbCategoryDo
	Not(conds ...gen.Condition) IEbCategoryDo
	Or(conds ...gen.Condition) IEbCategoryDo
	Select(conds ...field.Expr) IEbCategoryDo
	Where(conds ...gen.Condition) IEbCategoryDo
	Order(conds ...field.Expr) IEbCategoryDo
	Distinct(cols ...field.Expr) IEbCategoryDo
	Omit(cols ...field.Expr) IEbCategoryDo
	Join(table schema.Tabler, on ...field.Expr) IEbCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbCategoryDo
	Group(cols ...field.Expr) IEbCategoryDo
	Having(conds ...gen.Condition) IEbCategoryDo
	Limit(limit int) IEbCategoryDo
	Offset(offset int) IEbCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbCategoryDo
	Unscoped() IEbCategoryDo
	Create(values ...*model.EbCategory) error
	CreateInBatches(values []*model.EbCategory, batchSize int) error
	Save(values ...*model.EbCategory) error
	First() (*model.EbCategory, error)
	Take() (*model.EbCategory, error)
	Last() (*model.EbCategory, error)
	Find() ([]*model.EbCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbCategory, err error)
	FindInBatches(result *[]*model.EbCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbCategoryDo
	Assign(attrs ...field.AssignExpr) IEbCategoryDo
	Joins(fields ...field.RelationField) IEbCategoryDo
	Preload(fields ...field.RelationField) IEbCategoryDo
	FirstOrInit() (*model.EbCategory, error)
	FirstOrCreate() (*model.EbCategory, error)
	FindByPage(offset int, limit int) (result []*model.EbCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebCategoryDo) Debug() IEbCategoryDo {
	return e.withDO(e.DO.Debug())
}

func (e ebCategoryDo) WithContext(ctx context.Context) IEbCategoryDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebCategoryDo) ReadDB() IEbCategoryDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebCategoryDo) WriteDB() IEbCategoryDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebCategoryDo) Session(config *gorm.Session) IEbCategoryDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebCategoryDo) Clauses(conds ...clause.Expression) IEbCategoryDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebCategoryDo) Returning(value interface{}, columns ...string) IEbCategoryDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebCategoryDo) Not(conds ...gen.Condition) IEbCategoryDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebCategoryDo) Or(conds ...gen.Condition) IEbCategoryDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebCategoryDo) Select(conds ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebCategoryDo) Where(conds ...gen.Condition) IEbCategoryDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebCategoryDo) Order(conds ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebCategoryDo) Distinct(cols ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebCategoryDo) Omit(cols ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebCategoryDo) Join(table schema.Tabler, on ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebCategoryDo) Group(cols ...field.Expr) IEbCategoryDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebCategoryDo) Having(conds ...gen.Condition) IEbCategoryDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebCategoryDo) Limit(limit int) IEbCategoryDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebCategoryDo) Offset(offset int) IEbCategoryDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbCategoryDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebCategoryDo) Unscoped() IEbCategoryDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebCategoryDo) Create(values ...*model.EbCategory) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebCategoryDo) CreateInBatches(values []*model.EbCategory, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebCategoryDo) Save(values ...*model.EbCategory) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebCategoryDo) First() (*model.EbCategory, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbCategory), nil
	}
}

func (e ebCategoryDo) Take() (*model.EbCategory, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbCategory), nil
	}
}

func (e ebCategoryDo) Last() (*model.EbCategory, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbCategory), nil
	}
}

func (e ebCategoryDo) Find() ([]*model.EbCategory, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbCategory), err
}

func (e ebCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbCategory, err error) {
	buf := make([]*model.EbCategory, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebCategoryDo) FindInBatches(result *[]*model.EbCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebCategoryDo) Attrs(attrs ...field.AssignExpr) IEbCategoryDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebCategoryDo) Assign(attrs ...field.AssignExpr) IEbCategoryDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebCategoryDo) Joins(fields ...field.RelationField) IEbCategoryDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebCategoryDo) Preload(fields ...field.RelationField) IEbCategoryDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebCategoryDo) FirstOrInit() (*model.EbCategory, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbCategory), nil
	}
}

func (e ebCategoryDo) FirstOrCreate() (*model.EbCategory, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbCategory), nil
	}
}

func (e ebCategoryDo) FindByPage(offset int, limit int) (result []*model.EbCategory, count int64, err error) {
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

func (e ebCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebCategoryDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebCategoryDo) Delete(models ...*model.EbCategory) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebCategoryDo) withDO(do gen.Dao) *ebCategoryDo {
	e.DO = *do.(*gen.DO)
	return e
}