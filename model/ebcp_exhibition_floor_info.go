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


Table: v_ebcp_exhibition_floor_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] short_name                                     VARCHAR(10)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 3] sort_order                                     INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "TLSAOSjvrPSJNTAvIQyUxTLsU",    "name": "byroprPGkkLIiFdKZSOyUKvdr",    "short_name": "RBTFQWcymQVrxIOlUirdIvxuD",    "sort_order": 2}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_floor_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_floor_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_floor_info_FIELD_NAME_id = "id"

	Ebcp_exhibition_floor_info_FIELD_NAME_name = "name"

	Ebcp_exhibition_floor_info_FIELD_NAME_short_name = "short_name"

	Ebcp_exhibition_floor_info_FIELD_NAME_sort_order = "sort_order"
)

// Ebcp_exhibition_floor_info struct is a row record of the v_ebcp_exhibition_floor_info table in the  database
type Ebcp_exhibition_floor_info struct {
	ID string `json:"id"` //id

	Name string `json:"name"` //name

	ShortName string `json:"short_name"` //short_name

	SortOrder int32 `json:"sort_order"` //sort_order

}

var Ebcp_exhibition_floor_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_floor_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_ebcp_exhibition_floor_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_floor_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "short_name",
			Comment:            `short_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10,
			GoFieldName:        "ShortName",
			GoFieldType:        "string",
			JSONFieldName:      "short_name",
			ProtobufFieldName:  "short_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "sort_order",
			Comment:            `sort_order`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SortOrder",
			GoFieldType:        "int32",
			JSONFieldName:      "sort_order",
			ProtobufFieldName:  "sort_order",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_floor_info) TableName() string {
	return "v_ebcp_exhibition_floor_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_floor_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_floor_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_floor_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_floor_info) TableInfo() *TableInfo {
	return Ebcp_exhibition_floor_infoTableInfo
}
