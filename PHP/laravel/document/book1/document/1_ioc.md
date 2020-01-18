依赖注入
=======

`整个Laravel框架的基石是一个功能强大的IoC容器（控制反转容器），如果你想真正从底层理解Laravel框架，就必须好好掌握它.`  
`而且要实现依赖注入并不一定非要通过 IoC 容器，只是使用IoC容器会更容易一点儿`

面向接口开发：`编写接口看上去好像要多写一些代码，但是磨刀不误砍柴工，对于大型项目而言实际上反而能提升你的开发效率，这就是软件设计领域经常说的面向接口开发，而不是面向对象开发。`

```php
//支付接口
interface paymentInterface{
    public function pay($amount);
}
//通知接口
interface notifyInterface{
    public function send(string $message);
}

//通知接口实现类
class ConsoleNotify implements notify {
    public function send(string $message)
    {
        echo $message;
    }
}

//支付接口实现类
class PaymentMethod implements paymentInterface {

    public function __construct(notifyInterface $notify)
    {
        $this->notify =  $notify;
    }

    public function pay($amount)
    {
        echo 'payed';
        $msg =  "pay $amount success!";
        $this->notify->send($msg);
    }

}

$notify = new ConsoleNotify();

$pay = new PaymentMethod($notify);
$pay->pay(1000);
```

`难道依赖注入不需要IoC容器了么？当然不需要！IoC容器使得依赖注入更易于管理，但是容器本身不是依赖注入所必须的.`