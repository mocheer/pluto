
# pprof

```bash
go tool pprof -inuse_space http://127.0.0.1:9912/debug/pprof/heap
```

### flag
-alloc_space
-inuse_space

### column
flat：给定函数上运行耗时
flat%：同上的 CPU 运行耗时总比例
sum%：给定函数累积使用 CPU 总比例
cum：当前函数加上它之上的调用运行总耗时
cum%：同上的 CPU 运行耗时总比例
-inuse_space：分析应用程序的常驻内存占用情况
-alloc_objects：分析应用程序的内存临时分配情况top