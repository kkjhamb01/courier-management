package config

import "time"

type OfferingData struct {
	DbName                     string        `yaml:"dbName"`
	GrpcPort                   string        `yaml:"grpcPort"`
	DaprPort                   string        `yaml:"daprPort"`
	MaxCouriersOnNewOffer      int           `yaml:"maxCouriersOnNewOffer"`
	NearbyCouriersToSearch     int           `yaml:"nearbyCouriersToSearch"`
	MaxOfferRetries            int           `yaml:"maxOfferRetries"`
	CouriersDistanceOnNewOffer int           `yaml:"couriersDistanceOnNewOffer"`
	MaxOffersPerCourier        int           `yaml:"maxOffersPerCourier"`
	CourierTimeToAnswerOffer   time.Duration `yaml:"courierTimeToAnswerOffer"`
}

type DeliveryData struct {
	CreateRequestScheduleDuration time.Duration `yaml:"createRequestScheduleDuration"`
	GrpcPort                      string        `yaml:"grpcPort"`
	ValidCities                   []string      `yaml:"validCities"`
	DbName                        string        `yaml:"dbName"`
}

type InternalHttpData struct {
	Port string `yaml:"port"`
}

type FinanceData struct {
	GrpcPort                      string `yaml:"grpcPort"`
	WebhookPort                   string `yaml:"webhookPort"`
	AccessToken                   string `yaml:"accessToken"`
	DbName                        string `yaml:"dbName"`
	CronJobSchedule               string `yaml:"cronJobSchedule"`
	RevenueStripeId               string `yaml:"revenueStripeId"`
	TaxStripeId                   string `yaml:"taxStripeId"`
	StripeSecretKey               string `yaml:"stripeSecretKey"`
	PaymentCourierSharePercent    int    `yaml:"paymentCourierSharePercent"`
	PaymentTaxSharePercent        int    `yaml:"paymentTaxSharePercent"`
	PaymentRevenuePercent         int    `yaml:"paymentRevenuePercent"`
	CreateStripeAccountMaxRetries int    `yaml:"createStripeAccountMaxRetries"`
}

type GoogleMapData struct {
	ApiKey string `yaml:"apiKey"`
}

type PricingData struct {
	GrpcPort 		string 		`yaml:"grpcPort"`
	Connection	 	string 		`yaml:"connection"`
}

type NatsData struct {
	Address string `yaml:"address"`
}

type Tile38Data struct {
	Address                   string `yaml:"address"`
	CourierLocationCollection string `yaml:"courierLocationCollection"`
}

type OfferRedisData struct {
	Address  string `yaml:"address"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type LogData struct {
	FilePath string `yaml:"filePath"`
}

type JwtData struct {
	SignKey                string        `yaml:"signKey"`
	VerifyKey              string        `yaml:"verifyKey"`
	Cert                   string        `yaml:"cert"`
	SigningMethod          string        `yaml:"signingMethod"`
	ExpirationMinutes      time.Duration `yaml:"expirationMinutes"`
	RefreshExpirationHours time.Duration `yaml:"refreshExpirationHours"`
	Issuer                 string        `yaml:"issuer"`
}

type Server struct {
	Address string `yaml:"address"`
}

type Database struct {
	Type                 string `yaml:"type"`
	DatabaseDriver       string `yaml:"databaseDriver"`
	DatabaseHost         string `yaml:"databaseHost"`
	DatabasePort         string `yaml:"databasePort"`
	Username             string `yaml:"username"`
	Password             string `yaml:"password"`
	DatabaseName         string `yaml:"databaseName"`
	MaxIdleConnections   int    `yaml:"maxIdleConnections"`
	MaxOpenConnections   int    `yaml:"maxOpenConnections"`
	MaxIdleTimeInMinutes int    `yaml:"maxIdleTimeInMinutes"`
	MaxLifetimeInMinutes int    `yaml:"maxLifetimeInMinutes"`
}

type PartyData struct {
	Server    Server    `yaml:"server"`
	Database  Database  `yaml:"database"`
	Download  Download  `yaml:"download"`
	Mot       MOT       `yaml:"mot"`
	Thumbnail Thumbnail `yaml:"thumbnail"`
}

type Thumbnail struct {
	MaxDimension int `yaml:"maxDimension"`
}

type MOT struct {
	ApiAddress string `yaml:"apiAddress"`
	ApiKey     string `yaml:"apiKey"`
	Timeout    int    `yaml:"timeout"`
}

type NotificationData struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Gorush   string   `yaml:"gorush"`
}

type Download struct {
	Expiration          int64  `yaml:"expiration"`
	ExpirationSecretKey string `yaml:"expirationSecretKey"`
}

type OauthByRole struct {
	ClientID     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	RedirectURL  string `yaml:"redirectUrl"`
}

type Google struct {
	Driver    OauthByRole `yaml:"driver"`
	Passenger OauthByRole `yaml:"passenger"`
}

type Redis struct {
	Address                   string `yaml:"address"`
	Db                        int    `yaml:"db"`
	Password                  string `yaml:"password"`
	OauthSessionExpireSeconds int    `yaml:"oauthSessionExpireSeconds"`
	OTPSessionExpireSeconds   int    `yaml:"otpSessionExpireSeconds"`
}

type Otp struct {
	Static   string `yaml:"static"`
	Length   int    `yaml:"length"`
	MaxRetry int32  `yaml:"maxRetry"`
}

type UaaData struct {
	Server           Server   `yaml:"server"`
	Redis            Redis    `yaml:"redis"`
	Google           Google   `yaml:"google"`
	Facebook         Facebook `yaml:"facebook"`
	Otp              Otp      `yaml:"otp"`
	PartyAddress     string   `yaml:"partyAddress"`
	DeepLinkTemplate string   `yaml:"deepLinkTemplate"`
}

type Facebook struct {
	Driver    OauthByRole `yaml:"driver"`
	Passenger OauthByRole `yaml:"passenger"`
}

type MariaDbData struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
}

type RatingData struct {
	Server      Server      `yaml:"server"`
	Database    Database    `yaml:"database"`
	RideService RideService `yaml:"rideService"`
}

type PromotionData struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Referral Referral `yaml:"referral"`
}

type AnnouncementData struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Referral Referral `yaml:"referral"`
}

type Referral struct {
	DiscountValue      float64 `yaml:"discountValue"`
	DiscountPercentage float64 `yaml:"discountPercentage"`
	PromotionName      string  `yaml:"promotionName"`
}

type RideService struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// The Data struct represents the configured values and needs to be initialized by the config solution
type Data struct {
	Log          		LogData          	`yaml:"log"`
	Offering     		OfferingData     	`yaml:"offering"`
	Tile38       		Tile38Data       	`yaml:"tile38"`
	Uaa          		UaaData          	`yaml:"uaa"`
	Party        		PartyData        	`yaml:"party"`
	Finance      		FinanceData      	`yaml:"finance"`
	Pricing      		PricingData      	`yaml:"pricing"`
	InternalHttp 		InternalHttpData 	`yaml:"internalHttp"`
	OfferRedis   		OfferRedisData   	`yaml:"offerRedis"`
	Delivery     		DeliveryData     	`yaml:"delivery"`
	Nats         		NatsData         	`yaml:"nats"`
	MariaDb      		MariaDbData      	`yaml:"mariadb"`
	GoogleMap    		GoogleMapData    	`yaml:"googleMap"`
	Jwt          		JwtData          	`yaml:"jwt"`
	Notification 		NotificationData 	`yaml:"notification"`
	Rating       		RatingData       	`yaml:"rating"`
	Promotion    		PromotionData    	`yaml:"promotion"`
	Announcement    	AnnouncementData    `yaml:"announcement"`
	Tracing		    	TracingData    		`yaml:"tracing"`
	DistanceApi		    DistanceApi    		`yaml:"distanceApi"`
}

type TracingData struct {
	DSN string `yaml:"dsn"`
}

type DistanceApi struct {
	Key string `yaml:"key"`
}

var data *Data

func Offering() OfferingData {
	// create a copy of the config to make sure the config will not be altered
	return data.Offering
}

func Finance() FinanceData {
	return data.Finance
}

func Pricing() PricingData {
	return data.Pricing
}

func MariaDb() MariaDbData {
	return data.MariaDb
}

func Delivery() DeliveryData {
	return data.Delivery
}

func Nats() NatsData {
	return data.Nats
}

func OfferRedis() OfferRedisData {
	return data.OfferRedis
}

func Log() LogData {
	return data.Log
}

func Tile38() Tile38Data {
	return data.Tile38
}

func GoogleMap() GoogleMapData {
	return data.GoogleMap
}

func Uaa() UaaData {
	return data.Uaa
}

func Party() PartyData {
	return data.Party
}

func Rating() RatingData {
	return data.Rating
}

func Promotion() PromotionData {
	return data.Promotion
}

func Announcement() AnnouncementData {
	return data.Announcement
}

func Tracing() TracingData {
	return data.Tracing
}

func Notification() NotificationData {
	return data.Notification
}

func Jwt() JwtData {
	return data.Jwt
}

func InternalHttp() InternalHttpData {
	return data.InternalHttp
}

func GetData() Data {
	return *data
}
