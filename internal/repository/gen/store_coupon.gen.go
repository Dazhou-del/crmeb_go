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

func newStoreCoupon(db *gorm.DB, opts ...gen.DOOption) storeCoupon {
	_storeCoupon := storeCoupon{}

	_storeCoupon.storeCouponDo.UseDB(db, opts...)
	_storeCoupon.storeCouponDo.UseModel(&model.StoreCoupon{})

	tableName := _storeCoupon.storeCouponDo.TableName()
	_storeCoupon.ALL = field.NewAsterisk(tableName)
	_storeCoupon.ID = field.NewInt64(tableName, "id")
	_storeCoupon.Name = field.NewString(tableName, "name")
	_storeCoupon.Money = field.NewString(tableName, "money")
	_storeCoupon.IsLimited = field.NewInt64(tableName, "is_limited")
	_storeCoupon.Total = field.NewInt64(tableName, "total")
	_storeCoupon.LastTotal = field.NewInt64(tableName, "last_total")
	_storeCoupon.UseType = field.NewInt64(tableName, "use_type")
	_storeCoupon.PrimaryKey = field.NewString(tableName, "primary_key")
	_storeCoupon.MinPrice = field.NewString(tableName, "min_price")
	_storeCoupon.ReceiveStartTime = field.NewInt64(tableName, "receive_start_time")
	_storeCoupon.ReceiveEndTime = field.NewInt64(tableName, "receive_end_time")
	_storeCoupon.IsFixedTime = field.NewInt64(tableName, "is_fixed_time")
	_storeCoupon.UseStartTime = field.NewInt64(tableName, "use_start_time")
	_storeCoupon.UseEndTime = field.NewInt64(tableName, "use_end_time")
	_storeCoupon.Day = field.NewInt64(tableName, "day")
	_storeCoupon.Type = field.NewInt64(tableName, "type")
	_storeCoupon.Sort = field.NewInt64(tableName, "sort")
	_storeCoupon.Status = field.NewInt64(tableName, "status")
	_storeCoupon.IsDel = field.NewInt64(tableName, "is_del")
	_storeCoupon.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeCoupon.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeCoupon.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeCoupon.fillFieldMap()

	return _storeCoupon
}

// storeCoupon 优惠券表
type storeCoupon struct {
	storeCouponDo storeCouponDo

	ALL              field.Asterisk
	ID               field.Int64  // 优惠券表ID
	Name             field.String // 优惠券名称
	Money            field.String // 兑换的优惠券面值
	IsLimited        field.Int64  // 是否限量, 默认0 不限量， 1限量
	Total            field.Int64  // 发放总数
	LastTotal        field.Int64  // 剩余数量
	UseType          field.Int64  // 使用类型 1 全场通用, 2 商品券, 3 品类券
	PrimaryKey       field.String // 所属商品id / 分类id
	MinPrice         field.String // 最低消费，0代表不限制
	ReceiveStartTime field.Int64
	ReceiveEndTime   field.Int64
	IsFixedTime      field.Int64 // 是否固定使用时间, 默认0 否， 1是
	UseStartTime     field.Int64
	UseEndTime       field.Int64
	Day              field.Int64 // 天数
	Type             field.Int64 // 优惠券类型 1 手动领取, 2 新人券, 3 赠送券
	Sort             field.Int64 // 排序
	Status           field.Int64 // 状态（0：关闭，1：开启）
	IsDel            field.Int64 // 是否删除 状态（0：否，1：是）
	CreatedAt        field.Int64
	UpdatedAt        field.Int64
	DeletedAt        field.Field

	fieldMap map[string]field.Expr
}

func (s storeCoupon) Table(newTableName string) *storeCoupon {
	s.storeCouponDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeCoupon) As(alias string) *storeCoupon {
	s.storeCouponDo.DO = *(s.storeCouponDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeCoupon) updateTableName(table string) *storeCoupon {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Name = field.NewString(table, "name")
	s.Money = field.NewString(table, "money")
	s.IsLimited = field.NewInt64(table, "is_limited")
	s.Total = field.NewInt64(table, "total")
	s.LastTotal = field.NewInt64(table, "last_total")
	s.UseType = field.NewInt64(table, "use_type")
	s.PrimaryKey = field.NewString(table, "primary_key")
	s.MinPrice = field.NewString(table, "min_price")
	s.ReceiveStartTime = field.NewInt64(table, "receive_start_time")
	s.ReceiveEndTime = field.NewInt64(table, "receive_end_time")
	s.IsFixedTime = field.NewInt64(table, "is_fixed_time")
	s.UseStartTime = field.NewInt64(table, "use_start_time")
	s.UseEndTime = field.NewInt64(table, "use_end_time")
	s.Day = field.NewInt64(table, "day")
	s.Type = field.NewInt64(table, "type")
	s.Sort = field.NewInt64(table, "sort")
	s.Status = field.NewInt64(table, "status")
	s.IsDel = field.NewInt64(table, "is_del")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeCoupon) WithContext(ctx context.Context) IStoreCouponDo {
	return s.storeCouponDo.WithContext(ctx)
}

func (s storeCoupon) TableName() string { return s.storeCouponDo.TableName() }

func (s storeCoupon) Alias() string { return s.storeCouponDo.Alias() }

func (s storeCoupon) Columns(cols ...field.Expr) gen.Columns { return s.storeCouponDo.Columns(cols...) }

func (s *storeCoupon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeCoupon) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 22)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["money"] = s.Money
	s.fieldMap["is_limited"] = s.IsLimited
	s.fieldMap["total"] = s.Total
	s.fieldMap["last_total"] = s.LastTotal
	s.fieldMap["use_type"] = s.UseType
	s.fieldMap["primary_key"] = s.PrimaryKey
	s.fieldMap["min_price"] = s.MinPrice
	s.fieldMap["receive_start_time"] = s.ReceiveStartTime
	s.fieldMap["receive_end_time"] = s.ReceiveEndTime
	s.fieldMap["is_fixed_time"] = s.IsFixedTime
	s.fieldMap["use_start_time"] = s.UseStartTime
	s.fieldMap["use_end_time"] = s.UseEndTime
	s.fieldMap["day"] = s.Day
	s.fieldMap["type"] = s.Type
	s.fieldMap["sort"] = s.Sort
	s.fieldMap["status"] = s.Status
	s.fieldMap["is_del"] = s.IsDel
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeCoupon) clone(db *gorm.DB) storeCoupon {
	s.storeCouponDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeCoupon) replaceDB(db *gorm.DB) storeCoupon {
	s.storeCouponDo.ReplaceDB(db)
	return s
}

type storeCouponDo struct{ gen.DO }

type IStoreCouponDo interface {
	gen.SubQuery
	Debug() IStoreCouponDo
	WithContext(ctx context.Context) IStoreCouponDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreCouponDo
	WriteDB() IStoreCouponDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreCouponDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreCouponDo
	Not(conds ...gen.Condition) IStoreCouponDo
	Or(conds ...gen.Condition) IStoreCouponDo
	Select(conds ...field.Expr) IStoreCouponDo
	Where(conds ...gen.Condition) IStoreCouponDo
	Order(conds ...field.Expr) IStoreCouponDo
	Distinct(cols ...field.Expr) IStoreCouponDo
	Omit(cols ...field.Expr) IStoreCouponDo
	Join(table schema.Tabler, on ...field.Expr) IStoreCouponDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreCouponDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreCouponDo
	Group(cols ...field.Expr) IStoreCouponDo
	Having(conds ...gen.Condition) IStoreCouponDo
	Limit(limit int) IStoreCouponDo
	Offset(offset int) IStoreCouponDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreCouponDo
	Unscoped() IStoreCouponDo
	Create(values ...*model.StoreCoupon) error
	CreateInBatches(values []*model.StoreCoupon, batchSize int) error
	Save(values ...*model.StoreCoupon) error
	First() (*model.StoreCoupon, error)
	Take() (*model.StoreCoupon, error)
	Last() (*model.StoreCoupon, error)
	Find() ([]*model.StoreCoupon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreCoupon, err error)
	FindInBatches(result *[]*model.StoreCoupon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreCoupon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreCouponDo
	Assign(attrs ...field.AssignExpr) IStoreCouponDo
	Joins(fields ...field.RelationField) IStoreCouponDo
	Preload(fields ...field.RelationField) IStoreCouponDo
	FirstOrInit() (*model.StoreCoupon, error)
	FirstOrCreate() (*model.StoreCoupon, error)
	FindByPage(offset int, limit int) (result []*model.StoreCoupon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreCouponDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeCouponDo) Debug() IStoreCouponDo {
	return s.withDO(s.DO.Debug())
}

func (s storeCouponDo) WithContext(ctx context.Context) IStoreCouponDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeCouponDo) ReadDB() IStoreCouponDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeCouponDo) WriteDB() IStoreCouponDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeCouponDo) Session(config *gorm.Session) IStoreCouponDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeCouponDo) Clauses(conds ...clause.Expression) IStoreCouponDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeCouponDo) Returning(value interface{}, columns ...string) IStoreCouponDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeCouponDo) Not(conds ...gen.Condition) IStoreCouponDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeCouponDo) Or(conds ...gen.Condition) IStoreCouponDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeCouponDo) Select(conds ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeCouponDo) Where(conds ...gen.Condition) IStoreCouponDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeCouponDo) Order(conds ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeCouponDo) Distinct(cols ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeCouponDo) Omit(cols ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeCouponDo) Join(table schema.Tabler, on ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeCouponDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeCouponDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeCouponDo) Group(cols ...field.Expr) IStoreCouponDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeCouponDo) Having(conds ...gen.Condition) IStoreCouponDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeCouponDo) Limit(limit int) IStoreCouponDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeCouponDo) Offset(offset int) IStoreCouponDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeCouponDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreCouponDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeCouponDo) Unscoped() IStoreCouponDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeCouponDo) Create(values ...*model.StoreCoupon) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeCouponDo) CreateInBatches(values []*model.StoreCoupon, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeCouponDo) Save(values ...*model.StoreCoupon) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeCouponDo) First() (*model.StoreCoupon, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCoupon), nil
	}
}

func (s storeCouponDo) Take() (*model.StoreCoupon, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCoupon), nil
	}
}

func (s storeCouponDo) Last() (*model.StoreCoupon, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCoupon), nil
	}
}

func (s storeCouponDo) Find() ([]*model.StoreCoupon, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreCoupon), err
}

func (s storeCouponDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreCoupon, err error) {
	buf := make([]*model.StoreCoupon, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeCouponDo) FindInBatches(result *[]*model.StoreCoupon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeCouponDo) Attrs(attrs ...field.AssignExpr) IStoreCouponDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeCouponDo) Assign(attrs ...field.AssignExpr) IStoreCouponDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeCouponDo) Joins(fields ...field.RelationField) IStoreCouponDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeCouponDo) Preload(fields ...field.RelationField) IStoreCouponDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeCouponDo) FirstOrInit() (*model.StoreCoupon, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCoupon), nil
	}
}

func (s storeCouponDo) FirstOrCreate() (*model.StoreCoupon, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreCoupon), nil
	}
}

func (s storeCouponDo) FindByPage(offset int, limit int) (result []*model.StoreCoupon, count int64, err error) {
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

func (s storeCouponDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeCouponDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeCouponDo) Delete(models ...*model.StoreCoupon) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeCouponDo) withDO(do gen.Dao) *storeCouponDo {
	s.DO = *do.(*gen.DO)
	return s
}
