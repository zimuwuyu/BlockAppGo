package model

import "time"

type BlockModel struct {
	ID           int       `gorm:"column:id;primaryKey"`
	Url          string    `gorm:"column:url;comment:'图片minio存储地址'"`
	Title        string    `gorm:"column:title;varchar(30);not null;comment:'模型名称'"`
	ModelType    int       `gorm:"column:model_type;not null;comment:'模型类别对模型分类表主键'"`
	ModelFile    string    `gorm:"column:model_file;comment:'图片minio存储地址'"`
	BucketName   string    `gorm:"column:bucket_name;varchar(20);comment:'图片minio存储地址'"`
	Hot          int       `gorm:"column:hot;default:0;comment:'图片minio存储地址'"`
	Introduction string    `gorm:"column:introduction;comment:'模型描述信息'"`
	IsRecommend  bool      `gorm:"column:is_recommend;default:false;comment:'是否是推荐模型'"`
	TotalCount   int       `gorm:"column:total_count;comment:'积木总数量'"`
	IsShow       bool      `gorm:"column:total_count;default:false;comment:'是否展示'"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}
