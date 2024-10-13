import React from "react";
import Router from "./router/Router";
import { AppProvider } from "./components";

function App() {
  return (
    <AppProvider>
      <Router />
    </AppProvider>
  );
}

export default App;
