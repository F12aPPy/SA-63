// Code generated by entc, DO NOT EDIT.

package antenatal

import (
	"time"
)

const (
	// Label holds the string label denoting the antenatal type in the database.
	Label = "antenatal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldADDEDTIME holds the string denoting the added_time field in the database.
	FieldADDEDTIME = "added_time"

	// EdgeGETMOM holds the string denoting the getmom edge name in mutations.
	EdgeGETMOM = "GETMOM"
	// EdgeTAKECARE holds the string denoting the takecare edge name in mutations.
	EdgeTAKECARE = "TAKECARE"
	// EdgeGETSTATUS holds the string denoting the getstatus edge name in mutations.
	EdgeGETSTATUS = "GETSTATUS"

	// Table holds the table name of the antenatal in the database.
	Table = "antenatals"
	// GETMOMTable is the table the holds the GETMOM relation/edge.
	GETMOMTable = "antenatals"
	// GETMOMInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	GETMOMInverseTable = "patients"
	// GETMOMColumn is the table column denoting the GETMOM relation/edge.
	GETMOMColumn = "patient_setmom"
	// TAKECARETable is the table the holds the TAKECARE relation/edge.
	TAKECARETable = "antenatals"
	// TAKECAREInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TAKECAREInverseTable = "users"
	// TAKECAREColumn is the table column denoting the TAKECARE relation/edge.
	TAKECAREColumn = "user_care"
	// GETSTATUSTable is the table the holds the GETSTATUS relation/edge.
	GETSTATUSTable = "antenatals"
	// GETSTATUSInverseTable is the table name for the Babystatus entity.
	// It exists in this package in order to avoid circular dependency with the "babystatus" package.
	GETSTATUSInverseTable = "babystatuses"
	// GETSTATUSColumn is the table column denoting the GETSTATUS relation/edge.
	GETSTATUSColumn = "babystatus_setstatus"
)

// Columns holds all SQL columns for antenatal fields.
var Columns = []string{
	FieldID,
	FieldADDEDTIME,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Antenatal type.
var ForeignKeys = []string{
	"babystatus_setstatus",
	"patient_setmom",
	"user_care",
}

var (
	// DefaultADDEDTIME holds the default value on creation for the ADDED_TIME field.
	DefaultADDEDTIME func() time.Time
)
