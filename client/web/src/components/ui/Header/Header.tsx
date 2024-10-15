import Cookies from "js-cookie";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

interface NavbarProps {
  isAuthenticated: boolean;
  username?: string;
}

const Navbar: React.FC<NavbarProps> = ({ isAuthenticated, username }) => {
  const [dropdownOpen, setDropdownOpen] = useState(false);
  const navigate = useNavigate();
  const toggleDropdown = () => {
    setDropdownOpen(!dropdownOpen);
  };

  return (
    <nav className="bg-gray-800 text-white px-12 py-4 shadow-lg">
      <div className="flex justify-between items-center w-full">
        {/* Left: Navigation Links */}
        <div className="flex gap-8 text-xl">
          <Link to="/" className="hover:text-gray-300 transition duration-300">
            خانه
          </Link>
          {/* <Link
            to="/register-charity"
            className="hover:text-gray-300 transition duration-300"
          >
            ثبت نام خیریه
          </Link> */}
          <Link
            to="/my-charity"
            className="hover:text-gray-300 transition duration-300"
          >
            خیریه من
          </Link>
        </div>

        {/* Right: Username Dropdown or Login/Register */}
        <div className="relative">
          {isAuthenticated ? (
            <div className="relative">
              <button
                onClick={toggleDropdown}
                className="focus:outline-none text-white hover:text-gray-300 transition duration-300"
              >
                {username} &#9662;
              </button>
              {dropdownOpen && (
                <div
                  className="absolute mt-2 w-48 bg-white text-gray-900 rounded-md shadow-lg py-2 z-50"
                  style={{
                    maxHeight: "200px",
                    overflowY: "auto",
                    left: "0", // Align the left side of the dropdown with the button
                  }}
                >
                  <Link
                    to="/profile"
                    className="block px-4 py-2 hover:bg-gray-200 transition duration-300"
                    onClick={() => setDropdownOpen(false)}
                  >
                    پروفایل
                  </Link>
                  <button
                    className="w-full  text-right block px-4 py-2 hover:bg-gray-200 transition duration-300"
                    onClick={() => {
                      Cookies.remove("isAuthenticated");
                      window.location.reload();
                    }}
                  >
                    خروج
                  </button>
                </div>
              )}
            </div>
          ) : (
            <div className="flex">
              <Link
                to="/signin"
                className="hover:text-gray-300 transition duration-300"
              >
                ورود
              </Link>
              <div className="mx-4">|</div>
              <Link
                to="/signup"
                className="hover:text-gray-300 transition duration-300"
              >
                ثبت نام
              </Link>
            </div>
          )}
        </div>
      </div>
    </nav>
  );
};

export { Navbar };
