import Cookies from "js-cookie";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { FaEye, FaEyeSlash } from "react-icons/fa"; // Import icons

export default function Profile() {
  const isAuthenticated = !!Cookies.get("isAuthenticated");

  const navigate = useNavigate();
  useEffect(() => {
    if (!isAuthenticated) navigate("/signin");
  }, []);
  return <ProfilePage />;
}

const ProfilePage: React.FC = () => {
  // State to toggle the visibility of "سهم کمک‌ها"
  const [showContributions, setShowContributions] = useState(false);
  const [contributions, _setContributions] = useState(50000);

  // Handle file input for resume upload
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files ? e.target.files[0] : null;
    if (file) {
      console.log("Uploaded file:", file.name);
    }
  };

  return (
    <div className="pb-16 flex flex-col items-center p-8">
      {/* Profile Header */}
      <div className="w-full max-w-md bg-white shadow-md rounded-md p-6">
        {/* Avatar and Username */}
        <div className="flex items-center justify-between mb-4">
          <div>
            <h2 className="text-lg font-semibold">نام کاربری: شهرت من</h2>
            <p className="text-sm text-gray-500">شماره تماس: ۰۹۱۲۵۸۳۴۵۶۲</p>
          </div>
          <div className="rounded-full h-16 w-16 bg-gray-300 flex items-center justify-center">
            {/* Placeholder for avatar */}
            <span>آواتار</span>
          </div>
        </div>

        {/* Stats */}
        <div className="bg-gray-50 p-4 rounded-md mb-4">
          <p className="text-sm font-semibold">سابقه‌ی من: 4.8/5</p>
          <p className="text-sm font-semibold">تراکنش‌های من: 127</p>
          <p className="text-sm font-semibold">درخواست‌های من: 15</p>

          {/* سهم کمک‌ها */}
          <div className="flex items-center">
            <span className="text-sm font-semibold">سهم کمک‌های من: </span>
            <div className="mr-2 px-2 py-1 bg-gray-200 rounded-md text-xs flex items-center">
              {/* Conditionally show/hide contributions value */}
              <span>{showContributions ? `${contributions}T` : "•••••••"}</span>
              <button
                className="mr-1 w-4 h-4 text-gray-500 focus:outline-none"
                onClick={() => setShowContributions(!showContributions)}
              >
                {/* Eye/eye-off icon */}
                {showContributions ? <FaEyeSlash /> : <FaEye />}
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* Resume Upload Section */}
      <div className="mt-8 w-full max-w-md">
        <h3 className="text-lg font-semibold mb-4">رزومه‌ی من</h3>
        <div className="bg-white shadow-md rounded-md p-6">
          {/* Dashed border area for file selection */}
          <div
            className="border-2 border-dashed border-gray-400 p-4 rounded-md cursor-pointer text-center"
            onClick={() => document.getElementById("file-input")?.click()}
          >
            <p className="text-gray-600">فایل رزومه خود را اینجا انتخاب کنید</p>
            <p className="text-sm text-gray-400">(فقط PDF و DOC مجاز هستند)</p>
          </div>

          {/* Hidden file input */}
          <input
            type="file"
            id="file-input"
            accept=".pdf,.doc,.docx"
            className="hidden"
            onChange={handleFileChange}
          />

          <button className="w-full bg-gray-500 text-white py-2 rounded-md mt-4">
            آپلود
          </button>
        </div>
      </div>
    </div>
  );
};
