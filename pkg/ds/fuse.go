package ds

// FUSE 是 Filesystem in Userspace 的缩写，中文常译作 “用户态文件系统”。
// https://github.com/hanwen/go-fuse ：目前维护最活跃的 Go FUSE 库，支持 FUSE 3 协议，文档完善
// Linux 内核本身没有提供 “带捕获功能的虚拟文件系统”，但有多种机制可以实现 “捕获文件读写” 的能力，FUSE 是最灵活、最常用的方式。
// tmpfs（内存文件系统）
// FUSE（Filesystem in Userspace）—— 可实现 “捕获文件读写” 的最强大、最灵活的用户态方案
