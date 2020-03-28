import React, { Fragment } from "react";
import {
  Table,
  Input,
  Button,
  Row,
  Col,
  Select,
  Divider,
  Icon,
  DatePicker,
  Tag,
  message
} from "antd";
import moment from "moment";
import { TaskStatus, ViewTask } from "../../models/tasks";

import { callbackWhenPressEnter } from '../../events/key'

import {AddTask} from '../../api/tasks'

interface HomeState {
  tasks: ViewTask[];
  editingTask: ViewTask;
}

interface HomeProps {}

class Home extends React.Component<HomeProps, HomeState> {
  state = {
    tasks: [] as ViewTask[],
    editingTask: {} as ViewTask,
  };
  // rowSelection object indicates the need for row selection
  rowSelection = {
    onChange: (selectedRowKeys: any[], selectedRows: ViewTask[]) => {
      console.log(
        `selectedRowKeys: ${selectedRowKeys}`,
        "selectedRows: ",
        selectedRows
      );
    },
    getCheckboxProps: (record: ViewTask) => ({
      disabled: record.name === "Disabled User", // Column configuration not to be checked
      name: record.name
    })
  };

  columns = [
    {
      width: 300,
      title: "Name",
      key: "name",
      render: (record: ViewTask) => {
        return (
          <div>
            {record.editing ? (
              <Input
                onPressEnter={this.handleAddTask(record)}
                onBlur={this.handleAddTask(record)}
                placeholder={"Input a name"}
                autoFocus={true}
                defaultValue={record.name}
                onChange={this.handleChangeEditingTaskName}
              />
            ) : (
              <Button type="dashed" block>
                {record.name} <Icon type="edit" />
              </Button>
            )}
          </div>
        );
      }
    },
    {
      width: 100,
      title: "Status",
      key: "status",
      dataIndex: "status",
      render: (status: TaskStatus) => (
        <Select defaultValue={status}>
          <Select.Option value={status}>{TaskStatus[status]}</Select.Option>
        </Select>
      )
    },
    {
      width: 200,
      title: "Plan",
      key: "plan",
      render: (record: ViewTask) => (
        <div>
          <Select defaultValue="week1"></Select>
          <Select defaultValue="Sunday"></Select>
        </div>
      )
    },
    {
      width: 200,
      title: "Due Time",
      key: "due_time",
      dataIndex: "due_time",
      render: (time: Date) => <DatePicker defaultValue={moment(time)} />
    },
    {
      width: 100,
      title: "Project",
      key: "project",
      dataIndex: "project",
      render: (plan: string) => <Tag color="cyan">{plan}</Tag>
    }
  ];

  handleChangeEditingTaskName = (e: any) => {
    let { editingTask } = this.state
    editingTask.name = e.target.value
    this.setState({ editingTask })
  }

  handleAddTask = (task: ViewTask) => () => {
    AddTask(task, (ok: boolean) => {
      if (ok) {
        let { editingTask, tasks } = this.state
        editingTask.editing = false
        tasks = [
          editingTask,
          ...tasks,
        ]
        this.setState({ editingTask, tasks })
      }
    })
  };

  handleAddEditingTask = () => {
    const { editingTask, tasks } = this.state;
    if (editingTask.editing) {
      message.error("there's still a new task editing")
      return
    }
    let tomorrow = new Date();
    tomorrow.setDate(new Date().getDate() + 1);
    const newTask = {
      name: ``,
      status: TaskStatus.TODO,
      due_time: tomorrow,
      project: "",
      editing: true
    };
    this.setState({
      tasks: [...tasks],
      editingTask: newTask,
    });
  };

  render() {
    let { editingTask, tasks } = this.state
    if (editingTask.editing) {
      tasks = [
        editingTask,
        ...tasks
      ]
    }
    tasks = [
      ...tasks,
      // mock:
      {
        name: "task1",
        status: TaskStatus.TODO,
        due_time: new Date(),
        project: "pj1"
      }
    ];
    return (
      <Fragment>
        <Row>
          <Col span={4}>
            {"Plan: "}
            <Select defaultValue="all" style={{ width: 90 }}>
              <Select.Option key="all">all</Select.Option>
              <Select.Option key="none">none</Select.Option>
              <Select.Option key="day">day</Select.Option>
              <Select.Option key="wk">week</Select.Option>
              <Select.Option key="mon">month</Select.Option>
            </Select>
          </Col>
          <Col span={6}>
            {"Date: "}
            <DatePicker />{" "}
            <Button shape="circle">
              <Icon type="left" />
            </Button>
            <Button shape="circle">
              <Icon type="right" />
            </Button>
          </Col>
          <Col span={1} />
          <Col span={4}>
            <Input.Search />
          </Col>
          <Col span={4}>
            <Button
              type="primary"
              shape="round"
              onClick={this.handleAddEditingTask}
            >
              <Icon type="plus" />
              Add
            </Button>
          </Col>
        </Row>
        <Row>
          <Divider
            orientation="left"
            style={{ color: "#333", fontWeight: "normal" }}
          />
        </Row>
        <Table
          rowKey="name"
          rowSelection={this.rowSelection}
          columns={this.columns}
          dataSource={tasks}
        />
      </Fragment>
    );
  }
}
export default Home;
