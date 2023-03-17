package models

import "time"

type Location string

const (
	LevTLV      Location = "LevTLV"
	OldNorth    Location = "OldNorth"
	NewNorth    Location = "NewNorth"
	Sarona      Location = "Sarona"
	NeveTzedek  Location = "NeveTzedek"
	NeveShaanan Location = "NeveShaanan"
	Florentin   Location = "Florentin"
	RamatAviv   Location = "RamatAviv"
)

type BoolPref int64

const (
	False        BoolPref = 0
	True         BoolPref = 1
	NotMandatory BoolPref = 2
	Mandatory    BoolPref = 4
)

type Range []int

type TimeArray []time.Time

type LocationArray []Location

type ApartmentArray []Apartment

type ApartmentPrefArray []ApartmentPref
