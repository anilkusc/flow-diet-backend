package material

type Material struct {
	//gorm.Model `json:"-" swaggerignore:"true"`
	Type       string `json:"type" example:"fruit"`
	Name       string `json:"name" example:"banana"`
	Photo_Urls string `json:"photo_urls" example:"[{'url':'exampleS3URL'}]"`
}
