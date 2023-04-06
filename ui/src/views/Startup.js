import React from "react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Form, Button } from "react-bulma-components";

const Startup = ({ sendMsg }) => {
  const [profile, setProfile] = useState("");
  const [profileOptions, setProfileOptions] = useState([]);

  useEffect(() => {
    (async () => {
      const multiboxProfileFormOpts = ["default"];
      setProfileOptions(multiboxProfileFormOpts);
    })();
  }, []);

  const navigate = useNavigate();
  const handleClickOpen = () => {
    const startMsg = {
      type: "engine",
      cmd: "start",
      args: ["default"],
    };
    sendMsg(startMsg);
    navigate("/activeProfiles");
  };

  return (
    <div className="Startup">
      <form>
        <Form.Field>
          <Form.Label>Multibox Profile</Form.Label>
          <Form.Field kind="group">
            <Form.Control>
              <Form.Select
                value={profile}
                onChange={(e) => {
                  return setProfile(e.target.value);
                }}
              >
                {profileOptions.map((opt) => (
                  <option key={opt} value={opt}>
                    {opt}
                  </option>
                ))}
              </Form.Select>
            </Form.Control>
          </Form.Field>
        </Form.Field>

        <Form.Field kind="group">
          <Form.Control>
            <Button color="black" onClick={handleClickOpen}>
              Open
            </Button>
          </Form.Control>
        </Form.Field>
      </form>
    </div>
  );
};
export default Startup;
