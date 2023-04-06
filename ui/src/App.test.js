import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders Multibox Profile header", () => {
  render(<App />);
  const linkElement = screen.getByText(/Running Profiles/i);
  expect(linkElement).toBeInTheDocument();
});
