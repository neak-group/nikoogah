import React from "react";
import { FaGithub, FaTwitter, FaLinkedin } from "react-icons/fa";

const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-900 text-white py-8">
      <div className="container mx-auto px-4 flex flex-col md:flex-row justify-between items-center">
        {/* Left column: Project summary */}
        <div className="mb-6 md:mb-0 md:w-1/2 text-right">
          <h2 className="text-lg font-semibold">عنوان خلاصه</h2>
          <p className="text-gray-400 mt-2">
            توضیحات خلاصه خلاصه توضیحات خلاصه خلاصه توضیحات خلاصه خلاصه توضیحات
            خلاصه خلاصه توضیحات خلاصه خلاصه توضیحات خلاصه خلاصه توضیحات خلاصه
            خلاصه توضیحات خلاصه خلاصه توضیحات خلاصه خلاصه توضیحات خلاصه خلاصه
          </p>
        </div>

        {/* Right column: Social media icons */}
        <div className="md:w-1/2 flex justify-center md:justify-end">
          <a
            href="https://github.com/your-profile"
            target="_blank"
            rel="noopener noreferrer"
            className="text-gray-400 hover:text-white mx-1"
          >
            <FaGithub size={24} />
          </a>
          <a
            href="https://twitter.com/your-profile"
            target="_blank"
            rel="noopener noreferrer"
            className="text-gray-400 hover:text-white mx-1"
          >
            <FaTwitter size={24} />
          </a>
          <a
            href="https://linkedin.com/in/your-profile"
            target="_blank"
            rel="noopener noreferrer"
            className="text-gray-400 hover:text-white mx-1"
          >
            <FaLinkedin size={24} />
          </a>
        </div>
      </div>
    </footer>
  );
};

export { Footer };
