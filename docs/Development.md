# Development

## Quick Start

### Windows

1. Configure prerequisites

    - Go v1.19+
    - Node.js v18.1.0+
    - [Webview](https://github.com/webview/webview)

2. Build Webview DLLs

    1. Run build.bat in `dist/windows/build.bat`
       Alternatively see [Webview](https://github.com/webview/webview) for DLL
       compilation instructions
    2. Place webview.dll & Webview2Loader.dll inside `build` folder

3. Set environment variables then start UI development server provided by
    [Create React App](https://create-react-app.dev/)

    ```ps
    # PowerShell
    $env:PORT = "7777"
    cd ui
    $env:BROWSER = "none"
    npm start
    ```

4. Provide a minimal config at `build/settings.json`

   ```json
   {
     "MultiboxGroups": [
      {
        "Name": "default",
        "GameProfiles": [
          {
            "Name": "Game Client 1",
            "Password": "",
            "Path": "%UserProfile%\\PATH\\TO\\GAME",
            "BinPath": "",
            "BinFileName": "main.exe",
            "StartCmd": "Game.exe",
            "StartArgs": null
          },
        ]
      }
    ],
    "InteropOptions": {},
    "UiOptions": {}
   }
   ```

5. Run application

    _NOTE: Boxer Buddy must be run with administrator privileges.
    This allows the program to access the Windows User32 API for process &
    window management. Your IDE must be run in Administrator mode when
    debugging._

    1. Run with debugger

        A debug configuration is provided for Visual Studio Code at the path `.vscode/launch.json`
        and will be detected automatically by the IDE.

    2. Build & run manually

        ```ps
        # PowerShell
        go build -o build/boxer-buddy.exe
        .\build\boxer-buddy.exe
        ```

### macOS & Linux

Currently unsupported
