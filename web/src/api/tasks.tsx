import { Task, ViewTask } from '../models/tasks'
import { axiosPost } from './axios'

const API_PREFIX = "/api/v1"

export async function AddTask(task: ViewTask, fn: (success: boolean)=>(void)) {
  let data = viewTaskToTask(task)
  let res = await axiosPost(`${API_PREFIX}/tasks`, data)
  fn(res)
}

function dateToString(date: Date): string {
  return date.toISOString()
} 

function viewTaskToTask(task: ViewTask): Task {
  let ret = {
    name: task.name,
    status: task.status,
    due_time: dateToString(task.due_time),
    project: task.project
  }
  return ret
}