@echo off
minesweeper.exe
if %errorlevel% == 1 (echo "You lost...") else (echo "You won!")