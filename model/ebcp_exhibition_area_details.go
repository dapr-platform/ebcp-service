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


Table: v_ebcp_exhibition_area_details
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] area_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] current_exhibition_name                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] exhibition_room_id                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] exhibition_items                               JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "qFPyHaxHHZJiTylSfJuJCqfno",    "area_name": "PZSeOBivytADCvMdaIhuMaQEU",    "current_exhibition_name": "ioWRRjkHARaLxEggNDTWXAGCT",    "exhibition_room_id": "geCwnILQsOeNjbETyOjaJsSWg",    "exhibition_items": 72}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_area_details does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_area_details primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_area_details_FIELD_NAME_id = "id"

	Ebcp_exhibition_area_details_FIELD_NAME_area_name = "area_name"

	Ebcp_exhibition_area_details_FIELD_NAME_current_exhibition_name = "current_exhibition_name"

	Ebcp_exhibition_area_details_FIELD_NAME_exhibition_room_id = "exhibition_room_id"

	Ebcp_exhibition_area_details_FIELD_NAME_exhibition_items = "exhibition_items"
)

// Ebcp_exhibition_area_details struct is a row record of the v_ebcp_exhibition_area_details table in the  database
type Ebcp_exhibition_area_details struct {
	ID                    string `json:"id"`                      //id
	AreaName              string `json:"area_name"`               //area_name
	CurrentExhibitionName string `json:"current_exhibition_name"` //current_exhibition_name
	ExhibitionRoomID      string `json:"exhibition_room_id"`      //exhibition_room_id
	ExhibitionItems       any    `json:"exhibition_items"`        //exhibition_items

}

var Ebcp_exhibition_area_detailsTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_area_details",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_ebcp_exhibition_area_details does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_area_details primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "area_name",
			Comment:            `area_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AreaName",
			GoFieldType:        "string",
			JSONFieldName:      "area_name",
			ProtobufFieldName:  "area_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "current_exhibition_name",
			Comment:            `current_exhibition_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "CurrentExhibitionName",
			GoFieldType:        "string",
			JSONFieldName:      "current_exhibition_name",
			ProtobufFieldName:  "current_exhibition_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "exhibition_room_id",
			Comment:            `exhibition_room_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionRoomID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_room_id",
			ProtobufFieldName:  "exhibition_room_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "exhibition_items",
			Comment:            `exhibition_items`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "ExhibitionItems",
			GoFieldType:        "any",
			JSONFieldName:      "exhibition_items",
			ProtobufFieldName:  "exhibition_items",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_area_details) TableName() string {
	return "v_ebcp_exhibition_area_details"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_area_details) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_area_details) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_area_details) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_area_details) TableInfo() *TableInfo {
	return Ebcp_exhibition_area_detailsTableInfo
}
