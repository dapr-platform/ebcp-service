dp-cli gen --connstr "postgresql://things:things2024@ali4:37054/thingsdb?sslmode=disable" \
--tables=v_ebcp_player_info,v_ebcp_exhibition_info,v_ebcp_exhibition_area_info,v_ebcp_exhibition_hall_info,v_ebcp_exhibition_room_info,v_ebcp_exhibition_item_info,v_ebcp_player_program_info --model_naming "{{ toUpperCamelCase ( replace . \"v_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"v_\" \"\") }}" \
--module ebcp-service --api R

