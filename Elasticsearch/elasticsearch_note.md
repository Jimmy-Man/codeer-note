## Elasticsearch开发文档



### 运行环境

#### Elasticsearch 安装

> 下载镜像

```bash
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.8.0
```

> 运行容器

```bash
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.8.0
```

----

####  Elasticsearch head 插件安装

​	[git地址](https://github.com/mobz/elasticsearch-head)

* `git clone git://github.com/mobz/elasticsearch-head.git`
* `cd elasticsearch-head`
* `npm install`
* `npm run start`
* `open` http://localhost:9100/



#### Elasticsearch IK分词器插件安装

[github地址](https://github.com/medcl/elasticsearch-analysis-ik)

> elasticsearch-plugin 方式安装,cd到elasticsearch目录下执行以下命令

```bash
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.8.0/elasticsearch-analysis-ik-7.8.0.zip
```

注意: 选择对应的版本号的插件版本

#### Elasticsearch 拼音分词器插件安装

[git地址](https://github.com/medcl/elasticsearch-analysis-pinyin)

```bash
./bin/elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-pinyin/releases/download/v7.8.0/elasticsearch-analysis-pinyin-7.8.0.zip
```

注意: 选择对应的版本号的插件版本

#### Kibana安装

1. 官网下载或者[国内镜像下载](https://mirrors.huaweicloud.com/kibana/)

2. 解压

3. 配置`config/kibana.yml`配置文件  

   ```yml
   server.port: 5601
   server.host: "0.0.0.0"
   elasticsearch.hosts: ["http://localhost:9200"]
   kibana.index: ".kibana"
   i18n.locale: "zh-CN"
   ```

4. bin/kibana运行

5. 浏览器访问http://localhost:5601

![image-20200721143516496](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20200721143516496.png)

### ES 添加索引

> 添加索引

```json
PUT /problems/_doc/1 
{
  "title" : "发放5月工资，  借：应付工资5000",
  "tags":"工资,计提工资,发放工资,",
   "sort": 2
}
```

> 批量添加索引

```json
POST _bulk
{ "index" : { "_index" : "test", "_id" : "1" } }
{ "field1" : "value1" }
{ "delete" : { "_index" : "test", "_id" : "2" } }
{ "create" : { "_index" : "test", "_id" : "3" } }
{ "field1" : "value3" }
{ "update" : {"_id" : "1", "_index" : "test"} }
{ "doc" : {"field2" : "value2"} }
```

> 搜索索引

```

```

- `took` – Elasticsearch运行查询所需的时间（以毫秒为单位）
- `timed_out` –搜索请求是否超时
- `_shards` –搜索了多少个分片，以及成功，失败或跳过了多少个分片。
- `max_score` –找到的最相关文件的分数
- `hits.total.value` -找到了多少个匹配的文档
- `hits.sort` -文档的排序位置（不按相关性得分排序时）
- `hits._score`-文档的相关性得分（使用时不适用`match_all`）



### IK分词器

```json
GET _analyze
{
  "analyzer": "ik_smart",
  "text" : "开发工具"
}

```

`ik_smart`为最少切分

```json
GET _analyze
{
  "analyzer": "ik_max_word",
  "text" : "开发工具"
}
```

`ik_max_word`为最细粒度划分,穷尽词库的可能.

> 自定义分词

配置文件目录 `elasticsearch/config/analysis-ik/IKAnalyzer.cfg.xml`

在这个配置目录中配置自己的词典文件



###  基本操作命令

| 请求方法 | URL地址                                           | 说明                 |
| :------: | :------------------------------------------------ | :------------------- |
|   PUT    | localhost:9200/索引名称/类型名称/文档id           | 创建文档(指定文档id) |
|   POST   | Localhsot:9200/索引名称/类型名称                  | 创建文档(随机文档id) |
|   POST   | `Localhost:9200/索引名称/类型名称/文档id/_update` | 修改文档             |
|  DELETE  | `localhost:9200/索引名称/类型名称/文档id`         | 删除文档             |
|   GET    | `localhost:9200/索引名称/类型名称/文档id`         | 查询文档通过文档id   |
|   POST   | `localhost:9200/索引名称/类型名称/_search`        | 查询所有数据         |



#### 指定类型

* 字符类型
  + `text`,`keyword`
* 数值类型
  + `long`,` integer`,`short`,`byte`,`double`,`float`,`half float`,`scaled float`
* 日期类型
  + `date`
* 布尔类型
  + `boolean`
* 二进制类型
  + `binary`



> #### 快速创建索引与文档

```json
PUT /course3/type/2
{
  "title": "会计基础",
  "age": 18,
  "tags":"ab,cd,ef"
}
```

> 创建索引结构[创建数据结构]

```json
PUT /course2
{
  "mappings": {
    "properties": {
      "title":{
        "type": "text"
      },
      "age":
      {
        "type": "integer"
      },
      "birthday":
      {
        "type": "date"
      }
    }
  }
}
```

> 查看创建的信息

```
GET /索引名称
```

> GET _cat/获取es当前许多信息

```json
GET _cat/health
```

> 修改文档

```json
POST /user/_doc/1/_update
{
  "doc":{
    "title" : "无悬念霜"
  }
}
```

> 删除

```json
DELETE /user/_doc/1
```

通过`DELETE`命令实现删除,可以根据参数是索引还是文档来实现删除索引或者文档.

> 简单查询 

```json
GET /course3/_search?q=title:基础
```

> 复杂查询 

```json
GET /course3/_search
{
  "query": {
    "match": {
      "title": "会计"
    }
  },
  //"_source": ["age",20],
  "sort": [
    {
      "age": {
        "order": "desc"
      }
    }
  ]
}
```

> 过滤 `_source`

> 排序 `sort`

> 分页 `from`,`size`

> 布尔筛选

`must`(and),所有的条件都要符合,相当于sql `where a = ** and b = ***`

```json
GET /course3/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "title": "会计"
          }
        },
        {
          "match": {
            "age": 20
          }
        }
      ]
    }
  }
}
```

`should`(or) 期中一个条件符合,相当于SQL `where a= *** or b = *** `

```json
GET /course3/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "title": "会计"
          }
        },
        {
          "match": {
            "age": 20
          }
        }
      ]
    }
  }
}
```

`must_not`(not) 除了条件之外的数据,相当于SQL `where a not in ()`

```json
GET /course3/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "match": {
            "title": "会计"
          }
        },
        {
          "match": {
            "age": 20
          }
        }
      ]
    }
  }
}
```

`filter`过滤器

```json
GET /course3/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "match": {
            "title": "会计"
          }
        }
      ],
      "filter": [
        {
          "range": {
            "age": {
              "gte": 10,
              "lte": 20
            }
          }
        }
      ]
    }
  }
}
```

* `gt`大于
* `get` 大于等于
* `lt`小于
* `lte`小于等于

> 多个条件

多个条件使用空格隔开,只要满足期中一个条件便可查询

```json
GET /course3/_search
{
  "query": {
    "match": {
      "tags": "条件1 条件2"
    }
  }
}
```

> 精确 查询 

`term`查询是直接通过倒排索引指定的词条进程精确查找的

关于分词:

Term,直接查询精确的

Match,会使用分词器解析!(先分析文档,然后在通过分析的文档进行查询)

*两个类型* `keyword`,`text`

```json
GET /course3/_search
{
  "query": {
   "term": {
     "title": {
       "value": "基础"
     }
   }
  }
}
```

> 多个匹配值查询 

> 高亮查询 `highlight`

```json
GET /course3/_search
{
  "query": {
    "match": {
      "title": "基础"
    }
  },
  "highlight": {
    "pre_tags": "<p class='key' style='color:red'", 
    "post_tags": "</p>", 
    "fields": {
      "title": {}
    }
  }
}
```



