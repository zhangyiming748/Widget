#!/bin/bash
echo 删除多余隐藏文件
find . -name ".DS_Store" -exec rm {} \;
echo 删除外部驱动器产生的挂载点文件
find . -size -4k -type f -name ".*"
# shellcheck disable=SC2039
read -p "输入yes确认删除:" -t 30 confirm
# shellcheck disable=SC2039
if [ $confirm == yes ]; then
  find . -size -4k -type f -name ".*" -exec {} \;
fi
