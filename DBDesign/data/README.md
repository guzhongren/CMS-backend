# 数据库文件操作

## 备份

> 以 postgres 用户备份 cms 数据库
```shell
$ cd /Library/PostgreSQL/11/bin
$ ./pg_dump -h 47.95.247.139 -U postgres cms>~/Desktop/Temp/cms.backup
```

## 恢复
> 需在要恢复的数据库上建立 cms 数据库

> 以 postgres 用户恢复备份的数据到本地 cms 数据库中
```shell
$ cd /Library/PostgreSQL/11/bin
./psql -h localhost -U postgres -d cms -f ~/Desktop/Temp/cms.backup
```