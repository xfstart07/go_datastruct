# 栈

```
ADT 栈（stack）

Data

Operation：

InitStack(*S) 初始化
DestroyStack(*S) 若栈存在，则销毁它
ClearStack(*S) 将栈清空
StackEmpty(S) 判断栈为空
GetTop(S, *e) 返回栈顶元素
Push(*S, e) 入栈
Pop(*S, *e) 出栈
StackLength(S) 返回栈 S 的元素个数
```

### 例子

* [ch4stack/sqstack] 栈的顺序实现
* [ch4stack/stack_link] 栈的链式实现
* [ch4stack/fbi] 斐波那契数列
* [ch4stack/suffix_expression] 后缀表达式
* [ch4stack/infix2suffix] 中缀转后缀表达式