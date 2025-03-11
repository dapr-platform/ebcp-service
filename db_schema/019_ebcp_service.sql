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

-- 展览表
CREATE TABLE o_ebcp_exhibition (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    remarks TEXT,
    hall_id VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition IS '展览表';
COMMENT ON COLUMN o_ebcp_exhibition.name IS '展览名称';
COMMENT ON COLUMN o_ebcp_exhibition.start_time IS '开始时间';
COMMENT ON COLUMN o_ebcp_exhibition.end_time IS '结束时间';
COMMENT ON COLUMN o_ebcp_exhibition.remarks IS '备注';
COMMENT ON COLUMN o_ebcp_exhibition.hall_id IS '所属展馆ID';
COMMENT ON COLUMN o_ebcp_exhibition.status IS '状态（1: 运行中, 2: 筹备中, 3: 已结束）';

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
    exhibition_id VARCHAR(32),
    status INTEGER NOT NULL DEFAULT 1,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_room IS '展厅表';
COMMENT ON COLUMN o_ebcp_exhibition_room.name IS '展厅名称';
COMMENT ON COLUMN o_ebcp_exhibition_room.location IS '展厅位置(西侧，西北侧)';
COMMENT ON COLUMN o_ebcp_exhibition_room.exhibition_hall_id IS '所属展馆ID';
COMMENT ON COLUMN o_ebcp_exhibition_room.floor IS '楼层(B1,F1,F2,F3...)';
COMMENT ON COLUMN o_ebcp_exhibition_room.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN o_ebcp_exhibition_room.status IS '状态（1: 正常, 2: 未使用, 3: 维修）';
COMMENT ON COLUMN o_ebcp_exhibition_room.remarks IS '备注';

-- 展项表
CREATE TABLE o_ebcp_exhibition_item (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    exhibition_id VARCHAR(32) NOT NULL,
    exhibition_room_id VARCHAR(32) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status INTEGER NOT NULL DEFAULT 1,
    remarks TEXT,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_item IS '展项表';
COMMENT ON COLUMN o_ebcp_exhibition_item.name IS '展项名称';
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_room_id IS '所属展厅ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.type IS '展项类型（media、static）';
COMMENT ON COLUMN o_ebcp_exhibition_item.status IS '状态（1: 启动, 2: 停止, 3: 故障）';
COMMENT ON COLUMN o_ebcp_exhibition_item.remarks IS '备注';

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
    item_id VARCHAR(32),
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player IS '播放设备表';
COMMENT ON COLUMN o_ebcp_player.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_player.ip_address IS 'IP地址';
COMMENT ON COLUMN o_ebcp_player.port IS '端口';
COMMENT ON COLUMN o_ebcp_player.version IS '版本';
COMMENT ON COLUMN o_ebcp_player.item_id IS '所属展项ID';
COMMENT ON COLUMN o_ebcp_player.status IS '状态（1: 正常, 2: 离线, 3: 故障）';

-- 播放设备节目表
CREATE TABLE o_ebcp_player_program (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    player_id VARCHAR(32) NOT NULL,
    program_id VARCHAR(32) NOT NULL,
    program_index INTEGER NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player_program IS '播放设备节目表';
COMMENT ON COLUMN o_ebcp_player_program.name IS '节目名称';
COMMENT ON COLUMN o_ebcp_player_program.player_id IS '播放设备ID';
COMMENT ON COLUMN o_ebcp_player_program.program_id IS '节目ID';
COMMENT ON COLUMN o_ebcp_player_program.program_index IS '节目序号';

-- 中控设备表
CREATE TABLE o_ebcp_control_device (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    device_type VARCHAR(50) NOT NULL,
    item_id VARCHAR(32),
    room_id VARCHAR(32),
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_control_device IS '中控设备表';
COMMENT ON COLUMN o_ebcp_control_device.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_control_device.device_type IS '设备类型';
COMMENT ON COLUMN o_ebcp_control_device.item_id IS '所属展项ID';
COMMENT ON COLUMN o_ebcp_control_device.room_id IS '所属展厅ID';
COMMENT ON COLUMN o_ebcp_control_device.status IS '状态(1: 正常, 2: 故障)';

-- 展项关联配置表
CREATE TABLE o_ebcp_item_device_relation (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    exhibition_item_id VARCHAR(32) NOT NULL,
    device_type INTEGER NOT NULL,
    device_sub_type VARCHAR(50),
    device_id VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_item_device_relation IS '展项关联配置表';
COMMENT ON COLUMN o_ebcp_item_device_relation.exhibition_item_id IS '展项ID';
COMMENT ON COLUMN o_ebcp_item_device_relation.device_type IS '关联设备类型(1: 播放设备, 2: 中控设备)';
COMMENT ON COLUMN o_ebcp_item_device_relation.device_sub_type IS '关联设备子类型(中控设备时需要)';
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

-- 创建视图
CREATE VIEW v_ebcp_exhibition_info AS
SELECT 
    e.id AS id,
    e.name AS name,
    e.start_time AS start_time,
    e.end_time AS end_time,
    e.status AS status,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_room WHERE exhibition_id = e.id) AS total_room_count,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE exhibition_id = e.id) AS total_item_count,
    (
        SELECT json_agg(
            json_build_object(
                'id', r.id,
                'name', r.name,
                'floor', r.floor,
                'floor_value', (SELECT dict_value FROM o_ops_dict WHERE id = r.floor),
                'location', r.location,
                'location_value', (SELECT dict_value FROM o_ops_dict WHERE id = r.location),
                'status', r.status
            )
        )
        FROM o_ebcp_exhibition_room r
        WHERE r.exhibition_id = e.id
    ) AS rooms,
    (
        SELECT json_agg(
            json_build_object(
                'id', i.id,
                'name', i.name,
                'type', i.type,
                'status', i.status,
                'room_id', i.exhibition_room_id
            )
        )
        FROM o_ebcp_exhibition_item i
        WHERE i.exhibition_id = e.id
    ) AS items
FROM o_ebcp_exhibition e;

COMMENT ON VIEW v_ebcp_exhibition_info IS '展览信息视图';
COMMENT ON COLUMN v_ebcp_exhibition_info.id IS '展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_info.name IS '展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_info.start_time IS '展览开始时间';
COMMENT ON COLUMN v_ebcp_exhibition_info.end_time IS '展览结束时间';
COMMENT ON COLUMN v_ebcp_exhibition_info.status IS '展览状态（1: 运行中, 2: 筹备中, 3: 已结束）';    
COMMENT ON COLUMN v_ebcp_exhibition_info.total_room_count IS '展厅总数';
COMMENT ON COLUMN v_ebcp_exhibition_info.total_item_count IS '展项总数';
COMMENT ON COLUMN v_ebcp_exhibition_info.rooms IS '展览使用的所有展厅';
COMMENT ON COLUMN v_ebcp_exhibition_info.items IS '展览的所有展项';

CREATE VIEW v_ebcp_exhibition_area_info AS
SELECT 
    e.id AS id,
    e.name AS name,
    e.start_time AS start_time,
    e.end_time AS end_time,
    h.id AS hall_id,
    h.name AS hall_name,
    r.id AS room_id,
    r.name AS room_name,
    r.floor AS room_floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = r.floor) AS room_floor_value,
    r.location AS room_location,
    (SELECT dict_value FROM o_ops_dict WHERE id = r.location) AS room_location_value,
    r.status AS room_status,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_room WHERE exhibition_id = e.id) AS total_room_count,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE exhibition_id = e.id) AS total_item_count,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE exhibition_room_id = r.id) AS room_item_count
FROM o_ebcp_exhibition e
JOIN o_ebcp_exhibition_room r ON r.exhibition_id = e.id
JOIN o_ebcp_exhibition_hall h ON h.id = r.exhibition_hall_id;

COMMENT ON VIEW v_ebcp_exhibition_area_info IS '展览区域信息视图';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.id IS '展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.name IS '展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.start_time IS '展览开始时间';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.end_time IS '展览结束时间';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.hall_id IS '展馆ID';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.hall_name IS '展馆名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_id IS '展厅ID';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_name IS '展厅名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_floor IS '展厅楼层';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_floor_value IS '展厅楼层值';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_location IS '展厅位置';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_location_value IS '展厅位置值';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_status IS '展厅状态';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.total_room_count IS '展厅总数';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.total_item_count IS '展项总数';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_item_count IS '展厅展项总数';

-- 展馆详细视图
CREATE VIEW v_ebcp_exhibition_hall_info AS
SELECT 
    eh.id AS id,
    eh.name AS name,
    eh.remarks AS remarks,
    json_agg(
        json_build_object(
            'id', er.id,
            'name', er.name,
            'floor', er.floor,
            'floor_value', (SELECT dict_value FROM o_ops_dict WHERE id = er.floor),
            'location', er.location,
            'location_value', (SELECT dict_value FROM o_ops_dict WHERE id = er.location),
            'status', er.status,
            'exhibition_id', er.exhibition_id,
            'exhibition_name', e.name,
            'exhibition_start_time', e.start_time,
            'exhibition_end_time', e.end_time,
            'items', (
                SELECT json_agg(
                    json_build_object(
                        'id', ei.id,
                        'name', ei.name,
                        'type', ei.type,
                        'status', ei.status,
                        'remarks', ei.remarks
                    )
                )
                FROM o_ebcp_exhibition_item ei
                WHERE ei.exhibition_room_id = er.id
            )
        )
    ) AS rooms
FROM 
    o_ebcp_exhibition_hall eh
LEFT JOIN 
    o_ebcp_exhibition_room er ON er.exhibition_hall_id = eh.id
LEFT JOIN 
    o_ebcp_exhibition e ON e.id = er.exhibition_id
GROUP BY 
    eh.id, eh.name, eh.remarks;

COMMENT ON VIEW v_ebcp_exhibition_hall_info IS '展馆详细视图，包含展馆信息及其关联的展厅和展项信息（JSON格式）';

-- 展厅详细视图
CREATE VIEW v_ebcp_exhibition_room_info AS
SELECT 
    er.id AS id,
    er.name AS name,
    er.floor AS floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS floor_value,
    er.location AS location,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.location) AS location_value,
    er.status AS status,
    er.remarks AS remarks,
    eh.id AS hall_id,
    eh.name AS hall_name,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    e.start_time AS exhibition_start_time,
    e.end_time AS exhibition_end_time,
    e.status AS exhibition_status,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE exhibition_room_id = er.id) AS item_count,
    json_agg(
        json_build_object(
            'id', ei.id,
            'name', ei.name,
            'type', ei.type,
            'status', ei.status,
            'remarks', ei.remarks
        )
    ) AS items
FROM 
    o_ebcp_exhibition_room er
JOIN 
    o_ebcp_exhibition_hall eh ON er.exhibition_hall_id = eh.id
LEFT JOIN 
    o_ebcp_exhibition e ON er.exhibition_id = e.id
LEFT JOIN 
    o_ebcp_exhibition_item ei ON ei.exhibition_room_id = er.id
GROUP BY 
    er.id, er.name, er.floor, er.location, er.status, er.remarks, 
    eh.id, eh.name, e.id, e.name, e.start_time, e.end_time, e.status;

COMMENT ON VIEW v_ebcp_exhibition_room_info IS '展厅详细视图，包含展厅信息及其关联的展馆、展览和展项信息（JSON格式）';

-- 展项详细视图
CREATE VIEW v_ebcp_exhibition_item_info AS
SELECT 
    ei.id AS id,
    ei.name AS name,
    ei.type AS type,
    ei.status AS status,
    ei.remarks AS remarks,
    er.id AS room_id,
    er.name AS room_name,
    er.floor AS room_floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS room_floor_value,
    er.location AS room_location,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.location) AS room_location_value,
    eh.id AS hall_id,
    eh.name AS hall_name,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    (
        SELECT json_agg(
            json_build_object(
                'device_id', p.id,
                'device_name', p.name,
                'device_type', 'player',
                'ip_address', p.ip_address,
                'port', p.port,
                'status', p.status
            )
        )
        FROM o_ebcp_player p
        WHERE p.item_id = ei.id
    ) AS player_devices,
    (
        SELECT json_agg(
            json_build_object(
                'device_id', cd.id,
                'device_name', cd.name,
                'device_type', cd.device_type,
                'status', cd.status
            )
        )
        FROM o_ebcp_control_device cd
        WHERE cd.item_id = ei.id
    ) AS control_devices,
    (
        SELECT json_agg(
            json_build_object(
                'schedule_id', s.id,
                'schedule_time', s.schedule_time,
                'task_type', s.task_type,
                'cycle_type', s.cycle_type
            )
        )
        FROM o_ebcp_item_schedule s
        WHERE s.exhibition_item_id = ei.id
    ) AS schedules
FROM 
    o_ebcp_exhibition_item ei
JOIN 
    o_ebcp_exhibition_room er ON ei.exhibition_room_id = er.id
JOIN 
    o_ebcp_exhibition_hall eh ON er.exhibition_hall_id = eh.id
JOIN 
    o_ebcp_exhibition e ON ei.exhibition_id = e.id;

COMMENT ON VIEW v_ebcp_exhibition_item_info IS '展项详细视图，包含展项信息及其关联的展厅、展馆、展览、设备和定时任务信息（JSON格式）';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP VIEW IF EXISTS v_ebcp_exhibition_item_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_room_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_hall_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_area_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_info;

DROP TABLE IF EXISTS o_ebcp_item_schedule;
DROP TABLE IF EXISTS o_ebcp_item_device_relation;
DROP TABLE IF EXISTS o_ebcp_player_program;
DROP TABLE IF EXISTS o_ebcp_control_device;
DROP TABLE IF EXISTS o_ebcp_player;
DROP TABLE IF EXISTS o_ebcp_exhibition_item;
DROP TABLE IF EXISTS o_ebcp_exhibition_room;
DROP TABLE IF EXISTS o_ebcp_exhibition;
DROP TABLE IF EXISTS o_ebcp_exhibition_hall;

-- +goose StatementEnd
