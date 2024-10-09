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


Table: o_ebcp_schedule_action
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] action_type                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] target_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] operation_details                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "WlhgTWwFxqDxvCXyLCeEEqNZM",    "action_type": 87,    "target_id": "tthPqXcEqrfOKwYXJWIKGRIZd",    "operation_details": "sUgUhaVKjsQLhxBPwsVAHvpoD"}



*/

var (
	Ebcp_schedule_action_FIELD_NAME_id = "id"

	Ebcp_schedule_action_FIELD_NAME_action_type = "action_type"

	Ebcp_schedule_action_FIELD_NAME_target_id = "target_id"

	Ebcp_schedule_action_FIELD_NAME_operation_details = "operation_details"
)

// Ebcp_schedule_action struct is a row record of the o_ebcp_schedule_action table in the  database
type Ebcp_schedule_action struct {
	ID               string `json:"id"`                //动作唯一标识
	ActionType       int32  `json:"action_type"`       //动作类型
	TargetID         string `json:"target_id"`         //目标设备或展项
	OperationDetails string `json:"operation_details"` //操作细节

}

var Ebcp_schedule_actionTableInfo = &TableInfo{
	Name: "o_ebcp_schedule_action",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `动作唯一标识`,
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
			Name:               "action_type",
			Comment:            `动作类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ActionType",
			GoFieldType:        "int32",
			JSONFieldName:      "action_type",
			ProtobufFieldName:  "action_type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "target_id",
			Comment:            `目标设备或展项`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "TargetID",
			GoFieldType:        "string",
			JSONFieldName:      "target_id",
			ProtobufFieldName:  "target_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "operation_details",
			Comment:            `操作细节`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "OperationDetails",
			GoFieldType:        "string",
			JSONFieldName:      "operation_details",
			ProtobufFieldName:  "operation_details",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_schedule_action) TableName() string {
	return "o_ebcp_schedule_action"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_schedule_action) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_schedule_action) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_schedule_action) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_schedule_action) TableInfo() *TableInfo {
	return Ebcp_schedule_actionTableInfo
}
