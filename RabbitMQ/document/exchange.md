Exchange 交换机
--------------

Exchange 接收消息,并根据路由键转发消息所绑定的队列.

#### Exchange 交换机属性
* Name : 交换机名称
* Type : 次的机类型: `direct`,`topic`,`fanout`,`headers`
* Durability : 是否需要持久化,ture为持久化
* Auto Delete: 当最后一个绑定到`Exchange`上的队列删除后,自动删除该Exchange
* Internal : 当前Exchange是否用于RabbitMQ内部使用,默认为false [很小使用]
* Arguments : 扩展参数,用于扩展AMQP协议可制定化使用

#### Exchange Type类型
##### `Direct Exchange`
* 所有发送到`Direct Exchange`的消息被转发到`RouteKey`中指定的`Queue`
注意: `Direct`模式可以使用RabbitMQ自带的`Exchange:default` Exchange,所以不需要将`Exchange`进行任何绑定(`binding`)操作,消息传递时,`RouteKey`必须完全匹配才会被队列接收,否则该消息会被抛弃.

##### `Topic Exchange`
* 所有发送到`Topic Exchange`的消息被转发到所有关心`RouteKey`中指定`Topic`的`Queue`上
* `Exchange`将`RouteKey`和某`Topic`进行模糊匹配,此时队列需要绑定一个`Topic`
* 可以使用通配符进行模糊匹配
  + `#`: 匹配一个或多个
  + `*`: 匹配不多不少一个词