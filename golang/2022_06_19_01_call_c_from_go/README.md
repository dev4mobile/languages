查看动态链接库
1. Linux： ldd
2. Mac OS X: otool


gcc 

gcc -print-search-dirs  // 查看查找库的路径

gcc -dynamiclib -o libgreeter.dylib greeter.c  // Mac OS X 编译动态链接库
