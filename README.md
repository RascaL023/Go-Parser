Title: Go Parser

Description:
Parse JSON-based theme definition into multiple tools configuration using SSOT (Single Source of Truth) approach.

Features:
- Parse theme JSON into internal structured data
- Render theme into tools-specific config formats
- Modular renderer design (easy to add new tools)
- Custom input and output path
- Deterministic output (same input → same config)
- Minimal runtime dependency

Supported / Planned Tools:
- Hyprland
- Kitty
- Cava
- Waybar (in progress)

Future Plans:
- Hot-reload support
- Theme validation & schema checking
- CLI flags for selective rendering
- Template-based renderer (Mustache / Go template)
- Documentation & examples


Last Update: Add waybar, improving structure, update README.md
