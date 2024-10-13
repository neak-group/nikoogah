import React, { useState } from "react";
import { useFormik } from "formik";
import * as Yup from "yup";

// Validation schema using Yup
const validationSchema = Yup.object({
  name: Yup.string().required("نام ضروری است"),
  email: Yup.string().email("ایمیل معتبر نیست").required("ایمیل ضروری است"),
  address: Yup.string().required("آدرس ضروری است"),
  phone: Yup.string().required("شماره تلفن ضروری است"),
  nationalID: Yup.string().required("شماره ملی ضروری است"),
  managerName: Yup.string().required("نام مدیرعامل ضروری است"),
  position: Yup.string().required("سمت ضروری است"),
});

const RegisterCharity: React.FC = () => {
  const [logo, setLogo] = useState<string | ArrayBuffer | null>(null);

  const formik = useFormik({
    initialValues: {
      name: "",
      email: "",
      address: "",
      phone: "",
      nationalID: "",
      managerName: "",
      position: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      console.log("Form data:", values);
    },
    onReset: () => {
      setLogo(null); // Clear logo on form reset
    },
  });

  // Handle file upload
  const handleLogoUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = () => {
        setLogo(reader.result);
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 py-12">
      <form
        onSubmit={formik.handleSubmit}
        onReset={formik.handleReset}
        className="bg-white shadow-md p-8 rounded-md w-full max-w-lg"
      >
        <h2 className="text-xl font-semibold mb-4 text-center">
          ثبت نام خیریه
        </h2>

        {/* Logo Upload */}
        <div className="flex justify-center mb-4">
          <label className="cursor-pointer">
            <div className="w-24 h-24 bg-gray-200 rounded-full flex items-center justify-center overflow-hidden border">
              {logo ? (
                <img
                  src={logo as string}
                  alt="logo"
                  className="object-cover w-full h-full"
                />
              ) : (
                <span className="text-gray-500">لوگو</span>
              )}
            </div>
            <input
              type="file"
              accept="image/*"
              className="hidden"
              onChange={handleLogoUpload}
            />
          </label>
        </div>

        {/* Name */}
        <div className="mb-4">
          <label className="block text-sm mb-2">نام:</label>
          <input
            type="text"
            name="name"
            className={`w-full p-2 border ${
              formik.touched.name && formik.errors.name
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.name}
          />
          {formik.touched.name && formik.errors.name && (
            <p className="text-red-500 text-sm">{formik.errors.name}</p>
          )}
        </div>

        {/* Address */}
        <div className="mb-4">
          <label className="block text-sm mb-2">آدرس:</label>
          <input
            type="text"
            name="address"
            className={`w-full p-2 border ${
              formik.touched.address && formik.errors.address
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.address}
          />
          {formik.touched.address && formik.errors.address && (
            <p className="text-red-500 text-sm">{formik.errors.address}</p>
          )}
        </div>

        {/* Phone */}
        <div className="mb-4">
          <label className="block text-sm mb-2">شماره تلفن:</label>
          <input
            type="text"
            name="phone"
            className={`w-full p-2 border ${
              formik.touched.phone && formik.errors.phone
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.phone}
          />
          {formik.touched.phone && formik.errors.phone && (
            <p className="text-red-500 text-sm">{formik.errors.phone}</p>
          )}
        </div>

        {/* Email */}
        <div className="mb-4">
          <label className="block text-sm mb-2">ایمیل:</label>
          <input
            type="email"
            name="email"
            className={`w-full p-2 border ${
              formik.touched.email && formik.errors.email
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.email}
          />
          {formik.touched.email && formik.errors.email && (
            <p className="text-red-500 text-sm">{formik.errors.email}</p>
          )}
        </div>

        {/* National ID */}
        <div className="mb-4">
          <label className="block text-sm mb-2">شماره ملی:</label>
          <input
            type="text"
            name="nationalID"
            className={`w-full p-2 border ${
              formik.touched.nationalID && formik.errors.nationalID
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.nationalID}
          />
          {formik.touched.nationalID && formik.errors.nationalID && (
            <p className="text-red-500 text-sm">{formik.errors.nationalID}</p>
          )}
        </div>

        {/* Manager Name */}
        <div className="mb-4">
          <label className="block text-sm mb-2">نام مدیرعامل:</label>
          <input
            type="text"
            name="managerName"
            className={`w-full p-2 border ${
              formik.touched.managerName && formik.errors.managerName
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.managerName}
          />
          {formik.touched.managerName && formik.errors.managerName && (
            <p className="text-red-500 text-sm">{formik.errors.managerName}</p>
          )}
        </div>

        {/* Position */}
        <div className="mb-4">
          <label className="block text-sm mb-2">سمت:</label>
          <input
            type="text"
            name="position"
            className={`w-full p-2 border ${
              formik.touched.position && formik.errors.position
                ? "border-red-500"
                : "border-gray-300"
            }`}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.position}
          />
          {formik.touched.position && formik.errors.position && (
            <p className="text-red-500 text-sm">{formik.errors.position}</p>
          )}
        </div>

        {/* Submit and Reset Buttons */}
        <div className="flex justify-between">
          <button
            type="submit"
            className="bg-blue-500 text-white py-2 px-4 rounded-md"
          >
            ثبت نام
          </button>
          <button
            type="reset"
            className="bg-gray-500 text-white py-2 px-4 rounded-md"
          >
            پاک کردن
          </button>
        </div>
      </form>
    </div>
  );
};

export default RegisterCharity;
