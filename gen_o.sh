dp-cli gen --connstr "postgresql://things:things2024@ali4:37054/thingsdb?sslmode=disable" \
--tables=o_ebcp_exhibition_hall,o_ebcp_exhibition_room,o_ebcp_exhibition_area,o_ebcp_exhibition_item,o_ebcp_camera,o_ebcp_player,o_ebcp_item_device_relation,o_ebcp_item_schedule --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module ebcp-service --api RUDB

