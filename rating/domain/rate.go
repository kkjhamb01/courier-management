package domain

type Rate struct {
	Rated                string                		 `gorm:"type:string;size:64;primaryKey;not null;column:rated"`
	RateTotal     	     int64               		 `gorm:"type:int64;not null;column:rate_total"`
	RateCount  	         int64               		 `gorm:"type:int64;not null;column:rate_count"`
}

func (Rate) TableName() string {
	return "rate"
}
