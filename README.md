# `fileserver` 工具使用指南

## 1. 安装方法

要安装 `fileserver` 工具，请执行以下命令：

```bash
go install github.com/zoulux/fileserver@latest
```

## 2. 使用说明

### 2.1 基本用法

安装完成后，您可以直接在命令行中使用 `fileserver` 命令：

```bash
sendfile [参数]
```

### 2.2 常用参数

| 参数 | 描述               |
|-----|------------------|
| `-port` | 指定端口号 (默认: 8080) |
| `-dir` | 路径（默认 ./）        |
| `-h` | 显示帮助信息           |

## 3. 示例

- 服务端
```bash
./fileserver 
```

- 客户端
```bash
wget http://remotehost:80/file.txt
```