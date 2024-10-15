import React, { useState } from "react";
import { useFormik } from "formik";
import * as Yup from "yup";
import { FaCalendar } from "react-icons/fa";
import { FaUpload, FaTrash } from "react-icons/fa";

import { Calendar } from "@amir04lm26/react-modern-calendar-date-picker";
import "@amir04lm26/react-modern-calendar-date-picker/lib/DatePicker.css"; // Date picker styles

// Validation schema using Yup
const validationSchema = Yup.object({
  name: Yup.string().required("نام فراخوان الزامی است"),
  date: Yup.object().nullable().required("تاریخ ایام فراخوان الزامی است"),
  description: Yup.string().required("تفصیلات فراخوان الزامی است"),
});

const MyForm = () => {
  // State for controlling the calendar visibility
  const [isCalendarOpen, setIsCalendarOpen] = useState(false);

  // Initial form values
  const formik = useFormik({
    initialValues: {
      name: "",
      date: null,
      description: "",
    },
    validationSchema: validationSchema,
    onSubmit: (values) => {
      console.log("Form values:", values);
    },
  });

  // Toggle calendar visibility
  const toggleCalendar = () => {
    setIsCalendarOpen((prev) => !prev);
  };

  return (
    <div className="max-w-lg mx-auto mt-10 p-4 border rounded-md shadow-md">
      <h1 className="text-xl mb-4">فراخوان جدید</h1>

      <form onSubmit={formik.handleSubmit}>
        {/* Name input */}
        <div className="mb-4">
          <label htmlFor="name" className="block mb-2 text-right">
            نام فراخوان:
          </label>
          <input
            type="text"
            id="name"
            name="name"
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.name}
            className={`w-full border p-2 rounded ${
              formik.touched.name && formik.errors.name ? "border-red-500" : ""
            }`}
            placeholder="نام فراخوان را وارد کنید"
          />
          {formik.touched.name && formik.errors.name ? (
            <div className="text-red-500 text-sm mt-1">
              {formik.errors.name}
            </div>
          ) : null}
        </div>

        {/* Persian Date Picker */}
        <div className="flex items-center mb-4">
          <label htmlFor="date" className="block text-right">
            تاریخ ایام فراخوان:
          </label>
          <div className="relative">
            <button
              type="button"
              onClick={toggleCalendar}
              className="mx-2 pt-2 text-gray-700"
            >
              <FaCalendar />
            </button>
            {isCalendarOpen && (
              <div className="absolute z-10">
                <Calendar
                  value={formik.values.date}
                  onChange={(value) => {
                    formik.setFieldValue("date", value);
                    setIsCalendarOpen(false); // Close the calendar after picking a date
                  }}
                  shouldHighlightWeekends
                  locale="fa" // For Persian calendar
                />
              </div>
            )}
          </div>
        </div>

        {/* Description */}
        <div className="mb-4">
          <label htmlFor="description" className="block mb-2 text-right">
            توضیحات فراخوان:
          </label>
          <textarea
            id="description"
            name="description"
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            value={formik.values.description}
            className={`w-full border p-2 rounded h-24 ${
              formik.touched.description && formik.errors.description
                ? "border-red-500"
                : ""
            }`}
            placeholder="تفصیلات فراخوان را وارد کنید"
          ></textarea>
          {formik.touched.description && formik.errors.description ? (
            <div className="text-red-500 text-sm mt-1">
              {formik.errors.description}
            </div>
          ) : null}
        </div>

        {/* Submit Button */}
        {/* <button
          type="submit"
          className="bg-blue-500 text-white px-4 py-2 rounded disabled:opacity-25"
          disabled={!formik.isValid || formik.isSubmitting}
        >
          ایجاد فراخوان
        </button> */}
      </form>
    </div>
  );
};

const PhotoUpload = () => {
  const [photos, setPhotos] = useState([]);

  // Handle file input change
  const handleFileChange = (e: any) => {
    const newPhotos = Array.from(e.target.files).slice(0, 6 - photos.length);
    //@ts-ignore
    setPhotos((prev) => [...prev, ...newPhotos]);
  };

  // Remove photo
  const removePhoto = (index: any) => {
    setPhotos((prev) => prev.filter((_, i) => i !== index));
  };

  return (
    <div className="max-w-lg mx-auto mt-10 p-6 border border-gray-300 rounded-lg shadow-lg bg-white">
      <h1 className="text-xl mb-4 text-center">بارگذاری تصاویر</h1>

      {/* Upload Box */}
      <div className="flex flex-col items-center mb-4">
        <div
          className="border-2 border-dashed border-gray-400 p-4 rounded-lg flex flex-col items-center justify-center cursor-pointer hover:bg-gray-100 transition"
          onClick={() => document.getElementById("fileInput")?.click()}
        >
          <FaUpload className="text-2xl mb-2" />
          <p className="text-gray-500">برای بارگذاری تصویر کلیک کنید</p>
        </div>
        <input
          type="file"
          id="fileInput"
          onChange={handleFileChange}
          accept="image/*"
          multiple
          className="hidden"
        />
      </div>

      {/* Display Uploaded Photos */}
      <div className="grid grid-cols-3 gap-4 mb-4">
        {photos.map((photo, index) => (
          <div key={index} className="relative">
            <img
              src={URL.createObjectURL(photo)}
              alt={`Uploaded ${index + 1}`}
              className="w-full h-32 object-cover rounded-lg"
            />
            <button
              className="absolute top-1 right-1 bg-red-500 text-white rounded-full p-1 hover:bg-red-600"
              onClick={() => removePhoto(index)}
            >
              <FaTrash />
            </button>
          </div>
        ))}
      </div>

      {/* Main Image Requirement */}
      <div className="mb-4">
        {photos.length === 0 && (
          <div className="text-red-500 text-sm text-center">
            لطفا حداقل یک تصویر (تصویر اصلی) بارگذاری کنید.
          </div>
        )}
        {photos.length > 0 && (
          <p className="text-gray-600 text-sm text-center">
            تصویر اصلی:{" "}
            {photos.length > 0
              ? //@ts-ignore
                photos[0].name
              : "نامشخص"}
          </p>
        )}
      </div>

      {/* Add Photo Button */}
      {photos.length < 6 && (
        <button
          onClick={() => document.getElementById("fileInput")?.click()}
          className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition"
        >
          اضافه کردن تصویر
        </button>
      )}
    </div>
  );
};

const FinancialAssistanceForm = () => {
  const [isFormActive, setIsFormActive] = useState(false);
  const [assistanceAmount, setAssistanceAmount] = useState("");
  const [leaveOpen, setLeaveOpen] = useState(false);

  const handleCheckboxChange = () => {
    setIsFormActive((prev) => !prev);
  };

  const handleAmountChange = (e: any) => {
    const value = e.target.value.replace(/\D/g, ""); // Remove non-digit characters
    setAssistanceAmount(value);
  };

  const handleLeaveOpenChange = (e: any) => {
    setLeaveOpen(e.target.checked);
  };

  return (
    <div className="max-w-lg mx-auto mt-10 p-6 border border-gray-300 rounded-lg shadow-lg bg-white">
      {/* Checkbox to activate the form */}
      <div className="mb-4">
        <label className="flex items-center">
          <span>فعال کردن فرم کمک مالی</span>
          <input
            type="checkbox"
            checked={isFormActive}
            onChange={handleCheckboxChange}
            className="mr-2"
          />
        </label>
      </div>

      {/* Form Fields */}
      <form>
        <div
          className={`mb-4 ${
            isFormActive ? "" : "opacity-25 pointer-events-none"
          }`}
        >
          <label htmlFor="assistanceAmount" className="block mb-2 text-right">
            مقدار کمک مالی (تومان):
          </label>
          <input
            type="text"
            id="assistanceAmount"
            value={assistanceAmount}
            onChange={handleAmountChange}
            placeholder="مقدار را وارد کنید"
            className="w-full border p-2 rounded"
          />
        </div>

        <label
          className={`flex items-center mb-4 ${
            isFormActive ? "" : "opacity-25 pointer-events-none"
          }`}
        >
          <span>بازگذاشتن کمک مالی پس از رسیدن به هدف مالی</span>
          <input
            type="checkbox"
            checked={leaveOpen}
            onChange={handleLeaveOpenChange}
            className="mr-2"
          />
        </label>
      </form>
    </div>
  );
};

const VolunteerNeedsForm = () => {
  const [isFormActive, setIsFormActive] = useState(false);
  const [numberOfVolunteers, setNumberOfVolunteers] = useState("");
  const [needsResume, setNeedsResume] = useState(false);
  const [volunteerConditions, setVolunteerConditions] = useState("");

  const handleCheckboxChange = () => {
    setIsFormActive((prev) => !prev);
  };

  const handleNumberChange = (e: any) => {
    const value = e.target.value.replace(/\D/g, ""); // Remove non-digit characters
    setNumberOfVolunteers(value);
  };

  const handleResumeChange = (e: any) => {
    setNeedsResume(e.target.checked);
  };

  const handleConditionsChange = (e: any) => {
    setVolunteerConditions(e.target.value);
  };

  return (
    <div className="max-w-lg mx-auto mt-10 p-6 border border-gray-300 rounded-lg shadow-lg bg-white">
      {/* Checkbox to activate the form */}
      <div className="mb-4">
        <label className="flex items-center">
          <span>نیاز به داوطلب</span>
          <input
            type="checkbox"
            checked={isFormActive}
            onChange={handleCheckboxChange}
            className="mr-2"
          />
        </label>
      </div>

      {/* Form Fields */}
      <form>
        <div
          className={`mb-4 ${
            isFormActive ? "" : "opacity-25 pointer-events-none"
          }`}
        >
          <label htmlFor="numberOfVolunteers" className="block mb-2 text-right">
            تعداد داوطلبان:
          </label>
          <input
            type="text"
            id="numberOfVolunteers"
            value={numberOfVolunteers}
            onChange={handleNumberChange}
            placeholder="تعداد را وارد کنید"
            className="w-full border p-2 rounded"
          />
        </div>

        <label
          className={`flex items-center mb-4 ${
            isFormActive ? "" : "opacity-25 pointer-events-none"
          }`}
        >
          <span>نیاز به رزومه</span>
          <input
            type="checkbox"
            checked={needsResume}
            onChange={handleResumeChange}
            className="mr-2"
          />
        </label>

        <div
          className={`mb-4 ${
            isFormActive ? "" : "opacity-25 pointer-events-none"
          }`}
        >
          <label
            htmlFor="volunteerConditions"
            className="block mb-2 text-right"
          >
            شرایط داوطلب:
          </label>
          <textarea
            id="volunteerConditions"
            value={volunteerConditions}
            onChange={handleConditionsChange}
            placeholder="شرایط داوطلب را وارد کنید"
            className="w-full border p-2 rounded h-24"
          ></textarea>
        </div>
      </form>
    </div>
  );
};

export default function NewRally() {
  return (
    <div>
      <MyForm />
      <PhotoUpload />
      <FinancialAssistanceForm />
      <VolunteerNeedsForm />
      <div className="flex justify-center my-6">
        <button className="bg-blue-500 text-white rounded shadow-md px-4 py-2">
          ثبت فراخوان
        </button>
      </div>
    </div>
  );
}
