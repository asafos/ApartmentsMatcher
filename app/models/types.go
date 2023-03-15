package models

import "time"

type Area string

const (
	LevTLV      Area = "LevTLV"
	OldNorth    Area = "OldNorth"
	NewNorth    Area = "NewNorth"
	Sarona      Area = "Sarona"
	NeveTzedek  Area = "NeveTzedek"
	NeveShaanan Area = "NeveShaanan"
	Florentin   Area = "Florentin"
	RamatAviv   Area = "RamatAviv"
)

type BoolPref int64

const (
	False        BoolPref = 0
	True         BoolPref = 1
	NotMandatory BoolPref = 2
	Mandatory    BoolPref = 4
)

type Range []int64

type TimeArray []time.Time

type AreaArray []Area

type ApartmentArray []Apartment

type ApartmentPrefArray []ApartmentPref
