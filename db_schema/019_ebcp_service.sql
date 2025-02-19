-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 展馆表
CREATE TABLE o_ebcp_exhibition_hall (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
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
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    exhibition_hall_id VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1, -- 1: 正常, 2: 未使用, 3: 维修
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_room IS '展厅表';
COMMENT ON COLUMN o_ebcp_exhibition_room.name IS '展厅名称';
COMMENT ON COLUMN o_ebcp_exhibition_room.location IS '展厅位置';
COMMENT ON COLUMN o_ebcp_exhibition_room.exhibition_hall_id IS '所属展馆ID';
COMMENT ON COLUMN o_ebcp_exhibition_room.status IS '状态（1: 正常, 2: 未使用, 3: 维修）';
COMMENT ON COLUMN o_ebcp_exhibition_room.remarks IS '备注';

-- 展区表
CREATE TABLE o_ebcp_exhibition_area (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    exhibition_room_id VARCHAR(32) NOT NULL,
    current_exhibition_name VARCHAR(255),
    current_exhibition_start_time TIMESTAMP,
    current_exhibition_end_time TIMESTAMP,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_area IS '展区表';
COMMENT ON COLUMN o_ebcp_exhibition_area.name IS '展区名称';
COMMENT ON COLUMN o_ebcp_exhibition_area.exhibition_room_id IS '所属展厅ID';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_name IS '当前展览名称';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_start_time IS '当前展览开始时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_end_time IS '当前展览结束时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.remarks IS '备注';

-- 展项表
CREATE TABLE o_ebcp_exhibition_item (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    exhibition_area_id VARCHAR(32) NOT NULL,
    type VARCHAR(50) NOT NULL, -- 多媒体、static
    status INTEGER NOT NULL DEFAULT 1, -- 1: 启动, 2: 停止, 3: 故障
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_item IS '展项表';
COMMENT ON COLUMN o_ebcp_exhibition_item.name IS '展项名称';
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_area_id IS '所属展区ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.type IS '展项类型（多媒体、static）';
COMMENT ON COLUMN o_ebcp_exhibition_item.status IS '状态（1: 启动, 2: 停止, 3: 故障）';
COMMENT ON COLUMN o_ebcp_exhibition_item.remarks IS '备注';

-- 摄像头表
CREATE TABLE o_ebcp_camera (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    device_no VARCHAR(255) NOT NULL,
    main_stream_url VARCHAR(1024) NOT NULL,
    sub_stream_url VARCHAR(1024) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1, -- 1: 正常, 2: 故障
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
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player IS '播放设备表';
COMMENT ON COLUMN o_ebcp_player.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_player.ip_address IS 'IP地址';
COMMENT ON COLUMN o_ebcp_player.port IS '端口';

-- 展项关联配置表
CREATE TABLE o_ebcp_item_device_relation (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    exhibition_item_id VARCHAR(32) NOT NULL,
    device_type INTEGER NOT NULL, -- 1: 播放设备, 2: 摄像头, 3: 照明回路
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
    created_time TIMESTAMP NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL,
    exhibition_item_id VARCHAR(32) NOT NULL,
    schedule_time TIME NOT NULL,
    task_type INTEGER NOT NULL, -- 1: 启动, 2: 停止
    cycle_type INTEGER NOT NULL, -- 1:工作日, 2:周末, 3:节假日, 4:闭馆日, 5:每天
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
    r.status AS room_status,
    a.id AS area_id,
    a.name AS area_name,
    a.current_exhibition_name,
    i.id AS item_id,
    i.name AS item_name,
    i.type AS item_type,
    i.status AS item_status
FROM o_ebcp_exhibition_hall h
LEFT JOIN o_ebcp_exhibition_room r ON r.exhibition_hall_id = h.id
LEFT JOIN o_ebcp_exhibition_area a ON a.exhibition_room_id = r.id
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
            'room_location', er.location,
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
                WHERE ea.exhibition_room_id = er.id
            )
        )
    ) AS rooms
FROM 
    o_ebcp_exhibition_hall eh
LEFT JOIN 
    o_ebcp_exhibition_room er ON er.exhibition_hall_id = eh.id
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
    ea.exhibition_room_id,
    er.location,
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
    o_ebcp_exhibition_item ei ON ei.exhibition_area_id = ea.id
LEFT JOIN 
    o_ebcp_exhibition_room er ON er.id = ea.exhibition_room_id
GROUP BY 
    ea.id, ea.name, ea.current_exhibition_name, ea.exhibition_room_id, ea.current_exhibition_start_time, ea.current_exhibition_end_time, er.location;

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
DROP TABLE IF EXISTS o_ebcp_player;
DROP TABLE IF EXISTS o_ebcp_camera;
DROP TABLE IF EXISTS o_ebcp_exhibition_item;
DROP TABLE IF EXISTS o_ebcp_exhibition_area;
DROP TABLE IF EXISTS o_ebcp_exhibition_room;
DROP TABLE IF EXISTS o_ebcp_exhibition_hall;

-- +goose StatementEnd
