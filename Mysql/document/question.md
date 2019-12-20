常见问题
========

### Mysql5.7 报this is incompatible with sql_mode=only_full_group_by错误
Mysql5.7 默认是开启`ONLY_FULL_GROUP_BY`模式的,这会导致原有SQL语句报错
```
### 查看sql_mode的值
select @@GLOBAL.sql_mode
```
从查询到的值中开掉`ONLY_FULL_GROUP_BY`然后保存
```
set @@GLOBAL.sql_model='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'
```

