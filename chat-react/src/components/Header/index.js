import React from "react";
import { Component } from "react";
import "./index.css";
import event from "../event";

class Header extends Component {
  logout=()=>{
    event.emit("eventMsg", "")
    localStorage.setItem("uid", "")
  }
  render() {
    return (
      <div className="header">
        <h2>Chat App</h2>
        <button className="logout" onClick={()=>this.logout()}>
          Logout
        </button>
      </div>
    );
  }
}

export default Header;
