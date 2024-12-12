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

func newUserAddress(db *gorm.DB, opts ...gen.DOOption) userAddress {
	_userAddress := userAddress{}

	_userAddress.userAddressDo.UseDB(db, opts...)
	_userAddress.userAddressDo.UseModel(&model.UserAddress{})

	tableName := _userAddress.userAddressDo.TableName()
	_userAddress.ALL = field.NewAsterisk(tableName)
	_userAddress.ID = field.NewInt64(tableName, "id")
	_userAddress.UID = field.NewInt64(tableName, "uid")
	_userAddress.RealName = field.NewString(tableName, "real_name")
	_userAddress.Phone = field.NewString(tableName, "phone")
	_userAddress.Province = field.NewString(tableName, "province")
	_userAddress.City = field.NewString(tableName, "city")
	_userAddress.CityID = field.NewInt64(tableName, "city_id")
	_userAddress.District = field.NewString(tableName, "district")
	_userAddress.Detail = field.NewString(tableName, "detail")
	_userAddress.PostCode = field.NewInt64(tableName, "post_code")
	_userAddress.Longitude = field.NewString(tableName, "longitude")
	_userAddress.Latitude = field.NewString(tableName, "latitude")
	_userAddress.IsDefault = field.NewInt64(tableName, "is_default")
	_userAddress.IsDel = field.NewInt64(tableName, "is_del")
	_userAddress.CreatedAt = field.NewInt64(tableName, "created_at")
	_userAddress.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_userAddress.DeletedAt = field.NewField(tableName, "deleted_at")

	_userAddress.fillFieldMap()

	return _userAddress
}

// userAddress 用户地址表
type userAddress struct {
	userAddressDo userAddressDo

	ALL       field.Asterisk
	ID        field.Int64  // 用户地址id
	UID       field.Int64  // 用户id
	RealName  field.String // 收货人姓名
	Phone     field.String // 收货人电话
	Province  field.String // 收货人所在省
	City      field.String // 收货人所在市
	CityID    field.Int64  // 城市id
	District  field.String // 收货人所在区
	Detail    field.String // 收货人详细地址
	PostCode  field.Int64  // 邮编
	Longitude field.String // 经度
	Latitude  field.String // 纬度
	IsDefault field.Int64  // 是否默认
	IsDel     field.Int64  // 是否删除
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (u userAddress) Table(newTableName string) *userAddress {
	u.userAddressDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userAddress) As(alias string) *userAddress {
	u.userAddressDo.DO = *(u.userAddressDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userAddress) updateTableName(table string) *userAddress {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewInt64(table, "uid")
	u.RealName = field.NewString(table, "real_name")
	u.Phone = field.NewString(table, "phone")
	u.Province = field.NewString(table, "province")
	u.City = field.NewString(table, "city")
	u.CityID = field.NewInt64(table, "city_id")
	u.District = field.NewString(table, "district")
	u.Detail = field.NewString(table, "detail")
	u.PostCode = field.NewInt64(table, "post_code")
	u.Longitude = field.NewString(table, "longitude")
	u.Latitude = field.NewString(table, "latitude")
	u.IsDefault = field.NewInt64(table, "is_default")
	u.IsDel = field.NewInt64(table, "is_del")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userAddress) WithContext(ctx context.Context) IUserAddressDo {
	return u.userAddressDo.WithContext(ctx)
}

func (u userAddress) TableName() string { return u.userAddressDo.TableName() }

func (u userAddress) Alias() string { return u.userAddressDo.Alias() }

func (u userAddress) Columns(cols ...field.Expr) gen.Columns { return u.userAddressDo.Columns(cols...) }

func (u *userAddress) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userAddress) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 17)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["real_name"] = u.RealName
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["province"] = u.Province
	u.fieldMap["city"] = u.City
	u.fieldMap["city_id"] = u.CityID
	u.fieldMap["district"] = u.District
	u.fieldMap["detail"] = u.Detail
	u.fieldMap["post_code"] = u.PostCode
	u.fieldMap["longitude"] = u.Longitude
	u.fieldMap["latitude"] = u.Latitude
	u.fieldMap["is_default"] = u.IsDefault
	u.fieldMap["is_del"] = u.IsDel
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userAddress) clone(db *gorm.DB) userAddress {
	u.userAddressDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userAddress) replaceDB(db *gorm.DB) userAddress {
	u.userAddressDo.ReplaceDB(db)
	return u
}

type userAddressDo struct{ gen.DO }

type IUserAddressDo interface {
	gen.SubQuery
	Debug() IUserAddressDo
	WithContext(ctx context.Context) IUserAddressDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserAddressDo
	WriteDB() IUserAddressDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserAddressDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserAddressDo
	Not(conds ...gen.Condition) IUserAddressDo
	Or(conds ...gen.Condition) IUserAddressDo
	Select(conds ...field.Expr) IUserAddressDo
	Where(conds ...gen.Condition) IUserAddressDo
	Order(conds ...field.Expr) IUserAddressDo
	Distinct(cols ...field.Expr) IUserAddressDo
	Omit(cols ...field.Expr) IUserAddressDo
	Join(table schema.Tabler, on ...field.Expr) IUserAddressDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserAddressDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserAddressDo
	Group(cols ...field.Expr) IUserAddressDo
	Having(conds ...gen.Condition) IUserAddressDo
	Limit(limit int) IUserAddressDo
	Offset(offset int) IUserAddressDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAddressDo
	Unscoped() IUserAddressDo
	Create(values ...*model.UserAddress) error
	CreateInBatches(values []*model.UserAddress, batchSize int) error
	Save(values ...*model.UserAddress) error
	First() (*model.UserAddress, error)
	Take() (*model.UserAddress, error)
	Last() (*model.UserAddress, error)
	Find() ([]*model.UserAddress, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAddress, err error)
	FindInBatches(result *[]*model.UserAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserAddress) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserAddressDo
	Assign(attrs ...field.AssignExpr) IUserAddressDo
	Joins(fields ...field.RelationField) IUserAddressDo
	Preload(fields ...field.RelationField) IUserAddressDo
	FirstOrInit() (*model.UserAddress, error)
	FirstOrCreate() (*model.UserAddress, error)
	FindByPage(offset int, limit int) (result []*model.UserAddress, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserAddressDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userAddressDo) Debug() IUserAddressDo {
	return u.withDO(u.DO.Debug())
}

func (u userAddressDo) WithContext(ctx context.Context) IUserAddressDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userAddressDo) ReadDB() IUserAddressDo {
	return u.Clauses(dbresolver.Read)
}

func (u userAddressDo) WriteDB() IUserAddressDo {
	return u.Clauses(dbresolver.Write)
}

func (u userAddressDo) Session(config *gorm.Session) IUserAddressDo {
	return u.withDO(u.DO.Session(config))
}

func (u userAddressDo) Clauses(conds ...clause.Expression) IUserAddressDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userAddressDo) Returning(value interface{}, columns ...string) IUserAddressDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userAddressDo) Not(conds ...gen.Condition) IUserAddressDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userAddressDo) Or(conds ...gen.Condition) IUserAddressDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userAddressDo) Select(conds ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userAddressDo) Where(conds ...gen.Condition) IUserAddressDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userAddressDo) Order(conds ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userAddressDo) Distinct(cols ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userAddressDo) Omit(cols ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userAddressDo) Join(table schema.Tabler, on ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userAddressDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userAddressDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userAddressDo) Group(cols ...field.Expr) IUserAddressDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userAddressDo) Having(conds ...gen.Condition) IUserAddressDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userAddressDo) Limit(limit int) IUserAddressDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userAddressDo) Offset(offset int) IUserAddressDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userAddressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAddressDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userAddressDo) Unscoped() IUserAddressDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userAddressDo) Create(values ...*model.UserAddress) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userAddressDo) CreateInBatches(values []*model.UserAddress, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userAddressDo) Save(values ...*model.UserAddress) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userAddressDo) First() (*model.UserAddress, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAddress), nil
	}
}

func (u userAddressDo) Take() (*model.UserAddress, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAddress), nil
	}
}

func (u userAddressDo) Last() (*model.UserAddress, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAddress), nil
	}
}

func (u userAddressDo) Find() ([]*model.UserAddress, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserAddress), err
}

func (u userAddressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAddress, err error) {
	buf := make([]*model.UserAddress, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userAddressDo) FindInBatches(result *[]*model.UserAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userAddressDo) Attrs(attrs ...field.AssignExpr) IUserAddressDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userAddressDo) Assign(attrs ...field.AssignExpr) IUserAddressDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userAddressDo) Joins(fields ...field.RelationField) IUserAddressDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userAddressDo) Preload(fields ...field.RelationField) IUserAddressDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userAddressDo) FirstOrInit() (*model.UserAddress, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAddress), nil
	}
}

func (u userAddressDo) FirstOrCreate() (*model.UserAddress, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAddress), nil
	}
}

func (u userAddressDo) FindByPage(offset int, limit int) (result []*model.UserAddress, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userAddressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userAddressDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userAddressDo) Delete(models ...*model.UserAddress) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userAddressDo) withDO(do gen.Dao) *userAddressDo {
	u.DO = *do.(*gen.DO)
	return u
}