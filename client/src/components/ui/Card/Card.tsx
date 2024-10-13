import React, { ReactNode } from "react";

interface CardProps {
  imageUrl: string;
  altText: string;
  children: ReactNode;
}

const Card: React.FC<CardProps> = ({ imageUrl, altText, children }) => {
  return (
    <div className="max-w-sm rounded-lg overflow-hidden shadow-lg bg-white hover:shadow-xl transition-shadow duration-300 text-right">
      <img className="w-full h-48 object-cover" src={imageUrl} alt={altText} />
      <div className="px-6 py-4">{children}</div>
    </div>
  );
};

export { Card };
