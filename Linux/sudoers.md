sudo后定义PATH环境变量

```
/etc/sudoers :11
Defaults	!env_reset
```
把`Defaults	env_reset` 改为`Defaults	!env_reset`