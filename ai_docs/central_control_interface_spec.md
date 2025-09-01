# Central Control Device Interface Specification

This document outlines the network communication protocol for Central Control Devices that need to integrate with the `ebcp-service`. The service communicates with these devices to control static exhibition items, such as power and lighting.

## 1. Communication Protocol

-   **Protocol:** UDP (User Datagram Protocol)
-   **IP Address & Port:** The IP address and port for each central control device are configured within the `o_ebcp_control_device` table in the system's database. The `ebcp-service` will send commands to the configured IP and port.
-   **Encoding:** Command data is sent as plain ASCII strings.

## 2. Command Structure

The service sends simple, string-based commands. Each command is a short string, and the device is expected to perform an action based on the received command string.

### Command Format

The command is a raw string sent in the UDP packet payload. For example, to turn on all devices, the payload would simply be `A1`.

## 3. Command List

The following is a list of commands the `ebcp-service` will send. Third-party devices must implement the logic to handle these commands. The commands are derived from the `service/hall.go` and `db_schema/` files.

| Command String | Description (Translated) | Device Type | Sub-Type | Notes |
| :--- | :--- | :--- | :--- | :--- |
| `F1 01 01` | Master Control: All On | `master_control` | - | Turns on all associated systems. |
| `F1 01 02` | Master Control: All Off | `master_control` | - | Turns off all associated systems. |
| `F1 01 03` | Master Control: All On (Exclude Interactive) | `master_control` | - | Turns on systems, excluding interactive elements. |
| `F2 01 01` | Lighting Control: All On | `lighting_control` | - | Turns on all lighting circuits. |
| `F2 01 02` | Lighting Control: All Off | `lighting_control` | - | Turns off all lighting circuits. |
| `F3 01 01` | Power Control: All On | `power_control` | - | Turns on power for all devices. |
| `F3 01 02` | Power Control: All Off | `power_control` | - | Turns off power for all devices. |
| `FB 01 01` | Distributed Power: On | `power` | - | Example command for a specific power device. |
| `FB 01 02` | Distributed Power: Off | `power` | - | Example command for a specific power device. |
| `FC 01 01` | East Aisle Light: On | `lighting` | - | Example command for a specific lighting device. |
| `FC 01 02` | East Aisle Light: Off | `lighting` | - | Example command for a specific lighting device. |
| `F7 01 01` | Computer Control: On | `device` | - | Example command for a specific computer. |
| `F7 01 02` | Computer Control: Off | `device` | - | Example command for a specific computer. |

**Note:** The `commands` field in the `o_ebcp_exhibition_item` and `o_ebcp_control_device` tables contains a JSON array that maps descriptive names (e.g., "开启", "关闭") to specific command strings (e.g., "FA 01 01"). The service reads these strings and sends them directly via UDP. A third-party device must be programmed to recognize these specific command strings and act accordingly.
