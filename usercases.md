# 任务管理系统需求：
- 有项目名、时间、是否杂事
- 按项目纬度添加
- 分配到周计划
- 分配到日计划
- 项目状态：todo,done,pending，特定时间执行

## 新需求
- 一个大任务可以添加多个子任务
- 子任务的完成度和完成个数反应给大任务的完成度
- 每个任务可以估算耗时
- 每个任务可以记录实际使用时间
- 每日、每周、每月、每年有一个任务视图

# detail
## task
#### done
- AddTask
    - add by name

- QueryAllTasks

- UpdateTask
    - update status
    - update name
    - update day plan
    - update week plan
    - change order in table
        - how: position column

- DeleteTask

#### TODO
- UpdateTask
    - add/update task description
    - update task group (devide life and work)


