// PrivateRoute.tsx
import React from "react";
// import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";
import { Navbar, Footer } from "../components";

// A wrapper for <Route> that redirects to the sign-in screen if not authenticated
const PrivateRoute: React.FC<{ children: JSX.Element }> = ({ children }) => {
  // Check if the user is authenticated by looking for a cookie (e.g., "access_token")
  const isAuthenticated = !!Cookies.get("isAuthenticated");

  return (
    <>
      <Navbar isAuthenticated={isAuthenticated} username="lklk" />
      {children}
      <Footer />
    </>
  );
};

export default PrivateRoute;
