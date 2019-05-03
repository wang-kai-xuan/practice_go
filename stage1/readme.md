# demo简介
借助tcp实现文件的传输

# 实现思路
1. client向server发送文件名，server保存文件名
1. server向client返回一个消息ok，确认文件名保存成功
1. client收到消息后先server发送文件数据
1. server读取文件内容，写入到之前保存的文件中