## Capistrano Note

Capistrano 远程服务器自动化工具

```Ruby
# 列出所有可用的任务
$ bundle exec cap -T

# 部署到暂存环境 deploy to the staging environment 
$ bundle exec cap staging deploy

# 部署到生产环境 deploy to the production environment 
$ bundle exec cap production deploy

# 可模拟部署到生产环境，实际上并没有做任何事情
$ bundle exec cap production deploy --dry-run

# 列表任务相关  list task dependencies
$ bundle exec cap production deploy --prereqs

# 通过任务调用跟踪  trace through task invocations
$ bundle exec cap production deploy --trace

# 在部署任务之前列出所有配置变量  lists all config variable before deployment tasks
$ bundle exec cap production deploy --print-config-variables

```