import React, { useState } from "react";
import { Carousel } from "../../components";

const imageUrls = [
  "https://d1csarkz8obe9u.cloudfront.net/posterpreviews/charity-template-banner-design-326a2d18518aadf7713faf82710e7d03_screen.jpg?ts=1662626212",
  "https://d1csarkz8obe9u.cloudfront.net/posterpreviews/charity-template-banner-design-326a2d18518aadf7713faf82710e7d03_screen.jpg?ts=1662626212",
  "https://d1csarkz8obe9u.cloudfront.net/posterpreviews/charity-template-banner-design-326a2d18518aadf7713faf82710e7d03_screen.jpg?ts=1662626212",
];

const RallyDetail: React.FC = () => {
  const [phoneNumber, setPhoneNumber] = useState("");
  const [donationAmount, setDonationAmount] = useState<number | undefined>();
  const [totalRaised, setTotalRaised] = useState<number>(500000); // Example total raised
  const goalAmount = 1000000; // Example goal amount

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file && ["application/pdf", "application/msword"].includes(file.type)) {
      console.log("File uploaded:", file.name);
    } else {
      alert("Please upload a valid PDF or DOC file.");
    }
  };

  const handlePhoneChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const phone = event.target.value;
    setPhoneNumber(phone);
  };

  const handleDonationChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const donation = Number(event.target.value);
    setDonationAmount(donation);
  };

  let projectedPercentage;
  const progressPercentage = Math.min((totalRaised / goalAmount) * 100, 100);
  if (donationAmount) {
    projectedPercentage = Math.min(
      ((totalRaised + donationAmount) / goalAmount) * 100,
      100
    );
  } else {
    projectedPercentage = 0;
  }

  return (
    <div className="container mx-auto p-4">
      {/* Title */}
      <h1 className="text-center text-4xl m-6 ">نام فراخوان</h1>

      {/* Slideshow */}
      <Carousel height="400px">
        {imageUrls.map((image, index) => (
          <div key={index} className="h-full">
            <img
              src={image}
              alt={`Slide ${index + 1}`}
              className="object-cover w-full h-full rounded-lg"
              style={{ height: "400px" }}
            />
          </div>
        ))}
      </Carousel>

      {/* Description Box */}
      <div className="border border-gray-300 p-4 rounded-lg my-12">
        <h2 className="text-xl mb-2">توضیحات فراخوان</h2>
        <p className="text-gray-700">
          این خیریه برای کمک به اهداف خاصی راه‌اندازی شده است. شما می‌توانید با
          کمک‌های نقدی و یا داوطلبانه در این امر سهیم باشید. برای اطلاعات بیشتر
          می‌توانید به شماره تماس‌های موجود مراجعه کنید.
        </p>
      </div>

      <div className="flex gap-6">
        <div className="my-6 border border-gray-300 p-4 rounded-lg w-[50%]">
          <h2 className="text-xl mb-2">شرایط داوطلبی</h2>
          <p className="text-gray-700">
            افرادی که مایل به شرکت در فعالیت‌های داوطلبانه هستند، می‌بایست رزومه
            خود را به همراه شماره تماس معتبر ارسال نمایند. همچنین شما می‌توانید
            بر اساس تجربه و توانایی‌های خود در بخش‌های مختلف کمک کنید.
          </p>
        </div>

        {/* Resume Upload */}
        <div className="w-[50%] my-6 border border-dashed border-gray-400 p-4 rounded-lg">
          <label className="block mb-2">بارگذاری رزومه:</label>
          <input
            type="tel"
            className="block w-full p-2 mb-4 border border-gray-400"
            placeholder="شماره تماس"
            value={phoneNumber}
            onChange={handlePhoneChange}
            pattern="[0-9]{11}"
            required
          />
          <div className="bg-white shadow-md rounded-md p-6">
            {/* Dashed border area for file selection */}
            <div
              className="border-2 border-dashed border-gray-400 p-4 rounded-md cursor-pointer text-center"
              onClick={() => document.getElementById("file-input")?.click()}
            >
              <p className="text-gray-600">
                فایل رزومه خود را اینجا انتخاب کنید
              </p>
              <p className="text-sm text-gray-400">
                (فقط PDF و DOC مجاز هستند)
              </p>
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
              داوطلب می‌شوم
            </button>
          </div>
        </div>
      </div>

      {/* Donation Progress Bar */}
      <div className="border border-gray-300 rounded-lg my-6 p-4">
        <label className="block mb-2">مبلغ اهدا:</label>
        <input
          type="number"
          className="block w-full p-2 mb-4 border border-gray-400"
          placeholder="مبلغ"
          value={donationAmount}
          onChange={handleDonationChange}
        />

        <div dir="ltr" className="h-4 bg-gray-200 rounded-full mb-2 relative">
          {/* Total raised portion */}
          <div
            className="bg-yellow-600 h-full rounded-full absolute"
            style={{ width: `${progressPercentage}%`, zIndex: 1 }}
          />
          {/* User donation portion */}
          <div
            className="bg-yellow-300 h-full rounded-full absolute"
            style={{
              width: `${projectedPercentage}%`,
              zIndex: 0,
            }}
          />
        </div>

        <p className="text-gray-700">
          تا کنون: {totalRaised.toLocaleString()} تومان از{" "}
          {goalAmount.toLocaleString()}
        </p>
        <div className="mt-6 flex justify-center">
          <button className="bg-green-500 text-white p-2 rounded">
            پرداخت
          </button>
        </div>
      </div>

      {/* Conditions of Volunteering */}

      {/* Payment Button */}
    </div>
  );
};

export default RallyDetail;
