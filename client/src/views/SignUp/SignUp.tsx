import React, { useState, useRef } from "react";
import { useFormik } from "formik";
import * as Yup from "yup";
import Cookies from "js-cookie";
import { useNavigate } from "react-router-dom";

// Validation schema for the first step
const validationSchema = Yup.object().shape({
  firstName: Yup.string().required("نام الزامی است"),
  lastName: Yup.string().required("نام خانوادگی الزامی است"),
  nationalId: Yup.string()
    .length(10, "کد ملی باید 10 رقم باشد")
    .required("کد ملی الزامی است"),
  phoneNumber: Yup.string().required("شماره تماس الزامی است"),
});

// OTP validation function
const validateOTP = (otp: string) => {
  return otp === "111111";
};

const SignUp = () => {
  const [step, setStep] = useState(1);
  const [otp, setOtp] = useState(["", "", "", "", "", ""]);
  const navigate = useNavigate();
  const inputRefs = useRef<(HTMLInputElement | null)[]>([]); // Ref to store input elements

  const formik = useFormik({
    initialValues: {
      firstName: "",
      lastName: "",
      nationalId: "",
      phoneNumber: "",
    },
    validationSchema,
    onSubmit: (values) => {
      console.log(values); // Handle form values
      // Move to the next step (OTP)
      setStep(2);
      // Simulate sending OTP and show the OTP in the UI
    },
  });

  const handleOtpChange = (index: number, value: string) => {
    const newOtp = [...otp];
    newOtp[index] = value;

    setOtp(newOtp);

    // Move to the next input
    if (value && index < otp.length - 1) {
      inputRefs.current[index + 1]?.focus();
    }

    // Move to the previous input on backspace
    if (!value && index > 0) {
      inputRefs.current[index - 1]?.focus();
    }
  };

  const handleOtpSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const otpInput = otp.join(""); // Join the OTP array into a string

    if (validateOTP(otpInput)) {
      // Set the authenticated cookie
      Cookies.set("isAuthenticated", "true");
      // Redirect to the home page
      navigate("/");
    } else {
      alert("کد تایید نامعتبر است");
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 border border-gray-300 rounded-lg shadow-lg bg-white">
      {step === 1 ? (
        <form onSubmit={formik.handleSubmit}>
          <h2 className="text-lg mb-4 text-right">فرم ورود</h2>
          <div className="mb-4">
            <label htmlFor="firstName" className="block mb-2 text-right">
              نام:
            </label>
            <input
              type="text"
              id="firstName"
              name="firstName"
              onChange={formik.handleChange}
              value={formik.values.firstName}
              className={`w-full border p-2 rounded ${
                formik.touched.firstName && formik.errors.firstName
                  ? "border-red-500"
                  : ""
              }`}
              placeholder="نام خود را وارد کنید"
            />
            {formik.touched.firstName && formik.errors.firstName && (
              <div className="text-red-500 text-sm mt-1">
                {formik.errors.firstName}
              </div>
            )}
          </div>

          <div className="mb-4">
            <label htmlFor="lastName" className="block mb-2 text-right">
              نام خانوادگی:
            </label>
            <input
              type="text"
              id="lastName"
              name="lastName"
              onChange={formik.handleChange}
              value={formik.values.lastName}
              className={`w-full border p-2 rounded ${
                formik.touched.lastName && formik.errors.lastName
                  ? "border-red-500"
                  : ""
              }`}
              placeholder="نام خانوادگی خود را وارد کنید"
            />
            {formik.touched.lastName && formik.errors.lastName && (
              <div className="text-red-500 text-sm mt-1">
                {formik.errors.lastName}
              </div>
            )}
          </div>

          <div className="mb-4">
            <label htmlFor="nationalId" className="block mb-2 text-right">
              کد ملی:
            </label>
            <input
              type="text"
              id="nationalId"
              name="nationalId"
              onChange={formik.handleChange}
              value={formik.values.nationalId}
              className={`w-full border p-2 rounded ${
                formik.touched.nationalId && formik.errors.nationalId
                  ? "border-red-500"
                  : ""
              }`}
              placeholder="کد ملی خود را وارد کنید"
            />
            {formik.touched.nationalId && formik.errors.nationalId && (
              <div className="text-red-500 text-sm mt-1">
                {formik.errors.nationalId}
              </div>
            )}
          </div>

          <div className="mb-4">
            <label htmlFor="phoneNumber" className="block mb-2 text-right">
              شماره تماس:
            </label>
            <input
              type="text"
              id="phoneNumber"
              name="phoneNumber"
              onChange={formik.handleChange}
              value={formik.values.phoneNumber}
              className={`w-full border p-2 rounded ${
                formik.touched.phoneNumber && formik.errors.phoneNumber
                  ? "border-red-500"
                  : ""
              }`}
              placeholder="شماره تماس خود را وارد کنید"
            />
            {formik.touched.phoneNumber && formik.errors.phoneNumber && (
              <div className="text-red-500 text-sm mt-1">
                {formik.errors.phoneNumber}
              </div>
            )}
          </div>

          <button
            type="submit"
            className="mt-4 w-full bg-blue-500 text-white p-2 rounded"
          >
            ارسال
          </button>
        </form>
      ) : (
        <div className="text-center flex flex-col items-center gap-4">
          <h2 className="text-lg mb-4 text-right">کد تایید</h2>
          <p className="mb-4 text-right">کد تایید خود ارسالی را وارد کنید:</p>
          <form
            dir="ltr"
            onSubmit={handleOtpSubmit}
            className="flex justify-center"
          >
            {otp.map((digit, index) => (
              <input
                key={index}
                type="text"
                maxLength={1}
                value={digit}
                onChange={(e) => handleOtpChange(index, e.target.value)}
                onKeyDown={(e) => {
                  // Handle backspace
                  if (e.key === "Backspace" && digit === "") {
                    inputRefs.current[index - 1]?.focus();
                  }
                }}
                ref={(el) => (inputRefs.current[index] = el)}
                className="border p-2 rounded mx-1 w-12 text-center"
                placeholder="-"
              />
            ))}
          </form>
          <button
            onClick={() => {
              Cookies.set("isAuthenticated", "true");
              navigate("/");
            }}
            type="submit"
            className="mt-4 w-full bg-green-500 text-white p-2 rounded"
          >
            تایید
          </button>
        </div>
      )}
    </div>
  );
};

export default SignUp;
