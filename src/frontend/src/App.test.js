import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders idk yet", () => {
  render(<App />);
  const linkElement = screen.getByText(/OVERVIEW/i);
  expect(linkElement).toBeInTheDocument();
});
