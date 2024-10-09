dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=o_ebcp_device,o_ebcp_exhibition_hall,o_ebcp_exhibition_room,o_ebcp_exhibition_area,o_ebcp_exhibition_item,o_ebcp_schedule_task,o_ebcp_schedule_time,o_ebcp_schedule_action --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module ebcp-service --api true

