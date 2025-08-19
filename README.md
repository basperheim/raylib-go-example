# Raylib Golang (raylib-go) "Hello Cube" Example

Working **raylib-go** 3D example (Go + raylib bindings). Opens a window and renders a rotating cube.

![go-game-engine-raylib-examples-raylib-golang-raylib-go](https://github.com/user-attachments/assets/00166c0c-38bd-4dab-87ae-fbf1c30966b9)

Check out my [learnprogramming.us blog post](https://learnprogramming.us/blogs/go-game-engine-raylib-examples-raylib-golang-raylib-go) for more details.

## Requirements
- Go 1.21+ (CGO enabled)
- A C toolchain (`gcc`) and the native **raylib** library installed
- `pkg-config` available on your system

## Install raylib (native library)

Raylib itself is a **C library**, so you'll need to install the native [raylib](https://www.raylib.com/) library first.

### macOS (Homebrew, incl. Apple Silicon)

```bash
brew install cmake pkg-config
brew install raylib
````

> If `pkg-config` can't find raylib on Apple Silicon, try:
>
> ```bash
> export PKG_CONFIG_PATH="/opt/homebrew/lib/pkgconfig"
> ```

### Debian/Ubuntu

```bash
sudo apt-get update
sudo apt-get install -y build-essential git pkg-config cmake
sudo apt-get install -y libraylib-dev
```

### Windows (MSYS2 / MinGW64)

1. Install MSYS2: [https://www.msys2.org/](https://www.msys2.org/)
2. Open **MSYS2 MinGW64** shell and install:

   ```bash
   pacman -S --needed mingw-w64-x86_64-gcc mingw-w64-x86_64-pkgconf mingw-w64-x86_64-raylib
   ```
3. Ensure `C:\msys64\mingw64\bin` is on your Windows `PATH` when running `go run`.

## Project setup

If you're starting fresh:

```bash
go mod init github.com/yourname/raylib-go-example
go get github.com/gen2brain/raylib-go/raylib
go mod tidy
```

## Run

```bash
go run .
```

## Build

```bash
go build -o raylib-go-example .
./raylib-go-example   # (Windows: .\raylib-go-example.exe)
```

## Troubleshooting

* **linker/undefined reference**: your system likely can't find `raylib`. Confirm `pkg-config --libs raylib` works.
* **CGO disabled**: enable it with `go env -w CGO_ENABLED=1`.
* **Windows**: run commands from the **MSYS2 MinGW64** shell or ensure the MinGW64 `bin` directory is on `PATH`.

## Useful links

* Raylib Go bindings: [https://github.com/gen2brain/raylib-go](https://github.com/gen2brain/raylib-go)
* Raylib cheatsheet (math & transforms): [https://www.raylib.com/cheatsheet/raymath\_cheatsheet.html](https://www.raylib.com/cheatsheet/raymath_cheatsheet.html)
* Raylib (C) repo/issues: [https://github.com/raysan5/raylib](https://github.com/raysan5/raylib)


