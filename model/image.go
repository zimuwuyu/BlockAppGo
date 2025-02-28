package model

type ImageType string

const (
	Carousel   ImageType = "CAROUSEL"
	Process    ImageType = "PROCESS"
	ProcessGif ImageType = "PROCESS_GIF"
	Show3D     ImageType = "SHOW_3D"
	ShowGif    ImageType = "SHOW_GIF"
	Step       ImageType = "STEP"
)

type StorageType string

const (
	BlockMinio StorageType = "BLOCKMINIO"
	BlockYhs   StorageType = "BLOCKYHS"
)

type PictureStorage struct {
	ID          int         `gorm:"column:id;type:integer;primaryKey"`
	ModelId     int         `gorm:"column:model_id;type:integer;not null;comment:'模型表主键id'"`
	FileName    string      `gorm:"column:file_name;type:varchar(30);not null;comment:'文件名'"`
	ImagePath   string      `gorm:"column:image_path;type:varchar(100);not null;comment:'在minio的存储路径'"`
	BucketName  string      `gorm:"column:bucket_name;type:varchar(30);not null;comment:'在minio的存储的桶名'"`
	Url         string      `gorm:"column:url;type:varchar(100);comment:'存在公共桶的此自动才有值'"`
	ImageType   ImageType   `gorm:"column:image_type;not null;comment:'图片类型'"`
	StorageType StorageType `gorm:"column:storage_type;not null;comment:'存储类型'"`
	CreateTime  string      `gorm:"column:create_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime  string      `gorm:"column:update_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}

func (PictureStorage) TableName() string {
	return "picture_storage"
}
