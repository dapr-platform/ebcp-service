# Media Player Device Interface Specification

This document outlines the network communication protocol for Media Player Devices that need to integrate with the `ebcp-service`. The service acts as a client, sending commands to the Media Player to control playback, volume, and other media-related functions.

## 1. Communication Protocol

-   **Protocol:** TCP (Transmission Control Protocol)
-   **IP Address & Port:** The IP address and port for each player device are configured within the `o_ebcp_player` table in the system's database. The `ebcp-service` will establish a TCP connection to this address.
-   **Endianness:** All multi-byte numerical values are encoded in **Little-Endian** format.

## 2. Frame Structure

All data packets, both from the service to the player and from the player to the service, follow a unified frame structure.

### Packet Header (12 bytes)

| Field | Size (bytes) | Default Value | Description |
| :--- | :--- | :--- | :--- |
| `Head` | 4 | `0x55CC55CC` | A fixed header to identify the protocol. Sent as `CC 55 CC 55` in Little-Endian. |
| `Packet Type` | 2 | `0x0001` | Reserved. Should be set to 1. |
| `Protocol Version` | 2 | `0x0001` | The protocol version, e.g., `0x0101` for v1.1. |
| `Sequence` | 2 | Auto-increment | A sequence number for the packet, incremented by the sender. |
| `Content Length` | 2 | Calculated | The total length of the content area, which includes the TLV Header and TLV Value. |

### Content Area (TLV Structure)

The content area follows the header and uses a Type-Length-Value (TLV) format.

| Field | Size (bytes) | Description |
| :--- | :--- | :--- |
| `Tag` | 2 | A unique identifier for the command or response type. |
| `Length` | 2 | The length of the `Value` field in bytes. |
| `Value` | N | The actual data payload for the command or response. Its structure depends on the `Tag`. |

## 3. Command List

The `ebcp-service` will send the following commands to the Media Player. The player must listen for these commands, parse them, and perform the corresponding actions.

### Program & Media Control

| Command | Tag (Decimal) | Tag (Hex) | Value (Payload) Structure | Description |
| :--- | :--- | :--- | :--- | :--- |
| **Get Program List** | 129 | `0x0081` | (Empty) | Requests the list of all available programs from the player. |
| **Fade Program** | 131 | `0x0083` | `programID` (4 bytes, uint32) | Fades to the specified program. |
| **Cut Program** | 132 | `0x0084` | `programID` (4 bytes, uint32) | Cuts directly to the specified program. |
| **Pause Program** | 133 | `0x0085` | `programID` (4 bytes, uint32) | Pauses the specified program. |
| **Play Program** | 271 | `0x010F` | `programID` (4 bytes, uint32) | Plays the specified program. |
| **Stop Program** | 272 | `0x0110` | `programID` (4 bytes, uint32) | Stops the specified program. |
| **Play Layer Media** | 273 | `0x0111` | `layerIndex` (4 bytes, uint32) | Plays the media in the specified layer of the current program. |
| **Pause Layer Media** | 274 | `0x0112` | `layerIndex` (4 bytes, uint32) | Pauses the media in the specified layer of the current program. |
| **Get All Program Media** | 276 | `0x0114` | `programID` (4 bytes, uint32) | Requests the list of all media files within a specific program. |
| **Query Layer Progress** | 293 | `0x0125` | `layerIndex` (2 bytes, uint16) | Queries the playback progress of a specific layer. |
| **Get Current Program** | 294 | `0x0126` | (Empty) | Queries the ID and state of the currently playing program. |

### Sound Control

| Command | Tag (Decimal) | Tag (Hex) | Value (Payload) Structure | Description |
| :--- | :--- | :--- | :--- | :--- |
| **Open Global Sound** | 262 | `0x0106` | (Empty) | Unmutes the global sound. |
| **Close Global Sound** | 263 | `0x0107` | (Empty) | Mutes the global sound. |
| **Set Global Volume** | 264 | `0x0108` | `volume` (4 bytes, uint32, 0-100) | Sets the global volume. |
| **Increase Volume** | 328 | `0x0148` | `step` (4 bytes, uint32) | Increases the global volume by a specified step. |
| **Decrease Volume** | 329 | `0x0149` | `step` (4 bytes, uint32) | Decreases the global volume by a specified step. |
| **Query Layer Volume/Mute** | 322 | `0x0142` | `layerIndex` (2 bytes, uint16) | Queries the volume and mute status of a specific layer. |

## 4. Response Format

For commands that require a response, the Media Player should send back a packet with the same structure. The `Tag` in the response should match the `Tag` of the request. The `Value` field will contain the response data.

### Example Response: Get Program List

The player should respond with one or more packets. The first packet's value contains the total number of programs, and subsequent data contains the details for each program.

-   **Response TLV - Value[]**
    -   `ProgramCount` (4 bytes, uint32): Total number of programs.
    -   `ProgramIndex` (4 bytes, uint32): Index of the program in the list (0-based).
    -   `ProgramId` (4 bytes, uint32): The unique ID of the program.
    -   `ProgramName` (128 bytes, string, UTF-8): The name of the program.
    -   `IsEmpty` (1 byte, uint8): `0` for an empty program, `1` for a non-empty program.

### Example Response: Query Layer Progress

-   **Response TLV - Value[]**
    -   `Success` (1 byte, uint8): `1` for success, `0` for failure.
    -   `LayerIndex` (2 bytes, uint16): The index of the layer.
    -   `RemainTime` (4 bytes, uint32): Remaining playback time in seconds.
    -   `TotalTime` (4 bytes, uint32): Total playback time in seconds.
