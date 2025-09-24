# mbw

一个嵌入mbw依赖的golang库(A golang library with embedded mbw dependencies)

## 关于 About

这个库提供了对MBW内存带宽基准测试工具的Go语言封装。MBW是一个简单的内存带宽基准测试程序，用于测量可持续的内存带宽（MB/s）和相应的计算速率。它通过三种不同的测试方法（MEMCPY、DUMB、MCBLOCK）来评估内存子系统的性能。

This library provides a Go wrapper for the MBW memory bandwidth benchmark tool. MBW is a simple memory bandwidth benchmark program that measures sustainable memory bandwidth (in MB/s) and the corresponding computation rate. It evaluates memory subsystem performance through three different test methods (MEMCPY, DUMB, MCBLOCK).

## 特性 Features

- 支持多平台：Linux (amd64, arm64, 386, arm等), macOS (amd64, arm64), FreeBSD, OpenBSD, NetBSD等
- 自动检测并使用系统安装的mbw，或使用嵌入的二进制文件
- 自动清理临时文件
- 三种内存带宽测试方法：MEMCPY、DUMB（简单循环）、MCBLOCK（块复制）
- Multi-platform support: Linux (amd64, arm64, 386, arm, etc.), macOS (amd64, arm64), FreeBSD, OpenBSD, NetBSD, etc.
- Automatically detects and uses system-installed mbw, or uses embedded binaries
- Automatic cleanup of temporary files
- Three memory bandwidth test methods: MEMCPY, DUMB (simple loop), MCBLOCK (block copy)

## 安装 Installation

```bash
go get github.com/oneclickvirt/mbw@v0.0.2-20250924041746
```

## 使用方法 Usage

```go
package main

import (
    "log"
    "github.com/oneclickvirt/mbw"
)

func main() {
    // 获取mbw命令路径
    // Get mbw command path
    mbwCmd, tempFile, err := mbw.GetMBW()
    if err != nil {
        log.Fatalf("Failed to get mbw: %v", err)
    }

    // 如果使用了临时文件，确保清理
    // Clean up temporary files if used
    if tempFile != "" {
        defer mbw.CleanMBW(tempFile)
    }

    // 执行mbw基准测试，测试64MB内存
    // Execute mbw benchmark, testing 64MB memory
    err = mbw.ExecuteMBW(mbwCmd, []string{"64"})
    if err != nil {
        log.Fatalf("Failed to execute mbw: %v", err)
    }
}
```

### 高级用法示例 Advanced Usage Examples

```go
// 运行指定次数的测试
// Run specific number of tests
err = mbw.ExecuteMBW(mbwCmd, []string{"-n", "5", "128"})

// 只运行memcpy测试
// Run only memcpy test
err = mbw.ExecuteMBW(mbwCmd, []string{"-t0", "64"})

// 静默模式，只显示统计信息
// Quiet mode, show statistics only
err = mbw.ExecuteMBW(mbwCmd, []string{"-q", "256"})

// 自定义块大小
// Custom block size
err = mbw.ExecuteMBW(mbwCmd, []string{"-t2", "-b", "1048576", "64"})
```

## API文档 API Documentation

### 函数 Functions

#### `GetMBW() (string, string, error)`

获取可用的mbw命令路径。

Get the available mbw command path.

**返回值 Returns:**
- `string`: mbw命令的路径 (Path to the mbw command)
- `string`: 临时文件路径（如果使用了嵌入的二进制文件）(Temporary file path if using embedded binary)
- `error`: 错误信息 (Error information)

**行为 Behavior:**
1. 首先尝试在系统PATH中查找mbw命令
2. 如果未找到，则从嵌入的二进制文件中提取对应平台的mbw
3. 将二进制文件写入临时目录并设置执行权限

**Behavior:**
1. First attempts to find the mbw command in the system PATH
2. If not found, extracts the appropriate platform mbw from embedded binaries
3. Writes the binary to a temporary directory and sets execution permissions

#### `ExecuteMBW(mbwPath string, args []string) error`

执行mbw命令。

Execute the mbw command.

**参数 Parameters:**
- `mbwPath`: mbw命令的路径 (Path to the mbw command)
- `args`: 传递给mbw的参数 (Arguments to pass to mbw)

**常用参数 Common Arguments:**
- `数字`: 要测试的内存大小（MB） (Memory size to test in MB)
- `-n <num>`: 每个测试运行的次数 (Number of runs per test)
- `-t0`: 只运行memcpy测试 (Run only memcpy test)
- `-t1`: 只运行dumb测试 (Run only dumb test)
- `-t2`: 只运行固定块大小的memcpy测试 (Run only fixed block size memcpy test)
- `-b <size>`: 设置块大小（字节） (Set block size in bytes)
- `-q`: 静默模式 (Quiet mode)
- `-a`: 不显示平均值 (Don't display average)

#### `CleanMBW(tempFile string) error`

清理临时文件。

Clean up temporary files.

**参数 Parameters:**
- `tempFile`: 需要清理的临时文件路径 (Path to temporary file to clean up)

## 平台支持 Platform Support

库包含了以下平台的预编译二进制文件：

The library includes precompiled binaries for the following platforms:

### Linux
- amd64 (x86_64)
- 386 (x86 32-bit) 
- arm64 (ARMv8)
- armv7 (ARMv7)
- riscv64 (RISC-V 64-bit)
- ppc64le (PowerPC64 little-endian)
- ppc64 (PowerPC64 big-endian)
- mips64le (MIPS64 little-endian)
- mips64 (MIPS64 big-endian)
- mipsle (MIPS little-endian)
- mips (MIPS big-endian)

### macOS
- amd64 (Intel x86_64)
- arm64 (Apple Silicon)

### BSD系统 BSD Systems
- FreeBSD (amd64, arm64)
- OpenBSD (amd64)
- NetBSD (amd64)

## 测试方法说明 Test Methods

MBW提供三种不同的内存带宽测试方法：

MBW provides three different memory bandwidth test methods:

1. **MEMCPY**: 使用标准库的memcpy函数进行内存复制 (Uses standard library memcpy function for memory copying)
2. **DUMB**: 使用简单的循环逐个复制数组元素 (Uses simple loop to copy array elements one by one) 
3. **MCBLOCK**: 使用固定块大小的memcpy进行分块复制 (Uses fixed block size memcpy for block copying)

## 示例输出 Sample Output

```
Long uses 8 bytes. Allocating 2*8388608 elements = 134217728 bytes of memory.
Using 262144 bytes as blocks for memcpy block copy test.
Getting down to business... Doing 10 runs per test.
0	Method: MEMCPY	Elapsed: 0.00403	MiB: 64.00000	Copy: 15869.080 MiB/s
...
AVG	Method: MEMCPY	Elapsed: 0.00301	MiB: 64.00000	Copy: 21264.578 MiB/s
...
```

## 注意事项 Notes

- 本程序仅可在类Unix系统中运行 (This program can only run on Unix-like systems)
- 测试结果可能受到系统负载、内存频率、CPU缓存等因素影响 (Test results may be affected by system load, memory frequency, CPU cache, etc.)
- 建议在系统空闲时运行测试以获得更准确的结果 (It's recommended to run tests when the system is idle for more accurate results)
- 大内存测试可能会消耗大量系统内存，请根据系统配置选择合适的测试大小 (Large memory tests may consume significant system memory, please choose appropriate test size based on system configuration)

## 许可证 License

MIT License