package material

type Material struct {
	//gorm.Model `json:"-" swaggerignore:"true"`
	Type                string   `json:"type" example:"fruit"`
	Name                string   `json:"name" example:"banana"`
	Material_Photo_Urls []string `gorm:"-" json:"material_photo_urls" example:"S3URL1,S3URL2"`
}
