import "./App.css";
import "bulma/css/bulma.min.css";
import { Routes, Route } from "react-router-dom";
import Startup from "./views/Startup";
import ActiveProfiles from "./views/ActiveProfiles";
import Settings from "./views/Settings";
import StatusBar from "./components/StatusBar";
import { useRef, useEffect } from "react";
import { Container } from "react-bulma-components";
import { useLocalStorage } from "./hooks/state";
import SideBar from "./components/SideBar";

// setup interop connection string
const params = new URLSearchParams(window.location.search);
if (!params.has("interop_host") || !params.has("interop_port")) {
  console.log("interop host and/or port was not provided");
  localStorage.setItem("interop_addr", "ws://localhost:7778");
} else {
  localStorage.setItem(
    "interop_addr",
    "ws://" + params.get("interop_host") + ":" + params.get("interop_port")
  );
}

const App = () => {
  const [state, setState] = useLocalStorage("state", []);
  const ws = useRef(null);

  const sendMsg = (msg) => {
    if (!ws.current) {
      console.log("can't send message, websocket not connected");
      return;
    }
    console.log("sending message");
    console.log(msg);
    ws.current.send(JSON.stringify(msg));
    console.log("message sent");
  };
  useEffect(() => {
    // Connect if not connected
    if (!ws.current) {
      ws.current = new WebSocket(localStorage.getItem("interop_addr"));
    }

    // Handle initial connection event
    ws.current.onopen = (event) => {
      console.log("opened connection");
      console.log(event);
    };

    // Handle event pushed from backend
    ws.current.onmessage = (event) => {
      console.log("received message");
      console.log(event);

      let data = JSON.parse(event.data);
      console.log(data);

      // todo: handle data other than just activeProfiles
      if (data != null) {
        setState(data);
      }
    };

    return () => {};
  }, [state, setState]);

  return (
    <div className="App">
      <SideBar sendMsg={sendMsg}></SideBar>
      <Container className="main-container">
        <Routes>
          <Route path="/" element={<Startup sendMsg={sendMsg} />} />
          <Route
            path="/activeProfiles"
            element={<ActiveProfiles sendMsg={sendMsg} clientStates={state} />}
          />
          <Route path="/settings" element={<Settings />} />
        </Routes>
      </Container>
      <StatusBar></StatusBar>
    </div>
  );
};

export default App;
