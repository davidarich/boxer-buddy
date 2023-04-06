import "./SideBar.css";
import { Icon } from "@mdi/react";
import { mdiCog, mdiAccount } from "@mdi/js";
import { Button } from "react-bulma-components";

const SideBar = ({ sendMsg }) => {
  return (
    <div className="sidebar">
      <Button
        color="black"
        size="small"
        onClick={() => {
          window.location.href = "/settings";
        }}
      >
        <Icon path={mdiCog} title="Settings" size={"12px"} color="white" />
      </Button>
      <br />
      <Button.Group>
        <Button
          color="dark"
          size="small"
          onClick={() => {
            window.location.href = "/activeProfiles";
          }}
        >
          <Icon
            path={mdiAccount}
            title="Profiles"
            size={"12px"}
            color="white"
          />
        </Button>
      </Button.Group>
      <br />
    </div>
  );
};

export default SideBar;
