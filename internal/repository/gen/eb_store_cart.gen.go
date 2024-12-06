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

func newEbStoreCart(db *gorm.DB, opts ...gen.DOOption) ebStoreCart {
	_ebStoreCart := ebStoreCart{}

	_ebStoreCart.ebStoreCartDo.UseDB(db, opts...)
	_ebStoreCart.ebStoreCartDo.UseModel(&model.EbStoreCart{})

	tableName := _ebStoreCart.ebStoreCartDo.TableName()
	_ebStoreCart.ALL = field.NewAsterisk(tableName)
	_ebStoreCart.ID = field.NewInt64(tableName, "id")
	_ebStoreCart.UID = field.NewInt32(tableName, "uid")
	_ebStoreCart.Type = field.NewString(tableName, "type")
	_ebStoreCart.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreCart.ProductAttrUnique = field.NewString(tableName, "product_attr_unique")
	_ebStoreCart.CartNum = field.NewInt32(tableName, "cart_num")
	_ebStoreCart.IsNew = field.NewBool(tableName, "is_new")
	_ebStoreCart.CombinationID = field.NewInt32(tableName, "combination_id")
	_ebStoreCart.SeckillID = field.NewInt32(tableName, "seckill_id")
	_ebStoreCart.BargainID = field.NewInt32(tableName, "bargain_id")
	_ebStoreCart.CreateTime = field.NewTime(tableName, "create_time")
	_ebStoreCart.UpdateTime = field.NewTime(tableName, "update_time")
	_ebStoreCart.Status = field.NewBool(tableName, "status")

	_ebStoreCart.fillFieldMap()

	return _ebStoreCart
}

// ebStoreCart 购物车表
type ebStoreCart struct {
	ebStoreCartDo ebStoreCartDo

	ALL               field.Asterisk
	ID                field.Int64  // 购物车表ID
	UID               field.Int32  // 用户ID
	Type              field.String // 类型
	ProductID         field.Int32  // 商品ID
	ProductAttrUnique field.String // 商品属性
	CartNum           field.Int32  // 商品数量
	IsNew             field.Bool   // 是否为立即购买
	CombinationID     field.Int32  // 拼团id
	SeckillID         field.Int32  // 秒杀商品ID
	BargainID         field.Int32  // 砍价id
	CreateTime        field.Time   // 添加时间
	UpdateTime        field.Time   // g
	Status            field.Bool   // 购物车状态

	fieldMap map[string]field.Expr
}

func (e ebStoreCart) Table(newTableName string) *ebStoreCart {
	e.ebStoreCartDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreCart) As(alias string) *ebStoreCart {
	e.ebStoreCartDo.DO = *(e.ebStoreCartDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreCart) updateTableName(table string) *ebStoreCart {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.UID = field.NewInt32(table, "uid")
	e.Type = field.NewString(table, "type")
	e.ProductID = field.NewInt32(table, "product_id")
	e.ProductAttrUnique = field.NewString(table, "product_attr_unique")
	e.CartNum = field.NewInt32(table, "cart_num")
	e.IsNew = field.NewBool(table, "is_new")
	e.CombinationID = field.NewInt32(table, "combination_id")
	e.SeckillID = field.NewInt32(table, "seckill_id")
	e.BargainID = field.NewInt32(table, "bargain_id")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")
	e.Status = field.NewBool(table, "status")

	e.fillFieldMap()

	return e
}

func (e *ebStoreCart) WithContext(ctx context.Context) IEbStoreCartDo {
	return e.ebStoreCartDo.WithContext(ctx)
}

func (e ebStoreCart) TableName() string { return e.ebStoreCartDo.TableName() }

func (e ebStoreCart) Alias() string { return e.ebStoreCartDo.Alias() }

func (e ebStoreCart) Columns(cols ...field.Expr) gen.Columns { return e.ebStoreCartDo.Columns(cols...) }

func (e *ebStoreCart) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreCart) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 13)
	e.fieldMap["id"] = e.ID
	e.fieldMap["uid"] = e.UID
	e.fieldMap["type"] = e.Type
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["product_attr_unique"] = e.ProductAttrUnique
	e.fieldMap["cart_num"] = e.CartNum
	e.fieldMap["is_new"] = e.IsNew
	e.fieldMap["combination_id"] = e.CombinationID
	e.fieldMap["seckill_id"] = e.SeckillID
	e.fieldMap["bargain_id"] = e.BargainID
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
	e.fieldMap["status"] = e.Status
}

func (e ebStoreCart) clone(db *gorm.DB) ebStoreCart {
	e.ebStoreCartDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreCart) replaceDB(db *gorm.DB) ebStoreCart {
	e.ebStoreCartDo.ReplaceDB(db)
	return e
}

type ebStoreCartDo struct{ gen.DO }

type IEbStoreCartDo interface {
	gen.SubQuery
	Debug() IEbStoreCartDo
	WithContext(ctx context.Context) IEbStoreCartDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreCartDo
	WriteDB() IEbStoreCartDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreCartDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreCartDo
	Not(conds ...gen.Condition) IEbStoreCartDo
	Or(conds ...gen.Condition) IEbStoreCartDo
	Select(conds ...field.Expr) IEbStoreCartDo
	Where(conds ...gen.Condition) IEbStoreCartDo
	Order(conds ...field.Expr) IEbStoreCartDo
	Distinct(cols ...field.Expr) IEbStoreCartDo
	Omit(cols ...field.Expr) IEbStoreCartDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreCartDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreCartDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreCartDo
	Group(cols ...field.Expr) IEbStoreCartDo
	Having(conds ...gen.Condition) IEbStoreCartDo
	Limit(limit int) IEbStoreCartDo
	Offset(offset int) IEbStoreCartDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreCartDo
	Unscoped() IEbStoreCartDo
	Create(values ...*model.EbStoreCart) error
	CreateInBatches(values []*model.EbStoreCart, batchSize int) error
	Save(values ...*model.EbStoreCart) error
	First() (*model.EbStoreCart, error)
	Take() (*model.EbStoreCart, error)
	Last() (*model.EbStoreCart, error)
	Find() ([]*model.EbStoreCart, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreCart, err error)
	FindInBatches(result *[]*model.EbStoreCart, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreCart) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreCartDo
	Assign(attrs ...field.AssignExpr) IEbStoreCartDo
	Joins(fields ...field.RelationField) IEbStoreCartDo
	Preload(fields ...field.RelationField) IEbStoreCartDo
	FirstOrInit() (*model.EbStoreCart, error)
	FirstOrCreate() (*model.EbStoreCart, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreCart, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreCartDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreCartDo) Debug() IEbStoreCartDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreCartDo) WithContext(ctx context.Context) IEbStoreCartDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreCartDo) ReadDB() IEbStoreCartDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreCartDo) WriteDB() IEbStoreCartDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreCartDo) Session(config *gorm.Session) IEbStoreCartDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreCartDo) Clauses(conds ...clause.Expression) IEbStoreCartDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreCartDo) Returning(value interface{}, columns ...string) IEbStoreCartDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreCartDo) Not(conds ...gen.Condition) IEbStoreCartDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreCartDo) Or(conds ...gen.Condition) IEbStoreCartDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreCartDo) Select(conds ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreCartDo) Where(conds ...gen.Condition) IEbStoreCartDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreCartDo) Order(conds ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreCartDo) Distinct(cols ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreCartDo) Omit(cols ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreCartDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreCartDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreCartDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreCartDo) Group(cols ...field.Expr) IEbStoreCartDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreCartDo) Having(conds ...gen.Condition) IEbStoreCartDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreCartDo) Limit(limit int) IEbStoreCartDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreCartDo) Offset(offset int) IEbStoreCartDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreCartDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreCartDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreCartDo) Unscoped() IEbStoreCartDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreCartDo) Create(values ...*model.EbStoreCart) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreCartDo) CreateInBatches(values []*model.EbStoreCart, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreCartDo) Save(values ...*model.EbStoreCart) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreCartDo) First() (*model.EbStoreCart, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCart), nil
	}
}

func (e ebStoreCartDo) Take() (*model.EbStoreCart, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCart), nil
	}
}

func (e ebStoreCartDo) Last() (*model.EbStoreCart, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCart), nil
	}
}

func (e ebStoreCartDo) Find() ([]*model.EbStoreCart, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreCart), err
}

func (e ebStoreCartDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreCart, err error) {
	buf := make([]*model.EbStoreCart, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreCartDo) FindInBatches(result *[]*model.EbStoreCart, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreCartDo) Attrs(attrs ...field.AssignExpr) IEbStoreCartDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreCartDo) Assign(attrs ...field.AssignExpr) IEbStoreCartDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreCartDo) Joins(fields ...field.RelationField) IEbStoreCartDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreCartDo) Preload(fields ...field.RelationField) IEbStoreCartDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreCartDo) FirstOrInit() (*model.EbStoreCart, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCart), nil
	}
}

func (e ebStoreCartDo) FirstOrCreate() (*model.EbStoreCart, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCart), nil
	}
}

func (e ebStoreCartDo) FindByPage(offset int, limit int) (result []*model.EbStoreCart, count int64, err error) {
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

func (e ebStoreCartDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreCartDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreCartDo) Delete(models ...*model.EbStoreCart) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreCartDo) withDO(do gen.Dao) *ebStoreCartDo {
	e.DO = *do.(*gen.DO)
	return e
}
