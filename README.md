# TR Server Control (TRSC)

Minecraft Java 服务端本机管理面板。Go 后端 + 嵌入式网页。

仓库：https://github.com/trmccc/TRSC

## 运行

Windows：双击 `启动TRSC.bat`  
Linux：`./bin/trsc-linux-amd64`

默认地址 `http://127.0.0.1:8080`  
默认账号 `admin` / `admin`

## 更新

面板登录页和设置页可「检查更新」，对照本仓库 GitHub Releases。  
发布新版本时请打 tag（如 `v0.3.1`）并上传：

- `trsc-windows-amd64.exe`
- `trsc-linux-amd64`

设置里的「自动更新」会下载对应系统的二进制并重启面板。

## 构建

```
go build -o bin/trsc-linux-amd64 ./cmd/trsc
GOOS=windows GOARCH=amd64 go build -o bin/trsc-windows-amd64.exe ./cmd/trsc
```
