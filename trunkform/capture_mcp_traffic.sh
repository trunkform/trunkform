#!/bin/zsh
# Capture and display localhost traffic in real-time

PORT="${CAPTURE_MCP_TRAFFIC_PORT:-8080}"

# Check if tcpdump is available (comes with macOS)
if ! command -v tcpdump &> /dev/null; then
    echo "tcpdump not found. Install with: brew install tcpdump"
    exit 1
fi

echo "Capturing traffic on localhost:${PORT}..."
echo "Press Ctrl+C to stop"
echo ""

# Capture and display traffic in ASCII format
sudo tcpdump -i lo0 -A -s 0 "tcp port ${PORT}"
