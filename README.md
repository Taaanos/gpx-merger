# gpx-merger

Merges gpx files in a chronological order.

## Motivation

* Apple watch decides to stop an activity and there is no option to resume, ending up with mutliple seperate activities for 1 route.

## Usage

```bash
go build gpxmerge.go
./gpxmerge <FILE1> <FILE2> 
```

You should end up with a `merged.gpx`

## Details

Tracks are merged together in a single gpx xml.

## Prerequisites

* Go installed on your machine.
