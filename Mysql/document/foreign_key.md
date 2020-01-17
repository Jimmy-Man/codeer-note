外键
===


添加外键


添加外键约束
~~~SQL
ALTER TABLE `表名` add CONSTRAINT `外键名` FOREIGN KEY (`字段名`) REFERENCES `主表名`(`主表字段`);
ALTER TABLE `表名` add CONSTRAINT `外键名` FOREIGN KEY (`字段名`) REFERENCES `主表名`(`主表字段`) ON DELETE CASCADE ON UPDATE NO ACTION;

~~~

外键关联表联合删除选项  
[on delete {cascade | set null | no action| restrict}]  
[on update {cascade | set null | no action| restrict}]  
`restrict` 是默认操作，表示拒绝主表删除或修改外键关联列，这是最安全的设置  
`cascade` 表示删除包含与已删除键值有参考关系的所有记录  

删除外键约束
```sql
ALTER TABLE `表名` drop FOREIGN KEY `外键约束名`;
```