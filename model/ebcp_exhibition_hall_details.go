package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: v_ebcp_exhibition_hall_details
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] hall_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] hall_description                               TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] rooms                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "MiqiAdHMaVJYBJAhnJbgEdBYF",    "hall_name": "knIlomVwRpuHfKLOwsyWEPhDD",    "hall_description": "fBUCrJBYRIPbKRtvvubVQjoiN",    "rooms": 62}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_hall_details does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_hall_details primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_hall_details_FIELD_NAME_id = "id"

	Ebcp_exhibition_hall_details_FIELD_NAME_hall_name = "hall_name"

	Ebcp_exhibition_hall_details_FIELD_NAME_hall_description = "hall_description"

	Ebcp_exhibition_hall_details_FIELD_NAME_rooms = "rooms"
)

// Ebcp_exhibition_hall_details struct is a row record of the v_ebcp_exhibition_hall_details table in the  database
type Ebcp_exhibition_hall_details struct {
	ID string `json:"id"` //id

	HallName string `json:"hall_name"` //hall_name

	HallDescription string `json:"hall_description"` //hall_description

	Rooms any `json:"rooms"` //rooms

}

var Ebcp_exhibition_hall_detailsTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_hall_details",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_ebcp_exhibition_hall_details does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_hall_details primary key column id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "hall_name",
			Comment:            `hall_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "HallName",
			GoFieldType:        "string",
			JSONFieldName:      "hall_name",
			ProtobufFieldName:  "hall_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "hall_description",
			Comment:            `hall_description`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "HallDescription",
			GoFieldType:        "string",
			JSONFieldName:      "hall_description",
			ProtobufFieldName:  "hall_description",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "rooms",
			Comment:            `rooms`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Rooms",
			GoFieldType:        "any",
			JSONFieldName:      "rooms",
			ProtobufFieldName:  "rooms",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_hall_details) TableName() string {
	return "v_ebcp_exhibition_hall_details"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_hall_details) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_hall_details) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_hall_details) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_hall_details) TableInfo() *TableInfo {
	return Ebcp_exhibition_hall_detailsTableInfo
}
