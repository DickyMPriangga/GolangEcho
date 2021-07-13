import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8081";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: []
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    console.log("pRINTING task", this.state);
    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          {
            "task" : this.state.task
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getTask();
          this.setState({
            task: ""
          });
          console.log(res);
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then(res => {
      console.log(res);
      if (res.data) {
        this.setState({
          items: res.data.map(item => {
            let color = "yellow";

            if (item.status) {
              color = "green";
            }
            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="green"
                      onClick={() => this.updateTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Done</span>
                    <Icon
                      name="undo"
                      color="yellow"
                      onClick={() => this.undoTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Undo</span>
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => this.deleteTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };

  updateTask = id => {
    axios
      .put(endpoint + "/api/task/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  undoTask = id => {
    axios
      .put(endpoint + "/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };

  deleteTask = id => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        console.log(res);
        this.getTask();
      });
  };
  render() {
    return (
      /*
      <div>
        <Grid.Row>
          <Header className="header" as="h2">
            TO DO LIST
          </Header>
        </Grid.Row>
        <Grid columns={3}>
          <Grid.Row>
            <Grid.Column width={12}>
              <Form onSubmit={this.onSubmit}>
                <Input
                  type="text"
                  name="task"
                  onChange={this.onChange}
                  value={this.state.task}
                  fluid
                  placeholder="Create Task"
                />
              </Form>
            </Grid.Column>
            <Grid.Column width={2}>
              <Button type='submit'>Create Task</Button>
            </Grid.Column>
            <Grid.Column width={2}>
              <Button >Delete All</Button>
            </Grid.Column>
          </Grid.Row>
        </Grid>
        <Grid.Row>
          <Card.Group>{this.state.items}</Card.Group>
        </Grid.Row>
      </div>
      */
      
      <div>
        <div className="row">
          <Header className="header" as="h2">
            TO DO LIST
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              action="Create Task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
          </Form>
        </div>
        <br></br>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
      
    );
  }
}

export default ToDoList;