# caura

**caura** is a cross-platform system information tool written in Go. It displays basic system details (OS, kernel, uptime, CPU, memory, disk, etc.) in a clean, colorful terminal output.

---

## English

### Features

- **Cross-platform**: Works on Linux and FreeBSD (Android and macOS planned)
- **No external dependencies**: Only requires Go to compile
- **Clean output**: Colored sections with readable information
- **Portable**: Single binary, no runtime dependencies

### Supported platforms

| Platform | Status |
|----------|--------|
| Linux    | ✅ Complete |
| FreeBSD  | ✅ Complete |
| Android  | 🔜 Planned |
| macOS    | 🔜 Planned |

### Output example

**Linux:**
```
lizardev@linux
   ------------------OS--------------------
   OS:       Arch Linux (Linux)
   Kernel:   7.0.9-arch2-1
   Uptime:   0d 14h 11m 25s
   Shell:    bash
   Terminal: kitty
   ------------------PC--------------------
   Host:     LENOVO ThinkPad T460 (20FMS2292A)
   CPU:      Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz (2 Cores / 4 Threads)
   Arch:     x86_64
   Graphics: Intel Corporation Skylake-U GT2 [HD Graphics 520] (rev 07)
   Disk:     64.51 / 237.46 (used: 27.17%)
   Memory:   2.83 GiB / 15.47 GiB
   Swap:     0.00 GiB / 4.00 GiB
   ----------------------------------------
```

**FreeBSD:**
```
lizardev@serverBSD
   ------------------OS--------------------
   OS:       FreeBSD (15.0-RELEASE-p8)
   Kernel:   15.0-RELEASE-p8
   Uptime:   0d 0h 0m 48s
   Shell:    bash
   Terminal: kitty
   ------------------PC--------------------
   Host:     Acer Aspire 5733Z (V1.07)
   CPU:      Intel(R) Pentium(R) CPU P6200 @ 2.13GHz (2 Cores / 2 Threads)
   Arch:     amd64
   Graphics: Intel Corporation Core Processor Integrated Graphics Controller
   Disk:     10.18 / 279.17 (used: 3.65%)
   Memory:   0.2 GiB / 1.7 GiB
   Swap:     0.0 GiB / 2.0 GiB
   ----------------------------------------
```

### Installation

#### Prerequisites

- [Go](https://go.dev/dl/) version 1.24 or higher

#### Install with Go

 ```bash
go install github.com/soylizardev/caura@latest
```

> **Note:** Make sure `~/go/bin` (or `$GOPATH/bin`) is in your `PATH`. Add this to your shell config if needed:
> ```bash
> # Bash (~/.bashrc)
> export PATH=$PATH:~/go/bin
> ```
> ```fish
> # Fish (~/.config/fish/fish_variables)
> fish_add_path ~/go/bin
> ```
> ```zsh
> # Zsh (~/.zshrc)
> export PATH=$PATH:~/go/bin
> ```

#### Install from source

```bash
# Clone the repository
git clone https://github.com/soylizardev/caura.git
cd caura

# Build the binary
go build -o caura .

# Move to a directory in your PATH (optional)
# Examples:
sudo mv caura /usr/local/bin/
# or: mv caura ~/.local/bin/

# Run it
caura
```

#### Cross-compile for another platform

```bash
# For FreeBSD (amd64)
GOOS=freebsd GOARCH=amd64 go build -o caura-freebsd .
```

### Usage

Just run `caura` in your terminal:

```bash
caura
```

No configuration files or arguments needed.

### Roadmap

- [x] Linux support
- [x] FreeBSD support
- [ ] Android support
- [ ] macOS support
- [ ] ASCII logo support
- [ ] PNG/image logo support
- [ ] Customizable output via TOML/YAML configuration file
- [ ] Plugin system for custom modules

---

## Español

### Características

- **Multiplataforma**: Funciona en Linux y FreeBSD (Android y macOS planificados)
- **Sin dependencias externas**: Solo requiere Go para compilar
- **Salida limpia**: Secciones coloreadas con información legible
- **Portátil**: Un solo binario, sin dependencias en tiempo de ejecución

### Plataformas soportadas

| Plataforma | Estado |
|------------|--------|
| Linux      | ✅ Completo |
| FreeBSD    | ✅ Completo |
| Android    | 🔜 Planificado |
| macOS      | 🔜 Planificado |

### Ejemplo de salida

**Linux:**
```
lizardev@omarchy
   ------------------OS--------------------
   OS:       Arch Linux (Linux)
   Kernel:   7.0.9-arch2-1
   Uptime:   0d 14h 11m 25s
   Shell:    bash
   Terminal: kitty
   ------------------PC--------------------
   Host:     LENOVO ThinkPad T460 (20FMS2292A)
   CPU:      Intel(R) Core(TM) i5-6300U CPU @ 2.40GHz (2 Cores / 4 Threads)
   Arch:     x86_64
   Graphics: Intel Corporation Skylake-U GT2 [HD Graphics 520] (rev 07)
   Disk:     64.51 / 237.46 (used: 27.17%)
   Memory:   2.83 GiB / 15.47 GiB
   Swap:     0.00 GiB / 4.00 GiB
   ----------------------------------------
```

**FreeBSD:**
```
lizardev@serverBSD
   ------------------OS--------------------
   OS:       FreeBSD (15.0-RELEASE-p8)
   Kernel:   15.0-RELEASE-p8
   Uptime:   0d 0h 0m 48s
   Shell:    bash
   Terminal: kitty
   ------------------PC--------------------
   Host:     Acer Aspire 5733Z (V1.07)
   CPU:      Intel(R) Pentium(R) CPU P6200 @ 2.13GHz (2 Cores / 2 Threads)
   Arch:     amd64
   Graphics: Intel Corporation Core Processor Integrated Graphics Controller
   Disk:     10.18 / 279.17 (used: 3.65%)
   Memory:   0.2 GiB / 1.7 GiB
   Swap:     0.0 GiB / 2.0 GiB
   ----------------------------------------
```

### Instalación

#### Requisitos

- [Go](https://go.dev/dl/) versión 1.24 o superior

#### Instalar con Go

 ```bash
go install github.com/soylizardev/caura@latest
```

> **Nota:** Asegúrate de que `~/go/bin` (o `$GOPATH/bin`) esté en tu `PATH`. Agrega esto a la configuración de tu shell si hace falta:
> ```bash
> # Bash (~/.bashrc)
> export PATH=$PATH:~/go/bin
> ```
> ```fish
> # Fish (~/.config/fish/config.fish)
> fish_add_path ~/go/bin
> ```
> ```zsh
> # Zsh (~/.zshrc)
> export PATH=$PATH:~/go/bin
> ```

#### Instalar desde el código fuente

```bash
# Clonar el repositorio
git clone https://github.com/soylizardev/caura.git
cd caura

# Compilar el binario
go build -o caura .

# Mover a un directorio en tu PATH (opcional)
# Ejemplos:
sudo mv caura /usr/local/bin/
# o: mv caura ~/.local/bin/

# Ejecutar
caura
```

#### Compilación cruzada para otra plataforma

```bash
# Para FreeBSD (amd64)
GOOS=freebsd GOARCH=amd64 go build -o caura-freebsd .
```

### Uso

Simplemente ejecuta `caura` en tu terminal:

```bash
caura
```

No necesita archivos de configuración ni argumentos.

### Hoja de ruta

- [x] Soporte para Linux
- [x] Soporte para FreeBSD
- [ ] Soporte para Android
- [ ] Soporte para macOS
- [ ] Logo en ASCII
- [ ] Logo en PNG/imagen
- [ ] Personalización de salida via archivo de configuración TOML/YAML
- [ ] Sistema de plugins para módulos personalizados

---

## Contributing / Contribuir

Contributions are welcome! Feel free to:

- Open an [issue](https://github.com/soylizardev/caura/issues) for bugs or feature requests
- Submit a [pull request](https://github.com/soylizardev/caura/pulls)
- Suggest improvements

Las contribuciones son bienvenidas. Puedes:

- Abrir un [issue](https://github.com/soylizardev/caura/issues) para reportar errores o sugerir funciones
- Enviar un [pull request](https://github.com/soylizardev/caura/pulls)
- Sugerir mejoras

---

## License / Licencia

GPL-3.0
