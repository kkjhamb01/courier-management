package domain

import (
	"bytes"
	"encoding/json"
	"strings"
)

type CourierMot struct {
	UserID           				string		`gorm:"type:string;size:64;primaryKey;not null;column:user_id" json:"user_id"`
	RegistrationNumber           	string		`gorm:"type:string;size:64;primaryKey;not null;column:registration_number" json:"registrationNumber"`
	Co2Emissions  	        		int32		`gorm:"type:int;column:co2_emissions" json:"co2Emissions"`
	EngineCapacity  	        	int32		`gorm:"type:int;column:engine_capacity" json:"engineCapacity"`
	EuroStatus    					string		`gorm:"type:string;size:64;column:euro_status" json:"euroStatus"`
	MarkedForExport  	        	bool		`gorm:"type:int;column:marked_for_export" json:"markedForExport"`
	FuelType    					string		`gorm:"type:string;size:64;column:fuel_type" json:"fuelType"`
	MotStatus    					MotStatus	`gorm:"type:string;size:64;column:mot_status" json:"motStatus"`
	RevenueWeight  	   		     	int32		`gorm:"type:int;column:revenue_weight" json:"revenueWeight"`
	Colour	    					string		`gorm:"type:string;size:64;column:colour" json:"colour"`
	Make	    					string		`gorm:"type:string;size:64;column:make" json:"make"`
	TypeApproval    				string		`gorm:"type:string;size:64;column:type_approval" json:"typeApproval"`
	YearOfManufacture  	   		    int32		`gorm:"type:int;column:year_of_manufacture" json:"yearOfManufacture"`
	TaxDueDate 		   				string		`gorm:"type:string;size:64;column:tax_due_date" json:"taxDueDate"`
	TaxStatus	    				TaxStatus	`gorm:"type:string;size:64;column:tax_status" json:"taxStatus"`
	DateOfLastV5CIssued    			string		`gorm:"type:string;size:64;column:date_of_last_v5c_issued" json:"dateOfLastV5CIssued"`
	RealDrivingEmissions    		string		`gorm:"type:string;size:64;column:real_driving_emissions" json:"realDrivingEmissions"`
	Wheelplan  		  				string		`gorm:"type:string;size:64;column:wheelplan" json:"wheelplan"`
	MonthOfFirstRegistration    	string		`gorm:"type:string;size:64;column:month_of_first_registration" json:"monthOfFirstRegistration"`
}

func (CourierMot) TableName() string {
	return "courier_mot"
}



type MotStatus int32

const (
	MOT_STATUS_UNKNOWN MotStatus = iota
	MOT_STATUS_NO_DETAILS_HELD_BY_DVLA
	MOT_STATUS_NO_RESULTS_RETURNED
	MOT_STATUS_NOT_VALID
	MOT_STATUS_VALID
)

func (s MotStatus) String() string {
	return motToString[s]
}

var motToString = map[MotStatus]string{
	MOT_STATUS_UNKNOWN:  "",
	MOT_STATUS_NO_DETAILS_HELD_BY_DVLA:  "no details held by dvla",
	MOT_STATUS_NO_RESULTS_RETURNED: "no results returned",
	MOT_STATUS_NOT_VALID: "not valid",
	MOT_STATUS_VALID: "valid",
}

var motToID = map[string]MotStatus{
	"":  MOT_STATUS_UNKNOWN,
	"no details held by dvla":  MOT_STATUS_NO_DETAILS_HELD_BY_DVLA,
	"no results returned": MOT_STATUS_NO_RESULTS_RETURNED,
	"not valid": MOT_STATUS_NOT_VALID,
	"valid": MOT_STATUS_VALID,
}

func (s MotStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(motToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *MotStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = motToID[strings.ToLower(j)]
	return nil
}


type TaxStatus int32

const (
	TAX_STATUS_UNKNOWN TaxStatus = iota
	TAX_STATUS_NOT_TAXED_FOR_ON_ROAD_USE
	TAX_STATUS_SORN
	TAX_STATUS_TAXED
	TAX_STATUS_UNTAXED
)

func (s TaxStatus) String() string {
	return taxToString[s]
}

var taxToString = map[TaxStatus]string{
	TAX_STATUS_UNKNOWN:  "",
	TAX_STATUS_NOT_TAXED_FOR_ON_ROAD_USE:  "not taxed for on road use",
	TAX_STATUS_SORN: "sorn",
	TAX_STATUS_TAXED: "taxed",
	TAX_STATUS_UNTAXED: "untaxed",
}

var taxToID = map[string]TaxStatus{
	"":  TAX_STATUS_UNKNOWN,
	"not taxed for on road use":  TAX_STATUS_NOT_TAXED_FOR_ON_ROAD_USE,
	"sorn": TAX_STATUS_SORN,
	"taxed": TAX_STATUS_TAXED,
	"untaxed": TAX_STATUS_UNTAXED,
}

func (s TaxStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(taxToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *TaxStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = taxToID[strings.ToLower(j)]
	return nil
}