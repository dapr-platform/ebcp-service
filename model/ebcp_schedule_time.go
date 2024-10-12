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


Table: o_ebcp_schedule_time
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] specific_time                                  TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] repeat_pattern                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "rNNAVbSkEZFukoPcJKLvFtWFU",    "type": 64,    "specific_time": 41,    "repeat_pattern": "uFEGEbhQbrTCoYcMGDBevkUtF"}



*/

var (
	Ebcp_schedule_time_FIELD_NAME_id = "id"

	Ebcp_schedule_time_FIELD_NAME_type = "type"

	Ebcp_schedule_time_FIELD_NAME_specific_time = "specific_time"

	Ebcp_schedule_time_FIELD_NAME_repeat_pattern = "repeat_pattern"
)

// Ebcp_schedule_time struct is a row record of the o_ebcp_schedule_time table in the  database
type Ebcp_schedule_time struct {
	ID            string           `json:"id"`             //时间配置唯一标识
	Type          int32            `json:"type"`           //时间类型
	SpecificTime  common.LocalTime `json:"specific_time"`  //具体时间
	RepeatPattern string           `json:"repeat_pattern"` //重复模式

}

var Ebcp_schedule_timeTableInfo = &TableInfo{
	Name: "o_ebcp_schedule_time",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `时间配置唯一标识`,
			Notes:              ``,
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
			Name:               "type",
			Comment:            `时间类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "specific_time",
			Comment:            `具体时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "SpecificTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "specific_time",
			ProtobufFieldName:  "specific_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "repeat_pattern",
			Comment:            `重复模式`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RepeatPattern",
			GoFieldType:        "string",
			JSONFieldName:      "repeat_pattern",
			ProtobufFieldName:  "repeat_pattern",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_schedule_time) TableName() string {
	return "o_ebcp_schedule_time"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_schedule_time) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_schedule_time) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_schedule_time) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_schedule_time) TableInfo() *TableInfo {
	return Ebcp_schedule_timeTableInfo
}
