@echo off
title Building VTetris for Windows
go build -x -ldflags="-H=windowsgui"  -o ./build/VTetris.exe
