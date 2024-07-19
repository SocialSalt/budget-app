import React from "react";
import "./App.css";
import TabBox from "./components/TabBox";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <div className="App">
        <div>
          <TabBox
            numTabs={3}
            tabNames={["overview", "budget", "transactions", "categories"]}
          />
        </div>
      </div>
    </ThemeProvider>
  );
}

export default App;
