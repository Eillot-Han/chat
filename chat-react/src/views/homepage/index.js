import React, { Component } from "react";
import Header from "../../components/Header";
import Login from "../../components/loginbox";
import event from "../../components/event";
import ChatList from "../../components/ChatList";
import request from "../../api/request";
import "./index.css";

class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      account: localStorage.getItem("uid"),
      password: "",
      group: true,
      groupList: [],
      friendList: [],
    };
  }

  componentDidMount() {
    event.addListener("eventMsg", (val) => {
      this.setState({
        account: val,
      });
    });
    if (this.state.account) {
      request
        .get("/group/userPartGroup", {
          account: this.state.account,
        })
        .then((res) => {
          this.setState({ groupList: res.data });
        });
      request
        .get("/relationship/allFriend", {
          account: this.state.account,
        })
        .then((res) => {
          this.setState({ friendList: res.data });
        });
    }
  }
  componentWillUnmount() {
    event.removeListener("eventMsg", () => console.log("closed"));
  }
  create() {
    request
      .post(
        "/group/createGroup",
        {
          account: this.state.account,
        },
        "form"
      )
      .then((res) => {
        if (res.code === 200) {
          alert("添加成功");
          request
            .get("/group/userPartGroup", {
              account: this.state.account,
            })
            .then((res) => {
              this.setState({ groupList: res.data });
            });
        }
      });
  }
  render() {
    return (
      <div>
        <Header />
        {this.state.account ? (
          <div>
            <div className="button-group">
              <button
                onClick={() => this.setState({ group: false })}
                style={{ background: this.state.group ? "#b1c7f1" : "#405b8d" }}
              >
                friends
              </button>
              <button
                onClick={() => this.setState({ group: true })}
                style={{ background: !this.state.group ? "#b1c7f1" : "#405b8d" }}
              >
                groups
              </button>
            </div>
            {this.state.group ? (
              <div>
                <div className="button-group">
                  <button onClick={() => this.create()}>create group</button>
                </div>
                {this.state.groupList.map((val) => (
                  <ChatList name={val.name} id={val.account} group={true} />
                ))}
              </div>
            ) : (
              <div>
                {this.state.friendList.map((val) => (
                  <ChatList name={val.big_id} id={val.ship_id} group={false} />
                ))}
              </div>
            )}
          </div>
        ) : (
          <Login />
        )}
      </div>
    );
  }
}

export default Home;
