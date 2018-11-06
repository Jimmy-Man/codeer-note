Symfony 笔记：

### 安装 

composer 安装
    --website-skeleton是web应用优化版
```ssh
composer create-project symfony/website-skeleton project-name
```
    -- 徽服务 或者 API
```
composer create-project symfony/skeleton peoject-name
cd project-name
composer require symfony/web-server-bundle --dev
```
-----------------------------------------------------------------
### 运行
```
cd project-name
php bin/console server:run
```
Lnmp或者LAMP环境下：
待继.....

---------------------------------------------------------------
#### 检查服务器环境要求

```
cd your-project/
composer require symfony/requirements-checker
```
该组件会在Public目录下创建check.php文件，在浏览器访问该文件排查服务器环境问题
检查完问题后记得删除该组件
```
cd your-project/
composer remove symfony/requirements-checker
```
-----------------------------------------------------------------

#### 检查安全漏洞
```
cd my-project/
composer require sensiolabs/security-checker --dev
```

#### Symfony演示应用程序
https://github.com/symfony/demo
```
composer create-project symfony/symfony-demo
```


--------------------------------------------------------------
## 数据库和Doctrine ORM

symfony框架并未整合任何需要使用数据库的组件，但是却紧密集成了一个名为 Doctrine 的三方类库。

#### 安装Doctrine
```
composer require symfony/orm-pack
composer require symfony/maker-bundle --dev
```

#### 配置数据库
symfony2、3：
```yaml
# app/config/parameters.yml
parameters:
    database_host:      localhost
    database_name:      test_project
    database_user:      root
    database_password:  password
 
# ...
```
symfony4 
```
# .env

# customize this line!
DATABASE_URL="mysql://db_user:db_password@127.0.0.1:3306/db_name"

# to use sqlite:
# DATABASE_URL="sqlite:///%kernel.project_dir%/var/app.db"
```
##### 创建数据库
```
php bin/console doctrine:database:create
```
#### Doctrine 完整命令列表
```
php bin/console list doctrine
```
### 创建一个Entity类
命令行创建：
symfony2,3
```
php bin/console doctrine:generate:entity
```

或者手动创建：
```php
// src/AppBundle/Entity/Product.php
namespace AppBundle\Entity;
 
class Product
{
    private $name;
    private $price;
    private $description;
}
...
```
### 验证映射
创建entity之后，你应该使用以下命令来验证映射（mappings）：
```
php bin/console doctrine:schema:validate
```
#### 生成Getters和Setters
生成简单Getters和Setters
```
php bin/console doctrine:generate:entities AppBundle/Entity/Product
```
复杂的功能，添加逻辑进去，以满足程序之需求

### 创建数据表/Schema 
生成Entity类后，就要在数据库生成对应的数据表
```
php bin/console doctrine:schema:update --force
```
注意：doctrine:schema:update 命令只适合在开发环境中使用。它不应该被用于生产环境。

××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××
symfony4 可以自动添加Entity类与Getters,Setters
```
php bin/console make:entity
```
#### 创建迁移SQL
```
php bin/console make:migration
```
该命令会创建 src/Migrations/Version2018*****.php 的迁移文件。

#### 执行变更的文件
```
php bin/console doctrine:migrations:migrate
```
或者打开生成的迁移文件中的SQL,在数据库中执行该SQL。
需要添加新字段或者修改重复以上Symfony4的流程即可

××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××××

### 进行增、删、改、查操作
#### 将对象保存到数据库
```php
// src/AppBundle/Controller/DefaultController.php
 
// ...
use AppBundle\Entity\Product;
use Symfony\Component\HttpFoundation\Response;
 
// ...
public function createAction()
{
    $product = new Product();
    $product->setName('Keyboard');
    $product->setPrice(19.99);
    $product->setDescription('Ergonomic and stylish!');
 
    $em = $this->getDoctrine()->getManager();
 
    // tells Doctrine you want to (eventually) save the Product (no queries yet)
    // 告诉Doctrine你希望（最终）存储Product对象（还没有语句执行）
    $em->persist($product);
 
    // actually executes the queries (i.e. the INSERT query)
    // 真正执行语句（如，INSERT 查询）
    $em->flush();
 
    return new Response('Saved new product with id '.$product->getId());
}
```
####  从数据库中获取对象
```php
public function showAction($productId)
{
    $product = $this->getDoctrine()
        //->getRepository(Product::class)
        ->getRepository('AppBundle:Product')
        ->find($productId);
 
    if (!$product) {
        throw $this->createNotFoundException(
            'No product found for id '.$productId
        );
    }
 
    // ... do something, like pass the $product object into a template
    // ... 做一些事，比如把 $product 对象传入模板
}
```
当你要查询某个特定类型的对象时，你总是要使用它的”respository”,一旦有了Repository对象，你就可以访问它的全部有用的方法了。
```php
$repository = $this->getDoctrine()->getRepository('AppBundle:Product');
 
// query for a single product by its primary key (usually "id")
// 通过主键（通常是id）查询一件产品
$product = $repository->find($productId);
 
// dynamic method names to find a single product based on a column value
// 动态方法名称，基于字段的值来找到一件产品
$product = $repository->findOneById($productId);
$product = $repository->findOneByName('Keyboard');
 
// dynamic method names to find a group of products based on a column value
// 动态方法名称，基于字段值来找出一组产品
$products = $repository->findByPrice(19.99);
 
// find *all* products / 查出 *全部* 产品
$products = $repository->findAll();
```
你也可以有效利用 findBy 和 findOneBy 方法，基于多个条件来轻松获取对象：
```php
$repository = $this->getDoctrine()->getRepository('AppBundle:Product');
 
// query for a single product matching the given name and price
// 查询一件产品，要匹配给定的名称和价格
$product = $repository->findOneBy(
    array('name' => 'Keyboard', 'price' => 19.99)
);
 
// query for multiple products matching the given name, ordered by price
// 查询多件产品，要匹配给定的名称和价格
$products = $repository->findBy(
    array('name' => 'Keyboard'),
    array('price' => 'ASC')
);
```
#### 对象更新
```php
public function updateAction($productId)
{
    $em = $this->getDoctrine()->getManager();
    $product = $em->getRepository('AppBundle:Product')->find($productId);
 
    if (!$product) {
        throw $this->createNotFoundException(
            'No product found for id '.$productId
        );
    }
 
    $product->setName('New product name!');
    $em->flush();
 
    return $this->redirectToRoute('homepage');
}
```
#### 删除对象
```php
$em->remove($product);
$em->flush();
```
### 使用DQL进行对象查询
```php
$em = $this->getDoctrine()->getManager();
$query = $em->createQuery(
    'SELECT p
    FROM AppBundle:Product p
    WHERE p.price > :price
    ORDER BY p.price ASC'
)->setParameter('price', 19.99);
 
$products = $query->getResult();
```
getResult() 方法返回一个结果数组。要得到一个结果，可以使用getSingleResult()（这个方法在没有结果时会抛出一个异常）或者 getOneOrNullResult() ：
```php
$product = $query->setMaxResults(1)->getOneOrNullResult();
```

#### 使用Doctrine's Query Builder进行对象查询
```php
$repository = $this->getDoctrine()
    ->getRepository('AppBundle:Product');
 
// createQueryBuilder() automatically selects FROM AppBundle:Product
// and aliases it to "p"
// createQueryBuilder() 自动从 AppBundle:Product 进行 select 并赋予 p 假名
$query = $repository->createQueryBuilder('p')
    ->where('p.price > :price')
    ->setParameter('price', '19.99')
    ->orderBy('p.price', 'ASC')
    ->getQuery();
 
$products = $query->getResult();
// to get just one result: / 要得到一个结果：
// $product = $query->setMaxResults(1)->getOneOrNullResult();
```
--------------------------------------------------------------------------

