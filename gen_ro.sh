dp-cli gen --connstr "postgresql://things:things2024@ali4:37054/thingsdb?sslmode=disable" \
--tables=o_ebcp_player_program,o_ebcp_player_program_media --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module ebcp-service --api R

