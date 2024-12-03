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

func newEbStoreProductDescription(db *gorm.DB, opts ...gen.DOOption) ebStoreProductDescription {
	_ebStoreProductDescription := ebStoreProductDescription{}

	_ebStoreProductDescription.ebStoreProductDescriptionDo.UseDB(db, opts...)
	_ebStoreProductDescription.ebStoreProductDescriptionDo.UseModel(&model.EbStoreProductDescription{})

	tableName := _ebStoreProductDescription.ebStoreProductDescriptionDo.TableName()
	_ebStoreProductDescription.ALL = field.NewAsterisk(tableName)
	_ebStoreProductDescription.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreProductDescription.Description = field.NewString(tableName, "description")
	_ebStoreProductDescription.Type = field.NewBool(tableName, "type")
	_ebStoreProductDescription.ID = field.NewInt32(tableName, "id")

	_ebStoreProductDescription.fillFieldMap()

	return _ebStoreProductDescription
}

// ebStoreProductDescription 商品描述表
type ebStoreProductDescription struct {
	ebStoreProductDescriptionDo ebStoreProductDescriptionDo

	ALL         field.Asterisk
	ProductID   field.Int32  // 商品ID
	Description field.String // 商品详情
	Type        field.Bool   // 商品类型
	ID          field.Int32

	fieldMap map[string]field.Expr
}

func (e ebStoreProductDescription) Table(newTableName string) *ebStoreProductDescription {
	e.ebStoreProductDescriptionDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreProductDescription) As(alias string) *ebStoreProductDescription {
	e.ebStoreProductDescriptionDo.DO = *(e.ebStoreProductDescriptionDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreProductDescription) updateTableName(table string) *ebStoreProductDescription {
	e.ALL = field.NewAsterisk(table)
	e.ProductID = field.NewInt32(table, "product_id")
	e.Description = field.NewString(table, "description")
	e.Type = field.NewBool(table, "type")
	e.ID = field.NewInt32(table, "id")

	e.fillFieldMap()

	return e
}

func (e *ebStoreProductDescription) WithContext(ctx context.Context) IEbStoreProductDescriptionDo {
	return e.ebStoreProductDescriptionDo.WithContext(ctx)
}

func (e ebStoreProductDescription) TableName() string {
	return e.ebStoreProductDescriptionDo.TableName()
}

func (e ebStoreProductDescription) Alias() string { return e.ebStoreProductDescriptionDo.Alias() }

func (e ebStoreProductDescription) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreProductDescriptionDo.Columns(cols...)
}

func (e *ebStoreProductDescription) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreProductDescription) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 4)
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["description"] = e.Description
	e.fieldMap["type"] = e.Type
	e.fieldMap["id"] = e.ID
}

func (e ebStoreProductDescription) clone(db *gorm.DB) ebStoreProductDescription {
	e.ebStoreProductDescriptionDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreProductDescription) replaceDB(db *gorm.DB) ebStoreProductDescription {
	e.ebStoreProductDescriptionDo.ReplaceDB(db)
	return e
}

type ebStoreProductDescriptionDo struct{ gen.DO }

type IEbStoreProductDescriptionDo interface {
	gen.SubQuery
	Debug() IEbStoreProductDescriptionDo
	WithContext(ctx context.Context) IEbStoreProductDescriptionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreProductDescriptionDo
	WriteDB() IEbStoreProductDescriptionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreProductDescriptionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreProductDescriptionDo
	Not(conds ...gen.Condition) IEbStoreProductDescriptionDo
	Or(conds ...gen.Condition) IEbStoreProductDescriptionDo
	Select(conds ...field.Expr) IEbStoreProductDescriptionDo
	Where(conds ...gen.Condition) IEbStoreProductDescriptionDo
	Order(conds ...field.Expr) IEbStoreProductDescriptionDo
	Distinct(cols ...field.Expr) IEbStoreProductDescriptionDo
	Omit(cols ...field.Expr) IEbStoreProductDescriptionDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreProductDescriptionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductDescriptionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductDescriptionDo
	Group(cols ...field.Expr) IEbStoreProductDescriptionDo
	Having(conds ...gen.Condition) IEbStoreProductDescriptionDo
	Limit(limit int) IEbStoreProductDescriptionDo
	Offset(offset int) IEbStoreProductDescriptionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductDescriptionDo
	Unscoped() IEbStoreProductDescriptionDo
	Create(values ...*model.EbStoreProductDescription) error
	CreateInBatches(values []*model.EbStoreProductDescription, batchSize int) error
	Save(values ...*model.EbStoreProductDescription) error
	First() (*model.EbStoreProductDescription, error)
	Take() (*model.EbStoreProductDescription, error)
	Last() (*model.EbStoreProductDescription, error)
	Find() ([]*model.EbStoreProductDescription, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductDescription, err error)
	FindInBatches(result *[]*model.EbStoreProductDescription, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreProductDescription) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreProductDescriptionDo
	Assign(attrs ...field.AssignExpr) IEbStoreProductDescriptionDo
	Joins(fields ...field.RelationField) IEbStoreProductDescriptionDo
	Preload(fields ...field.RelationField) IEbStoreProductDescriptionDo
	FirstOrInit() (*model.EbStoreProductDescription, error)
	FirstOrCreate() (*model.EbStoreProductDescription, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreProductDescription, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreProductDescriptionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreProductDescriptionDo) Debug() IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreProductDescriptionDo) WithContext(ctx context.Context) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreProductDescriptionDo) ReadDB() IEbStoreProductDescriptionDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreProductDescriptionDo) WriteDB() IEbStoreProductDescriptionDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreProductDescriptionDo) Session(config *gorm.Session) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreProductDescriptionDo) Clauses(conds ...clause.Expression) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreProductDescriptionDo) Returning(value interface{}, columns ...string) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreProductDescriptionDo) Not(conds ...gen.Condition) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreProductDescriptionDo) Or(conds ...gen.Condition) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreProductDescriptionDo) Select(conds ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreProductDescriptionDo) Where(conds ...gen.Condition) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreProductDescriptionDo) Order(conds ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreProductDescriptionDo) Distinct(cols ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreProductDescriptionDo) Omit(cols ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreProductDescriptionDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreProductDescriptionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreProductDescriptionDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreProductDescriptionDo) Group(cols ...field.Expr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreProductDescriptionDo) Having(conds ...gen.Condition) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreProductDescriptionDo) Limit(limit int) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreProductDescriptionDo) Offset(offset int) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreProductDescriptionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreProductDescriptionDo) Unscoped() IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreProductDescriptionDo) Create(values ...*model.EbStoreProductDescription) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreProductDescriptionDo) CreateInBatches(values []*model.EbStoreProductDescription, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreProductDescriptionDo) Save(values ...*model.EbStoreProductDescription) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreProductDescriptionDo) First() (*model.EbStoreProductDescription, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductDescription), nil
	}
}

func (e ebStoreProductDescriptionDo) Take() (*model.EbStoreProductDescription, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductDescription), nil
	}
}

func (e ebStoreProductDescriptionDo) Last() (*model.EbStoreProductDescription, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductDescription), nil
	}
}

func (e ebStoreProductDescriptionDo) Find() ([]*model.EbStoreProductDescription, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreProductDescription), err
}

func (e ebStoreProductDescriptionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductDescription, err error) {
	buf := make([]*model.EbStoreProductDescription, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreProductDescriptionDo) FindInBatches(result *[]*model.EbStoreProductDescription, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreProductDescriptionDo) Attrs(attrs ...field.AssignExpr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreProductDescriptionDo) Assign(attrs ...field.AssignExpr) IEbStoreProductDescriptionDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreProductDescriptionDo) Joins(fields ...field.RelationField) IEbStoreProductDescriptionDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreProductDescriptionDo) Preload(fields ...field.RelationField) IEbStoreProductDescriptionDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreProductDescriptionDo) FirstOrInit() (*model.EbStoreProductDescription, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductDescription), nil
	}
}

func (e ebStoreProductDescriptionDo) FirstOrCreate() (*model.EbStoreProductDescription, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductDescription), nil
	}
}

func (e ebStoreProductDescriptionDo) FindByPage(offset int, limit int) (result []*model.EbStoreProductDescription, count int64, err error) {
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

func (e ebStoreProductDescriptionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreProductDescriptionDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreProductDescriptionDo) Delete(models ...*model.EbStoreProductDescription) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreProductDescriptionDo) withDO(do gen.Dao) *ebStoreProductDescriptionDo {
	e.DO = *do.(*gen.DO)
	return e
}
