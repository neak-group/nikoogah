import React from "react";

interface InputProps {
  label?: string;
  placeholder?: string;
  type?: string;
  value?: string;
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
  errorMessage?: string;
  icon?: React.ReactNode; // Accepts an icon component
}

const Input: React.FC<InputProps> = ({
  label,
  placeholder,
  type = "text",
  value,
  onChange,
  errorMessage,
  icon,
}) => {
  return (
    <div className="mb-4">
      {label && (
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor={label}
        >
          {label}
        </label>
      )}
      <div className="relative">
        {icon && (
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            {icon}
          </div>
        )}
        <input
          id={label}
          type={type}
          value={value}
          onChange={onChange}
          placeholder={placeholder}
          className={`w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 ${
            errorMessage ? "border-red-500" : "border-gray-300"
          } ${icon ? "pl-10" : ""}`}
        />
      </div>
      {errorMessage && (
        <p className="text-red-500 text-xs italic mt-2">{errorMessage}</p>
      )}
    </div>
  );
};

export { Input };
