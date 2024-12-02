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

func newEbStoreSeckill(db *gorm.DB, opts ...gen.DOOption) ebStoreSeckill {
	_ebStoreSeckill := ebStoreSeckill{}

	_ebStoreSeckill.ebStoreSeckillDo.UseDB(db, opts...)
	_ebStoreSeckill.ebStoreSeckillDo.UseModel(&model.EbStoreSeckill{})

	tableName := _ebStoreSeckill.ebStoreSeckillDo.TableName()
	_ebStoreSeckill.ALL = field.NewAsterisk(tableName)
	_ebStoreSeckill.ID = field.NewInt32(tableName, "id")
	_ebStoreSeckill.ProductID = field.NewInt32(tableName, "product_id")
	_ebStoreSeckill.Image = field.NewString(tableName, "image")
	_ebStoreSeckill.Images = field.NewString(tableName, "images")
	_ebStoreSeckill.Title = field.NewString(tableName, "title")
	_ebStoreSeckill.Info = field.NewString(tableName, "info")
	_ebStoreSeckill.Price = field.NewFloat64(tableName, "price")
	_ebStoreSeckill.Cost = field.NewFloat64(tableName, "cost")
	_ebStoreSeckill.OtPrice = field.NewFloat64(tableName, "ot_price")
	_ebStoreSeckill.GiveIntegral = field.NewFloat64(tableName, "give_integral")
	_ebStoreSeckill.Sort = field.NewInt32(tableName, "sort")
	_ebStoreSeckill.Stock = field.NewInt32(tableName, "stock")
	_ebStoreSeckill.Sales = field.NewInt32(tableName, "sales")
	_ebStoreSeckill.UnitName = field.NewString(tableName, "unit_name")
	_ebStoreSeckill.Postage = field.NewFloat64(tableName, "postage")
	_ebStoreSeckill.Description = field.NewString(tableName, "description")
	_ebStoreSeckill.StartTime = field.NewTime(tableName, "start_time")
	_ebStoreSeckill.StopTime = field.NewTime(tableName, "stop_time")
	_ebStoreSeckill.CreateTime = field.NewTime(tableName, "create_time")
	_ebStoreSeckill.Status = field.NewInt32(tableName, "status")
	_ebStoreSeckill.IsPostage = field.NewInt32(tableName, "is_postage")
	_ebStoreSeckill.IsDel = field.NewInt32(tableName, "is_del")
	_ebStoreSeckill.Num = field.NewInt32(tableName, "num")
	_ebStoreSeckill.IsShow = field.NewInt32(tableName, "is_show")
	_ebStoreSeckill.TimeID = field.NewInt32(tableName, "time_id")
	_ebStoreSeckill.TempID = field.NewInt32(tableName, "temp_id")
	_ebStoreSeckill.Weight = field.NewFloat64(tableName, "weight")
	_ebStoreSeckill.Volume = field.NewFloat64(tableName, "volume")
	_ebStoreSeckill.Quota = field.NewInt32(tableName, "quota")
	_ebStoreSeckill.QuotaShow = field.NewInt32(tableName, "quota_show")
	_ebStoreSeckill.SpecType = field.NewBool(tableName, "spec_type")

	_ebStoreSeckill.fillFieldMap()

	return _ebStoreSeckill
}

// ebStoreSeckill 商品秒杀产品表
type ebStoreSeckill struct {
	ebStoreSeckillDo ebStoreSeckillDo

	ALL          field.Asterisk
	ID           field.Int32   // 商品秒杀产品表id
	ProductID    field.Int32   // 商品id
	Image        field.String  // 推荐图
	Images       field.String  // 轮播图
	Title        field.String  // 活动标题
	Info         field.String  // 简介
	Price        field.Float64 // 价格
	Cost         field.Float64 // 成本
	OtPrice      field.Float64 // 原价
	GiveIntegral field.Float64 // 返多少积分
	Sort         field.Int32   // 排序
	Stock        field.Int32   // 库存
	Sales        field.Int32   // 销量
	UnitName     field.String  // 单位名
	Postage      field.Float64 // 邮费
	Description  field.String  // 内容
	StartTime    field.Time    // 开始时间
	StopTime     field.Time    // 结束时间
	CreateTime   field.Time    // 添加时间
	Status       field.Int32   // 秒杀状态 0=关闭 1=开启
	IsPostage    field.Int32   // 是否包邮
	IsDel        field.Int32   // 删除 0未删除1已删除
	Num          field.Int32   // 当天参与活动次数
	IsShow       field.Int32   // 显示
	TimeID       field.Int32   // 时间段ID
	TempID       field.Int32   // 运费模板ID
	Weight       field.Float64 // 重量
	Volume       field.Float64 // 体积
	Quota        field.Int32   // 限购总数,随减
	QuotaShow    field.Int32   // 限购总数显示.不变
	SpecType     field.Bool    // 规格 0=单 1=多

	fieldMap map[string]field.Expr
}

func (e ebStoreSeckill) Table(newTableName string) *ebStoreSeckill {
	e.ebStoreSeckillDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebStoreSeckill) As(alias string) *ebStoreSeckill {
	e.ebStoreSeckillDo.DO = *(e.ebStoreSeckillDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebStoreSeckill) updateTableName(table string) *ebStoreSeckill {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.ProductID = field.NewInt32(table, "product_id")
	e.Image = field.NewString(table, "image")
	e.Images = field.NewString(table, "images")
	e.Title = field.NewString(table, "title")
	e.Info = field.NewString(table, "info")
	e.Price = field.NewFloat64(table, "price")
	e.Cost = field.NewFloat64(table, "cost")
	e.OtPrice = field.NewFloat64(table, "ot_price")
	e.GiveIntegral = field.NewFloat64(table, "give_integral")
	e.Sort = field.NewInt32(table, "sort")
	e.Stock = field.NewInt32(table, "stock")
	e.Sales = field.NewInt32(table, "sales")
	e.UnitName = field.NewString(table, "unit_name")
	e.Postage = field.NewFloat64(table, "postage")
	e.Description = field.NewString(table, "description")
	e.StartTime = field.NewTime(table, "start_time")
	e.StopTime = field.NewTime(table, "stop_time")
	e.CreateTime = field.NewTime(table, "create_time")
	e.Status = field.NewInt32(table, "status")
	e.IsPostage = field.NewInt32(table, "is_postage")
	e.IsDel = field.NewInt32(table, "is_del")
	e.Num = field.NewInt32(table, "num")
	e.IsShow = field.NewInt32(table, "is_show")
	e.TimeID = field.NewInt32(table, "time_id")
	e.TempID = field.NewInt32(table, "temp_id")
	e.Weight = field.NewFloat64(table, "weight")
	e.Volume = field.NewFloat64(table, "volume")
	e.Quota = field.NewInt32(table, "quota")
	e.QuotaShow = field.NewInt32(table, "quota_show")
	e.SpecType = field.NewBool(table, "spec_type")

	e.fillFieldMap()

	return e
}

func (e *ebStoreSeckill) WithContext(ctx context.Context) IEbStoreSeckillDo {
	return e.ebStoreSeckillDo.WithContext(ctx)
}

func (e ebStoreSeckill) TableName() string { return e.ebStoreSeckillDo.TableName() }

func (e ebStoreSeckill) Alias() string { return e.ebStoreSeckillDo.Alias() }

func (e ebStoreSeckill) Columns(cols ...field.Expr) gen.Columns {
	return e.ebStoreSeckillDo.Columns(cols...)
}

func (e *ebStoreSeckill) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebStoreSeckill) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 31)
	e.fieldMap["id"] = e.ID
	e.fieldMap["product_id"] = e.ProductID
	e.fieldMap["image"] = e.Image
	e.fieldMap["images"] = e.Images
	e.fieldMap["title"] = e.Title
	e.fieldMap["info"] = e.Info
	e.fieldMap["price"] = e.Price
	e.fieldMap["cost"] = e.Cost
	e.fieldMap["ot_price"] = e.OtPrice
	e.fieldMap["give_integral"] = e.GiveIntegral
	e.fieldMap["sort"] = e.Sort
	e.fieldMap["stock"] = e.Stock
	e.fieldMap["sales"] = e.Sales
	e.fieldMap["unit_name"] = e.UnitName
	e.fieldMap["postage"] = e.Postage
	e.fieldMap["description"] = e.Description
	e.fieldMap["start_time"] = e.StartTime
	e.fieldMap["stop_time"] = e.StopTime
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["status"] = e.Status
	e.fieldMap["is_postage"] = e.IsPostage
	e.fieldMap["is_del"] = e.IsDel
	e.fieldMap["num"] = e.Num
	e.fieldMap["is_show"] = e.IsShow
	e.fieldMap["time_id"] = e.TimeID
	e.fieldMap["temp_id"] = e.TempID
	e.fieldMap["weight"] = e.Weight
	e.fieldMap["volume"] = e.Volume
	e.fieldMap["quota"] = e.Quota
	e.fieldMap["quota_show"] = e.QuotaShow
	e.fieldMap["spec_type"] = e.SpecType
}

func (e ebStoreSeckill) clone(db *gorm.DB) ebStoreSeckill {
	e.ebStoreSeckillDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebStoreSeckill) replaceDB(db *gorm.DB) ebStoreSeckill {
	e.ebStoreSeckillDo.ReplaceDB(db)
	return e
}

type ebStoreSeckillDo struct{ gen.DO }

type IEbStoreSeckillDo interface {
	gen.SubQuery
	Debug() IEbStoreSeckillDo
	WithContext(ctx context.Context) IEbStoreSeckillDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbStoreSeckillDo
	WriteDB() IEbStoreSeckillDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbStoreSeckillDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbStoreSeckillDo
	Not(conds ...gen.Condition) IEbStoreSeckillDo
	Or(conds ...gen.Condition) IEbStoreSeckillDo
	Select(conds ...field.Expr) IEbStoreSeckillDo
	Where(conds ...gen.Condition) IEbStoreSeckillDo
	Order(conds ...field.Expr) IEbStoreSeckillDo
	Distinct(cols ...field.Expr) IEbStoreSeckillDo
	Omit(cols ...field.Expr) IEbStoreSeckillDo
	Join(table schema.Tabler, on ...field.Expr) IEbStoreSeckillDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreSeckillDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreSeckillDo
	Group(cols ...field.Expr) IEbStoreSeckillDo
	Having(conds ...gen.Condition) IEbStoreSeckillDo
	Limit(limit int) IEbStoreSeckillDo
	Offset(offset int) IEbStoreSeckillDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreSeckillDo
	Unscoped() IEbStoreSeckillDo
	Create(values ...*model.EbStoreSeckill) error
	CreateInBatches(values []*model.EbStoreSeckill, batchSize int) error
	Save(values ...*model.EbStoreSeckill) error
	First() (*model.EbStoreSeckill, error)
	Take() (*model.EbStoreSeckill, error)
	Last() (*model.EbStoreSeckill, error)
	Find() ([]*model.EbStoreSeckill, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreSeckill, err error)
	FindInBatches(result *[]*model.EbStoreSeckill, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbStoreSeckill) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbStoreSeckillDo
	Assign(attrs ...field.AssignExpr) IEbStoreSeckillDo
	Joins(fields ...field.RelationField) IEbStoreSeckillDo
	Preload(fields ...field.RelationField) IEbStoreSeckillDo
	FirstOrInit() (*model.EbStoreSeckill, error)
	FirstOrCreate() (*model.EbStoreSeckill, error)
	FindByPage(offset int, limit int) (result []*model.EbStoreSeckill, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbStoreSeckillDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebStoreSeckillDo) Debug() IEbStoreSeckillDo {
	return e.withDO(e.DO.Debug())
}

func (e ebStoreSeckillDo) WithContext(ctx context.Context) IEbStoreSeckillDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebStoreSeckillDo) ReadDB() IEbStoreSeckillDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebStoreSeckillDo) WriteDB() IEbStoreSeckillDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebStoreSeckillDo) Session(config *gorm.Session) IEbStoreSeckillDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebStoreSeckillDo) Clauses(conds ...clause.Expression) IEbStoreSeckillDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebStoreSeckillDo) Returning(value interface{}, columns ...string) IEbStoreSeckillDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebStoreSeckillDo) Not(conds ...gen.Condition) IEbStoreSeckillDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebStoreSeckillDo) Or(conds ...gen.Condition) IEbStoreSeckillDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebStoreSeckillDo) Select(conds ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebStoreSeckillDo) Where(conds ...gen.Condition) IEbStoreSeckillDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebStoreSeckillDo) Order(conds ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebStoreSeckillDo) Distinct(cols ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebStoreSeckillDo) Omit(cols ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebStoreSeckillDo) Join(table schema.Tabler, on ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebStoreSeckillDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebStoreSeckillDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebStoreSeckillDo) Group(cols ...field.Expr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebStoreSeckillDo) Having(conds ...gen.Condition) IEbStoreSeckillDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebStoreSeckillDo) Limit(limit int) IEbStoreSeckillDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebStoreSeckillDo) Offset(offset int) IEbStoreSeckillDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebStoreSeckillDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbStoreSeckillDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebStoreSeckillDo) Unscoped() IEbStoreSeckillDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebStoreSeckillDo) Create(values ...*model.EbStoreSeckill) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebStoreSeckillDo) CreateInBatches(values []*model.EbStoreSeckill, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebStoreSeckillDo) Save(values ...*model.EbStoreSeckill) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebStoreSeckillDo) First() (*model.EbStoreSeckill, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreSeckill), nil
	}
}

func (e ebStoreSeckillDo) Take() (*model.EbStoreSeckill, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreSeckill), nil
	}
}

func (e ebStoreSeckillDo) Last() (*model.EbStoreSeckill, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreSeckill), nil
	}
}

func (e ebStoreSeckillDo) Find() ([]*model.EbStoreSeckill, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbStoreSeckill), err
}

func (e ebStoreSeckillDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbStoreSeckill, err error) {
	buf := make([]*model.EbStoreSeckill, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebStoreSeckillDo) FindInBatches(result *[]*model.EbStoreSeckill, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebStoreSeckillDo) Attrs(attrs ...field.AssignExpr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebStoreSeckillDo) Assign(attrs ...field.AssignExpr) IEbStoreSeckillDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebStoreSeckillDo) Joins(fields ...field.RelationField) IEbStoreSeckillDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebStoreSeckillDo) Preload(fields ...field.RelationField) IEbStoreSeckillDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebStoreSeckillDo) FirstOrInit() (*model.EbStoreSeckill, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreSeckill), nil
	}
}

func (e ebStoreSeckillDo) FirstOrCreate() (*model.EbStoreSeckill, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbStoreSeckill), nil
	}
}

func (e ebStoreSeckillDo) FindByPage(offset int, limit int) (result []*model.EbStoreSeckill, count int64, err error) {
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

func (e ebStoreSeckillDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebStoreSeckillDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebStoreSeckillDo) Delete(models ...*model.EbStoreSeckill) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebStoreSeckillDo) withDO(do gen.Dao) *ebStoreSeckillDo {
	e.DO = *do.(*gen.DO)
	return e
}
