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

func newEbStoreCombination(db *gorm.DB, opts ...gen.DOOption) ebStoreCombination {
	_ebStoreCombination := ebStoreCombination{}

	_ebStoreCombination.ebStoreCombinationDo.UseDB(db, opts...)
	_ebStoreCombination.ebStoreCombinationDo.UseModel(&model.EbStoreCombination{})

	tableName := _ebStoreCombination.ebStoreCombinationDo.TableName()
	_ebStoreCombination.ALL = field.NewAsterisk(tableName)
	_ebStoreCombination.ID = field.NewInt32(tableName, "id")
	_ebStoreCombination.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreCombination.MerID = field.NewInt32(tableName, "mer_id")
	_ebStoreCombination.Image = field.NewString(tableName, "image")
	_ebStoreCombination.Images = field.NewString(tableName, "images")
	_ebStoreCombination.Title = field.NewString(tableName, "title")
	_ebStoreCombination.Attr = field.NewString(tableName, "attr")
	_ebStoreCombination.People = field.NewInt32(tableName, "people")
	_ebStoreCombination.Info = field.NewString(tableName, "info")
	_ebStoreCombination.Price = field.NewFloat64(tableName, "price")
	_ebStoreCombination.Sort = field.NewInt32(tableName, "sort")
	_ebStoreCombination.Sales = field.NewInt32(tableName, "sales")
	_ebStoreCombination.Stock = field.NewInt32(tableName, "stock")
	_ebStoreCombination.AddTime = field.NewInt64(tableName, "add_time")
	_ebStoreCombination.IsHost = field.NewInt32(tableName, "is_host")
	_ebStoreCombination.IsShow = field.NewInt32(tableName, "is_show")
	_ebStoreCombination.IsDel = field.NewInt32(tableName, "is_del")
	_ebStoreCombination.Combination = field.NewInt32(tableName, "combination")
	_ebStoreCombination.MerUse = field.NewInt32(tableName, "mer_use")
	_ebStoreCombination.IsPostage = field.NewInt32(tableName, "is_postage")
	_ebStoreCombination.Postage = field.NewFloat64(tableName, "postage")
	_ebStoreCombination.StartTime = field.NewInt64(tableName, "start_time")
	_ebStoreCombination.StopTime = field.NewInt64(tableName, "stop_time")
	_ebStoreCombination.EffectiveTime = field.NewInt32(tableName, "effective_time")
	_ebStoreCombination.Cost = field.NewFloat64(tableName, "cost")
	_ebStoreCombination.Browse = field.NewInt32(tableName, "browse")
	_ebStoreCombination.UnitName = field.NewString(tableName, "unit_name")
	_ebStoreCombination.TempID = field.NewInt32(tableName, "temp_id")
	_ebStoreCombination.Weight = field.NewFloat64(tableName, "weight")
	_ebStoreCombination.Volume = field.NewFloat64(tableName, "volume")
	_ebStoreCombination.Num = field.NewInt32(tableName, "num")
	_ebStoreCombination.Quota = field.NewInt32(tableName, "quota")
	_ebStoreCombination.QuotaShow = field.NewInt32(tableName, "quota_show")
	_ebStoreCombination.OtPrice = field.NewFloat64(tableName, "ot_price")
	_ebStoreCombination.OnceNum = field.NewInt32(tableName, "once_num")
	_ebStoreCombination.VirtualRation = field.NewInt32(tableName, "virtual_ration")

	_ebStoreCombination.fillFieldMap()

	return _ebStoreCombination
}

// ebStoreCombination 拼团商品表
type ebStoreCombination struct {
	ebStoreCombinationDo ebStoreCombinationDo

	ALL           field.Asterisk
	ID            field.Int32   // 拼团商品ID
	ProductID     field.Int32   // 商品id
	MerID         field.Int32   // 商户id
	Image         field.String  // 推荐图
	Images        field.String  // 轮播图
	Title         field.String  // 活动标题
	Attr          field.String  // 活动属性
	People        field.Int32   // 参团人数
	Info          field.String  // 简介
	Price         field.Float64 // 价格
	Sort          field.Int32   // 排序
	Sales         field.Int32   // 销量
	Stock         field.Int32   // 库存
	AddTime       field.Int64   // 添加时间
	IsHost        field.Int32   // 推荐
	IsShow        field.Int32   // 商品状态
	IsDel         field.Int32
	Combination   field.Int32
	MerUse        field.Int32   // 商户是否可用1可用0不可用
	IsPostage     field.Int32   // 是否包邮1是0否
	Postage       field.Float64 // 邮费
	StartTime     field.Int64   // 拼团开始时间
	StopTime      field.Int64   // 拼团结束时间
	EffectiveTime field.Int32   // 拼团订单有效时间(小时)
	Cost          field.Float64 // 拼图商品成本
	Browse        field.Int32   // 浏览量
	UnitName      field.String  // 单位名
	TempID        field.Int32   // 运费模板ID
	Weight        field.Float64 // 重量
	Volume        field.Float64 // 体积
	Num           field.Int32   // 单次购买数量
	Quota         field.Int32   // 限购总数
	QuotaShow     field.Int32   // 限量总数显示
	OtPrice       field.Float64 // 原价
	OnceNum       field.Int32   // 每个订单可购买数量
	VirtualRation field.Int32   // 虚拟成团百分比

	fieldMap map[string]field.Expr
}

func (e ebStoreCombination) Table(newTableName string) *ebStoreCombination {
	e.ebStoreCombinationDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreCombination) As(alias string) *ebStoreCombination {
	e.ebStoreCombinationDo.DO = *(e.ebStoreCombinationDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreCombination) updateTableName(table string) *ebStoreCombination {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.ProductID = field.NewInt32(table, "product_id")
	e.MerID = field.NewInt32(table, "mer_id")
	e.Image = field.NewString(table, "image")
	e.Images = field.NewString(table, "images")
	e.Title = field.NewString(table, "title")
	e.Attr = field.NewString(table, "attr")
	e.People = field.NewInt32(table, "people")
	e.Info = field.NewString(table, "info")
	e.Price = field.NewFloat64(table, "price")
	e.Sort = field.NewInt32(table, "sort")
	e.Sales = field.NewInt32(table, "sales")
	e.Stock = field.NewInt32(table, "stock")
	e.AddTime = field.NewInt64(table, "add_time")
	e.IsHost = field.NewInt32(table, "is_host")
	e.IsShow = field.NewInt32(table, "is_show")
	e.IsDel = field.NewInt32(table, "is_del")
	e.Combination = field.NewInt32(table, "combination")
	e.MerUse = field.NewInt32(table, "mer_use")
	e.IsPostage = field.NewInt32(table, "is_postage")
	e.Postage = field.NewFloat64(table, "postage")
	e.StartTime = field.NewInt64(table, "start_time")
	e.StopTime = field.NewInt64(table, "stop_time")
	e.EffectiveTime = field.NewInt32(table, "effective_time")
	e.Cost = field.NewFloat64(table, "cost")
	e.Browse = field.NewInt32(table, "browse")
	e.UnitName = field.NewString(table, "unit_name")
	e.TempID = field.NewInt32(table, "temp_id")
	e.Weight = field.NewFloat64(table, "weight")
	e.Volume = field.NewFloat64(table, "volume")
	e.Num = field.NewInt32(table, "num")
	e.Quota = field.NewInt32(table, "quota")
	e.QuotaShow = field.NewInt32(table, "quota_show")
	e.OtPrice = field.NewFloat64(table, "ot_price")
	e.OnceNum = field.NewInt32(table, "once_num")
	e.VirtualRation = field.NewInt32(table, "virtual_ration")

	e.fillFieldMap()

	return e
}

func (e *ebStoreCombination) WithContext(ctx context.Context) IEbStoreCombinationDo {
	return e.ebStoreCombinationDo.WithContext(ctx)
}

func (e ebStoreCombination) TableName() string { return e.ebStoreCombinationDo.TableName() }

func (e ebStoreCombination) Alias() string { return e.ebStoreCombinationDo.Alias() }

func (e ebStoreCombination) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreCombinationDo.Columns(cols...)
}

func (e *ebStoreCombination) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreCombination) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 36)
	e.fieldMap["id"] = e.ID
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["mer_id"] = e.MerID
	e.fieldMap["image"] = e.Image
	e.fieldMap["images"] = e.Images
	e.fieldMap["title"] = e.Title
	e.fieldMap["attr"] = e.Attr
	e.fieldMap["people"] = e.People
	e.fieldMap["info"] = e.Info
	e.fieldMap["price"] = e.Price
	e.fieldMap["sort"] = e.Sort
	e.fieldMap["sales"] = e.Sales
	e.fieldMap["stock"] = e.Stock
	e.fieldMap["add_time"] = e.AddTime
	e.fieldMap["is_host"] = e.IsHost
	e.fieldMap["is_show"] = e.IsShow
	e.fieldMap["is_del"] = e.IsDel
	e.fieldMap["combination"] = e.Combination
	e.fieldMap["mer_use"] = e.MerUse
	e.fieldMap["is_postage"] = e.IsPostage
	e.fieldMap["postage"] = e.Postage
	e.fieldMap["start_time"] = e.StartTime
	e.fieldMap["stop_time"] = e.StopTime
	e.fieldMap["effective_time"] = e.EffectiveTime
	e.fieldMap["cost"] = e.Cost
	e.fieldMap["browse"] = e.Browse
	e.fieldMap["unit_name"] = e.UnitName
	e.fieldMap["temp_id"] = e.TempID
	e.fieldMap["weight"] = e.Weight
	e.fieldMap["volume"] = e.Volume
	e.fieldMap["num"] = e.Num
	e.fieldMap["quota"] = e.Quota
	e.fieldMap["quota_show"] = e.QuotaShow
	e.fieldMap["ot_price"] = e.OtPrice
	e.fieldMap["once_num"] = e.OnceNum
	e.fieldMap["virtual_ration"] = e.VirtualRation
}

func (e ebStoreCombination) clone(db *gorm.DB) ebStoreCombination {
	e.ebStoreCombinationDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreCombination) replaceDB(db *gorm.DB) ebStoreCombination {
	e.ebStoreCombinationDo.ReplaceDB(db)
	return e
}

type ebStoreCombinationDo struct{ gen.DO }

type IEbStoreCombinationDo interface {
	gen.SubQuery
	Debug() IEbStoreCombinationDo
	WithContext(ctx context.Context) IEbStoreCombinationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreCombinationDo
	WriteDB() IEbStoreCombinationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreCombinationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreCombinationDo
	Not(conds ...gen.Condition) IEbStoreCombinationDo
	Or(conds ...gen.Condition) IEbStoreCombinationDo
	Select(conds ...field.Expr) IEbStoreCombinationDo
	Where(conds ...gen.Condition) IEbStoreCombinationDo
	Order(conds ...field.Expr) IEbStoreCombinationDo
	Distinct(cols ...field.Expr) IEbStoreCombinationDo
	Omit(cols ...field.Expr) IEbStoreCombinationDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreCombinationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreCombinationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreCombinationDo
	Group(cols ...field.Expr) IEbStoreCombinationDo
	Having(conds ...gen.Condition) IEbStoreCombinationDo
	Limit(limit int) IEbStoreCombinationDo
	Offset(offset int) IEbStoreCombinationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreCombinationDo
	Unscoped() IEbStoreCombinationDo
	Create(values ...*model.EbStoreCombination) error
	CreateInBatches(values []*model.EbStoreCombination, batchSize int) error
	Save(values ...*model.EbStoreCombination) error
	First() (*model.EbStoreCombination, error)
	Take() (*model.EbStoreCombination, error)
	Last() (*model.EbStoreCombination, error)
	Find() ([]*model.EbStoreCombination, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreCombination, err error)
	FindInBatches(result *[]*model.EbStoreCombination, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreCombination) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreCombinationDo
	Assign(attrs ...field.AssignExpr) IEbStoreCombinationDo
	Joins(fields ...field.RelationField) IEbStoreCombinationDo
	Preload(fields ...field.RelationField) IEbStoreCombinationDo
	FirstOrInit() (*model.EbStoreCombination, error)
	FirstOrCreate() (*model.EbStoreCombination, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreCombination, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreCombinationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreCombinationDo) Debug() IEbStoreCombinationDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreCombinationDo) WithContext(ctx context.Context) IEbStoreCombinationDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreCombinationDo) ReadDB() IEbStoreCombinationDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreCombinationDo) WriteDB() IEbStoreCombinationDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreCombinationDo) Session(config *gorm.Session) IEbStoreCombinationDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreCombinationDo) Clauses(conds ...clause.Expression) IEbStoreCombinationDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreCombinationDo) Returning(value interface{}, columns ...string) IEbStoreCombinationDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreCombinationDo) Not(conds ...gen.Condition) IEbStoreCombinationDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreCombinationDo) Or(conds ...gen.Condition) IEbStoreCombinationDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreCombinationDo) Select(conds ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreCombinationDo) Where(conds ...gen.Condition) IEbStoreCombinationDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreCombinationDo) Order(conds ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreCombinationDo) Distinct(cols ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreCombinationDo) Omit(cols ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreCombinationDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreCombinationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreCombinationDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreCombinationDo) Group(cols ...field.Expr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreCombinationDo) Having(conds ...gen.Condition) IEbStoreCombinationDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreCombinationDo) Limit(limit int) IEbStoreCombinationDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreCombinationDo) Offset(offset int) IEbStoreCombinationDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreCombinationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreCombinationDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreCombinationDo) Unscoped() IEbStoreCombinationDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreCombinationDo) Create(values ...*model.EbStoreCombination) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreCombinationDo) CreateInBatches(values []*model.EbStoreCombination, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreCombinationDo) Save(values ...*model.EbStoreCombination) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreCombinationDo) First() (*model.EbStoreCombination, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCombination), nil
	}
}

func (e ebStoreCombinationDo) Take() (*model.EbStoreCombination, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCombination), nil
	}
}

func (e ebStoreCombinationDo) Last() (*model.EbStoreCombination, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCombination), nil
	}
}

func (e ebStoreCombinationDo) Find() ([]*model.EbStoreCombination, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreCombination), err
}

func (e ebStoreCombinationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreCombination, err error) {
	buf := make([]*model.EbStoreCombination, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreCombinationDo) FindInBatches(result *[]*model.EbStoreCombination, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreCombinationDo) Attrs(attrs ...field.AssignExpr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreCombinationDo) Assign(attrs ...field.AssignExpr) IEbStoreCombinationDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreCombinationDo) Joins(fields ...field.RelationField) IEbStoreCombinationDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreCombinationDo) Preload(fields ...field.RelationField) IEbStoreCombinationDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreCombinationDo) FirstOrInit() (*model.EbStoreCombination, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCombination), nil
	}
}

func (e ebStoreCombinationDo) FirstOrCreate() (*model.EbStoreCombination, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreCombination), nil
	}
}

func (e ebStoreCombinationDo) FindByPage(offset int, limit int) (result []*model.EbStoreCombination, count int64, err error) {
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

func (e ebStoreCombinationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreCombinationDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreCombinationDo) Delete(models ...*model.EbStoreCombination) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreCombinationDo) withDO(do gen.Dao) *ebStoreCombinationDo {
	e.DO = *do.(*gen.DO)
	return e
}
