# cgo

cgo程序需要通过gcc编译，依赖环境[mingw](https://sourceforge.net/projects/mingw-w64/)

安装下载的`mingw-w64-install.exe`，选择x86_64代表安装64位，windows下选择win32

安装完成之后需要将 新的GCC的位置添加到环境变量path中：
- %MINGW64_HOME%
- %MINGW64_HOME%\bin

## 配置全局 make

进入 %MINGW64_HOME% 创建make.bat
```
@echo off
xxx\bin\mingw32-make.exe %1 %2 %3 %4 %5 %6 %7 %8 %9
```