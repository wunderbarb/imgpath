#!/bin/bash
go list -json -deps ./... | docker run --rm -i sonatypecommunity/nancy:v1.0.48 sleuth