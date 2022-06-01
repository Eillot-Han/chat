import React, { Component } from "react";
import { connect, sendMsg } from "../../api";
import Header from "../../components/Header";
import ChatHistory from "../../components/ChatHistory";
import Input from "../../components/Input";
import { Link } from "react-router-dom";
import { LeftOutlined } from "@ant-design/icons";
import request from "../../api/request";
import "./index.css";

class ChatPage extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send(event) {
    if (event.keyCode === 13) {
      var message = localStorage.getItem("uid") + ":" + event.target.value;
      sendMsg(message);
      if (localStorage.getItem("group") === true) {
        request
          .post(
            "/groupMessage/addMessage",
            {
              group_id: localStorage.getItem("roomid"),
              account: localStorage.getItem("uid"),
              type: 1,
              content: event.target.value,
            },
            "form"
          )
          .then((res) => {
            console.log(res);
          });
      } else {
        request
          .post(
            "/friendMessage/addMessage",
            {
              ship_id: localStorage.getItem("roomid"),
              account: localStorage.getItem("uid"),
              type: 1,
              content: event.target.value,
            },
            "form"
          )
          .then((res) => {
            console.log(res);
          });
      }
      event.target.value = "";
    }
  }

  componentWillMount() {
    this.setState((prevState) => ({
      chatHistory: [],
    }));
  }

  componentDidMount() {
    connect((msg) => {
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
    });
  }

  changeUrl = (e) => {
    const file = e.target.files[0];
    let formData = new FormData();
    formData.append("file", file);
    if (localStorage.getItem("group") === true) {
      console.log(formData);
      request
        .post("/groupMessage/addMessage", {
          group_id: localStorage.getItem("roomid"),
          account: localStorage.getItem("uid"),
          type: 0,
          upload_file: formData,
        })
        .then((res) => {
          console.log(res);
          if (res.code === 200) alert("删除成功,请刷新页面");
        });
    } else {
      request
        .post("/groupMessage/addMessage", {
          group_id: localStorage.getItem("roomid"),
          account: localStorage.getItem("uid"),
          type: 1,
          upload_file: formData,
        })
        .then((res) => {
          console.log(res);
          if (res.code === 200) alert("删除成功,请刷新页面");
        });
    }
  };

  render() {
    return (
      <div className="chat_page">
        <Header />
        <Link to={"/home"} className="link_button">
          <LeftOutlined color="black" />
          <p>Back</p>
        </Link>
        <ChatHistory chatHistory={this.state.chatHistory} />
        <div>
          <input className="fileinp" type="file" onChange={this.changeUrl} />
        </div>
        <Input send={this.send} />
      </div>
    );
  }
}

export default ChatPage;
