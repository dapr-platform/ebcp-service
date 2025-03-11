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


Table: v_ebcp_exhibition_hall_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] remarks                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] rooms                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "PHANbKITHikckZEaGdpdylaRJ",    "name": "QqKnNXYugSLnvlqTrkSuxuvEg",    "remarks": "AkLCZOcbnPgXiZHwVuRPxbCTe",    "rooms": 74}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_hall_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_hall_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_hall_info_FIELD_NAME_id = "id"

	Ebcp_exhibition_hall_info_FIELD_NAME_name = "name"

	Ebcp_exhibition_hall_info_FIELD_NAME_remarks = "remarks"

	Ebcp_exhibition_hall_info_FIELD_NAME_rooms = "rooms"
)

// Ebcp_exhibition_hall_info struct is a row record of the v_ebcp_exhibition_hall_info table in the  database
type Ebcp_exhibition_hall_info struct {
	ID string `json:"id"` //id

	Name string `json:"name"` //name

	Remarks string `json:"remarks"` //remarks

	Rooms any `json:"rooms"` //rooms

}

var Ebcp_exhibition_hall_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_hall_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_ebcp_exhibition_hall_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_hall_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "name",
			Comment:            `name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "remarks",
			Comment:            `remarks`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Remarks",
			GoFieldType:        "string",
			JSONFieldName:      "remarks",
			ProtobufFieldName:  "remarks",
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
func (e *Ebcp_exhibition_hall_info) TableName() string {
	return "v_ebcp_exhibition_hall_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_hall_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_hall_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_hall_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_hall_info) TableInfo() *TableInfo {
	return Ebcp_exhibition_hall_infoTableInfo
}
