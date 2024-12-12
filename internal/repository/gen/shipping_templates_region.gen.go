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

func newShippingTemplatesRegion(db *gorm.DB, opts ...gen.DOOption) shippingTemplatesRegion {
	_shippingTemplatesRegion := shippingTemplatesRegion{}

	_shippingTemplatesRegion.shippingTemplatesRegionDo.UseDB(db, opts...)
	_shippingTemplatesRegion.shippingTemplatesRegionDo.UseModel(&model.ShippingTemplatesRegion{})

	tableName := _shippingTemplatesRegion.shippingTemplatesRegionDo.TableName()
	_shippingTemplatesRegion.ALL = field.NewAsterisk(tableName)
	_shippingTemplatesRegion.ID = field.NewInt64(tableName, "id")
	_shippingTemplatesRegion.TempID = field.NewInt64(tableName, "temp_id")
	_shippingTemplatesRegion.CityID = field.NewInt64(tableName, "city_id")
	_shippingTemplatesRegion.Title = field.NewString(tableName, "title")
	_shippingTemplatesRegion.First = field.NewField(tableName, "first")
	_shippingTemplatesRegion.FirstPrice = field.NewField(tableName, "first_price")
	_shippingTemplatesRegion.Renewal = field.NewField(tableName, "renewal")
	_shippingTemplatesRegion.RenewalPrice = field.NewField(tableName, "renewal_price")
	_shippingTemplatesRegion.Type = field.NewInt64(tableName, "type")
	_shippingTemplatesRegion.Uniqid = field.NewString(tableName, "uniqid")
	_shippingTemplatesRegion.Status = field.NewInt64(tableName, "status")
	_shippingTemplatesRegion.CreatedAt = field.NewInt64(tableName, "created_at")
	_shippingTemplatesRegion.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_shippingTemplatesRegion.DeletedAt = field.NewField(tableName, "deleted_at")

	_shippingTemplatesRegion.fillFieldMap()

	return _shippingTemplatesRegion
}

// shippingTemplatesRegion 运费模板指定区域费用
type shippingTemplatesRegion struct {
	shippingTemplatesRegionDo shippingTemplatesRegionDo

	ALL          field.Asterisk
	ID           field.Int64  // 编号
	TempID       field.Int64  // 模板ID
	CityID       field.Int64  // 城市ID
	Title        field.String // 描述
	First        field.Field  // 首件
	FirstPrice   field.Field  // 首件运费
	Renewal      field.Field  // 续件
	RenewalPrice field.Field  // 续件运费
	Type         field.Int64  // 计费方式 1按件数 2按重量 3按体积
	Uniqid       field.String // 分组唯一值
	Status       field.Int64  // 是否无效
	CreatedAt    field.Int64
	UpdatedAt    field.Int64
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (s shippingTemplatesRegion) Table(newTableName string) *shippingTemplatesRegion {
	s.shippingTemplatesRegionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s shippingTemplatesRegion) As(alias string) *shippingTemplatesRegion {
	s.shippingTemplatesRegionDo.DO = *(s.shippingTemplatesRegionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *shippingTemplatesRegion) updateTableName(table string) *shippingTemplatesRegion {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.TempID = field.NewInt64(table, "temp_id")
	s.CityID = field.NewInt64(table, "city_id")
	s.Title = field.NewString(table, "title")
	s.First = field.NewField(table, "first")
	s.FirstPrice = field.NewField(table, "first_price")
	s.Renewal = field.NewField(table, "renewal")
	s.RenewalPrice = field.NewField(table, "renewal_price")
	s.Type = field.NewInt64(table, "type")
	s.Uniqid = field.NewString(table, "uniqid")
	s.Status = field.NewInt64(table, "status")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *shippingTemplatesRegion) WithContext(ctx context.Context) IShippingTemplatesRegionDo {
	return s.shippingTemplatesRegionDo.WithContext(ctx)
}

func (s shippingTemplatesRegion) TableName() string { return s.shippingTemplatesRegionDo.TableName() }

func (s shippingTemplatesRegion) Alias() string { return s.shippingTemplatesRegionDo.Alias() }

func (s shippingTemplatesRegion) Columns(cols ...field.Expr) gen.Columns {
	return s.shippingTemplatesRegionDo.Columns(cols...)
}

func (s *shippingTemplatesRegion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *shippingTemplatesRegion) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["temp_id"] = s.TempID
	s.fieldMap["city_id"] = s.CityID
	s.fieldMap["title"] = s.Title
	s.fieldMap["first"] = s.First
	s.fieldMap["first_price"] = s.FirstPrice
	s.fieldMap["renewal"] = s.Renewal
	s.fieldMap["renewal_price"] = s.RenewalPrice
	s.fieldMap["type"] = s.Type
	s.fieldMap["uniqid"] = s.Uniqid
	s.fieldMap["status"] = s.Status
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s shippingTemplatesRegion) clone(db *gorm.DB) shippingTemplatesRegion {
	s.shippingTemplatesRegionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s shippingTemplatesRegion) replaceDB(db *gorm.DB) shippingTemplatesRegion {
	s.shippingTemplatesRegionDo.ReplaceDB(db)
	return s
}

type shippingTemplatesRegionDo struct{ gen.DO }

type IShippingTemplatesRegionDo interface {
	gen.SubQuery
	Debug() IShippingTemplatesRegionDo
	WithContext(ctx context.Context) IShippingTemplatesRegionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IShippingTemplatesRegionDo
	WriteDB() IShippingTemplatesRegionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IShippingTemplatesRegionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IShippingTemplatesRegionDo
	Not(conds ...gen.Condition) IShippingTemplatesRegionDo
	Or(conds ...gen.Condition) IShippingTemplatesRegionDo
	Select(conds ...field.Expr) IShippingTemplatesRegionDo
	Where(conds ...gen.Condition) IShippingTemplatesRegionDo
	Order(conds ...field.Expr) IShippingTemplatesRegionDo
	Distinct(cols ...field.Expr) IShippingTemplatesRegionDo
	Omit(cols ...field.Expr) IShippingTemplatesRegionDo
	Join(table schema.Tabler, on ...field.Expr) IShippingTemplatesRegionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IShippingTemplatesRegionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IShippingTemplatesRegionDo
	Group(cols ...field.Expr) IShippingTemplatesRegionDo
	Having(conds ...gen.Condition) IShippingTemplatesRegionDo
	Limit(limit int) IShippingTemplatesRegionDo
	Offset(offset int) IShippingTemplatesRegionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IShippingTemplatesRegionDo
	Unscoped() IShippingTemplatesRegionDo
	Create(values ...*model.ShippingTemplatesRegion) error
	CreateInBatches(values []*model.ShippingTemplatesRegion, batchSize int) error
	Save(values ...*model.ShippingTemplatesRegion) error
	First() (*model.ShippingTemplatesRegion, error)
	Take() (*model.ShippingTemplatesRegion, error)
	Last() (*model.ShippingTemplatesRegion, error)
	Find() ([]*model.ShippingTemplatesRegion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ShippingTemplatesRegion, err error)
	FindInBatches(result *[]*model.ShippingTemplatesRegion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ShippingTemplatesRegion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IShippingTemplatesRegionDo
	Assign(attrs ...field.AssignExpr) IShippingTemplatesRegionDo
	Joins(fields ...field.RelationField) IShippingTemplatesRegionDo
	Preload(fields ...field.RelationField) IShippingTemplatesRegionDo
	FirstOrInit() (*model.ShippingTemplatesRegion, error)
	FirstOrCreate() (*model.ShippingTemplatesRegion, error)
	FindByPage(offset int, limit int) (result []*model.ShippingTemplatesRegion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IShippingTemplatesRegionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s shippingTemplatesRegionDo) Debug() IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Debug())
}

func (s shippingTemplatesRegionDo) WithContext(ctx context.Context) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s shippingTemplatesRegionDo) ReadDB() IShippingTemplatesRegionDo {
	return s.Clauses(dbresolver.Read)
}

func (s shippingTemplatesRegionDo) WriteDB() IShippingTemplatesRegionDo {
	return s.Clauses(dbresolver.Write)
}

func (s shippingTemplatesRegionDo) Session(config *gorm.Session) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Session(config))
}

func (s shippingTemplatesRegionDo) Clauses(conds ...clause.Expression) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s shippingTemplatesRegionDo) Returning(value interface{}, columns ...string) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s shippingTemplatesRegionDo) Not(conds ...gen.Condition) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s shippingTemplatesRegionDo) Or(conds ...gen.Condition) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s shippingTemplatesRegionDo) Select(conds ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s shippingTemplatesRegionDo) Where(conds ...gen.Condition) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s shippingTemplatesRegionDo) Order(conds ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s shippingTemplatesRegionDo) Distinct(cols ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s shippingTemplatesRegionDo) Omit(cols ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s shippingTemplatesRegionDo) Join(table schema.Tabler, on ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s shippingTemplatesRegionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s shippingTemplatesRegionDo) RightJoin(table schema.Tabler, on ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s shippingTemplatesRegionDo) Group(cols ...field.Expr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s shippingTemplatesRegionDo) Having(conds ...gen.Condition) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s shippingTemplatesRegionDo) Limit(limit int) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s shippingTemplatesRegionDo) Offset(offset int) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s shippingTemplatesRegionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s shippingTemplatesRegionDo) Unscoped() IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s shippingTemplatesRegionDo) Create(values ...*model.ShippingTemplatesRegion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s shippingTemplatesRegionDo) CreateInBatches(values []*model.ShippingTemplatesRegion, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s shippingTemplatesRegionDo) Save(values ...*model.ShippingTemplatesRegion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s shippingTemplatesRegionDo) First() (*model.ShippingTemplatesRegion, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingTemplatesRegion), nil
	}
}

func (s shippingTemplatesRegionDo) Take() (*model.ShippingTemplatesRegion, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingTemplatesRegion), nil
	}
}

func (s shippingTemplatesRegionDo) Last() (*model.ShippingTemplatesRegion, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingTemplatesRegion), nil
	}
}

func (s shippingTemplatesRegionDo) Find() ([]*model.ShippingTemplatesRegion, error) {
	result, err := s.DO.Find()
	return result.([]*model.ShippingTemplatesRegion), err
}

func (s shippingTemplatesRegionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ShippingTemplatesRegion, err error) {
	buf := make([]*model.ShippingTemplatesRegion, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s shippingTemplatesRegionDo) FindInBatches(result *[]*model.ShippingTemplatesRegion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s shippingTemplatesRegionDo) Attrs(attrs ...field.AssignExpr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s shippingTemplatesRegionDo) Assign(attrs ...field.AssignExpr) IShippingTemplatesRegionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s shippingTemplatesRegionDo) Joins(fields ...field.RelationField) IShippingTemplatesRegionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s shippingTemplatesRegionDo) Preload(fields ...field.RelationField) IShippingTemplatesRegionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s shippingTemplatesRegionDo) FirstOrInit() (*model.ShippingTemplatesRegion, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingTemplatesRegion), nil
	}
}

func (s shippingTemplatesRegionDo) FirstOrCreate() (*model.ShippingTemplatesRegion, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShippingTemplatesRegion), nil
	}
}

func (s shippingTemplatesRegionDo) FindByPage(offset int, limit int) (result []*model.ShippingTemplatesRegion, count int64, err error) {
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

func (s shippingTemplatesRegionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s shippingTemplatesRegionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s shippingTemplatesRegionDo) Delete(models ...*model.ShippingTemplatesRegion) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *shippingTemplatesRegionDo) withDO(do gen.Dao) *shippingTemplatesRegionDo {
	s.DO = *do.(*gen.DO)
	return s
}