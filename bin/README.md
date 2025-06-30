# mbw Static Binaries

This directory contains statically compiled binaries of [mbw](https://github.com/raas/mbw) for various platforms and architectures.

## About mbw

mbw is a memory bandwidth benchmark tool that measures the bandwidth of memory subsystems.

## Available Binaries

### Linux
- `mbw-linux-amd64` - Linux x86_64
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
- `mbw-darwin-amd64` - macOS x86_64
- `mbw-darwin-arm64` - macOS ARM64 (Apple Silicon)

### BSD
- `mbw-freebsd-amd64` - FreeBSD x86_64
- `mbw-freebsd-arm64` - FreeBSD ARM64
- `mbw-openbsd-amd64` - OpenBSD x86_64
- `mbw-openbsd-arm64` - OpenBSD ARM64
- `mbw-netbsd-amd64` - NetBSD x86_64

## Usage

Download the appropriate binary for your platform and run it:

```bash
# Make executable (if needed)
chmod +x mbw-linux-amd64

# Run memory bandwidth test
./mbw-linux-amd64 256
```

## Build Information

These binaries are automatically built using GitHub Actions from the [mbw source repository](https://github.com/raas/mbw).
All Linux binaries are statically linked for maximum compatibility.

Source: https://github.com/raas/mbw
Built in: oneclickvirt/mbw

