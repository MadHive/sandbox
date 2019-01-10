#!/usr/bin/env bash
BUILDVERSIONMAJOR=0
BUILDVERSIONMINOR=1
BUILDVERSIONPATCH=0
BUILDVERSIONFULL="$BUILDVERSIONMAJOR.$BUILDVERSIONMINOR.$BUILDVERSIONPATCH"
BUILDTIMESTAMP=$(date)

go build -o bin/mw -ldflags="-X 'github.com/madhive/sandbox/mimblewimble/version.BuildVersion=$BUILDVERSIONFULL' -X 'github.com/madhive/sandbox/mimblewimble/version.BuildTimestamp=$BUILDTIMESTAMP'" .