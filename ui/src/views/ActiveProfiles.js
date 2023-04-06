import { Button, Card } from "react-bulma-components";
import {
  mdiImageFilterCenterFocus,
  mdiClipboardText,
  mdiRestart,
  mdiStop,
  mdiPlay,
} from "@mdi/js";
import { Icon } from "@mdi/react";

const ActiveProfiles = (data) => {
  console.log(data);
  const buttonIconSize = "12px";

  const handleClickClientStart = (profileName) => {
    console.log("clicked client start");
    data.sendMsg({
      type: "game_client",
      cmd: "start",
      args: [profileName],
    });
  };
  const handleClickClientStop = (profileName) => {
    console.log("clicked client stop");
    data.sendMsg({
      type: "game_client",
      cmd: "stop",
      args: [profileName],
    });
  };
  const handleClickClientFocus = (profileName) => {
    console.log("clicked client focus");
    data.sendMsg({
      type: "game_client",
      cmd: "focus",
      args: [profileName],
    });
  };
  const handleClickClientCopyPass = (password) => {
    console.log("clicked copy password");
    if (password && password !== "") {
      navigator.clipboard.writeText(password);
      return;
    }
    console.log("password is empty, nothing to do");
  };
  const renderCards = (activeProfileState) => {
    return activeProfileState.map((activeProfile) => {
      // play/stop button toggle
      let playToggle = (
        <Button
          color="dark"
          size="small"
          onClick={() => {
            handleClickClientStart(activeProfile.GameProfile.Name);
          }}
        >
          <Icon
            path={mdiPlay}
            title="Start Client"
            size={buttonIconSize}
            color="black"
          />
        </Button>
      );
      if (activeProfile.GameClient.Process) {
        playToggle = (
          <Button
            color="black"
            size="small"
            onClick={() => {
              handleClickClientStop(activeProfile.GameProfile.Name);
            }}
          >
            <Icon
              path={mdiStop}
              title="Stop Client"
              size={buttonIconSize}
              color="darkred"
            />
          </Button>
        );
      }

      return (
        <div key={activeProfile.GameProfile.Name}>
          <Card
            style={{
              width: 350,
              margin: "auto",
              background: "#222",
              color: "#999",
            }}
          >
            <Card.Content>
              <h1>{activeProfile.GameProfile.Name}</h1>
              <br />
              <Button.Group>
                <Button
                  color="dark"
                  size="small"
                  onClick={() => {
                    handleClickClientFocus(activeProfile.GameProfile.Name);
                  }}
                >
                  <Icon
                    path={mdiImageFilterCenterFocus}
                    title="Focus Client"
                    size={buttonIconSize}
                    color="black"
                  />
                </Button>
                {playToggle}
                <Button color="dark" size="small" onClick={() => {}}>
                  <Icon
                    path={mdiRestart}
                    title="Restart Client"
                    size={buttonIconSize}
                    color="black"
                  />
                </Button>
                <Button
                  color="dark"
                  size="small"
                  onClick={() => {
                    handleClickClientCopyPass(
                      activeProfile.GameProfile.Password
                    );
                  }}
                >
                  <Icon
                    path={mdiClipboardText}
                    title="Copy Password"
                    size={buttonIconSize}
                    color="black"
                  />
                </Button>
              </Button.Group>
            </Card.Content>
          </Card>
          <br />
        </div>
      );
    });
  };

  return (
    <div key="Running" className="Running">
      <br />
      {renderCards(data.clientStates)}
      <br />
    </div>
  );
};

export default ActiveProfiles;
