import React from "react";
import { Icon } from "@mdi/react";
import { mdiRefresh, mdiStop } from "@mdi/js";
import { Button } from "react-bulma-components";
import "./StatusBar.css";

// todo: update status bar to reflect the correct current status
const StatusBar = ({ sendMsg }) => {
  return (
    <div className="statusbar">
      <div className="statusbar-content">
        <div className="statusbar-button-group">
          <Button size="small" color="dark">
            <Icon
              path={mdiRefresh}
              title="Refresh Status"
              size={"12px"}
              color="black"
              onClick={() => {
                sendMsg({
                  type: "engine",
                  cmd: "status",
                });
              }}
            />
          </Button>
        </div>
        <div className="statusbar-text">Status: Running</div>
        <div className="statusbar-button-group">
          <Button size="small" color="dark">
            <Icon
              path={mdiStop}
              title="Stop Engine"
              size={"12px"}
              color="black"
              onClick={() => {
                sendMsg({
                  type: "engine",
                  cmd: "stop",
                });
              }}
            />
          </Button>
        </div>
      </div>
    </div>
  );
};

export default StatusBar;
