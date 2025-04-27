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


Table: v_ebcp_player_program_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] program_id                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] program_index                                  INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] player_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] player_name                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] player_ip_address                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] player_port                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 8] player_status                                  INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 9] item_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] item_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[14] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[15] medias                                         JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "qORhJvpfZmTPSCguHyjrEtylU",    "name": "glZhTMTBafBbQMvgWfgWqCxGB",    "program_id": "bdSJrGbxBHPRAkNTBfiXtsxnD",    "program_index": 67,    "player_id": "VbQSRdPTpoXIhdGtiNlAooixZ",    "player_name": "JYmvQARfxsSdglLUblPsrXcsD",    "player_ip_address": "FhUXBynCFeDTDmZLoSjmXdGVA",    "player_port": 17,    "player_status": 13,    "item_id": "YkIYWaKnbZQuJyAUGiRWhFUSS",    "item_name": "RlVRhwyUmMevMMadiBaUdHckT",    "room_id": "mMmENCwmPMHhRJedhyXaKylJd",    "room_name": "AFyxudyuDjevFpxvpHkuDhIjN",    "exhibition_id": "xwiDfUuIDdgGeUmNihiYlUqXO",    "exhibition_name": "BcrdLAjRfvuLbIWTZUBcjxVhw",    "medias": 79}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_player_program_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_player_program_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_player_program_info_FIELD_NAME_id = "id"

	Ebcp_player_program_info_FIELD_NAME_name = "name"

	Ebcp_player_program_info_FIELD_NAME_program_id = "program_id"

	Ebcp_player_program_info_FIELD_NAME_program_index = "program_index"

	Ebcp_player_program_info_FIELD_NAME_player_id = "player_id"

	Ebcp_player_program_info_FIELD_NAME_player_name = "player_name"

	Ebcp_player_program_info_FIELD_NAME_player_ip_address = "player_ip_address"

	Ebcp_player_program_info_FIELD_NAME_player_port = "player_port"

	Ebcp_player_program_info_FIELD_NAME_player_status = "player_status"

	Ebcp_player_program_info_FIELD_NAME_item_id = "item_id"

	Ebcp_player_program_info_FIELD_NAME_item_name = "item_name"

	Ebcp_player_program_info_FIELD_NAME_room_id = "room_id"

	Ebcp_player_program_info_FIELD_NAME_room_name = "room_name"

	Ebcp_player_program_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_player_program_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_player_program_info_FIELD_NAME_medias = "medias"
)

// Ebcp_player_program_info struct is a row record of the v_ebcp_player_program_info table in the  database
type Ebcp_player_program_info struct {
	ID string `json:"id"` //id

	Name string `json:"name"` //name

	ProgramID string `json:"program_id"` //program_id

	ProgramIndex int32 `json:"program_index"` //program_index

	PlayerID string `json:"player_id"` //player_id

	PlayerName string `json:"player_name"` //player_name

	PlayerIPAddress string `json:"player_ip_address"` //player_ip_address

	PlayerPort int32 `json:"player_port"` //player_port

	PlayerStatus int32 `json:"player_status"` //player_status

	ItemID string `json:"item_id"` //item_id

	ItemName string `json:"item_name"` //item_name

	RoomID string `json:"room_id"` //room_id

	RoomName string `json:"room_name"` //room_name

	ExhibitionID string `json:"exhibition_id"` //exhibition_id

	ExhibitionName string `json:"exhibition_name"` //exhibition_name

	Medias any `json:"medias"` //medias

}

var Ebcp_player_program_infoTableInfo = &TableInfo{
	Name: "v_ebcp_player_program_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_ebcp_player_program_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_player_program_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "program_id",
			Comment:            `program_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ProgramID",
			GoFieldType:        "string",
			JSONFieldName:      "program_id",
			ProtobufFieldName:  "program_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "program_index",
			Comment:            `program_index`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ProgramIndex",
			GoFieldType:        "int32",
			JSONFieldName:      "program_index",
			ProtobufFieldName:  "program_index",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "player_id",
			Comment:            `player_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "PlayerID",
			GoFieldType:        "string",
			JSONFieldName:      "player_id",
			ProtobufFieldName:  "player_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "player_name",
			Comment:            `player_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "PlayerName",
			GoFieldType:        "string",
			JSONFieldName:      "player_name",
			ProtobufFieldName:  "player_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "player_ip_address",
			Comment:            `player_ip_address`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "PlayerIPAddress",
			GoFieldType:        "string",
			JSONFieldName:      "player_ip_address",
			ProtobufFieldName:  "player_ip_address",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "player_port",
			Comment:            `player_port`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "PlayerPort",
			GoFieldType:        "int32",
			JSONFieldName:      "player_port",
			ProtobufFieldName:  "player_port",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "player_status",
			Comment:            `player_status`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "PlayerStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "player_status",
			ProtobufFieldName:  "player_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "item_id",
			Comment:            `item_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ItemID",
			GoFieldType:        "string",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "item_name",
			Comment:            `item_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ItemName",
			GoFieldType:        "string",
			JSONFieldName:      "item_name",
			ProtobufFieldName:  "item_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "room_id",
			Comment:            `room_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RoomID",
			GoFieldType:        "string",
			JSONFieldName:      "room_id",
			ProtobufFieldName:  "room_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "room_name",
			Comment:            `room_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoomName",
			GoFieldType:        "string",
			JSONFieldName:      "room_name",
			ProtobufFieldName:  "room_name",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "exhibition_id",
			Comment:            `exhibition_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_id",
			ProtobufFieldName:  "exhibition_id",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "exhibition_name",
			Comment:            `exhibition_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ExhibitionName",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_name",
			ProtobufFieldName:  "exhibition_name",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "medias",
			Comment:            `medias`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Medias",
			GoFieldType:        "any",
			JSONFieldName:      "medias",
			ProtobufFieldName:  "medias",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_player_program_info) TableName() string {
	return "v_ebcp_player_program_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_player_program_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_player_program_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_player_program_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_player_program_info) TableInfo() *TableInfo {
	return Ebcp_player_program_infoTableInfo
}
