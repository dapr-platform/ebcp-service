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
    exhibition_hall_id VARCHAR(32),
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
    room_id VARCHAR(32) NOT NULL,
    type VARCHAR(50) NOT NULL,
    sub_type VARCHAR(50) NOT NULL DEFAULT '',
    export_info TEXT,
    status INTEGER NOT NULL DEFAULT 1,
    remarks TEXT,
    commands TEXT,
    device_id VARCHAR(32),
                                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_item IS '展项表';
COMMENT ON COLUMN o_ebcp_exhibition_item.name IS '展项名称';
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.room_id IS '所属展厅ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.type IS '展项类型（media、static）';
COMMENT ON COLUMN o_ebcp_exhibition_item.sub_type IS '展项子类型（static时需要,分为power,light）';
COMMENT ON COLUMN o_ebcp_exhibition_item.export_info IS '输出信息';
COMMENT ON COLUMN o_ebcp_exhibition_item.status IS '状态（0: 启动, 1: 暂停, 2: 停止）';
COMMENT ON COLUMN o_ebcp_exhibition_item.remarks IS '备注';
COMMENT ON COLUMN o_ebcp_exhibition_item.commands IS '命令列表,json格式,例如[{"name":"开启","command":"FA 01 01"},{"name":"关闭","command":"FA 01 02"}]';
COMMENT ON COLUMN o_ebcp_exhibition_item.device_id IS '中控设备ID';

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
    current_program_id VARCHAR(32),
    current_program_state INTEGER,
    volume INTEGER,
    sound_state INTEGER,
    status INTEGER NOT NULL DEFAULT 1,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player IS '播放设备表';
COMMENT ON COLUMN o_ebcp_player.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_player.ip_address IS 'IP地址';
COMMENT ON COLUMN o_ebcp_player.port IS '端口';
COMMENT ON COLUMN o_ebcp_player.version IS '版本';
COMMENT ON COLUMN o_ebcp_player.item_id IS '所属展项ID';
COMMENT ON COLUMN o_ebcp_player.current_program_id IS '当前节目ID';
COMMENT ON COLUMN o_ebcp_player.current_program_state IS '当前节目状态,0:播放,1:暂停,2:停止';
COMMENT ON COLUMN o_ebcp_player.volume IS '音量(0-100)';
COMMENT ON COLUMN o_ebcp_player.sound_state IS '音量状态,0:静音,1:非静音';
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
    state INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player_program IS '播放设备节目表';
COMMENT ON COLUMN o_ebcp_player_program.name IS '节目名称';
COMMENT ON COLUMN o_ebcp_player_program.player_id IS '播放设备ID';
COMMENT ON COLUMN o_ebcp_player_program.program_id IS '节目ID';
COMMENT ON COLUMN o_ebcp_player_program.program_index IS '节目序号';
COMMENT ON COLUMN o_ebcp_player_program.state IS '节目状态,0: 播放, 1: 暂停, 2: 停止';

-- 播放设备节目媒体表
CREATE TABLE o_ebcp_player_program_media (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    media_id VARCHAR(32) NOT NULL,
    media_name VARCHAR(255) NOT NULL,
    player_id VARCHAR(32) NOT NULL,
    program_id VARCHAR(32) NOT NULL,
    player_program_id VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_player_program_media IS '播放设备节目媒体表';
COMMENT ON COLUMN o_ebcp_player_program_media.media_id IS '媒体ID';
COMMENT ON COLUMN o_ebcp_player_program_media.media_name IS '媒体名称';
COMMENT ON COLUMN o_ebcp_player_program_media.player_id IS '播放设备ID(冗余)';
COMMENT ON COLUMN o_ebcp_player_program_media.program_id IS '节目ID(冗余)';
COMMENT ON COLUMN o_ebcp_player_program_media.player_program_id IS '播放设备节目ID(冗余)';

-- 中控设备表
CREATE TABLE o_ebcp_control_device (
    id VARCHAR(32) NOT NULL,
                                     created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    device_type VARCHAR(50) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    port INTEGER NOT NULL,
    version VARCHAR(255),
    room_id VARCHAR(32),
    status INTEGER NOT NULL DEFAULT 1,
    commands TEXT,
                                     PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_control_device IS '中控设备表';
COMMENT ON COLUMN o_ebcp_control_device.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_control_device.device_type IS '设备类型';
COMMENT ON COLUMN o_ebcp_control_device.room_id IS '所属展厅ID';
COMMENT ON COLUMN o_ebcp_control_device.status IS '状态(1: 正常, 2: 故障)';

-- 展项中控设备关联配置表
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
    item_id VARCHAR(32) NOT NULL,
    start_time VARCHAR(32) NOT NULL,
    stop_time VARCHAR(32) NOT NULL,
    start_date VARCHAR(32) NOT NULL DEFAULT '',
    stop_date VARCHAR(32) NOT NULL DEFAULT '',
    cycle_type INTEGER NOT NULL,
    action_type INTEGER NOT NULL,
    enabled INTEGER NOT NULL DEFAULT 1,
                                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_item_schedule IS '展项定时任务表';
COMMENT ON COLUMN o_ebcp_item_schedule.item_id IS '展项ID';
COMMENT ON COLUMN o_ebcp_item_schedule.start_time IS '开始时间';
COMMENT ON COLUMN o_ebcp_item_schedule.stop_time IS '停止时间'; 
COMMENT ON COLUMN o_ebcp_item_schedule.cycle_type IS '循环方式(1:工作日, 2:周末, 3:节假日, 4:闭馆日, 5:每天)';
COMMENT ON COLUMN o_ebcp_item_schedule.enabled IS '是否启用(0: 禁用, 1: 启用)';
COMMENT ON COLUMN o_ebcp_item_schedule.start_date IS '开始日期,暂时不用（预留寒暑假延长时间）';
COMMENT ON COLUMN o_ebcp_item_schedule.stop_date IS '停止日期,暂时不用（预留寒暑假延长时间）';
COMMENT ON COLUMN o_ebcp_item_schedule.action_type IS '动作类型(0: 停止, 1: 播放)';

-- 节假日日期表
CREATE TABLE o_ebcp_holiday_date (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date TIMESTAMP NOT NULL,
    name VARCHAR(100) NOT NULL,
    type INTEGER NOT NULL,
    year INTEGER NOT NULL,
    remarks VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX idx_holiday_date ON o_ebcp_holiday_date(date);

COMMENT ON TABLE o_ebcp_holiday_date IS '节假日日期表';
COMMENT ON COLUMN o_ebcp_holiday_date.date IS '日期';
COMMENT ON COLUMN o_ebcp_holiday_date.name IS '节假日名称';
COMMENT ON COLUMN o_ebcp_holiday_date.type IS '类型(1:法定节假日, 2:调休工作日, 3:周末调休, 4:闭馆日)';
COMMENT ON COLUMN o_ebcp_holiday_date.year IS '年份';
COMMENT ON COLUMN o_ebcp_holiday_date.remarks IS '备注';


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
                'floor_name', (SELECT dict_name FROM o_ops_dict WHERE id = r.floor),
                'location', r.location,
                'location_value', (SELECT dict_value FROM o_ops_dict WHERE id = r.location),
                'location_name', (SELECT dict_name FROM o_ops_dict WHERE id = r.location),
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
                'room_id', i.room_id
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
    r.id AS room_id,
    r.name AS room_name,
    r.floor AS room_floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = r.floor) AS room_floor_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = r.floor) AS room_floor_name,
    r.location AS room_location,
    (SELECT dict_value FROM o_ops_dict WHERE id = r.location) AS room_location_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = r.location) AS room_location_name,
    r.status AS room_status,
    r.remarks AS room_remarks,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    e.start_time AS exhibition_start_time,
    e.end_time AS exhibition_end_time,
    e.status AS exhibition_status,
    (
        SELECT json_agg(
            json_build_object(
                'id', i.id,
                'name', i.name,
                'type', i.type,
                'status', i.status,
                'remarks', i.remarks
            )
        )
        FROM o_ebcp_exhibition_item i
        WHERE i.room_id = r.id
    ) AS items
FROM o_ebcp_exhibition_room r
LEFT JOIN o_ebcp_exhibition e ON e.id = r.exhibition_id;

COMMENT ON VIEW v_ebcp_exhibition_area_info IS '展览区域信息视图';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_id IS '展厅ID';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_name IS '展厅名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_floor IS '展厅楼层';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_floor_value IS '展厅楼层值';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_floor_name IS '展厅楼层名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_location IS '展厅位置';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_location_value IS '展厅位置值';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_location_name IS '展厅位置名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_status IS '展厅状态';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.room_remarks IS '展厅备注';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.exhibition_id IS '展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.exhibition_name IS '展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.exhibition_start_time IS '展览开始时间';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.exhibition_end_time IS '展览结束时间';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.exhibition_status IS '展览状态';
COMMENT ON COLUMN v_ebcp_exhibition_area_info.items IS '展厅内的展项列表';

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
            'floor_name', (SELECT dict_name FROM o_ops_dict WHERE id = er.floor),
            'location', er.location,
            'location_value', (SELECT dict_value FROM o_ops_dict WHERE id = er.location),
            'location_name', (SELECT dict_name FROM o_ops_dict WHERE id = er.location),
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
                WHERE ei.room_id = er.id
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
COMMENT ON COLUMN v_ebcp_exhibition_hall_info.id IS '展馆ID';
COMMENT ON COLUMN v_ebcp_exhibition_hall_info.name IS '展馆名称';
COMMENT ON COLUMN v_ebcp_exhibition_hall_info.remarks IS '展馆备注';
COMMENT ON COLUMN v_ebcp_exhibition_hall_info.rooms IS '展馆下的展厅列表（JSON格式）';

-- 展厅详细视图
CREATE VIEW v_ebcp_exhibition_room_info AS
SELECT 
    er.id AS id,
    er.name AS name,
    er.floor AS floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS floor_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.floor) AS floor_name,
    er.location AS location,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.location) AS location_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.location) AS location_name,
    er.status AS status,
    er.remarks AS remarks,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    e.start_time AS exhibition_start_time,
    e.end_time AS exhibition_end_time,
    e.status AS exhibition_status,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE room_id = er.id) AS item_count,
    json_agg(
        json_build_object(
            'id', ei.id,
            'name', ei.name,
            'type', ei.type,
            'status', ei.status,
            'remarks', ei.remarks
        )
    ) AS items,
    (
        SELECT json_agg(json_build_object(
            'device_id', cd.id,
            'device_name', cd.name,
            'device_type', cd.device_type,
            'status', cd.status,
            'ip_address', cd.ip_address,
            'port', cd.port,
            'version', cd.version,
            'commands', cd.commands
        ))
        FROM o_ebcp_control_device cd 
        WHERE cd.room_id = er.id
    ) AS control_devices    
FROM 
    o_ebcp_exhibition_room er
LEFT JOIN 
    o_ebcp_exhibition e ON er.exhibition_id = e.id
LEFT JOIN 
    o_ebcp_exhibition_item ei ON ei.room_id = er.id
GROUP BY 
    er.id, er.name, er.floor, er.location, er.status, er.remarks, 
    e.id, e.name, e.start_time, e.end_time, e.status;

COMMENT ON VIEW v_ebcp_exhibition_room_info IS '展厅详细视图，包含展厅信息及其关联的展览和展项信息（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.id IS '展厅ID';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.name IS '展厅名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.floor IS '展厅楼层';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.floor_value IS '展厅楼层值';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.floor_name IS '展厅楼层名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.location IS '展厅位置';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.location_value IS '展厅位置值';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.location_name IS '展厅位置名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.status IS '展厅状态';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.remarks IS '展厅备注';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_name IS '所属展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_start_time IS '所属展览开始时间';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_end_time IS '所属展览结束时间';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_status IS '所属展览状态';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.item_count IS '展厅内展项数量';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.items IS '展厅内的展项列表（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.control_devices IS '展厅内的中控设备列表（JSON格式）';

-- 展项详细视图
CREATE VIEW v_ebcp_exhibition_item_info AS
SELECT 
    ei.id AS id,
    ei.name AS name,
    ei.type AS type,
    ei.status AS status,
    ei.remarks AS remarks,
    ei.export_info AS export_info,
    er.id AS room_id,
    er.name AS room_name,
    er.floor AS room_floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS room_floor_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.floor) AS room_floor_name,
    er.location AS room_location,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.location) AS room_location_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.location) AS room_location_name,
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
                'status', p.status,
                'current_program_id', p.current_program_id,
                'current_program_state', p.current_program_state,
                'volume', p.volume,
                'sound_state', p.sound_state,
                'programs', (
                    SELECT json_agg(
                        json_build_object(
                            'id', pp.id,
                            'name', pp.name,
                            'program_id', pp.program_id,
                            'program_index', pp.program_index,
                            'state', pp.state,
                            'medias', (
                                SELECT json_agg(
                                    json_build_object(
                                        'id', ppm.id,
                                        'media_id', ppm.media_id,
                                        'media_name', ppm.media_name
                                    )
                                )
                                FROM o_ebcp_player_program_media ppm
                                WHERE ppm.program_id = pp.id
                            )
                        )
                    )
                    FROM o_ebcp_player_program pp
                    WHERE pp.player_id = p.id
                )
            )
        )
        FROM o_ebcp_player p
        WHERE p.item_id = ei.id
    ) AS player_devices,
    (
        SELECT json_build_object(
            'device_id', cd.id,
            'device_name', cd.name,
            'device_type', cd.device_type,
            'status', cd.status,
            'ip_address', cd.ip_address,
            'port', cd.port,
            'version', cd.version,
            'commands', cd.commands
        )
        FROM o_ebcp_control_device cd 
        WHERE cd.id = ei.device_id
    ) AS control_device,
    (
        SELECT json_agg(
            json_build_object(
                'schedule_id', s.id,
                'start_time', s.start_time,
                'stop_time', s.stop_time,
                'cycle_type', s.cycle_type,
                'enabled', s.enabled,
                'action_type', s.action_type
            )
        )
        FROM o_ebcp_item_schedule s
        WHERE s.item_id = ei.id
    ) AS schedules,
    ei.commands AS commands,
    ei.sub_type AS sub_type
FROM 
    o_ebcp_exhibition_item ei
JOIN 
    o_ebcp_exhibition_room er ON ei.room_id = er.id
JOIN 
    o_ebcp_exhibition e ON ei.exhibition_id = e.id;

COMMENT ON VIEW v_ebcp_exhibition_item_info IS '展项详细视图，包含展项信息及其关联的展厅、展览、设备和定时任务信息（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.id IS '展项ID';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.name IS '展项名称';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.type IS '展项类型';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.status IS '展项状态';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.remarks IS '展项备注';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.export_info IS '展项输出信息';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_id IS '所属展厅ID';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_name IS '所属展厅名称';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_floor IS '所属展厅楼层';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_floor_value IS '所属展厅楼层值';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_floor_name IS '所属展厅楼层名称';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_location IS '所属展厅位置';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_location_value IS '所属展厅位置值';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.room_location_name IS '所属展厅位置名称';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.exhibition_name IS '所属展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.player_devices IS '关联的播放设备列表（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.control_device IS '关联的中控设备信息（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_item_info.schedules IS '关联的定时任务信息（JSON格式）';


-- 播放设备详细视图
CREATE VIEW v_ebcp_player_info AS
SELECT 
    p.id AS id,
    p.name AS name,
    p.ip_address AS ip_address,
    p.port AS port,
    p.version AS version,
    p.status AS status,
    p.current_program_id AS current_program_id,
    p.current_program_state AS current_program_state,
    p.volume AS volume,
    p.sound_state AS sound_state,
    ei.id AS item_id,
    ei.name AS item_name,
    ei.type AS item_type,
    er.id AS room_id,
    er.name AS room_name,
    er.floor AS room_floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS room_floor_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.floor) AS room_floor_name,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    (
        SELECT json_agg(
            json_build_object(
                'id', pp.id,
                'name', pp.name,
                'program_id', pp.program_id,
                'program_index', pp.program_index,
                'state', pp.state,
                'medias', (
                    SELECT json_agg(
                        json_build_object(
                            'id', ppm.id,
                            'media_id', ppm.media_id,
                            'media_name', ppm.media_name
                        )
                    )
                    FROM o_ebcp_player_program_media ppm
                    WHERE ppm.program_id = pp.id
                )
            )
        )
        FROM o_ebcp_player_program pp
        WHERE pp.player_id = p.id
    ) AS programs
FROM 
    o_ebcp_player p
LEFT JOIN 
    o_ebcp_exhibition_item ei ON p.item_id = ei.id
LEFT JOIN 
    o_ebcp_exhibition_room er ON ei.room_id = er.id
LEFT JOIN 
    o_ebcp_exhibition e ON ei.exhibition_id = e.id;

COMMENT ON VIEW v_ebcp_player_info IS '播放设备详细视图，包含设备信息及其关联的展项、展厅、展览和节目信息（JSON格式）';
COMMENT ON COLUMN v_ebcp_player_info.id IS '设备ID';
COMMENT ON COLUMN v_ebcp_player_info.name IS '设备名称';
COMMENT ON COLUMN v_ebcp_player_info.ip_address IS 'IP地址';
COMMENT ON COLUMN v_ebcp_player_info.port IS '端口';
COMMENT ON COLUMN v_ebcp_player_info.version IS '版本';
COMMENT ON COLUMN v_ebcp_player_info.status IS '状态,1:正常,2:离线,3:故障';
COMMENT ON COLUMN v_ebcp_player_info.current_program_id IS '当前节目ID';
COMMENT ON COLUMN v_ebcp_player_info.current_program_state IS '当前节目状态,0:播放,1:暂停,2:停止';
COMMENT ON COLUMN v_ebcp_player_info.volume IS '音量(0-100)';
COMMENT ON COLUMN v_ebcp_player_info.sound_state IS '音量状态,0:静音,1:非静音';
COMMENT ON COLUMN v_ebcp_player_info.item_id IS '所属展项ID';
COMMENT ON COLUMN v_ebcp_player_info.item_name IS '所属展项名称';
COMMENT ON COLUMN v_ebcp_player_info.item_type IS '所属展项类型';
COMMENT ON COLUMN v_ebcp_player_info.room_id IS '所属展厅ID';
COMMENT ON COLUMN v_ebcp_player_info.room_name IS '所属展厅名称';
COMMENT ON COLUMN v_ebcp_player_info.room_floor IS '所属展厅楼层';
COMMENT ON COLUMN v_ebcp_player_info.room_floor_value IS '所属展厅楼层值';
COMMENT ON COLUMN v_ebcp_player_info.room_floor_name IS '所属展厅楼层名称';
COMMENT ON COLUMN v_ebcp_player_info.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN v_ebcp_player_info.exhibition_name IS '所属展览名称';
COMMENT ON COLUMN v_ebcp_player_info.programs IS '关联的节目信息';

-- 节目详细视图
CREATE VIEW v_ebcp_player_program_info AS
SELECT 
    pp.id AS id,
    pp.name AS name,
    pp.program_id AS program_id,
    pp.program_index AS program_index,
    pp.state AS state,
    p.id AS player_id,
    p.name AS player_name,
    p.ip_address AS player_ip_address,
    p.port AS player_port,
    p.status AS player_status,
    p.current_program_id AS player_current_program_id,
    p.current_program_state AS player_current_program_state,
    p.volume AS player_volume,
    p.sound_state AS player_sound_state,
    ei.id AS item_id,
    ei.name AS item_name,
    er.id AS room_id,
    er.name AS room_name,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    (
        SELECT json_agg(
            json_build_object(
                'id', ppm.id,
                'media_id', ppm.media_id,
                'media_name', ppm.media_name
            )
        )
        FROM o_ebcp_player_program_media ppm
        WHERE ppm.program_id = pp.id
    ) AS medias
FROM 
    o_ebcp_player_program pp
LEFT JOIN 
    o_ebcp_player p ON pp.player_id = p.id
LEFT JOIN 
    o_ebcp_exhibition_item ei ON p.item_id = ei.id
LEFT JOIN 
    o_ebcp_exhibition_room er ON ei.room_id = er.id
LEFT JOIN 
    o_ebcp_exhibition e ON er.exhibition_id = e.id;

COMMENT ON VIEW v_ebcp_player_program_info IS '节目详细视图，包含节目信息及其关联的播放设备、展项、展厅和展览信息（JSON格式）';

-- 中控设备详细视图
CREATE VIEW v_ebcp_control_device_info AS
SELECT 
    cd.id AS id,
    cd.name AS name,
    cd.device_type AS device_type,
    cd.ip_address AS ip_address,
    cd.port AS port,
    cd.version AS version,
    cd.status AS status,
    cd.commands AS commands,
    cd.created_time AS created_time,
    cd.updated_time AS updated_time,
    er.id AS room_id,
    er.name AS room_name,
    er.status AS room_status,
    er.remarks AS room_remarks,
    er.floor AS room_floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS room_floor_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.floor) AS room_floor_name,
    er.location AS room_location,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.location) AS room_location_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.location) AS room_location_name,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    e.start_time AS exhibition_start_time,
    e.end_time AS exhibition_end_time,
    e.status AS exhibition_status,
    eh.id AS exhibition_hall_id,
    eh.name AS exhibition_hall_name,
    eh.remarks AS exhibition_hall_remarks,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE device_id = cd.id) AS linked_items_count,
    (
        SELECT json_agg(
            json_build_object(
                'id', ei.id,
                'name', ei.name,
                'type', ei.type,
                'sub_type', ei.sub_type,
                'status', ei.status,
                'remarks', ei.remarks,
                'export_info', ei.export_info,
                'commands', ei.commands
            )
        )
        FROM o_ebcp_exhibition_item ei
        WHERE ei.device_id = cd.id
    ) AS linked_items
FROM 
    o_ebcp_control_device cd
LEFT JOIN 
    o_ebcp_exhibition_room er ON cd.room_id = er.id
LEFT JOIN 
    o_ebcp_exhibition e ON er.exhibition_id = e.id
LEFT JOIN 
    o_ebcp_exhibition_hall eh ON er.exhibition_hall_id = eh.id;

COMMENT ON VIEW v_ebcp_control_device_info IS '中控设备详细视图，包含设备信息及其关联的展厅、展览、展馆和展项信息（JSON格式）';
COMMENT ON COLUMN v_ebcp_control_device_info.id IS '设备ID';
COMMENT ON COLUMN v_ebcp_control_device_info.name IS '设备名称';
COMMENT ON COLUMN v_ebcp_control_device_info.device_type IS '设备类型';
COMMENT ON COLUMN v_ebcp_control_device_info.ip_address IS 'IP地址';
COMMENT ON COLUMN v_ebcp_control_device_info.port IS '端口';
COMMENT ON COLUMN v_ebcp_control_device_info.version IS '版本';
COMMENT ON COLUMN v_ebcp_control_device_info.status IS '设备状态(1: 正常, 2: 故障)';
COMMENT ON COLUMN v_ebcp_control_device_info.commands IS '命令列表';
COMMENT ON COLUMN v_ebcp_control_device_info.created_time IS '创建时间';
COMMENT ON COLUMN v_ebcp_control_device_info.updated_time IS '更新时间';
COMMENT ON COLUMN v_ebcp_control_device_info.room_id IS '所属展厅ID';
COMMENT ON COLUMN v_ebcp_control_device_info.room_name IS '所属展厅名称';
COMMENT ON COLUMN v_ebcp_control_device_info.room_status IS '所属展厅状态';
COMMENT ON COLUMN v_ebcp_control_device_info.room_remarks IS '所属展厅备注';
COMMENT ON COLUMN v_ebcp_control_device_info.room_floor IS '所属展厅楼层';
COMMENT ON COLUMN v_ebcp_control_device_info.room_floor_value IS '所属展厅楼层值';
COMMENT ON COLUMN v_ebcp_control_device_info.room_floor_name IS '所属展厅楼层名称';
COMMENT ON COLUMN v_ebcp_control_device_info.room_location IS '所属展厅位置';
COMMENT ON COLUMN v_ebcp_control_device_info.room_location_value IS '所属展厅位置值';
COMMENT ON COLUMN v_ebcp_control_device_info.room_location_name IS '所属展厅位置名称';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_name IS '所属展览名称';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_start_time IS '所属展览开始时间';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_end_time IS '所属展览结束时间';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_status IS '所属展览状态';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_hall_id IS '所属展馆ID';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_hall_name IS '所属展馆名称';
COMMENT ON COLUMN v_ebcp_control_device_info.exhibition_hall_remarks IS '所属展馆备注';
COMMENT ON COLUMN v_ebcp_control_device_info.linked_items_count IS '直接关联的展项数量';
COMMENT ON COLUMN v_ebcp_control_device_info.linked_items IS '直接关联的展项列表（JSON格式）';




-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP VIEW IF EXISTS v_ebcp_control_device_info;
DROP VIEW IF EXISTS v_ebcp_player_program_info;
DROP VIEW IF EXISTS v_ebcp_player_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_item_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_room_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_hall_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_area_info;
DROP VIEW IF EXISTS v_ebcp_exhibition_info;
DROP TABLE IF EXISTS o_ebcp_holiday_date;
DROP TABLE IF EXISTS o_ebcp_item_schedule;
DROP TABLE IF EXISTS o_ebcp_item_device_relation;
DROP TABLE IF EXISTS o_ebcp_player_program_media;
DROP TABLE IF EXISTS o_ebcp_player_program;
DROP TABLE IF EXISTS o_ebcp_control_device;
DROP TABLE IF EXISTS o_ebcp_player;
DROP TABLE IF EXISTS o_ebcp_exhibition_item;
DROP TABLE IF EXISTS o_ebcp_exhibition_room;
DROP TABLE IF EXISTS o_ebcp_exhibition;
DROP TABLE IF EXISTS o_ebcp_exhibition_hall;

-- +goose StatementEnd
