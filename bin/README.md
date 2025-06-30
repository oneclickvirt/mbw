# mbw Static Binaries

This directory contains statically compiled binaries of [mbw](https://github.com/raas/mbw) for various platforms and architectures.

## About mbw

mbw is a memory bandwidth benchmark tool that measures the bandwidth of memory subsystems.

## Available Binaries

### Linux
- `mbw-linux-amd64` - Linux x86_64 (optimized)
- `mbw-linux-amd64-compat` - Linux x86_64 (maximum compatibility, no advanced CPU instructions)
- `mbw-linux-386` - Linux x86 (32-bit)
- `mbw-linux-arm64` - Linux ARM64
- `mbw-linux-armv7` - Linux ARMv7
- `mbw-linux-armv6` - Linux ARMv6
- `mbw-linux-riscv64` - Linux RISC-V 64-bit
- `mbw-linux-mips64` - Linux MIPS64 (big-endian)
- `mbw-linux-mips64le` - Linux MIPS64 (little-endian)
- `mbw-linux-mips` - Linux MIPS (big-endian)
- `mbw-linux-mipsle` - Linux MIPS (little-endian)
- `mbw-linux-ppc64` - Linux PowerPC64 (big-endian)
- `mbw-linux-ppc64le` - Linux PowerPC64 (little-endian)

### macOS
- `mbw-darwin-amd64` - macOS x86_64 (macOS 10.12+)
- `mbw-darwin-arm64` - macOS ARM64 (Apple Silicon, macOS 11.0+)

### Windows
- `mbw-windows-amd64.exe` - Windows x86_64
- `mbw-windows-386.exe` - Windows x86 (32-bit)
- `mbw-windows-arm64.exe` - Windows ARM64

### BSD
- `mbw-freebsd-amd64` - FreeBSD x86_64
- `mbw-freebsd-arm64` - FreeBSD ARM64
- `mbw-openbsd-amd64` - OpenBSD x86_64
- `mbw-openbsd-arm64` - OpenBSD ARM64
- `mbw-netbsd-amd64` - NetBSD x86_64

## Compatibility Notes

- **Linux amd64-compat**: Built without SSE4.2, AVX, and AVX2 instructions for maximum compatibility with older CPUs
- **Windows binaries**: Statically linked with MinGW-w64, should work on Windows 7+ without additional dependencies
- **macOS binaries**: Built with minimum version requirements for broader compatibility
- **All Linux binaries**: Statically linked for maximum portability

## Usage

### Linux/macOS/BSD
```bash
# Make executable (if needed)
chmod +x mbw-linux-amd64

# Run memory bandwidth test
./mbw-linux-amd64 256
```

### Windows
```cmd
# Run memory bandwidth test
mbw-windows-amd64.exe 256
```

## Build Information

These binaries are automatically built using GitHub Actions from the [mbw source repository](https://github.com/raas/mbw).
All binaries are optimized for compatibility and performance.

Source: https://github.com/raas/mbw
Built in: oneclickvirt/mbw
