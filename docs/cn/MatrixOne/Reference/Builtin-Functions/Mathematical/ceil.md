# **CEIL()**

## **函数说明**

CEIL(X)函数返回不小于X的最小整数。

## **函数语法**

```
> CEIL(X)
```

## **参数释义**

|  参数   | 说明  |
|  ----  | ----  |
| X | 必要参数，可取任意数值数据类型 |

对int类的绝对数值类型，返回值也是相同的绝对数值类型。对浮点数勒说，返回值也是浮点数。

## **示例**

```sql
> drop table if exists t1;
> create table t1(a int ,b float);
> insert into t1 values(1,0.5);
> insert into t1 values(2,0.499);
> insert into t1 values(3,0.501);
> insert into t1 values(4,20.5);
> insert into t1 values(5,20.499);
> insert into t1 values(6,13.500);
> insert into t1 values(7,-0.500);
> insert into t1 values(8,-0.499);
> insert into t1 values(9,-0.501);
> insert into t1 values(10,-20.499);
> insert into t1 values(11,-20.500);
> insert into t1 values(12,-13.500);
> select a,ceil(b) from t1;
+------+----------+
| a    | ceil(b)  |
+------+----------+
|    1 |   1.0000 |
|    2 |   1.0000 |
|    3 |   1.0000 |
|    4 |  21.0000 |
|    5 |  21.0000 |
|    6 |  14.0000 |
|    7 |  -0.0000 |
|    8 |  -0.0000 |
|    9 |  -0.0000 |
|   10 | -20.0000 |
|   11 | -20.0000 |
|   12 | -13.0000 |
+------+----------+
> select sum(ceil(b)) from t1;
+--------------+
| sum(ceil(b)) |
+--------------+
|       6.0000 |
+--------------+

```

## **限制**

MatrixOne目前只支持在查询表的时候使用函数，不支持单独使用函数。