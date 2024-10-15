/** @type {import('tailwindcss').Config} */
// tailwind.config.js
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}", // Include all relevant file types from the src directory
  ],
  theme: {
    extend: {
      fontFamily: {
        persian: ["IranYekan", "sans-serif"],
      },
    },
  },
  plugins: [require("tailwindcss-rtl")],
};
