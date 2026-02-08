#!/bin/zsh
# Capture and display localhost:8080 traffic in real-time

# Check if tcpdump is available (comes with macOS)
if ! command -v tcpdump &> /dev/null; then
    echo "tcpdump not found. Install with: brew install tcpdump"
    exit 1
fi

echo "Capturing traffic on localhost:8080..."
echo "Press Ctrl+C to stop"
echo ""

# Capture and display traffic in ASCII format
sudo tcpdump -i lo0 -A -s 0 "tcp port 8080"
