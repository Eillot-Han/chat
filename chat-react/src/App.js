import React, { Component } from "react";
import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./views/homepage";
import ChatPage from "./views/chatpage";
import ChatHis from "./views/chathis/index";

class App extends Component {
  componentWillMount(){
    localStorage.setItem("roomid", "")
  }
  render() {
    return (
      <div id="content">
      <BrowserRouter>
        <div id="content_right">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/home" element={<Home />} />
            <Route path="/chatroom/:id" element={<ChatPage />} />
            <Route path="/chathistory/:id" element={<ChatHis />} />
          </Routes>
        </div>
      </BrowserRouter>
    </div> 
    );
  }
}

export default App;

