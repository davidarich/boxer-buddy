import React from "react";
import { Container, Columns, Menu, Hero } from "react-bulma-components";

const Settings = () => {
  return (
    <Hero>
      <Hero.Header>Settings</Hero.Header>
      <Hero.Body>
        <Container>
          <Columns>
            <Columns.Column size={3}>
              <Menu>
                <Menu.List title="Multiboxing Config">
                  <Menu.List.Item>Clients</Menu.List.Item>
                  <Menu.List.Item>Groups</Menu.List.Item>
                  <Menu.List.Item>Layouts</Menu.List.Item>
                  <Menu.List.Item>Hotkeys</Menu.List.Item>
                </Menu.List>
                <Menu.List title="Interface">
                  <Menu.List.Item>Size</Menu.List.Item>
                  <Menu.List.Item>Theme</Menu.List.Item>
                  <Menu.List.Item>Behavior</Menu.List.Item>
                </Menu.List>
              </Menu>
            </Columns.Column>
            <Columns.Column size={8} backgroundColor="dark"></Columns.Column>
          </Columns>
        </Container>
      </Hero.Body>
      <Hero.Footer></Hero.Footer>
    </Hero>
  );
};

export default Settings;
