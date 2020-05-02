# change order in table

把 position 列 用来作为排序变量，新插入的数据，position默认值应该是之前的最大值+1 ，在变化时还要保持 position 的连续

eg: 一组 task id为(5,4,3,2,1) 他们的默认 $position 与 $id 相同。

eg1: 把 task.3 放到5的位置

需要的操作为：

``` SQL
START TRANSACTION;
SELECT position FROM task WHRERE id=3; 
-- 在 code 中检查 3.position 没变化
UPDATE task SET position = position - 1 WHERE position > 3 AND position <= 5;
UPDATE task SET position = 5 WHERE id=3;
COMMIT;
```


把 2 放到 4 的位置：

``` SQL
-- 已知 id 2 的 position 为 2, 需要把 position 变为 4
BEGIN;
UPDATE task SET position = position - 1 WHERE position > 2 AND position <= 4;
UPDATE task SET position = 4 WHERE id=2;
COMMIT;
```


 
