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

func newEbUser(db *gorm.DB, opts ...gen.DOOption) ebUser {
	_ebUser := ebUser{}

	_ebUser.ebUserDo.UseDB(db, opts...)
	_ebUser.ebUserDo.UseModel(&model.EbUser{})

	tableName := _ebUser.ebUserDo.TableName()
	_ebUser.ALL = field.NewAsterisk(tableName)
	_ebUser.UID = field.NewInt32(tableName, "uid")
	_ebUser.Account = field.NewString(tableName, "account")
	_ebUser.Pwd = field.NewString(tableName, "pwd")
	_ebUser.RealName = field.NewString(tableName, "real_name")
	_ebUser.Birthday = field.NewString(tableName, "birthday")
	_ebUser.CardID = field.NewString(tableName, "card_id")
	_ebUser.Mark = field.NewString(tableName, "mark")
	_ebUser.PartnerID = field.NewInt32(tableName, "partner_id")
	_ebUser.GroupID = field.NewString(tableName, "group_id")
	_ebUser.TagID = field.NewString(tableName, "tag_id")
	_ebUser.Nickname = field.NewString(tableName, "nickname")
	_ebUser.Avatar = field.NewString(tableName, "avatar")
	_ebUser.Phone = field.NewString(tableName, "phone")
	_ebUser.AddIP = field.NewString(tableName, "add_ip")
	_ebUser.LastIP = field.NewString(tableName, "last_ip")
	_ebUser.NowMoney = field.NewFloat64(tableName, "now_money")
	_ebUser.BrokeragePrice = field.NewFloat64(tableName, "brokerage_price")
	_ebUser.Integral = field.NewInt32(tableName, "integral")
	_ebUser.Experience = field.NewInt32(tableName, "experience")
	_ebUser.SignNum = field.NewInt32(tableName, "sign_num")
	_ebUser.Status = field.NewBool(tableName, "status")
	_ebUser.Level = field.NewInt32(tableName, "level")
	_ebUser.SpreadUID = field.NewInt32(tableName, "spread_uid")
	_ebUser.SpreadTime = field.NewTime(tableName, "spread_time")
	_ebUser.UserType = field.NewString(tableName, "user_type")
	_ebUser.IsPromoter = field.NewInt32(tableName, "is_promoter")
	_ebUser.PayCount = field.NewInt32(tableName, "pay_count")
	_ebUser.SpreadCount = field.NewInt32(tableName, "spread_count")
	_ebUser.Addres = field.NewString(tableName, "addres")
	_ebUser.Adminid = field.NewInt32(tableName, "adminid")
	_ebUser.LoginType = field.NewString(tableName, "login_type")
	_ebUser.CreateTime = field.NewTime(tableName, "create_time")
	_ebUser.UpdateTime = field.NewTime(tableName, "update_time")
	_ebUser.LastLoginTime = field.NewTime(tableName, "last_login_time")
	_ebUser.CleanTime = field.NewTime(tableName, "clean_time")
	_ebUser.Path = field.NewString(tableName, "path")
	_ebUser.Subscribe = field.NewInt32(tableName, "subscribe")
	_ebUser.SubscribeTime = field.NewTime(tableName, "subscribe_time")
	_ebUser.Sex = field.NewBool(tableName, "sex")
	_ebUser.Country = field.NewString(tableName, "country")
	_ebUser.PromoterTime = field.NewTime(tableName, "promoter_time")

	_ebUser.fillFieldMap()

	return _ebUser
}

// ebUser 用户表
type ebUser struct {
	ebUserDo ebUserDo

	ALL            field.Asterisk
	UID            field.Int32   // 用户id
	Account        field.String  // 用户账号
	Pwd            field.String  // 用户密码
	RealName       field.String  // 真实姓名
	Birthday       field.String  // 生日
	CardID         field.String  // 身份证号码
	Mark           field.String  // 用户备注
	PartnerID      field.Int32   // 合伙人id
	GroupID        field.String  // 用户分组id
	TagID          field.String  // 标签id
	Nickname       field.String  // 用户昵称
	Avatar         field.String  // 用户头像
	Phone          field.String  // 手机号码
	AddIP          field.String  // 添加ip
	LastIP         field.String  // 最后一次登录ip
	NowMoney       field.Float64 // 用户余额
	BrokeragePrice field.Float64 // 佣金金额
	Integral       field.Int32   // 用户剩余积分
	Experience     field.Int32   // 用户剩余经验
	SignNum        field.Int32   // 连续签到天数
	Status         field.Bool    // 1为正常，0为禁止
	Level          field.Int32   // 等级
	SpreadUID      field.Int32   // 推广员id
	SpreadTime     field.Time    // 推广员关联时间
	UserType       field.String  // 用户类型
	IsPromoter     field.Int32   // 是否为推广员
	PayCount       field.Int32   // 用户购买次数
	SpreadCount    field.Int32   // 下级人数
	Addres         field.String  // 详细地址
	Adminid        field.Int32   // 管理员编号
	LoginType      field.String  // 用户登陆类型，h5,wechat,routine
	CreateTime     field.Time    // 创建时间
	UpdateTime     field.Time    // 更新时间
	LastLoginTime  field.Time    // 最后一次登录时间
	CleanTime      field.Time    // 清除时间
	Path           field.String  // 推广等级记录
	Subscribe      field.Int32   // 是否关注公众号
	SubscribeTime  field.Time    // 关注公众号时间
	Sex            field.Bool    // 性别，0未知，1男，2女，3保密
	Country        field.String  // 国家，中国CN，其他OTHER
	PromoterTime   field.Time    // 成为分销员时间

	fieldMap map[string]field.Expr
}

func (e ebUser) Table(newTableName string) *ebUser {
	e.ebUserDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebUser) As(alias string) *ebUser {
	e.ebUserDo.DO = *(e.ebUserDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebUser) updateTableName(table string) *ebUser {
	e.ALL = field.NewAsterisk(table)
	e.UID = field.NewInt32(table, "uid")
	e.Account = field.NewString(table, "account")
	e.Pwd = field.NewString(table, "pwd")
	e.RealName = field.NewString(table, "real_name")
	e.Birthday = field.NewString(table, "birthday")
	e.CardID = field.NewString(table, "card_id")
	e.Mark = field.NewString(table, "mark")
	e.PartnerID = field.NewInt32(table, "partner_id")
	e.GroupID = field.NewString(table, "group_id")
	e.TagID = field.NewString(table, "tag_id")
	e.Nickname = field.NewString(table, "nickname")
	e.Avatar = field.NewString(table, "avatar")
	e.Phone = field.NewString(table, "phone")
	e.AddIP = field.NewString(table, "add_ip")
	e.LastIP = field.NewString(table, "last_ip")
	e.NowMoney = field.NewFloat64(table, "now_money")
	e.BrokeragePrice = field.NewFloat64(table, "brokerage_price")
	e.Integral = field.NewInt32(table, "integral")
	e.Experience = field.NewInt32(table, "experience")
	e.SignNum = field.NewInt32(table, "sign_num")
	e.Status = field.NewBool(table, "status")
	e.Level = field.NewInt32(table, "level")
	e.SpreadUID = field.NewInt32(table, "spread_uid")
	e.SpreadTime = field.NewTime(table, "spread_time")
	e.UserType = field.NewString(table, "user_type")
	e.IsPromoter = field.NewInt32(table, "is_promoter")
	e.PayCount = field.NewInt32(table, "pay_count")
	e.SpreadCount = field.NewInt32(table, "spread_count")
	e.Addres = field.NewString(table, "addres")
	e.Adminid = field.NewInt32(table, "adminid")
	e.LoginType = field.NewString(table, "login_type")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")
	e.LastLoginTime = field.NewTime(table, "last_login_time")
	e.CleanTime = field.NewTime(table, "clean_time")
	e.Path = field.NewString(table, "path")
	e.Subscribe = field.NewInt32(table, "subscribe")
	e.SubscribeTime = field.NewTime(table, "subscribe_time")
	e.Sex = field.NewBool(table, "sex")
	e.Country = field.NewString(table, "country")
	e.PromoterTime = field.NewTime(table, "promoter_time")

	e.fillFieldMap()

	return e
}

func (e *ebUser) WithContext(ctx context.Context) IEbUserDo { return e.ebUserDo.WithContext(ctx) }

func (e ebUser) TableName() string { return e.ebUserDo.TableName() }

func (e ebUser) Alias() string { return e.ebUserDo.Alias() }

func (e ebUser) Columns(cols ...field.Expr) gen.Columns { return e.ebUserDo.Columns(cols...) }

func (e *ebUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebUser) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 41)
	e.fieldMap["uid"] = e.UID
	e.fieldMap["account"] = e.Account
	e.fieldMap["pwd"] = e.Pwd
	e.fieldMap["real_name"] = e.RealName
	e.fieldMap["birthday"] = e.Birthday
	e.fieldMap["card_id"] = e.CardID
	e.fieldMap["mark"] = e.Mark
	e.fieldMap["partner_id"] = e.PartnerID
	e.fieldMap["group_id"] = e.GroupID
	e.fieldMap["tag_id"] = e.TagID
	e.fieldMap["nickname"] = e.Nickname
	e.fieldMap["avatar"] = e.Avatar
	e.fieldMap["phone"] = e.Phone
	e.fieldMap["add_ip"] = e.AddIP
	e.fieldMap["last_ip"] = e.LastIP
	e.fieldMap["now_money"] = e.NowMoney
	e.fieldMap["brokerage_price"] = e.BrokeragePrice
	e.fieldMap["integral"] = e.Integral
	e.fieldMap["experience"] = e.Experience
	e.fieldMap["sign_num"] = e.SignNum
	e.fieldMap["status"] = e.Status
	e.fieldMap["level"] = e.Level
	e.fieldMap["spread_uid"] = e.SpreadUID
	e.fieldMap["spread_time"] = e.SpreadTime
	e.fieldMap["user_type"] = e.UserType
	e.fieldMap["is_promoter"] = e.IsPromoter
	e.fieldMap["pay_count"] = e.PayCount
	e.fieldMap["spread_count"] = e.SpreadCount
	e.fieldMap["addres"] = e.Addres
	e.fieldMap["adminid"] = e.Adminid
	e.fieldMap["login_type"] = e.LoginType
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
	e.fieldMap["last_login_time"] = e.LastLoginTime
	e.fieldMap["clean_time"] = e.CleanTime
	e.fieldMap["path"] = e.Path
	e.fieldMap["subscribe"] = e.Subscribe
	e.fieldMap["subscribe_time"] = e.SubscribeTime
	e.fieldMap["sex"] = e.Sex
	e.fieldMap["country"] = e.Country
	e.fieldMap["promoter_time"] = e.PromoterTime
}

func (e ebUser) clone(db *gorm.DB) ebUser {
	e.ebUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebUser) replaceDB(db *gorm.DB) ebUser {
	e.ebUserDo.ReplaceDB(db)
	return e
}

type ebUserDo struct{ gen.DO }

type IEbUserDo interface {
	gen.SubQuery
	Debug() IEbUserDo
	WithContext(ctx context.Context) IEbUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbUserDo
	WriteDB() IEbUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbUserDo
	Not(conds ...gen.Condition) IEbUserDo
	Or(conds ...gen.Condition) IEbUserDo
	Select(conds ...field.Expr) IEbUserDo
	Where(conds ...gen.Condition) IEbUserDo
	Order(conds ...field.Expr) IEbUserDo
	Distinct(cols ...field.Expr) IEbUserDo
	Omit(cols ...field.Expr) IEbUserDo
	Join(table schema.Tabler, on ...field.Expr) IEbUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbUserDo
	Group(cols ...field.Expr) IEbUserDo
	Having(conds ...gen.Condition) IEbUserDo
	Limit(limit int) IEbUserDo
	Offset(offset int) IEbUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserDo
	Unscoped() IEbUserDo
	Create(values ...*model.EbUser) error
	CreateInBatches(values []*model.EbUser, batchSize int) error
	Save(values ...*model.EbUser) error
	First() (*model.EbUser, error)
	Take() (*model.EbUser, error)
	Last() (*model.EbUser, error)
	Find() ([]*model.EbUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUser, err error)
	FindInBatches(result *[]*model.EbUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbUserDo
	Assign(attrs ...field.AssignExpr) IEbUserDo
	Joins(fields ...field.RelationField) IEbUserDo
	Preload(fields ...field.RelationField) IEbUserDo
	FirstOrInit() (*model.EbUser, error)
	FirstOrCreate() (*model.EbUser, error)
	FindByPage(offset int, limit int) (result []*model.EbUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebUserDo) Debug() IEbUserDo {
	return e.withDO(e.DO.Debug())
}

func (e ebUserDo) WithContext(ctx context.Context) IEbUserDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebUserDo) ReadDB() IEbUserDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebUserDo) WriteDB() IEbUserDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebUserDo) Session(config *gorm.Session) IEbUserDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebUserDo) Clauses(conds ...clause.Expression) IEbUserDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebUserDo) Returning(value interface{}, columns ...string) IEbUserDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebUserDo) Not(conds ...gen.Condition) IEbUserDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebUserDo) Or(conds ...gen.Condition) IEbUserDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebUserDo) Select(conds ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebUserDo) Where(conds ...gen.Condition) IEbUserDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebUserDo) Order(conds ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebUserDo) Distinct(cols ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebUserDo) Omit(cols ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebUserDo) Join(table schema.Tabler, on ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebUserDo) Group(cols ...field.Expr) IEbUserDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebUserDo) Having(conds ...gen.Condition) IEbUserDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebUserDo) Limit(limit int) IEbUserDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebUserDo) Offset(offset int) IEbUserDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebUserDo) Unscoped() IEbUserDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebUserDo) Create(values ...*model.EbUser) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebUserDo) CreateInBatches(values []*model.EbUser, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebUserDo) Save(values ...*model.EbUser) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebUserDo) First() (*model.EbUser, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUser), nil
	}
}

func (e ebUserDo) Take() (*model.EbUser, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUser), nil
	}
}

func (e ebUserDo) Last() (*model.EbUser, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUser), nil
	}
}

func (e ebUserDo) Find() ([]*model.EbUser, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbUser), err
}

func (e ebUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUser, err error) {
	buf := make([]*model.EbUser, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebUserDo) FindInBatches(result *[]*model.EbUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebUserDo) Attrs(attrs ...field.AssignExpr) IEbUserDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebUserDo) Assign(attrs ...field.AssignExpr) IEbUserDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebUserDo) Joins(fields ...field.RelationField) IEbUserDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebUserDo) Preload(fields ...field.RelationField) IEbUserDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebUserDo) FirstOrInit() (*model.EbUser, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUser), nil
	}
}

func (e ebUserDo) FirstOrCreate() (*model.EbUser, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUser), nil
	}
}

func (e ebUserDo) FindByPage(offset int, limit int) (result []*model.EbUser, count int64, err error) {
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

func (e ebUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebUserDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebUserDo) Delete(models ...*model.EbUser) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebUserDo) withDO(do gen.Dao) *ebUserDo {
	e.DO = *do.(*gen.DO)
	return e
}
