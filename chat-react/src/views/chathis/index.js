import React, { Component } from "react";
import Header from "../../components/Header";
import Input from "../../components/Input";
import { Link } from "react-router-dom";
import { LeftOutlined } from "@ant-design/icons";
import request from "../../api/request";
import "./index.css";

class ChatHis extends Component {
  constructor(props) {
    super(props);
    this.state = {
      list: [],
    };
  }

  componentDidMount() {
    if (localStorage.getItem("group") === "true") {
      request
        .get("/groupMessage/allMessage", {
          group_id: localStorage.getItem("roomid"),
        })
        .then((res) => {
          this.setState({ list: res.data });
        });
    } else {
      request
        .get("/friendMessage/allMessage", {
          ship_id: localStorage.getItem("roomid"),
        })
        .then((res) => {
          this.setState({ list: res.data });
        });
    }
  }

  inquire = (event) => {
    if (event.keyCode === 13) {
      if (localStorage.getItem("group") === "true") {
        request
          .get("/groupMessage/inquireMessage", {
            group_id: localStorage.getItem("roomid"),
            content: event.target.value,
          })
          .then((res) => {
            this.setState({ list: res.data });
          });
      } else {
        request
          .get("/friendMessage/inquireMessage", {
            ship_id: localStorage.getItem("roomid"),
            account: localStorage.getItem("uid"),
            type: 1,
            content: event.target.value,
          })
          .then((res) => {
            this.setState({ list: res.data });
          });
      }
      event.target.value = "";
    }
  };
  render() {
    console.log(this.state.list);
    return (
      <div className="chat_page">
        <Header />
        <Link to={"/home"} className="link_button">
          <LeftOutlined color="black" />
          <p>Back</p>
        </Link>
        <div className="ChatHistory">
          <div>
            <h2 style={{ marginBottom: "40px" }}>{localStorage.getItem("roomid")}</h2>
            <div className="HisInput">
              <input onKeyDown={this.inquire}></input>
            </div>
            {this.state.list.map((val) => (
              <div className="msg_box">
                {localStorage.getItem("group") === "true" ? (
                  <p className="id">{val.user_id}:</p>
                ) : (
                  <p className="id">{val.from_id}:</p>
                )}
                <p className="content">{val.content}</p>
                <p className="time">{val.send_time}</p>
              </div>
            ))}
          </div>
        </div>
      </div>
    );
  }
}

export default ChatHis;
