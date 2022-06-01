import React from "react";
import { Component } from "react";
import { Link } from "react-router-dom";
import request from "../../api/request";
import "./index.css";

class ChatList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      id: 123,
      group: true,
    };
  }
  componentDidMount() {
    this.setState({
      name: this.props.name,
      id: this.props.id,
      group: this.props.group,
    });
  }
  UNSAFE_componentWillReceiveProps(nextProps) {
    this.setState({
      name: nextProps.name,
      id: nextProps.id,
      group: nextProps.group,
    });
  }
  delete() {
    if (this.state.group) {
      request
        .post(
          "/group/deleteGroup",
          {
            account: localStorage.getItem("uid"),
            group_id: this.state.id,
          },
          "form"
        )
        .then((res) => {
          console.log(res);
          if (res.code === 200) alert("删除成功,请刷新页面");
        });
    } else {
      request
        .post(
          "/relationship/deleteFriend",
          {
            account: localStorage.getItem("uid"),
            friend_id: this.state.id,
          },
          "form"
        )
        .then((res) => {
          if (res.code === 200) alert("删除成功,请刷新页面");
        });
    }
  }
  render() {
    return (
      <div className="ChatList">
        <p className="chat_name">{this.state.name}</p>
        <div>
          <Link
            onClick={() => {
              localStorage.setItem("roomid", this.state.id);
              localStorage.setItem("group", this.state.group);
            }}
            to={"/chatroom/" + this.state.id}
          >
            <button className="chat_button">Chat</button>
          </Link>
          <button
            className="chat_button"
            onClick={() => {
              this.delete();
            }}
          >
            delete
          </button>
        </div>
      </div>
    );
  }
}

export default ChatList;
