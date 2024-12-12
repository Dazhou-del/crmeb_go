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

func newStoreProduct(db *gorm.DB, opts ...gen.DOOption) storeProduct {
	_storeProduct := storeProduct{}

	_storeProduct.storeProductDo.UseDB(db, opts...)
	_storeProduct.storeProductDo.UseModel(&model.StoreProduct{})

	tableName := _storeProduct.storeProductDo.TableName()
	_storeProduct.ALL = field.NewAsterisk(tableName)
	_storeProduct.ID = field.NewInt64(tableName, "id")
	_storeProduct.MerID = field.NewInt64(tableName, "mer_id")
	_storeProduct.Image = field.NewString(tableName, "image")
	_storeProduct.SliderImage = field.NewString(tableName, "slider_image")
	_storeProduct.StoreName = field.NewString(tableName, "store_name")
	_storeProduct.StoreInfo = field.NewString(tableName, "store_info")
	_storeProduct.Keyword = field.NewString(tableName, "keyword")
	_storeProduct.BarCode = field.NewString(tableName, "bar_code")
	_storeProduct.CateID = field.NewString(tableName, "cate_id")
	_storeProduct.Price = field.NewField(tableName, "price")
	_storeProduct.VipPrice = field.NewField(tableName, "vip_price")
	_storeProduct.OtPrice = field.NewField(tableName, "ot_price")
	_storeProduct.Postage = field.NewField(tableName, "postage")
	_storeProduct.UnitName = field.NewString(tableName, "unit_name")
	_storeProduct.Sort = field.NewInt64(tableName, "sort")
	_storeProduct.Sales = field.NewInt64(tableName, "sales")
	_storeProduct.Stock = field.NewInt64(tableName, "stock")
	_storeProduct.IsShow = field.NewInt64(tableName, "is_show")
	_storeProduct.IsHot = field.NewInt64(tableName, "is_hot")
	_storeProduct.IsBenefit = field.NewInt64(tableName, "is_benefit")
	_storeProduct.IsBest = field.NewInt64(tableName, "is_best")
	_storeProduct.IsNew = field.NewInt64(tableName, "is_new")
	_storeProduct.AddTime = field.NewInt64(tableName, "add_time")
	_storeProduct.IsPostage = field.NewInt64(tableName, "is_postage")
	_storeProduct.IsDel = field.NewInt64(tableName, "is_del")
	_storeProduct.MerUse = field.NewInt64(tableName, "mer_use")
	_storeProduct.GiveIntegral = field.NewInt64(tableName, "give_integral")
	_storeProduct.Cost = field.NewField(tableName, "cost")
	_storeProduct.IsSeckill = field.NewInt64(tableName, "is_seckill")
	_storeProduct.IsBargain = field.NewInt64(tableName, "is_bargain")
	_storeProduct.IsGood = field.NewInt64(tableName, "is_good")
	_storeProduct.IsSub = field.NewInt64(tableName, "is_sub")
	_storeProduct.Ficti = field.NewInt64(tableName, "ficti")
	_storeProduct.Browse = field.NewInt64(tableName, "browse")
	_storeProduct.CodePath = field.NewString(tableName, "code_path")
	_storeProduct.SoureLink = field.NewString(tableName, "soure_link")
	_storeProduct.VideoLink = field.NewString(tableName, "video_link")
	_storeProduct.TempID = field.NewInt64(tableName, "temp_id")
	_storeProduct.SpecType = field.NewInt64(tableName, "spec_type")
	_storeProduct.Activity = field.NewString(tableName, "activity")
	_storeProduct.FlatPattern = field.NewString(tableName, "flat_pattern")
	_storeProduct.IsRecycle = field.NewInt64(tableName, "is_recycle")
	_storeProduct.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProduct.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProduct.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProduct.fillFieldMap()

	return _storeProduct
}

// storeProduct 商品表
type storeProduct struct {
	storeProductDo storeProductDo

	ALL          field.Asterisk
	ID           field.Int64  // 商品id
	MerID        field.Int64  // 商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image        field.String // 商品图片
	SliderImage  field.String // 轮播图
	StoreName    field.String // 商品名称
	StoreInfo    field.String // 商品简介
	Keyword      field.String // 关键字
	BarCode      field.String // 商品条码（一维码）
	CateID       field.String // 分类id
	Price        field.Field  // 商品价格
	VipPrice     field.Field  // 会员价格
	OtPrice      field.Field  // 市场价
	Postage      field.Field  // 邮费
	UnitName     field.String // 单位名
	Sort         field.Int64  // 排序
	Sales        field.Int64  // 销量
	Stock        field.Int64  // 库存
	IsShow       field.Int64  // 状态（0：未上架，1：上架）
	IsHot        field.Int64  // 是否热卖
	IsBenefit    field.Int64  // 是否优惠
	IsBest       field.Int64  // 是否精品
	IsNew        field.Int64  // 是否新品
	AddTime      field.Int64  // 添加时间
	IsPostage    field.Int64  // 是否包邮
	IsDel        field.Int64  // 是否删除
	MerUse       field.Int64  // 商户是否代理 0不可代理1可代理
	GiveIntegral field.Int64  // 获得积分
	Cost         field.Field  // 成本价
	IsSeckill    field.Int64  // 秒杀状态 0 未开启 1已开启
	IsBargain    field.Int64  // 砍价状态 0未开启 1开启
	IsGood       field.Int64  // 是否优品推荐
	IsSub        field.Int64  // 是否单独分佣
	Ficti        field.Int64  // 虚拟销量
	Browse       field.Int64  // 浏览量
	CodePath     field.String // 商品二维码地址(用户小程序海报)
	SoureLink    field.String // 淘宝京东1688类型
	VideoLink    field.String // 主图视频链接
	TempID       field.Int64  // 运费模板ID
	SpecType     field.Int64  // 规格 0单 1多
	Activity     field.String // 活动显示排序0=默认, 1=秒杀，2=砍价，3=拼团
	FlatPattern  field.String // 展示图
	IsRecycle    field.Int64  // 是否回收站
	CreatedAt    field.Int64
	UpdatedAt    field.Int64
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (s storeProduct) Table(newTableName string) *storeProduct {
	s.storeProductDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProduct) As(alias string) *storeProduct {
	s.storeProductDo.DO = *(s.storeProductDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProduct) updateTableName(table string) *storeProduct {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.MerID = field.NewInt64(table, "mer_id")
	s.Image = field.NewString(table, "image")
	s.SliderImage = field.NewString(table, "slider_image")
	s.StoreName = field.NewString(table, "store_name")
	s.StoreInfo = field.NewString(table, "store_info")
	s.Keyword = field.NewString(table, "keyword")
	s.BarCode = field.NewString(table, "bar_code")
	s.CateID = field.NewString(table, "cate_id")
	s.Price = field.NewField(table, "price")
	s.VipPrice = field.NewField(table, "vip_price")
	s.OtPrice = field.NewField(table, "ot_price")
	s.Postage = field.NewField(table, "postage")
	s.UnitName = field.NewString(table, "unit_name")
	s.Sort = field.NewInt64(table, "sort")
	s.Sales = field.NewInt64(table, "sales")
	s.Stock = field.NewInt64(table, "stock")
	s.IsShow = field.NewInt64(table, "is_show")
	s.IsHot = field.NewInt64(table, "is_hot")
	s.IsBenefit = field.NewInt64(table, "is_benefit")
	s.IsBest = field.NewInt64(table, "is_best")
	s.IsNew = field.NewInt64(table, "is_new")
	s.AddTime = field.NewInt64(table, "add_time")
	s.IsPostage = field.NewInt64(table, "is_postage")
	s.IsDel = field.NewInt64(table, "is_del")
	s.MerUse = field.NewInt64(table, "mer_use")
	s.GiveIntegral = field.NewInt64(table, "give_integral")
	s.Cost = field.NewField(table, "cost")
	s.IsSeckill = field.NewInt64(table, "is_seckill")
	s.IsBargain = field.NewInt64(table, "is_bargain")
	s.IsGood = field.NewInt64(table, "is_good")
	s.IsSub = field.NewInt64(table, "is_sub")
	s.Ficti = field.NewInt64(table, "ficti")
	s.Browse = field.NewInt64(table, "browse")
	s.CodePath = field.NewString(table, "code_path")
	s.SoureLink = field.NewString(table, "soure_link")
	s.VideoLink = field.NewString(table, "video_link")
	s.TempID = field.NewInt64(table, "temp_id")
	s.SpecType = field.NewInt64(table, "spec_type")
	s.Activity = field.NewString(table, "activity")
	s.FlatPattern = field.NewString(table, "flat_pattern")
	s.IsRecycle = field.NewInt64(table, "is_recycle")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProduct) WithContext(ctx context.Context) IStoreProductDo {
	return s.storeProductDo.WithContext(ctx)
}

func (s storeProduct) TableName() string { return s.storeProductDo.TableName() }

func (s storeProduct) Alias() string { return s.storeProductDo.Alias() }

func (s storeProduct) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductDo.Columns(cols...)
}

func (s *storeProduct) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProduct) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 45)
	s.fieldMap["id"] = s.ID
	s.fieldMap["mer_id"] = s.MerID
	s.fieldMap["image"] = s.Image
	s.fieldMap["slider_image"] = s.SliderImage
	s.fieldMap["store_name"] = s.StoreName
	s.fieldMap["store_info"] = s.StoreInfo
	s.fieldMap["keyword"] = s.Keyword
	s.fieldMap["bar_code"] = s.BarCode
	s.fieldMap["cate_id"] = s.CateID
	s.fieldMap["price"] = s.Price
	s.fieldMap["vip_price"] = s.VipPrice
	s.fieldMap["ot_price"] = s.OtPrice
	s.fieldMap["postage"] = s.Postage
	s.fieldMap["unit_name"] = s.UnitName
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["sales"] = s.Sales
	s.fieldMap["stock"] = s.Stock
	s.fieldMap["is_show"] = s.IsShow
	s.fieldMap["is_hot"] = s.IsHot
	s.fieldMap["is_benefit"] = s.IsBenefit
	s.fieldMap["is_best"] = s.IsBest
	s.fieldMap["is_new"] = s.IsNew
	s.fieldMap["add_time"] = s.AddTime
	s.fieldMap["is_postage"] = s.IsPostage
	s.fieldMap["is_del"] = s.IsDel
	s.fieldMap["mer_use"] = s.MerUse
	s.fieldMap["give_integral"] = s.GiveIntegral
	s.fieldMap["cost"] = s.Cost
	s.fieldMap["is_seckill"] = s.IsSeckill
	s.fieldMap["is_bargain"] = s.IsBargain
	s.fieldMap["is_good"] = s.IsGood
	s.fieldMap["is_sub"] = s.IsSub
	s.fieldMap["ficti"] = s.Ficti
	s.fieldMap["browse"] = s.Browse
	s.fieldMap["code_path"] = s.CodePath
	s.fieldMap["soure_link"] = s.SoureLink
	s.fieldMap["video_link"] = s.VideoLink
	s.fieldMap["temp_id"] = s.TempID
	s.fieldMap["spec_type"] = s.SpecType
	s.fieldMap["activity"] = s.Activity
	s.fieldMap["flat_pattern"] = s.FlatPattern
	s.fieldMap["is_recycle"] = s.IsRecycle
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProduct) clone(db *gorm.DB) storeProduct {
	s.storeProductDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProduct) replaceDB(db *gorm.DB) storeProduct {
	s.storeProductDo.ReplaceDB(db)
	return s
}

type storeProductDo struct{ gen.DO }

type IStoreProductDo interface {
	gen.SubQuery
	Debug() IStoreProductDo
	WithContext(ctx context.Context) IStoreProductDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductDo
	WriteDB() IStoreProductDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductDo
	Not(conds ...gen.Condition) IStoreProductDo
	Or(conds ...gen.Condition) IStoreProductDo
	Select(conds ...field.Expr) IStoreProductDo
	Where(conds ...gen.Condition) IStoreProductDo
	Order(conds ...field.Expr) IStoreProductDo
	Distinct(cols ...field.Expr) IStoreProductDo
	Omit(cols ...field.Expr) IStoreProductDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductDo
	Group(cols ...field.Expr) IStoreProductDo
	Having(conds ...gen.Condition) IStoreProductDo
	Limit(limit int) IStoreProductDo
	Offset(offset int) IStoreProductDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductDo
	Unscoped() IStoreProductDo
	Create(values ...*model.StoreProduct) error
	CreateInBatches(values []*model.StoreProduct, batchSize int) error
	Save(values ...*model.StoreProduct) error
	First() (*model.StoreProduct, error)
	Take() (*model.StoreProduct, error)
	Last() (*model.StoreProduct, error)
	Find() ([]*model.StoreProduct, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProduct, err error)
	FindInBatches(result *[]*model.StoreProduct, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProduct) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductDo
	Assign(attrs ...field.AssignExpr) IStoreProductDo
	Joins(fields ...field.RelationField) IStoreProductDo
	Preload(fields ...field.RelationField) IStoreProductDo
	FirstOrInit() (*model.StoreProduct, error)
	FirstOrCreate() (*model.StoreProduct, error)
	FindByPage(offset int, limit int) (result []*model.StoreProduct, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductDo) Debug() IStoreProductDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductDo) WithContext(ctx context.Context) IStoreProductDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductDo) ReadDB() IStoreProductDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductDo) WriteDB() IStoreProductDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductDo) Session(config *gorm.Session) IStoreProductDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductDo) Clauses(conds ...clause.Expression) IStoreProductDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductDo) Returning(value interface{}, columns ...string) IStoreProductDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductDo) Not(conds ...gen.Condition) IStoreProductDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductDo) Or(conds ...gen.Condition) IStoreProductDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductDo) Select(conds ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductDo) Where(conds ...gen.Condition) IStoreProductDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductDo) Order(conds ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductDo) Distinct(cols ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductDo) Omit(cols ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductDo) Group(cols ...field.Expr) IStoreProductDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductDo) Having(conds ...gen.Condition) IStoreProductDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductDo) Limit(limit int) IStoreProductDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductDo) Offset(offset int) IStoreProductDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductDo) Unscoped() IStoreProductDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductDo) Create(values ...*model.StoreProduct) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductDo) CreateInBatches(values []*model.StoreProduct, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductDo) Save(values ...*model.StoreProduct) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductDo) First() (*model.StoreProduct, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProduct), nil
	}
}

func (s storeProductDo) Take() (*model.StoreProduct, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProduct), nil
	}
}

func (s storeProductDo) Last() (*model.StoreProduct, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProduct), nil
	}
}

func (s storeProductDo) Find() ([]*model.StoreProduct, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProduct), err
}

func (s storeProductDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProduct, err error) {
	buf := make([]*model.StoreProduct, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductDo) FindInBatches(result *[]*model.StoreProduct, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductDo) Attrs(attrs ...field.AssignExpr) IStoreProductDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductDo) Assign(attrs ...field.AssignExpr) IStoreProductDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductDo) Joins(fields ...field.RelationField) IStoreProductDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductDo) Preload(fields ...field.RelationField) IStoreProductDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductDo) FirstOrInit() (*model.StoreProduct, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProduct), nil
	}
}

func (s storeProductDo) FirstOrCreate() (*model.StoreProduct, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProduct), nil
	}
}

func (s storeProductDo) FindByPage(offset int, limit int) (result []*model.StoreProduct, count int64, err error) {
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

func (s storeProductDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductDo) Delete(models ...*model.StoreProduct) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductDo) withDO(do gen.Dao) *storeProductDo {
	s.DO = *do.(*gen.DO)
	return s
}