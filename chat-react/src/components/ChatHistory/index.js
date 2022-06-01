import React, { Component } from "react";
import "./index.css";
import Message from "../Message";
import { Link } from "react-router-dom";

class ChatHistory extends Component {
  constructor(props){
    super(props)
    this.state = {
      name: "Chat History"
    }
  }
  componentDidMount() {
    this.setState({
      name : localStorage.getItem("roomid")
    })
  }
  render() {
    const messages = this.props.chatHistory.map((msg, index) => (
      <Message key={index} message={msg.data} />
    ));

    return (
      <div className="ChatHistory">
        <div><h2>{this.state.name}</h2><Link to={"/chathistory/" + this.state.name}>All History</Link></div>
        {messages}
      </div>
    );
  }
}

export default ChatHistory;
