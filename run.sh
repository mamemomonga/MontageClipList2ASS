#!/bin/bash
set -eu

rm -rf var
mkdir -p var

go run ./src/cmd/mcl2ass -c example/config.yaml -f example/ClipList.xml -o var/output.ass -x var/youtube.txt

if [ -f var/output.ass ]; then
    ffmpeg -f lavfi -i color=c=blue:s=1920x1080:d=5 \
          -vf "ass=var/output.ass" \
          -c:v libx264 -pix_fmt yuv420p -y var/output.mp4

    if [ -f var/output.mp4 ]; then
        open var/output.mp4
    fi
fi