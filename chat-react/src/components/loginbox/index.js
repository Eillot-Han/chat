import React, { Component } from "react";
import "./index.css";
import event from "../event";
import request from "../../api/request";

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      account: "",
      password: "",
    };
  }
  login = () => {
    console.log(this.state.account);
    request
      .get("/user/login", {
        account: this.state.account,
        password: this.state.password,
      })
      .then((res) => {
        console.log(res);
        if (res.code === 403) alert("密码错误");
        else {
          alert("登陆成功");
          event.emit("eventMsg", this.state.account);
          localStorage.setItem("uid", this.state.account);
        }
      });
  };

  render() {
    return (
      <div className="LoginInput">
        <p className="login">Login</p>
        <div className="inputbox">
          <p className="hint">Account</p>
          <input type="text" onChange={(e) => this.setState({ account: e.target.value })} />
        </div>
        <div className="inputbox">
          <p className="hint">Password</p>
          <input type="text" onChange={(e) => this.setState({ password: e.target.value })} />
        </div>
        <button onClick={() => this.login()}>Login Here</button>
      </div>
    );
  }
}

export default Login;
