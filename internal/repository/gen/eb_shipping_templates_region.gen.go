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

func newEbShippingTemplatesRegion(db *gorm.DB, opts ...gen.DOOption) ebShippingTemplatesRegion {
	_ebShippingTemplatesRegion := ebShippingTemplatesRegion{}

	_ebShippingTemplatesRegion.ebShippingTemplatesRegionDo.UseDB(db, opts...)
	_ebShippingTemplatesRegion.ebShippingTemplatesRegionDo.UseModel(&model.EbShippingTemplatesRegion{})

	tableName := _ebShippingTemplatesRegion.ebShippingTemplatesRegionDo.TableName()
	_ebShippingTemplatesRegion.ALL = field.NewAsterisk(tableName)
	_ebShippingTemplatesRegion.ID = field.NewInt32(tableName, "id")
	_ebShippingTemplatesRegion.TempID = field.NewInt32(tableName, "temp_id")
	_ebShippingTemplatesRegion.CityID = field.NewInt32(tableName, "city_id")
	_ebShippingTemplatesRegion.Title = field.NewString(tableName, "title")
	_ebShippingTemplatesRegion.First = field.NewFloat64(tableName, "first")
	_ebShippingTemplatesRegion.FirstPrice = field.NewFloat64(tableName, "first_price")
	_ebShippingTemplatesRegion.Renewal = field.NewFloat64(tableName, "renewal")
	_ebShippingTemplatesRegion.RenewalPrice = field.NewFloat64(tableName, "renewal_price")
	_ebShippingTemplatesRegion.Type = field.NewBool(tableName, "type")
	_ebShippingTemplatesRegion.Uniqid = field.NewString(tableName, "uniqid")
	_ebShippingTemplatesRegion.Status = field.NewBool(tableName, "status")
	_ebShippingTemplatesRegion.CreateTime = field.NewTime(tableName, "create_time")
	_ebShippingTemplatesRegion.UpdateTime = field.NewTime(tableName, "update_time")

	_ebShippingTemplatesRegion.fillFieldMap()

	return _ebShippingTemplatesRegion
}

// ebShippingTemplatesRegion 运费模板指定区域费用
type ebShippingTemplatesRegion struct {
	ebShippingTemplatesRegionDo ebShippingTemplatesRegionDo

	ALL          field.Asterisk
	ID           field.Int32   // 编号
	TempID       field.Int32   // 模板ID
	CityID       field.Int32   // 城市ID
	Title        field.String  // 描述
	First        field.Float64 // 首件
	FirstPrice   field.Float64 // 首件运费
	Renewal      field.Float64 // 续件
	RenewalPrice field.Float64 // 续件运费
	Type         field.Bool    // 计费方式 1按件数 2按重量 3按体积
	Uniqid       field.String  // 分组唯一值
	Status       field.Bool    // 是否无效
	CreateTime   field.Time    // 创建时间
	UpdateTime   field.Time    // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebShippingTemplatesRegion) Table(newTableName string) *ebShippingTemplatesRegion {
	e.ebShippingTemplatesRegionDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebShippingTemplatesRegion) As(alias string) *ebShippingTemplatesRegion {
	e.ebShippingTemplatesRegionDo.DO = *(e.ebShippingTemplatesRegionDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebShippingTemplatesRegion) updateTableName(table string) *ebShippingTemplatesRegion {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.TempID = field.NewInt32(table, "temp_id")
	e.CityID = field.NewInt32(table, "city_id")
	e.Title = field.NewString(table, "title")
	e.First = field.NewFloat64(table, "first")
	e.FirstPrice = field.NewFloat64(table, "first_price")
	e.Renewal = field.NewFloat64(table, "renewal")
	e.RenewalPrice = field.NewFloat64(table, "renewal_price")
	e.Type = field.NewBool(table, "type")
	e.Uniqid = field.NewString(table, "uniqid")
	e.Status = field.NewBool(table, "status")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebShippingTemplatesRegion) WithContext(ctx context.Context) IEbShippingTemplatesRegionDo {
	return e.ebShippingTemplatesRegionDo.WithContext(ctx)
}

func (e ebShippingTemplatesRegion) TableName() string {
	return e.ebShippingTemplatesRegionDo.TableName()
}

func (e ebShippingTemplatesRegion) Alias() string { return e.ebShippingTemplatesRegionDo.Alias() }

func (e ebShippingTemplatesRegion) Columns(cols ...field.Expr) gen.Columns {
	return e.ebShippingTemplatesRegionDo.Columns(cols...)
}

func (e *ebShippingTemplatesRegion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebShippingTemplatesRegion) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 13)
	e.fieldMap["id"] = e.ID
	e.fieldMap["temp_id"] = e.TempID
	e.fieldMap["city_id"] = e.CityID
	e.fieldMap["title"] = e.Title
	e.fieldMap["first"] = e.First
	e.fieldMap["first_price"] = e.FirstPrice
	e.fieldMap["renewal"] = e.Renewal
	e.fieldMap["renewal_price"] = e.RenewalPrice
	e.fieldMap["type"] = e.Type
	e.fieldMap["uniqid"] = e.Uniqid
	e.fieldMap["status"] = e.Status
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebShippingTemplatesRegion) clone(db *gorm.DB) ebShippingTemplatesRegion {
	e.ebShippingTemplatesRegionDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebShippingTemplatesRegion) replaceDB(db *gorm.DB) ebShippingTemplatesRegion {
	e.ebShippingTemplatesRegionDo.ReplaceDB(db)
	return e
}

type ebShippingTemplatesRegionDo struct{ gen.DO }

type IEbShippingTemplatesRegionDo interface {
	gen.SubQuery
	Debug() IEbShippingTemplatesRegionDo
	WithContext(ctx context.Context) IEbShippingTemplatesRegionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbShippingTemplatesRegionDo
	WriteDB() IEbShippingTemplatesRegionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbShippingTemplatesRegionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbShippingTemplatesRegionDo
	Not(conds ...gen.Condition) IEbShippingTemplatesRegionDo
	Or(conds ...gen.Condition) IEbShippingTemplatesRegionDo
	Select(conds ...field.Expr) IEbShippingTemplatesRegionDo
	Where(conds ...gen.Condition) IEbShippingTemplatesRegionDo
	Order(conds ...field.Expr) IEbShippingTemplatesRegionDo
	Distinct(cols ...field.Expr) IEbShippingTemplatesRegionDo
	Omit(cols ...field.Expr) IEbShippingTemplatesRegionDo
	Join(table schema.Tabler, on ...field.Expr) IEbShippingTemplatesRegionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbShippingTemplatesRegionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbShippingTemplatesRegionDo
	Group(cols ...field.Expr) IEbShippingTemplatesRegionDo
	Having(conds ...gen.Condition) IEbShippingTemplatesRegionDo
	Limit(limit int) IEbShippingTemplatesRegionDo
	Offset(offset int) IEbShippingTemplatesRegionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbShippingTemplatesRegionDo
	Unscoped() IEbShippingTemplatesRegionDo
	Create(values ...*model.EbShippingTemplatesRegion) error
	CreateInBatches(values []*model.EbShippingTemplatesRegion, batchSize int) error
	Save(values ...*model.EbShippingTemplatesRegion) error
	First() (*model.EbShippingTemplatesRegion, error)
	Take() (*model.EbShippingTemplatesRegion, error)
	Last() (*model.EbShippingTemplatesRegion, error)
	Find() ([]*model.EbShippingTemplatesRegion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbShippingTemplatesRegion, err error)
	FindInBatches(result *[]*model.EbShippingTemplatesRegion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbShippingTemplatesRegion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbShippingTemplatesRegionDo
	Assign(attrs ...field.AssignExpr) IEbShippingTemplatesRegionDo
	Joins(fields ...field.RelationField) IEbShippingTemplatesRegionDo
	Preload(fields ...field.RelationField) IEbShippingTemplatesRegionDo
	FirstOrInit() (*model.EbShippingTemplatesRegion, error)
	FirstOrCreate() (*model.EbShippingTemplatesRegion, error)
	FindByPage(offset int, limit int) (result []*model.EbShippingTemplatesRegion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbShippingTemplatesRegionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebShippingTemplatesRegionDo) Debug() IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Debug())
}

func (e ebShippingTemplatesRegionDo) WithContext(ctx context.Context) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebShippingTemplatesRegionDo) ReadDB() IEbShippingTemplatesRegionDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebShippingTemplatesRegionDo) WriteDB() IEbShippingTemplatesRegionDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebShippingTemplatesRegionDo) Session(config *gorm.Session) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebShippingTemplatesRegionDo) Clauses(conds ...clause.Expression) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebShippingTemplatesRegionDo) Returning(value interface{}, columns ...string) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebShippingTemplatesRegionDo) Not(conds ...gen.Condition) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebShippingTemplatesRegionDo) Or(conds ...gen.Condition) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebShippingTemplatesRegionDo) Select(conds ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebShippingTemplatesRegionDo) Where(conds ...gen.Condition) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebShippingTemplatesRegionDo) Order(conds ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebShippingTemplatesRegionDo) Distinct(cols ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebShippingTemplatesRegionDo) Omit(cols ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebShippingTemplatesRegionDo) Join(table schema.Tabler, on ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebShippingTemplatesRegionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebShippingTemplatesRegionDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebShippingTemplatesRegionDo) Group(cols ...field.Expr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebShippingTemplatesRegionDo) Having(conds ...gen.Condition) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebShippingTemplatesRegionDo) Limit(limit int) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebShippingTemplatesRegionDo) Offset(offset int) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebShippingTemplatesRegionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebShippingTemplatesRegionDo) Unscoped() IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebShippingTemplatesRegionDo) Create(values ...*model.EbShippingTemplatesRegion) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebShippingTemplatesRegionDo) CreateInBatches(values []*model.EbShippingTemplatesRegion, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebShippingTemplatesRegionDo) Save(values ...*model.EbShippingTemplatesRegion) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebShippingTemplatesRegionDo) First() (*model.EbShippingTemplatesRegion, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbShippingTemplatesRegion), nil
	}
}

func (e ebShippingTemplatesRegionDo) Take() (*model.EbShippingTemplatesRegion, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbShippingTemplatesRegion), nil
	}
}

func (e ebShippingTemplatesRegionDo) Last() (*model.EbShippingTemplatesRegion, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbShippingTemplatesRegion), nil
	}
}

func (e ebShippingTemplatesRegionDo) Find() ([]*model.EbShippingTemplatesRegion, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbShippingTemplatesRegion), err
}

func (e ebShippingTemplatesRegionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbShippingTemplatesRegion, err error) {
	buf := make([]*model.EbShippingTemplatesRegion, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebShippingTemplatesRegionDo) FindInBatches(result *[]*model.EbShippingTemplatesRegion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebShippingTemplatesRegionDo) Attrs(attrs ...field.AssignExpr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebShippingTemplatesRegionDo) Assign(attrs ...field.AssignExpr) IEbShippingTemplatesRegionDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebShippingTemplatesRegionDo) Joins(fields ...field.RelationField) IEbShippingTemplatesRegionDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebShippingTemplatesRegionDo) Preload(fields ...field.RelationField) IEbShippingTemplatesRegionDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebShippingTemplatesRegionDo) FirstOrInit() (*model.EbShippingTemplatesRegion, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbShippingTemplatesRegion), nil
	}
}

func (e ebShippingTemplatesRegionDo) FirstOrCreate() (*model.EbShippingTemplatesRegion, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbShippingTemplatesRegion), nil
	}
}

func (e ebShippingTemplatesRegionDo) FindByPage(offset int, limit int) (result []*model.EbShippingTemplatesRegion, count int64, err error) {
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

func (e ebShippingTemplatesRegionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebShippingTemplatesRegionDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebShippingTemplatesRegionDo) Delete(models ...*model.EbShippingTemplatesRegion) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebShippingTemplatesRegionDo) withDO(do gen.Dao) *ebShippingTemplatesRegionDo {
	e.DO = *do.(*gen.DO)
	return e
}
