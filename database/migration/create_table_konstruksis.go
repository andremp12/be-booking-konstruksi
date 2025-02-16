package database

import (
	"fmt"
	"time"
)

type Status string
type Tipe string

//const (
//	Booking  Status = "Booking"
//	Proses   Status = "Proses"
//	Finished Status = "Finished"
//	Canceled Status = "Canceled"
//	Payment  Status = "Payment"
//	Full     Tipe   = "Full"
//	Jasa     Tipe   = "Jasa"
//)

type Konstruksi struct {
	ID               int            `gorm:"type:int(11);primary_key;auto_increment"`
	NoBooking        string         `gorm:"type:varchar(50);"`
	TipeKonstruksiID int            `gorm:"type:int(11);not null"`
	TipeKonstruksi   TipeKonstruksi `gorm:"references:ID;foreignKey:TipeKonstruksiID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ClientID         *int           `gorm:"type:int(11);not null"`
	Client           User           `gorm:"references:ID;foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	MandorID         *int           `gorm:"type:int(11);null"`
	Mandor           User           `gorm:"foreignKey:MandorID;references:ID;constraint:OnUpdate:CASCADE,onDelete:RESTRICT;"`
	Tipe             Tipe           `gorm:"type:ENUM('Full', 'Jasa')"`
	Address          string         `gorm:"type:text;not null"`
	Keterangan       string         `gorm:"type:text"`
	Status           Status         `gorm:"type:ENUM('Booking','Payment','Proses','Finished','Canceled','Closed')"`
	Total            int            `gorm:"type:bigInt;not null"`
	Paid             int            `gorm:"type:bigInt;null"`
	RemainingPayment *int           `gorm:"type:bigInt;null"`
	StartDate        *time.Time     `gorm:"type:date; null"`
	EndDate          *time.Time     `gorm:"type:date; null"`
	CreatedAt        time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt        time.Time      `gorm:"type:timestamp;not null"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	//Parsed time with format YYYY-MM-DD
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		fmt.Println(err)
		return err
	}

	t.Time = parsedTime
	return nil
}
