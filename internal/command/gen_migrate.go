package command

import (
	"fmt"
	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/pkg/helper"
	"os"
	"time"

	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/bootstrap"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/command"
	"github.com/MQEnergy/mqshop/pkg/config"
	"github.com/urfave/cli/v2"
)

type GenMigrate struct{}

// Command ...
func (g *GenMigrate) Command() *cli.Command {
	var dbName string
	return &cli.Command{
		Name:  "genMigrate",
		Usage: "GenMigrate命令工具",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | test | prod]",
				Destination: &config.ConfEnv,
			},
			&cli.StringFlag{
				Name:        "dbName",
				Aliases:     []string{"db"},
				Value:       "default",
				Usage:       "请输入数据库的名称",
				Destination: &dbName,
				Required:    false,
			},
		},
		Before: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.MysqlService)
			return nil
		},
		Action: func(ctx *cli.Context) error {
			return handleGenMigrate(dbName)
		},
	}
}

var _ command.Interface = (*GenMigrate)(nil)

type MyStruct struct {
	Name  string
	Size  int64
	Mode  os.FileMode
	IsDir bool
}

// handleGenMigrate ...
func handleGenMigrate(dbName string) error {
	fmt.Println("handleGenMigrate")
	if err := vars.MDB[dbName].Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.Banner{},
		&model.BannerCate{},
		&model.CAdmin{},
		&model.CAreas{},
		&model.CAttachment{},
		&model.CCasbinRule{},
		&model.CCities{},
		&model.CProvinces{},
		&model.CResource{},
		&model.CRole{},
		&model.CRoleAuth{},
		&model.CStreets{},
		&model.CVillages{},
		&model.CommentLike{},
		&model.CommentReply{},
		&model.Member{},
		&model.MemberAddress{},
		&model.MemberCart{},
		&model.MemberFollowRelation{},
		&model.MemberGoodsFavorite{},
		&model.MemberInfo{},
		&model.MemberPost{},
		&model.MemberPostTagRelation{},
		&model.MemberThirdAuth{},
		&model.Order{},
		&model.OrderGoods{},
		&model.OrderPayLog{},
		&model.PostComment{},
		&model.PostLike{},
		&model.ProductBrand{},
		&model.ProductCategory{},
		&model.ProductComment{},
		&model.ProductGoods{},
		&model.ProductGoodsAttr{},
		&model.ProductGoodsAttrCate{},
		&model.ProductGoodsAttrRelation{},
		&model.ProductGoodsInfo{},
		&model.ProductGoodsSku{},
		&model.ProductGoodsSkuAttrRelation{},
		&model.ProductGoodsTagRelation{},
		&model.SmsCode{},
		&model.Suggestion{},
		&model.Tag{},
	); err != nil {
		return err
	}
	// 生成一个admin数据
	salt := helper.GenerateUuid(32)
	if err := dao.CAdmin.Create(&model.CAdmin{
		UUID:         helper.GenerateBaseSnowId(0, nil),
		Account:      "admin",
		Password:     helper.GeneratePasswordHash("admin888", salt),
		Phone:        "12345678910",
		Avatar:       "",
		Salt:         salt,
		RealName:     "Admin",
		RegisterTime: time.Now().Unix(),
		RegisterIP:   "127.0.0.1",
		LoginTime:    time.Now().Unix(),
		LoginIP:      "127.0.0.1",
		RoleIds:      "1",
		Status:       1,
	}); err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", dbName+" 数据表迁移完成"))
	return nil
}
