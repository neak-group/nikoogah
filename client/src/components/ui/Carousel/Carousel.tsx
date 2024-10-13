import React from "react";
import Slider from "react-slick"; // You need to install react-slick and slick-carousel

// Import the CSS for react-slick
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

interface ImageSliderProps {
  height?: string; // You can pass height as a prop
  width?: string; // You can pass width as a prop
  children?: React.ReactNode;
  settings?: any;
}

const Carousel: React.FC<ImageSliderProps> = ({
  height = "400px",
  children,
  width = "100%",
  settings,
}) => {
  const _settings = {
    dots: true,
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    ...settings,
  };

  return (
    <div className="mx-auto" style={{ width }}>
      <Slider {..._settings}>{children}</Slider>
    </div>
  );
};

export { Carousel };
