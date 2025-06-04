#!/bin/bash
set -eu
go run ./src/cmd/montage_cliplist2ass $@

if [ -f output.ass ]; then
    ffmpeg -f lavfi -i color=c=blue:s=1920x1080:d=5 \
          -vf "ass=output.ass" \
          -c:v libx264 -pix_fmt yuv420p -y output.mp4
    open output.mp4
fi