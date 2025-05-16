package test_model

type Doc struct {
	ID            string   `gorm:"type:varchar(60);uniqueIndex"`
	Title         string   `gorm:"type:varchar(255);"`
	DocURL        string   `gorm:"type:varchar(255);"`
	Location      string   `gorm:"type:varchar(255);"`
	Schools       []string `gorm:"type:text[]"`
	Labels        []string `gorm:"type:text[]"`
	ImageURLs     []string `gorm:"type:text[]"`
	Description   string   `gorm:"type:text"`
	DownloadTotal int      `gorm:"type:int"`
	ViewTotal     int      `gorm:"type:int"`
}

func (d *Doc) TableName() string {
	return "doc"
}
