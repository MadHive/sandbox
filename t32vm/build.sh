#!/usr/bin/env bash
BUILDVERSIONMAJOR=0
BUILDVERSIONMINOR=1
BUILDVERSIONPATCH=0
BUILDVERSIONFULL="$BUILDVERSIONMAJOR.$BUILDVERSIONMINOR.$BUILDVERSIONPATCH"
BUILDTIMESTAMP=$(date)

go build -o bin/t32vm -ldflags="-X 'github.com/madhive/sandbox/t32vm/version.BuildVersion=$BUILDVERSIONFULL' -X 'github.com/madhive/sandbox/t32vm/version.BuildTimestamp=$BUILDTIMESTAMP'" .