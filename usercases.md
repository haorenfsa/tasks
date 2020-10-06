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


# 数据模型
``` typescript
// use online tool to change code into golang: https://app.quicktype.io/ , then change float64 to int64

// task
export interface Task {
  id: number
  name: string
  status: TaskStatus
  position: number
  startAfter: LeveledTime
  endBefore: LeveledTime
  createdAt: number
  updatedAt: number
  // description: string
  // planId: number
  // subTasks: TaskEntry[]
  // parentTask: TaskEntry
}

export interface TaskEntry {
  id: number
  name: string
}

export enum TaskStatus {
  TODO,
  Doing,
  Done,
  Pending,
  Closed
}

// time
export interface LeveledTime {
  level: TimeLevel
  time: number
}

export enum TimeLevel {
  None,
  Year,
  Month,
  Week,
  Day,
}
// plan
// export interface Plan {
//   id: number
//   name: string
//   level: TimeLevel
//   parentPlanId: number
//   subPlanIds: number[]
// }
```

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

