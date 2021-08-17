## 单元测试Android

1. 步骤
setup->执行->验证结果
2. 目的
优化代码结构，验证
3. 要求
容易实现，覆盖核心语句
4. Junit4
![执行顺序](./pictures/Junit4执行顺序.png)
assert结果
5. Mockito
mock对象，打桩，验证行为
mock空方法返回默认值，对外部较多依赖，关心少数函数实现
spy非空可以打桩，对外部依赖少
验证结果：verify结果，次数，模糊结果
不能发文private，final类，static方法

6. PowerMock
增强mockito
private,static方法，final类，whitebox
和mockito版本一致
@Runwith
@PrepareForTest

7. Robolectric
JVM上的android APi实现
和powermock冲突
安装方法

8. 可测代码设计
测试简短，少使用mock，可维护性好
一个函数只做一件事
查询与行为分离，方便测试查询
逻辑和UI拆开好测试
函数依赖尽量通过参数传递，方便传递不同参数测试
尽量不要把复杂结构作为参数传递给函数


9. 组织单侧代码
