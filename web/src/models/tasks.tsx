// common
export enum TaskStatus {
  TODO = 0,
  Doing,
  Done,
  Pending,
  Closed
}

// common
export interface Task {
  name: string;
  status: TaskStatus;
  due_time: string;
  project: string;
}

// view
export interface ViewTask {
  name: string;
  status: TaskStatus;
  due_time: Date;
  project: string;
  editing?: boolean;
}