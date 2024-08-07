// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/MQEnergy/mqshop/internal/app/model"
)

func newMemberThirdAuth(db *gorm.DB, opts ...gen.DOOption) memberThirdAuth {
	_memberThirdAuth := memberThirdAuth{}

	_memberThirdAuth.memberThirdAuthDo.UseDB(db, opts...)
	_memberThirdAuth.memberThirdAuthDo.UseModel(&model.MemberThirdAuth{})

	tableName := _memberThirdAuth.memberThirdAuthDo.TableName()
	_memberThirdAuth.ALL = field.NewAsterisk(tableName)
	_memberThirdAuth.ID = field.NewInt64(tableName, "id")
	_memberThirdAuth.MemberID = field.NewInt64(tableName, "member_id")
	_memberThirdAuth.NickName = field.NewString(tableName, "nick_name")
	_memberThirdAuth.Avatar = field.NewString(tableName, "avatar")
	_memberThirdAuth.Gender = field.NewInt8(tableName, "gender")
	_memberThirdAuth.City = field.NewString(tableName, "city")
	_memberThirdAuth.Province = field.NewString(tableName, "province")
	_memberThirdAuth.Country = field.NewString(tableName, "country")
	_memberThirdAuth.Openid = field.NewString(tableName, "openid")
	_memberThirdAuth.SessionKey = field.NewString(tableName, "session_key")
	_memberThirdAuth.Unionid = field.NewString(tableName, "unionid")
	_memberThirdAuth.PlatformType = field.NewInt8(tableName, "platform_type")
	_memberThirdAuth.CreatedAt = field.NewInt64(tableName, "created_at")

	_memberThirdAuth.fillFieldMap()

	return _memberThirdAuth
}

// memberThirdAuth 用户三方授权表
type memberThirdAuth struct {
	memberThirdAuthDo

	ALL          field.Asterisk
	ID           field.Int64
	MemberID     field.Int64  // 用户ID
	NickName     field.String // 昵称
	Avatar       field.String // 用户头像
	Gender       field.Int8   // 性别 0：未设置 1：男 2：女
	City         field.String // 城市
	Province     field.String // 省份
	Country      field.String // 国家
	Openid       field.String // 用户唯一标识
	SessionKey   field.String // 会话秘钥
	Unionid      field.String // 用户在开放平台的唯一标识符
	PlatformType field.Int8   // 平台标识 1：微信小程序 2：微信公众号
	CreatedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (m memberThirdAuth) Table(newTableName string) *memberThirdAuth {
	m.memberThirdAuthDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m memberThirdAuth) As(alias string) *memberThirdAuth {
	m.memberThirdAuthDo.DO = *(m.memberThirdAuthDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *memberThirdAuth) updateTableName(table string) *memberThirdAuth {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.MemberID = field.NewInt64(table, "member_id")
	m.NickName = field.NewString(table, "nick_name")
	m.Avatar = field.NewString(table, "avatar")
	m.Gender = field.NewInt8(table, "gender")
	m.City = field.NewString(table, "city")
	m.Province = field.NewString(table, "province")
	m.Country = field.NewString(table, "country")
	m.Openid = field.NewString(table, "openid")
	m.SessionKey = field.NewString(table, "session_key")
	m.Unionid = field.NewString(table, "unionid")
	m.PlatformType = field.NewInt8(table, "platform_type")
	m.CreatedAt = field.NewInt64(table, "created_at")

	m.fillFieldMap()

	return m
}

func (m *memberThirdAuth) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *memberThirdAuth) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 13)
	m.fieldMap["id"] = m.ID
	m.fieldMap["member_id"] = m.MemberID
	m.fieldMap["nick_name"] = m.NickName
	m.fieldMap["avatar"] = m.Avatar
	m.fieldMap["gender"] = m.Gender
	m.fieldMap["city"] = m.City
	m.fieldMap["province"] = m.Province
	m.fieldMap["country"] = m.Country
	m.fieldMap["openid"] = m.Openid
	m.fieldMap["session_key"] = m.SessionKey
	m.fieldMap["unionid"] = m.Unionid
	m.fieldMap["platform_type"] = m.PlatformType
	m.fieldMap["created_at"] = m.CreatedAt
}

func (m memberThirdAuth) clone(db *gorm.DB) memberThirdAuth {
	m.memberThirdAuthDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m memberThirdAuth) replaceDB(db *gorm.DB) memberThirdAuth {
	m.memberThirdAuthDo.ReplaceDB(db)
	return m
}

type memberThirdAuthDo struct{ gen.DO }

type IMemberThirdAuthDo interface {
	gen.SubQuery
	Debug() IMemberThirdAuthDo
	WithContext(ctx context.Context) IMemberThirdAuthDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMemberThirdAuthDo
	WriteDB() IMemberThirdAuthDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMemberThirdAuthDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMemberThirdAuthDo
	Not(conds ...gen.Condition) IMemberThirdAuthDo
	Or(conds ...gen.Condition) IMemberThirdAuthDo
	Select(conds ...field.Expr) IMemberThirdAuthDo
	Where(conds ...gen.Condition) IMemberThirdAuthDo
	Order(conds ...field.Expr) IMemberThirdAuthDo
	Distinct(cols ...field.Expr) IMemberThirdAuthDo
	Omit(cols ...field.Expr) IMemberThirdAuthDo
	Join(table schema.Tabler, on ...field.Expr) IMemberThirdAuthDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMemberThirdAuthDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMemberThirdAuthDo
	Group(cols ...field.Expr) IMemberThirdAuthDo
	Having(conds ...gen.Condition) IMemberThirdAuthDo
	Limit(limit int) IMemberThirdAuthDo
	Offset(offset int) IMemberThirdAuthDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMemberThirdAuthDo
	Unscoped() IMemberThirdAuthDo
	Create(values ...*model.MemberThirdAuth) error
	CreateInBatches(values []*model.MemberThirdAuth, batchSize int) error
	Save(values ...*model.MemberThirdAuth) error
	First() (*model.MemberThirdAuth, error)
	Take() (*model.MemberThirdAuth, error)
	Last() (*model.MemberThirdAuth, error)
	Find() ([]*model.MemberThirdAuth, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MemberThirdAuth, err error)
	FindInBatches(result *[]*model.MemberThirdAuth, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MemberThirdAuth) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMemberThirdAuthDo
	Assign(attrs ...field.AssignExpr) IMemberThirdAuthDo
	Joins(fields ...field.RelationField) IMemberThirdAuthDo
	Preload(fields ...field.RelationField) IMemberThirdAuthDo
	FirstOrInit() (*model.MemberThirdAuth, error)
	FirstOrCreate() (*model.MemberThirdAuth, error)
	FindByPage(offset int, limit int) (result []*model.MemberThirdAuth, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMemberThirdAuthDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m memberThirdAuthDo) Debug() IMemberThirdAuthDo {
	return m.withDO(m.DO.Debug())
}

func (m memberThirdAuthDo) WithContext(ctx context.Context) IMemberThirdAuthDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memberThirdAuthDo) ReadDB() IMemberThirdAuthDo {
	return m.Clauses(dbresolver.Read)
}

func (m memberThirdAuthDo) WriteDB() IMemberThirdAuthDo {
	return m.Clauses(dbresolver.Write)
}

func (m memberThirdAuthDo) Session(config *gorm.Session) IMemberThirdAuthDo {
	return m.withDO(m.DO.Session(config))
}

func (m memberThirdAuthDo) Clauses(conds ...clause.Expression) IMemberThirdAuthDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memberThirdAuthDo) Returning(value interface{}, columns ...string) IMemberThirdAuthDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memberThirdAuthDo) Not(conds ...gen.Condition) IMemberThirdAuthDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memberThirdAuthDo) Or(conds ...gen.Condition) IMemberThirdAuthDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memberThirdAuthDo) Select(conds ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memberThirdAuthDo) Where(conds ...gen.Condition) IMemberThirdAuthDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memberThirdAuthDo) Order(conds ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memberThirdAuthDo) Distinct(cols ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memberThirdAuthDo) Omit(cols ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memberThirdAuthDo) Join(table schema.Tabler, on ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memberThirdAuthDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memberThirdAuthDo) RightJoin(table schema.Tabler, on ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memberThirdAuthDo) Group(cols ...field.Expr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memberThirdAuthDo) Having(conds ...gen.Condition) IMemberThirdAuthDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memberThirdAuthDo) Limit(limit int) IMemberThirdAuthDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memberThirdAuthDo) Offset(offset int) IMemberThirdAuthDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memberThirdAuthDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMemberThirdAuthDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memberThirdAuthDo) Unscoped() IMemberThirdAuthDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memberThirdAuthDo) Create(values ...*model.MemberThirdAuth) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memberThirdAuthDo) CreateInBatches(values []*model.MemberThirdAuth, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memberThirdAuthDo) Save(values ...*model.MemberThirdAuth) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memberThirdAuthDo) First() (*model.MemberThirdAuth, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberThirdAuth), nil
	}
}

func (m memberThirdAuthDo) Take() (*model.MemberThirdAuth, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberThirdAuth), nil
	}
}

func (m memberThirdAuthDo) Last() (*model.MemberThirdAuth, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberThirdAuth), nil
	}
}

func (m memberThirdAuthDo) Find() ([]*model.MemberThirdAuth, error) {
	result, err := m.DO.Find()
	return result.([]*model.MemberThirdAuth), err
}

func (m memberThirdAuthDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MemberThirdAuth, err error) {
	buf := make([]*model.MemberThirdAuth, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memberThirdAuthDo) FindInBatches(result *[]*model.MemberThirdAuth, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memberThirdAuthDo) Attrs(attrs ...field.AssignExpr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memberThirdAuthDo) Assign(attrs ...field.AssignExpr) IMemberThirdAuthDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memberThirdAuthDo) Joins(fields ...field.RelationField) IMemberThirdAuthDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memberThirdAuthDo) Preload(fields ...field.RelationField) IMemberThirdAuthDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memberThirdAuthDo) FirstOrInit() (*model.MemberThirdAuth, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberThirdAuth), nil
	}
}

func (m memberThirdAuthDo) FirstOrCreate() (*model.MemberThirdAuth, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberThirdAuth), nil
	}
}

func (m memberThirdAuthDo) FindByPage(offset int, limit int) (result []*model.MemberThirdAuth, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memberThirdAuthDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memberThirdAuthDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memberThirdAuthDo) Delete(models ...*model.MemberThirdAuth) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memberThirdAuthDo) withDO(do gen.Dao) *memberThirdAuthDo {
	m.DO = *do.(*gen.DO)
	return m
}
