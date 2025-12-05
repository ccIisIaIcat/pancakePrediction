@echo off
setlocal

rem Check if target folder name is provided
if "%1"=="" (
    echo Usage: build.bat [folder_name]
    echo Example: build.bat demo
    exit /b 1
)

set TARGET=%1
set BASE_DIR=%~dp0
set CMD_DIR=%BASE_DIR%cmd\%TARGET%

rem Check main.go exists
if not exist "%CMD_DIR%\main.go" (
    echo Error: %CMD_DIR%\main.go not found.
    exit /b 1
)

echo Building %TARGET% ...
go build -o "%CMD_DIR%\%TARGET%.exe" "%CMD_DIR%\main.go"

if %errorlevel% neq 0 (
    echo Build failed.
    exit /b 1
)

echo Build succeeded: %CMD_DIR%\%TARGET%.exe
endlocal