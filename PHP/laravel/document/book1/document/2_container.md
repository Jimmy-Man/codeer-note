服务容器(ICO容器)
===============


## 容器的使用
通过IoC容器可以帮助我们更方便地管理类依赖，而且Laravel提供了一个功能强大的IoC容器。这个IoC容器在Laravel中被称作服务容器，是整个Laravel框架最核心的部分，在它的调度下，框架各个组件可以很好的组合在一起工作。实际上，Laravel的`Application`类就是一个继承自`Container`的容器类，它就是整个`Laravel`应用的服务容器。

IoC 容器：`控制反转容器让依赖注入更方便，它负责在整个应用生命周期内解析和注入那些定义在容器中的类和接口。`  

`Laravel`应用中,通过`App`|`app()`来访问服务容器,在服务提供者中使用`$this->app`访问服务容器.  
`Laravel`中服务容器的操作[类与接口绑定到容器]在`app/Providers/`目录下
```php
    //app/Providers/TestProvider.php
    public function register()
    {
        $this->app->bind(paymentInterface::class,function ($app){
            return new PaymentMethod($app->make(notifyInterface::class));
        });

        $this->app->bind(notifyInterface::class,function($app){
            return new ConsoleNotify();
        });
        //
        $this->app->singleton(notifyInterface::class,function ($app){
            return new ConsoleNotify();
        });
    }
```  
注：`注意到我们在定义绑定关系的时候使用的是匿名函数，这样做的好处是用到该依赖时才会实例化，从而提升了应用的性能。`  

服务容器就是个用来注册各种接口与实现绑定的地方。一旦一个类在容器里注册了以后，就可以很容易地在应用的任何位置解析并调用它.  

有时候，你可能想在整个应用生命周期中只实例化某类一次，类似单例模式，可以通过`singleton`方法来注册接口与实现类:
```php
$this->app->singleton(notifyInterface::class,function ($app){
            return new ConsoleNotify();
        });
```
只要服务容器解析过这个账单通知对象实例一次，在剩余的请求生命周期中都会使用同一个实例。

服务容器还提供了和`singleton`方法很类似的`instance`方法，区别是`instance`方法可以绑定一个已经存在的对象实例。然后容器每次解析的时候都会返回这个对象实例
```php
        $notify = new ConsoleNotify;
        $this->app->instance(notifyInterface::class,$notify);
```

单独使用容器：即使你的项目不是基于 Laravel 框架的，依然可以使用Laravel 的服务容器，只要通过 Composer 安装 illuminate/container 就好了.

## 反射解决方案

Laravel 服务容器中最强大的功能之一就是通过反射来自动解析类的依赖。　　
反射是一种在运行时检查类和方法的能力，比如PHP的 ReflectionClass 类可以动态检查给定类的所有方法，PHP 函数 method_exists 从某种意义上说也是一种反射  
[PHP手册中的反射文档](https://www.php.net/manual/zh/book.reflection.php)  
