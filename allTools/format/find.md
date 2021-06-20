---
layout:     post                    # 使用的布局(不需要改)
title:      find命令详解             # 标题
subtitle:   我是一只被禁足的安小鸟 # 副标题
date:       2021-06-09 00:00:00 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: False                     # 是否归档
tags:                               #标签
    - shell
---
“”‘’
# 批量删除macOS读取外部驱动器产生的4k挂载点文件
`find ./ -size 4k -exec rm -f {} \;`

# 批量删除macOS自动建立的`.DS_Store`文件
`find ./ -name ".DS_Store" -depth -exec rm {} \;`

# 查找并显示文件
`find ./ -name '#.txt' -print`

# 查找指定范围文件
`find . -type f -mtime -1 -size +100k`

# 查找空文件
`find -type d -empty`

# 查询出所有的空文件夹
`find -type d -empty`

# 查询所有/root/下的空文件夹
`find /root -type d -empty` 

# 列出搜索到的文件/删除文件
`find . -name "shuaige.txt" -exec ls {} ;`

# 批量删除搜索到的文件
`find . -name "shuaige.txt" -exec rm -f {} ;`

# 删除前有提示
`find . -name "shuaige.txt" -ok rm -rf {} ;`

# 删除当前目录下面所有 test 文件夹下面的文件
`find . -name "test" -type d -exec rm -rf {} ;`

# 删除文件夹下面的所有的.svn文件
`find . -name '.svn' -exec rm -rf {} ;`

1. {}和之间有一个空格 注:
2. find . -name 之间也有空格 
3. exec 是一个后续的命令,{}内的内容代表前面查找出来的文件