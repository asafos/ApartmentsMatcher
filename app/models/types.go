package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type RoleEnum uint

const (
	UserRole  RoleEnum = 1
	AdminRole RoleEnum = 1
)

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

type LocationSlice []Location

type Range []int

type TimeSlice []time.Time

type ApartmentSlice []Apartment

type ApartmentPrefSlice []ApartmentPref

// Valuer converts []time.Time to a format that can be stored in a database.
func (ts TimeSlice) Value() (driver.Value, error) {
	var sb strings.Builder
	for i, t := range ts {
		sb.WriteString(t.Format("2006-01-02"))
		if i != len(ts)-1 {
			sb.WriteRune(',')
		}
	}
	return sb.String(), nil
}

// Scanner converts a database value to []time.Time.
func (ts *TimeSlice) Scan(value interface{}) error {
	if value == nil {
		*ts = nil
		return nil
	}

	str, ok := value.([]uint8)
	if !ok {
		return fmt.Errorf("failed to parse TimeSlice value %v", value)
	}

	times := strings.Split(string(str), ",")
	result := make([]time.Time, len(times))

	for i, t := range times {
		parsedTime, err := time.Parse("2006-01-02", t)
		if err != nil {
			return fmt.Errorf("failed to parse time value %v: %v", t, err)
		}
		result[i] = parsedTime
	}

	*ts = result
	return nil
}

// Value returns the database/sql/driver value of this Range.
func (r Range) Value() (driver.Value, error) {
	// Convert the range of ints to a comma-separated string.
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r)), ","), "[]")
	return str, nil
}

// Scan assigns a value from a database driver.
func (r *Range) Scan(value interface{}) error {
	if value == nil {
		*r = nil
		return nil
	}
	// Convert the value from the database driver to a string.
	str, ok := value.([]uint8)
	if !ok {
		return fmt.Errorf("failed to scan value to Range: unexpected type %T", value)
	}
	// Split the comma-separated string into an array of string values.
	strValues := strings.Split(string(str), ",")
	// Convert the array of string values to an array of int values.
	intValues := make([]int, len(strValues))
	for i, strValue := range strValues {
		if _, err := fmt.Sscanf(strValue, "%d", &intValues[i]); err != nil {
			return fmt.Errorf("failed to parse int value from string %q: %w", strValue, err)
		}
	}
	*r = Range(intValues)
	return nil
}

// Value returns the database/sql/driver value of this LocationSlice.
func (l LocationSlice) Value() (driver.Value, error) {
	// Convert the slice of Location values to a comma-separated string.
	strValues := make([]string, len(l))
	for i, loc := range l {
		strValues[i] = string(loc)
	}
	str := strings.Join(strValues, ",")
	return str, nil
}

// Scan assigns a value from a database driver.
func (l *LocationSlice) Scan(value interface{}) error {
	if value == nil {
		*l = nil
		return nil
	}
	// Convert the value from the database driver to a string.
	str, ok := value.([]uint8)
	if !ok {
		return fmt.Errorf("failed to scan value to LocationSlice: unexpected type %T", value)
	}
	// Split the comma-separated string into an array of string values.
	strValues := strings.Split(string(str), ",")
	// Convert the array of string values to an array of Location values.
	locValues := make([]Location, len(strValues))
	for i, strValue := range strValues {
		locValues[i] = Location(strValue)
	}
	*l = LocationSlice(locValues)
	return nil
}
