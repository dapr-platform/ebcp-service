package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitEbcp_exhibition_hallRoute(r)
	InitEbcp_exhibitionRoute(r)
	InitEbcp_exhibition_roomRoute(r)
	InitEbcp_exhibition_itemRoute(r)
	InitEbcp_exhibition_infoRoute(r)
	InitEbcp_playerRoute(r)
	InitEbcp_playerExtRoute(r)
	InitEbcp_item_device_relationRoute(r)
	InitEbcp_item_scheduleRoute(r)
	InitEbcpDashboardRoute(r)
	InitEbcp_exhibition_itemExtRoute(r)
	InitEbcp_player_programRoute(r)
	InitEbcp_control_deviceRoute(r)
	InitEbcp_exhibition_area_infoRoute(r)
	InitEbcp_exhibition_hall_infoRoute(r)
	InitEbcp_exhibition_room_infoRoute(r)
	InitEbcp_exhibition_item_infoRoute(r)
	InitEbcp_player_infoRoute(r)
	InitEbcp_player_program_infoRoute(r)
	InitEbcp_holiday_dateRoute(r)
	InitEbcp_exhibition_hallExtRoute(r)
	InitDebugRoute(r)
	InitEbcp_item_schedule_extRoute(r)
}
