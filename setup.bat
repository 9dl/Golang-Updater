@echo off
setlocal

set "STARTUP_FOLDER=%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup"
set "SOURCE_FILE=golangUpdater.exe"

if not exist "%STARTUP_FOLDER%\golangUpdater.exe" (
    copy "%~dp0%SOURCE_FILE%" "%STARTUP_FOLDER%\golangUpdater.exe"
    echo golangUpdater.exe has been added to Startup.
) else (
    echo golangUpdater.exe already exists in Startup.
)

endlocal
