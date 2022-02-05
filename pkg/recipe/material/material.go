package material

type Material struct {
	Name string `json:"name" example:"banana"`
	//Tags []string `json:"tags" example:"vegan,fruit"`
	//Material_Diet_Level uint     `json:"material_diet_level" example:"1"`
	Material_Photo_Urls []string `gorm:"-" json:"material_photo_urls" example:"S3URL1,S3URL2"`
}
