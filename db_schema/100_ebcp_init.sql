-- +goose Up
-- +goose StatementBegin
-- 初始化字典数据
-- 楼层字典
INSERT INTO o_ops_dict (
    id, created_by, created_time, updated_by, updated_time,
    dict_type, dict_code, dict_name, dict_value, sort_order, remarks
) VALUES
    (md5('floor_typeB1'), 'admin', NOW(), 'admin', NOW(), 'floor_type', 'B1', '地下一层', 'B1', 1, NULL),
    (md5('floor_type1F'), 'admin', NOW(), 'admin', NOW(), 'floor_type', '1F', '一层', '1F', 2, NULL),
    (md5('floor_type2F'), 'admin', NOW(), 'admin', NOW(), 'floor_type', '2F', '二层', '2F', 3, NULL),
    (md5('floor_type3F'), 'admin', NOW(), 'admin', NOW(), 'floor_type', '3F', '三层', '3F', 4, NULL);

-- 位置字典
INSERT INTO o_ops_dict (
    id, created_by, created_time, updated_by, updated_time,
    dict_type, dict_code, dict_name, dict_value, sort_order, remarks
) VALUES
    (md5('location_typeWEST'), 'admin', NOW(), 'admin', NOW(), 'location_type', 'WEST', '西侧', 'WEST', 1, NULL),
    (md5('location_typeSOUTHWEST'), 'admin', NOW(), 'admin', NOW(), 'location_type', 'SOUTHWEST', '西南侧', 'SOUTHWEST', 2, NULL),
    (md5('location_typeNORTHWEST'), 'admin', NOW(), 'admin', NOW(), 'location_type', 'NORTHWEST', '西北侧', 'NORTHWEST', 3, NULL),
    (md5('location_typeNORTHEAST'), 'admin', NOW(), 'admin', NOW(), 'location_type', 'NORTHEAST', '东北侧', 'NORTHEAST', 4, NULL),
    (md5('location_typeNORTH'), 'admin', NOW(), 'admin', NOW(), 'location_type', 'NORTH', '北侧', 'NORTH', 5, NULL),
    (md5('location_typeSOUTHEAST'), 'admin', NOW(), 'admin', NOW(), 'location_type', 'SOUTHEAST', '东南侧', 'SOUTHEAST', 6, NULL);

-- 初始化展馆数据
INSERT INTO o_ebcp_exhibition_hall (
    id, created_by, created_time, updated_by, updated_time, name, remarks
) VALUES (
    md5('军事博物馆'), 'admin', NOW(), 'admin', NOW(), '军事博物馆', NULL
);

-- 初始化展览数据
INSERT INTO o_ebcp_exhibition (
    id, created_by, created_time, updated_by, updated_time,
    name, start_time, end_time, remarks, hall_id, status
) VALUES
    (md5('陆军轻武器装备技术厅&陆军航空兵武器装备技术厅'), 'admin', NOW(), 'admin', NOW(), 
    '陆军轻武器装备技术厅&陆军航空兵武器装备技术厅', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('陆军重武器装备技术厅'), 'admin', NOW(), 'admin', NOW(), 
    '陆军重武器装备技术厅', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('领航强军向复兴-新时代国防和军队建设成就展'), 'admin', NOW(), 'admin', NOW(), 
    '领航强军向复兴-新时代国防和军队建设成就展', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('在党的旗帜下前进-人民军队庆祝中国共产党成立100周年主题展览'), 'admin', NOW(), 'admin', NOW(),
    '在党的旗帜下前进-人民军队庆祝中国共产党成立100周年主题展览', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('空军武器装备技术厅'), 'admin', NOW(), 'admin', NOW(), 
    '空军武器装备技术厅', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('导弹武器装备技术厅'), 'admin', NOW(), 'admin', NOW(), 
    '导弹武器装备技术厅', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('核武器与核技术和平利用厅'), 'admin', NOW(), 'admin', NOW(), 
    '核武器与核技术和平利用厅', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('海军武器装备技术厅'), 'admin', NOW(), 'admin', NOW(), 
    '海军武器装备技术厅', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('抗美援朝战争馆'), 'admin', NOW(), 'admin', NOW(), 
    '抗美援朝战争馆', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1),
    
    (md5('新民主主义革命时期陈列'), 'admin', NOW(), 'admin', NOW(), 
    '新民主主义革命时期陈列', NOW(), NOW() + INTERVAL '10 year', NULL, md5('军事博物馆'), 1);

-- 初始化展厅数据
INSERT INTO o_ebcp_exhibition_room (
    id, created_by, created_time, updated_by, updated_time,
    name, location, exhibition_hall_id, floor, exhibition_id, status, remarks
) VALUES
    (md5('B1西侧'), 'admin', NOW(), 'admin', NOW(), 'B1西侧', md5('location_typeWEST'), md5('军事博物馆'), md5('floor_typeB1'), md5('陆军轻武器装备技术厅&陆军航空兵武器装备技术厅'), 1, '陆军轻武器装备技术厅&陆军航空兵武器装备技术厅'),
    (md5('B1西北侧'), 'admin', NOW(), 'admin', NOW(), 'B1西北侧', md5('location_typeNORTHWEST'), md5('军事博物馆'), md5('floor_typeB1'), md5('陆军重武器装备技术厅'), 1, '陆军重武器装备技术厅'),
    (md5('1F西南侧'), 'admin', NOW(), 'admin', NOW(), '1F西南侧', md5('location_typeSOUTHWEST'), md5('军事博物馆'), md5('floor_type1F'), md5('领航强军向复兴-新时代国防和军队建设成就展'), 1, '领航强军向复兴-新时代国防和军队建设成就展'),
    (md5('1F西侧'), 'admin', NOW(), 'admin', NOW(), '1F西侧', md5('location_typeWEST'), md5('军事博物馆'), md5('floor_type1F'), md5('领航强军向复兴-新时代国防和军队建设成就展'), 1, '领航强军向复兴-新时代国防和军队建设成就展'),
    (md5('1F西北侧'), 'admin', NOW(), 'admin', NOW(), '1F西北侧', md5('location_typeNORTHWEST'), md5('军事博物馆'), md5('floor_type1F'), md5('领航强军向复兴-新时代国防和军队建设成就展'), 1, '领航强军向复兴-新时代国防和军队建设成就展'),
    (md5('1F东南侧'), 'admin', NOW(), 'admin', NOW(), '1F东南侧', md5('location_typeSOUTHEAST'), md5('军事博物馆'), md5('floor_type1F'), md5('在党的旗帜下前进-人民军队庆祝中国共产党成立100周年主题展览'), 1, '在党的旗帜下前进-人民军队庆祝中国共产党成立100周年主题展览'),
    (md5('1F东北侧'), 'admin', NOW(), 'admin', NOW(), '1F东北侧', md5('location_typeNORTHEAST'), md5('军事博物馆'), md5('floor_type1F'), md5('空军武器装备技术厅'), 1, '空军武器装备技术厅'),
    (md5('2F西南侧'), 'admin', NOW(), 'admin', NOW(), '2F西南侧', md5('location_typeSOUTHWEST'), md5('军事博物馆'), md5('floor_type2F'), md5('新民主主义革命时期陈列'), 1, '新民主主义革命时期陈列'),
    (md5('2F西北侧'), 'admin', NOW(), 'admin', NOW(), '2F西北侧', md5('location_typeNORTHWEST'), md5('军事博物馆'), md5('floor_type2F'), md5('导弹武器装备技术厅'), 1, '导弹武器装备技术厅'),
    (md5('2F北侧'), 'admin', NOW(), 'admin', NOW(), '2F北侧', md5('location_typeNORTH'), md5('军事博物馆'), md5('floor_type2F'), md5('核武器与核技术和平利用厅'), 1, '核武器与核技术和平利用厅'),
    (md5('2F东北侧'), 'admin', NOW(), 'admin', NOW(), '2F东北侧', md5('location_typeNORTHEAST'), md5('军事博物馆'), md5('floor_type2F'), md5('海军武器装备技术厅'), 1, '海军武器装备技术厅'),
    (md5('3F东南侧'), 'admin', NOW(), 'admin', NOW(), '3F东南侧', md5('location_typeSOUTHEAST'), md5('军事博物馆'), md5('floor_type3F'), md5('抗美援朝战争馆'), 1, '抗美援朝战争馆');

-- 初始化展项数据
INSERT INTO o_ebcp_exhibition_item (
    id, created_by, created_time, updated_by, updated_time,
    name, exhibition_id, exhibition_room_id, type, status, remarks
) VALUES
    (md5('九五枪族'), 'admin', NOW(), 'admin', NOW(), '九五枪族', 
    md5('陆军轻武器装备技术厅&陆军航空兵武器装备技术厅'), md5('B1西侧'), 'media', 1, 'PC工作站+播控软件'),
    
    (md5('重武器厅前言'), 'admin', NOW(), 'admin', NOW(), '重武器厅前言', 
    md5('陆军重武器装备技术厅'), md5('B1西北侧'), 'media', 1, 'PC工作站+播控软件'),
    
    (md5('领航强军向复兴'), 'admin', NOW(), 'admin', NOW(), '领航强军向复兴', 
    md5('领航强军向复兴-新时代国防和军队建设成就展'), md5('1F西南侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('学强军思想建强军事业'), 'admin', NOW(), 'admin', NOW(), '学强军思想建强军事业', 
    md5('领航强军向复兴-新时代国防和军队建设成就展'), md5('1F西南侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('合成营'), 'admin', NOW(), 'admin', NOW(), '合成营', 
    md5('领航强军向复兴-新时代国防和军队建设成就展'), md5('1F西南侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('阅兵'), 'admin', NOW(), 'admin', NOW(), '阅兵', 
    md5('领航强军向复兴-新时代国防和军队建设成就展'), md5('1F西侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('号令'), 'admin', NOW(), 'admin', NOW(), '号令', 
    md5('领航强军向复兴-新时代国防和军队建设成就展'), md5('1F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('统帅与士兵'), 'admin', NOW(), 'admin', NOW(), '统帅与士兵', 
    md5('领航强军向复兴-新时代国防和军队建设成就展'), md5('1F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('在党的旗帜下前进序厅'), 'admin', NOW(), 'admin', NOW(), '在党的旗帜下前进序厅',
    md5('在党的旗帜下前进-人民军队庆祝中国共产党成立100周年主题展览'), md5('1F东南侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('空军主战装备体系'), 'admin', NOW(), 'admin', NOW(), '空军主战装备体系', 
    md5('空军武器装备技术厅'), md5('1F东北侧'), 'media', 1, 'PC工作站+播控软件'),
    
    (md5('飞机隐身技术'), 'admin', NOW(), 'admin', NOW(), '飞机隐身技术', 
    md5('空军武器装备技术厅'), md5('1F东北侧'), 'media', 1, 'PC工作站+融合拼接软件+播控软件'),
    
    (md5('空中打击力'), 'admin', NOW(), 'admin', NOW(), '空中打击力', 
    md5('空军武器装备技术厅'), md5('1F东北侧'), 'media', 3, 'PC工作站+播控软件'),
    
    (md5('全域作战'), 'admin', NOW(), 'admin', NOW(), '全域作战', 
    md5('导弹武器装备技术厅'), md5('2F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('全空域作战'), 'admin', NOW(), 'admin', NOW(), '全空域作战', 
    md5('导弹武器装备技术厅'), md5('2F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('全领域作战'), 'admin', NOW(), 'admin', NOW(), '全领域作战', 
    md5('导弹武器装备技术厅'), md5('2F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('全疆域作战'), 'admin', NOW(), 'admin', NOW(), '全疆域作战', 
    md5('导弹武器装备技术厅'), md5('2F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('利剑受阅'), 'admin', NOW(), 'admin', NOW(), '利剑受阅', 
    md5('导弹武器装备技术厅'), md5('2F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('时刻准备着'), 'admin', NOW(), 'admin', NOW(), '时刻准备着', 
    md5('导弹武器装备技术厅'), md5('2F西北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('核铸强国'), 'admin', NOW(), 'admin', NOW(), '核铸强国', 
    md5('核武器与核技术和平利用厅'), md5('2F北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('环幕影院'), 'admin', NOW(), 'admin', NOW(), '环幕影院', 
    md5('核武器与核技术和平利用厅'), md5('2F北侧'), 'media', 1, '服务器+融合拼接软件+播控软件'),
    
    (md5('双龙'), 'admin', NOW(), 'admin', NOW(), '双龙', 
    md5('核武器与核技术和平利用厅'), md5('2F北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('深海隧道'), 'admin', NOW(), 'admin', NOW(), '深海隧道', 
    md5('海军武器装备技术厅'), md5('2F东北侧'), 'media', 1, 'PC工作站+播控软件'),
    
    (md5('航母之路'), 'admin', NOW(), 'admin', NOW(), '航母之路', 
    md5('海军武器装备技术厅'), md5('2F东北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('现代海洋装备体系'), 'admin', NOW(), 'admin', NOW(), '现代海洋装备体系', 
    md5('海军武器装备技术厅'), md5('2F东北侧'), 'media', 1, '服务器+播控软件'),
    
    (md5('抗美援朝序厅'), 'admin', NOW(), 'admin', NOW(), '抗美援朝序厅', 
    md5('抗美援朝战争馆'), md5('3F东南侧'), 'media', 1, 'PC工作站+播控软件'),
    
    (md5('抗美援朝地图展示'), 'admin', NOW(), 'admin', NOW(), '抗美援朝地图展示', 
    md5('抗美援朝战争馆'), md5('3F东南侧'), 'media', 1, 'PC工作站+播控软件');

-- 初始化播放器数据  
INSERT INTO o_ebcp_player (
    id, created_by, created_time, updated_by, updated_time,
    name, ip_address, port, status, item_id
) VALUES
    -- B1层播放设备
    (md5('player_B1_west_1'), 'admin', NOW(), 'admin', NOW(), 'B1西侧播放工作站1', '182.92.117.41', 40306, 1, md5('九五枪族'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM o_ebcp_player_program;
DELETE FROM o_ebcp_player;
DELETE FROM o_ebcp_exhibition_item;
DELETE FROM o_ebcp_exhibition_room;
DELETE FROM o_ebcp_exhibition;
DELETE FROM o_ebcp_exhibition_hall;
DELETE FROM o_ops_dict;
-- +goose StatementEnd