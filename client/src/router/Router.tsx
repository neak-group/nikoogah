// App.tsx
import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import PrivateRoute from "./PrivateRoute";
import { Home } from "../views";
import Cookies from "js-cookie";
import Profile from "../views/Profile/Profile";
import RegisterCharity from "../views/RegisterCharity/RegisterCharity";
import RallyDetail from "../views/RallyDetail/RallyDetail";
import MyCharity from "../views/MyCharity/MyCharity";
import NewRally from "../views/NewRally/NewRally";
import SignUp from "../views/SignUp/SignUp";
import Login from "../views/Login/Login";

const App: React.FC = () => {
  const isAuthenticated = !!Cookies.get("isAuthenticated");

  return (
    <Router>
      <Routes>
        {/* Public Routes */}
        {!isAuthenticated && <Route path="/signin" element={<Login />} />}
        {!isAuthenticated && <Route path="/signup" element={<SignUp />} />}

        {/* Private Routes */}
        <Route
          path="/"
          element={
            <PrivateRoute>
              <Home />
            </PrivateRoute>
          }
        />
        <Route
          path="/register-charity"
          element={
            <PrivateRoute>
              <RegisterCharity />
            </PrivateRoute>
          }
        />
        <Route
          path="/rallydetail/:id"
          element={
            <PrivateRoute>
              <RallyDetail />
            </PrivateRoute>
          }
        />
        <Route
          path="/my-charity"
          element={
            <PrivateRoute>
              <MyCharity />
            </PrivateRoute>
          }
        />
        <Route
          path="/new-rally"
          element={
            <PrivateRoute>
              <NewRally />
            </PrivateRoute>
          }
        />
        <Route
          path="/profile"
          element={
            <PrivateRoute>
              <Profile />
            </PrivateRoute>
          }
        />
        {/* 404 Handler */}
        <Route path="*" element={<h1>Not Found</h1>} />
      </Routes>
    </Router>
  );
};

export default App;
