package repository

import (
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/system_admin_repository"
	"crmeb_go/internal/repository/system_config_repository"
	"crmeb_go/internal/repository/system_group_data_repository"
	"crmeb_go/internal/repository/system_menu_repository"
	"crmeb_go/internal/repository/user_repository"
	"gorm.io/gorm"
)

type Container struct {
	UserRepository            *user_repository.UserRepository                  // 用户表 模型
	SystemConfigRepository    *system_config_repository.SystemConfigRepository // 配置表
	SystemGroupDataRepository *system_group_data_repository.SystemGroupDataRepository
	SystemAdminRepository     *system_admin_repository.SystemAdminRepository
	SystemMenuRepository      *system_menu_repository.SystemMenuRepository
}

func Register(db *gorm.DB, gen *gen.Query) *Container {
	return &Container{
		UserRepository:            user_repository.NewUserRepository(db, gen),                         // 用户表
		SystemConfigRepository:    system_config_repository.NewSystemConfigRepository(db, gen),        // 配置表
		SystemGroupDataRepository: system_group_data_repository.NewSystemGroupDataRepository(db, gen), // 配置表
		SystemAdminRepository:     system_admin_repository.NewSystemAdminRepository(db, gen),          // 后台管理表
		SystemMenuRepository:      system_menu_repository.NewSystemMenuRepository(db, gen),            // 系统菜单表
	}
}
