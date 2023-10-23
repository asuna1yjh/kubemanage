package option

import (
	"context"
	"fmt"
	"gin_demo/dao"
	"gin_demo/logger"
	"gin_demo/models"
	"gin_demo/settings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Option struct {
	Engine *gin.Engine
	DB     *gorm.DB
	// 抽象数据库接口
	Factory *dao.ShareDaoFactory
	Config  *settings.AppConfig
	Ctx     context.Context
}

func NewOption() *Option {
	return &Option{}
}

func (o *Option) Complete() error {
	// 初始化配置文件
	conf := settings.InitConf()
	o.Config = conf
	// 初始化日志
	if err := o.RegisterLogger(); err != nil {
		return err
	}
	// 初始化数据库
	if err := o.InitDB(); err != nil {
		return err
	}
	// 初始化数据库接口
	if err := o.RegisterFactory(); err != nil {
		return err
	}

	return nil
}

func (o *Option) RegisterLogger() error {
	err := logger.Init(o.Config.LogConfig, o.Config.Mode)
	if err != nil {
		return err
	}
	return nil
}

func (o *Option) InitDB() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		o.Config.MysqlConfig.User,
		o.Config.MysqlConfig.Password,
		o.Config.MysqlConfig.Host,
		o.Config.MysqlConfig.Port,
		o.Config.MysqlConfig.DBname,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.User{}, &models.SysUser{}, &models.SysRole{}, &models.SysUserRole{}, &models.SysMenu{}, &models.SysRoleMenu{})
	if err != nil {
		zap.L().Error("auto migrate table failed", zap.Error(err))
		return err
	}
	o.DB = db
	return nil
}

func (o *Option) RegisterFactory() error {
	o.Factory = dao.NewShareDaoFactory(context.TODO(), o.DB)
	return nil
}
