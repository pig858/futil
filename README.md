# futil

futil 是一個用go編寫的CLI，提供了兩個主要功能：計算文件的校驗和（checksum）和計算文件的行數（linecount）。

## 應用設計

futil 的設計旨在簡化文件處理過程。用戶可以通過簡單的命令來獲取文件的校驗和或行數，且工具支持從標準輸入讀取數據，使其能夠與其他命令結合使用！

### 功能

1. **Checksum**: 計算指定文件的 MD5、SHA-1 或 SHA-256 校驗和。
2. **Line Count**: 計算指定文件中的行數。

## 使用的第三方套件

- **Cobra**: 用於創建命令行界面，提供簡單的命令和標誌管理。
- **mimetype**: 用於檢測文件類型，確保工具使用時的檔案類型是正確的。

## 如何構建和運行項目

### 前提條件

請確保已安裝 Go 語言環境（版本 1.22.5或更高）。

### 構建項目

在項目根目錄下運行以下命令來構建工具：
```bash
Windows
go build

Linux
go build -o futil
```
### 執行指令

#### checksum
```bash
Windows
./futil.exe checksum -f myfile.txt --md5

Linux
./futil checksum -f myfile.txt --md5
```
#### linecount
```bash
Windows
./futil.exe linecount -f myfile.txt

Linux
./futil linecount -f myfile.txt
```
#### 如果想要從標準輸入讀取數據可以使用：
```bash
Windows
type myfile.txt | ./futil.exe checksum -f - --md5

Linux
type myfile.txt | ./futil checksum -f - --md5
```

### 尚未實現的功能
- 支援更多的文件類型
- 提供更詳細的錯誤處理

### 已知的問題
