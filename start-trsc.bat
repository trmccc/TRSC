@echo off
chcp 65001 >nul
cd /d "%~dp0"
echo Starting TR Server Control...
if exist bin\trsc-windows-amd64.exe (
  bin\trsc-windows-amd64.exe
) else if exist trsc-windows-amd64.exe (
  trsc-windows-amd64.exe
) else (
  echo 找不到 trsc-windows-amd64.exe
)
echo.
echo 面板已退出。若刚闪退，请把 data\panel.log 发给开发者。
pause
