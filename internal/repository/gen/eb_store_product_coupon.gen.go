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

func newEbStoreProductCoupon(db *gorm.DB, opts ...gen.DOOption) ebStoreProductCoupon {
	_ebStoreProductCoupon := ebStoreProductCoupon{}

	_ebStoreProductCoupon.ebStoreProductCouponDo.UseDB(db, opts...)
	_ebStoreProductCoupon.ebStoreProductCouponDo.UseModel(&model.EbStoreProductCoupon{})

	tableName := _ebStoreProductCoupon.ebStoreProductCouponDo.TableName()
	_ebStoreProductCoupon.ALL = field.NewAsterisk(tableName)
	_ebStoreProductCoupon.ID = field.NewInt32(tableName, "id")
	_ebStoreProductCoupon.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreProductCoupon.IssueCouponID = field.NewInt32(tableName, "issue_coupon_id")
	_ebStoreProductCoupon.AddTime = field.NewInt32(tableName, "add_time")

	_ebStoreProductCoupon.fillFieldMap()

	return _ebStoreProductCoupon
}

// ebStoreProductCoupon 商品优惠券表
type ebStoreProductCoupon struct {
	ebStoreProductCouponDo ebStoreProductCouponDo

	ALL           field.Asterisk
	ID            field.Int32
	ProductID     field.Int32 // 商品id
	IssueCouponID field.Int32 // 优惠劵id
	AddTime       field.Int32 // 添加时间

	fieldMap map[string]field.Expr
}

func (e ebStoreProductCoupon) Table(newTableName string) *ebStoreProductCoupon {
	e.ebStoreProductCouponDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreProductCoupon) As(alias string) *ebStoreProductCoupon {
	e.ebStoreProductCouponDo.DO = *(e.ebStoreProductCouponDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreProductCoupon) updateTableName(table string) *ebStoreProductCoupon {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.ProductID = field.NewInt32(table, "product_id")
	e.IssueCouponID = field.NewInt32(table, "issue_coupon_id")
	e.AddTime = field.NewInt32(table, "add_time")

	e.fillFieldMap()

	return e
}

func (e *ebStoreProductCoupon) WithContext(ctx context.Context) IEbStoreProductCouponDo {
	return e.ebStoreProductCouponDo.WithContext(ctx)
}

func (e ebStoreProductCoupon) TableName() string { return e.ebStoreProductCouponDo.TableName() }

func (e ebStoreProductCoupon) Alias() string { return e.ebStoreProductCouponDo.Alias() }

func (e ebStoreProductCoupon) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreProductCouponDo.Columns(cols...)
}

func (e *ebStoreProductCoupon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreProductCoupon) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 4)
	e.fieldMap["id"] = e.ID
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["issue_coupon_id"] = e.IssueCouponID
	e.fieldMap["add_time"] = e.AddTime
}

func (e ebStoreProductCoupon) clone(db *gorm.DB) ebStoreProductCoupon {
	e.ebStoreProductCouponDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreProductCoupon) replaceDB(db *gorm.DB) ebStoreProductCoupon {
	e.ebStoreProductCouponDo.ReplaceDB(db)
	return e
}

type ebStoreProductCouponDo struct{ gen.DO }

type IEbStoreProductCouponDo interface {
	gen.SubQuery
	Debug() IEbStoreProductCouponDo
	WithContext(ctx context.Context) IEbStoreProductCouponDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreProductCouponDo
	WriteDB() IEbStoreProductCouponDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreProductCouponDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreProductCouponDo
	Not(conds ...gen.Condition) IEbStoreProductCouponDo
	Or(conds ...gen.Condition) IEbStoreProductCouponDo
	Select(conds ...field.Expr) IEbStoreProductCouponDo
	Where(conds ...gen.Condition) IEbStoreProductCouponDo
	Order(conds ...field.Expr) IEbStoreProductCouponDo
	Distinct(cols ...field.Expr) IEbStoreProductCouponDo
	Omit(cols ...field.Expr) IEbStoreProductCouponDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreProductCouponDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCouponDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCouponDo
	Group(cols ...field.Expr) IEbStoreProductCouponDo
	Having(conds ...gen.Condition) IEbStoreProductCouponDo
	Limit(limit int) IEbStoreProductCouponDo
	Offset(offset int) IEbStoreProductCouponDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductCouponDo
	Unscoped() IEbStoreProductCouponDo
	Create(values ...*model.EbStoreProductCoupon) error
	CreateInBatches(values []*model.EbStoreProductCoupon, batchSize int) error
	Save(values ...*model.EbStoreProductCoupon) error
	First() (*model.EbStoreProductCoupon, error)
	Take() (*model.EbStoreProductCoupon, error)
	Last() (*model.EbStoreProductCoupon, error)
	Find() ([]*model.EbStoreProductCoupon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductCoupon, err error)
	FindInBatches(result *[]*model.EbStoreProductCoupon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreProductCoupon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreProductCouponDo
	Assign(attrs ...field.AssignExpr) IEbStoreProductCouponDo
	Joins(fields ...field.RelationField) IEbStoreProductCouponDo
	Preload(fields ...field.RelationField) IEbStoreProductCouponDo
	FirstOrInit() (*model.EbStoreProductCoupon, error)
	FirstOrCreate() (*model.EbStoreProductCoupon, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreProductCoupon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreProductCouponDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreProductCouponDo) Debug() IEbStoreProductCouponDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreProductCouponDo) WithContext(ctx context.Context) IEbStoreProductCouponDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreProductCouponDo) ReadDB() IEbStoreProductCouponDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreProductCouponDo) WriteDB() IEbStoreProductCouponDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreProductCouponDo) Session(config *gorm.Session) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreProductCouponDo) Clauses(conds ...clause.Expression) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreProductCouponDo) Returning(value interface{}, columns ...string) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreProductCouponDo) Not(conds ...gen.Condition) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreProductCouponDo) Or(conds ...gen.Condition) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreProductCouponDo) Select(conds ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreProductCouponDo) Where(conds ...gen.Condition) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreProductCouponDo) Order(conds ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreProductCouponDo) Distinct(cols ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreProductCouponDo) Omit(cols ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreProductCouponDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreProductCouponDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreProductCouponDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreProductCouponDo) Group(cols ...field.Expr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreProductCouponDo) Having(conds ...gen.Condition) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreProductCouponDo) Limit(limit int) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreProductCouponDo) Offset(offset int) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreProductCouponDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreProductCouponDo) Unscoped() IEbStoreProductCouponDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreProductCouponDo) Create(values ...*model.EbStoreProductCoupon) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreProductCouponDo) CreateInBatches(values []*model.EbStoreProductCoupon, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreProductCouponDo) Save(values ...*model.EbStoreProductCoupon) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreProductCouponDo) First() (*model.EbStoreProductCoupon, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCoupon), nil
	}
}

func (e ebStoreProductCouponDo) Take() (*model.EbStoreProductCoupon, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCoupon), nil
	}
}

func (e ebStoreProductCouponDo) Last() (*model.EbStoreProductCoupon, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCoupon), nil
	}
}

func (e ebStoreProductCouponDo) Find() ([]*model.EbStoreProductCoupon, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreProductCoupon), err
}

func (e ebStoreProductCouponDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreProductCoupon, err error) {
	buf := make([]*model.EbStoreProductCoupon, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreProductCouponDo) FindInBatches(result *[]*model.EbStoreProductCoupon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreProductCouponDo) Attrs(attrs ...field.AssignExpr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreProductCouponDo) Assign(attrs ...field.AssignExpr) IEbStoreProductCouponDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreProductCouponDo) Joins(fields ...field.RelationField) IEbStoreProductCouponDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreProductCouponDo) Preload(fields ...field.RelationField) IEbStoreProductCouponDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreProductCouponDo) FirstOrInit() (*model.EbStoreProductCoupon, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCoupon), nil
	}
}

func (e ebStoreProductCouponDo) FirstOrCreate() (*model.EbStoreProductCoupon, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreProductCoupon), nil
	}
}

func (e ebStoreProductCouponDo) FindByPage(offset int, limit int) (result []*model.EbStoreProductCoupon, count int64, err error) {
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

func (e ebStoreProductCouponDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreProductCouponDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreProductCouponDo) Delete(models ...*model.EbStoreProductCoupon) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreProductCouponDo) withDO(do gen.Dao) *ebStoreProductCouponDo {
	e.DO = *do.(*gen.DO)
	return e
}
