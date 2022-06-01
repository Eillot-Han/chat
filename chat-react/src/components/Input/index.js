import React, { Component } from "react";
import "./index.css";

class Input extends Component {
  render() {
    return (
      <div className="ChatInput">
        <input onKeyDown={this.props.send} />
      </div>
    );
  }
}

export default Input;
