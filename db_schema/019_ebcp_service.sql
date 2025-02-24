-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 展馆表
CREATE TABLE o_ebcp_exhibition_hall (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_hall IS '展馆表';
COMMENT ON COLUMN o_ebcp_exhibition_hall.name IS '展馆名称';
COMMENT ON COLUMN o_ebcp_exhibition_hall.remarks IS '备注';

-- 展厅表
CREATE TABLE o_ebcp_exhibition_room (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(32) NOT NULL,
    exhibition_hall_id VARCHAR(32) NOT NULL,
    floor VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_room IS '展厅表';
COMMENT ON COLUMN o_ebcp_exhibition_room.name IS '展厅名称';
COMMENT ON COLUMN o_ebcp_exhibition_room.location IS '展厅位置(西侧，西北侧)';
COMMENT ON COLUMN o_ebcp_exhibition_room.exhibition_hall_id IS '所属展馆ID';
COMMENT ON COLUMN o_ebcp_exhibition_room.floor IS '楼层(B1,F1,F2,F3...)';
COMMENT ON COLUMN o_ebcp_exhibition_room.status IS '状态（1: 正常, 2: 未使用, 3: 维修）';
COMMENT ON COLUMN o_ebcp_exhibition_room.remarks IS '备注';

-- 展区表
CREATE TABLE o_ebcp_exhibition_area (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    current_exhibition_name VARCHAR(255),
    current_exhibition_start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    current_exhibition_end_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_area IS '展区表';
COMMENT ON COLUMN o_ebcp_exhibition_area.name IS '展区名称';
COMMENT ON COLUMN o_ebcp_exhibition_area.location IS '展区位置';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_name IS '当前展览名称';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_start_time IS '当前展览开始时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_end_time IS '当前展览结束时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.remarks IS '备注';

-- 展厅展区关联表
CREATE TABLE o_ebcp_room_area_relation (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    exhibition_room_id VARCHAR(32) NOT NULL,
    exhibition_area_id VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_room_area_relation IS '展厅展区关联表';
COMMENT ON COLUMN o_ebcp_room_area_relation.exhibition_room_id IS '展厅ID';
COMMENT ON COLUMN o_ebcp_room_area_relation.exhibition_area_id IS '展区ID';

-- 展项表
CREATE TABLE o_ebcp_exhibition_item (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    exhibition_area_id VARCHAR(32) NOT NULL,
    type VARCHAR(50) NOT NULL,
    display_format VARCHAR(100) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_item IS '展项表';
COMMENT ON COLUMN o_ebcp_exhibition_item.name IS '展项名称';
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_area_id IS '所属展区ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.type IS '展项类型（media、static）';
COMMENT ON COLUMN o_ebcp_exhibition_item.display_format IS '展示格式';
COMMENT ON COLUMN o_ebcp_exhibition_item.status IS '状态（1: 启动, 2: 停止, 3: 故障）';
COMMENT ON COLUMN o_ebcp_exhibition_item.remarks IS '备注';

-- 摄像头表
CREATE TABLE o_ebcp_camera (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    device_no VARCHAR(255) NOT NULL,
    main_stream_url VARCHAR(1024) NOT NULL,
    sub_stream_url VARCHAR(1024) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_camera IS '摄像头表';
COMMENT ON COLUMN o_ebcp_camera.name IS '摄像头名称';
COMMENT ON COLUMN o_ebcp_camera.device_no IS '设备编号';
COMMENT ON COLUMN o_ebcp_camera.main_stream_url IS '主码流URL';
COMMENT ON COLUMN o_ebcp_camera.sub_stream_url IS '辅码流URL';
COMMENT ON COLUMN o_ebcp_camera.status IS '状态(1: 正常, 2: 故障)';

-- 播放设备表
CREATE TABLE o_ebcp_player (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    version VARCHAR(255),
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player IS '播放设备表';
COMMENT ON COLUMN o_ebcp_player.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_player.ip_address IS 'IP地址';
COMMENT ON COLUMN o_ebcp_player.port IS '端口';
COMMENT ON COLUMN o_ebcp_player.version IS '版本号';
COMMENT ON COLUMN o_ebcp_player.status IS '状态(1: 正常, 2: 故障)';

-- 播放设备节目表
CREATE TABLE o_ebcp_player_program (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    program_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    player_id VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player_program IS '播放设备节目表';
COMMENT ON COLUMN o_ebcp_player_program.program_id IS '节目ID';
COMMENT ON COLUMN o_ebcp_player_program.name IS '节目名称';
COMMENT ON COLUMN o_ebcp_player_program.player_id IS '播放设备ID';

-- 展项关联配置表
CREATE TABLE o_ebcp_item_device_relation (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    exhibition_item_id VARCHAR(32) NOT NULL,
    device_type INTEGER NOT NULL,
    device_id VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_item_device_relation IS '展项关联配置表';
COMMENT ON COLUMN o_ebcp_item_device_relation.exhibition_item_id IS '展项ID';
COMMENT ON COLUMN o_ebcp_item_device_relation.device_type IS '关联设备类型(1: 播放设备, 2: 摄像头, 3: 照明回路)';
COMMENT ON COLUMN o_ebcp_item_device_relation.device_id IS '关联设备ID';

-- 展项定时任务表
CREATE TABLE o_ebcp_item_schedule (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    exhibition_item_id VARCHAR(32) NOT NULL,
    schedule_time TIME NOT NULL,
    task_type INTEGER NOT NULL,
    cycle_type INTEGER NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_item_schedule IS '展项定时任务表';
COMMENT ON COLUMN o_ebcp_item_schedule.exhibition_item_id IS '展项ID';
COMMENT ON COLUMN o_ebcp_item_schedule.schedule_time IS '任务时间';
COMMENT ON COLUMN o_ebcp_item_schedule.task_type IS '任务类型(1: 启动, 2: 停止)';
COMMENT ON COLUMN o_ebcp_item_schedule.cycle_type IS '循环方式(1:工作日, 2:周末, 3:节假日, 4:闭馆日, 5:每天)';

-- 创建视图
CREATE VIEW v_ebcp_exhibition_info AS
SELECT 
    h.id AS hall_id,
    h.name AS hall_name,
    r.id AS room_id,
    r.name AS room_name,
    f.dict_name AS room_floor,
    l.dict_name AS room_location,
    r.status AS room_status,
    a.id AS area_id,
    a.name AS area_name,
    a.location AS area_location,
    a.current_exhibition_name,
    i.id AS item_id,
    i.name AS item_name,
    i.type AS item_type,
    i.status AS item_status
FROM o_ebcp_exhibition_hall h
LEFT JOIN o_ebcp_exhibition_room r ON r.exhibition_hall_id = h.id
LEFT JOIN o_ops_dict f ON f.id = r.floor AND f.dict_type = 'floor_type'
LEFT JOIN o_ops_dict l ON l.id = r.location AND l.dict_type = 'location_type'
LEFT JOIN o_ebcp_room_area_relation ra ON ra.exhibition_room_id = r.id
LEFT JOIN o_ebcp_exhibition_area a ON a.id = ra.exhibition_area_id
LEFT JOIN o_ebcp_exhibition_item i ON i.exhibition_area_id = a.id;

COMMENT ON VIEW v_ebcp_exhibition_info IS '展览信息视图';

-- 新增展馆详细视图
CREATE VIEW v_ebcp_exhibition_hall_details AS
SELECT 
    eh.id AS id,
    eh.name AS hall_name,
    eh.remarks AS hall_description,
    json_agg(
        json_build_object(
            'room_id', er.id,
            'room_name', er.name,
            'room_location', l.dict_name,
            'room_status', er.status,
            'areas', (
                SELECT json_agg(
                    json_build_object(
                        'area_id', ea.id,
                        'area_name', ea.name,
                        'current_exhibition_name', ea.current_exhibition_name
                    )
                )
                FROM o_ebcp_exhibition_area ea
                WHERE ea.location = er.location
            )
        )
    ) AS rooms
FROM 
    o_ebcp_exhibition_hall eh
LEFT JOIN 
    o_ebcp_exhibition_room er ON er.exhibition_hall_id = eh.id
LEFT JOIN 
    o_ops_dict l ON l.id = er.location AND l.dict_type = 'location_type'
GROUP BY 
    eh.id, eh.name, eh.remarks;

COMMENT ON VIEW v_ebcp_exhibition_hall_details IS '展馆详细视图，包含展馆信息及其关联的展厅和展区信息（JSON格式）';

-- 新增展区详细视图
CREATE VIEW v_ebcp_exhibition_area_details AS
SELECT 
    ea.id AS id,
    ea.name AS area_name,
    ea.current_exhibition_name,
    ea.current_exhibition_start_time,
    ea.current_exhibition_end_time,
    l.dict_name AS location,
    json_agg(
        json_build_object(
            'item_id', ei.id,
            'item_name', ei.name,
            'item_status', ei.status,
            'item_type', ei.type,
            'item_remarks', ei.remarks
        )
    ) AS exhibition_items
FROM 
    o_ebcp_exhibition_area ea
LEFT JOIN 
    o_ops_dict l ON l.id = ea.location AND l.dict_type = 'location_type'
LEFT JOIN 
    o_ebcp_exhibition_item ei ON ei.exhibition_area_id = ea.id
GROUP BY 
    ea.id, ea.name, ea.current_exhibition_name, l.dict_name, ea.current_exhibition_start_time, ea.current_exhibition_end_time;

COMMENT ON VIEW v_ebcp_exhibition_area_details IS '展区详细视图，包含展区信息及其关联的所有展项信息（JSON格式），展项信息包括名字、状态、类型和备注';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP VIEW IF EXISTS v_ebcp_exhibition_area_details;
DROP VIEW IF EXISTS v_ebcp_exhibition_hall_details;
DROP VIEW IF EXISTS v_ebcp_exhibition_info;

DROP TABLE IF EXISTS o_ebcp_item_schedule;
DROP TABLE IF EXISTS o_ebcp_item_device_relation;
DROP TABLE IF EXISTS o_ebcp_player_program;
DROP TABLE IF EXISTS o_ebcp_player;
DROP TABLE IF EXISTS o_ebcp_camera;
DROP TABLE IF EXISTS o_ebcp_exhibition_item;
DROP TABLE IF EXISTS o_ebcp_room_area_relation;
DROP TABLE IF EXISTS o_ebcp_exhibition_area;
DROP TABLE IF EXISTS o_ebcp_exhibition_room;
DROP TABLE IF EXISTS o_ebcp_exhibition_hall;

-- +goose StatementEnd
