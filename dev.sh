#!/bin/bash

# 1. Start Tailwind watcher in the background
# We use '>/dev/null 2>&1' to hide the "Rebuilding..." noise you don't want
tailwindcss -i ./static/input.css -o ./static/output.css --minify --watch > /dev/null 2>&1 &

# 2. Start Air in the foreground
# This keeps the terminal active so you can see your Go logs/errors
air

# 3. Clean up: Kill the Tailwind process when you stop Air
trap "kill $!" EXIT
