import React, { useState } from "react";
import { IoMdClose } from "react-icons/io";
interface CharityTableRow {
  name: string;
  holders: string;
  capacity: string;
  deadline: string;
}

interface ModalProps {
  onClose: () => void;
  row: CharityTableRow | null;
}

const Modal: React.FC<ModalProps> = ({ onClose, row }) => {
  if (!row) return null;

  const tableData: TableRow[] = [
    { mobileNumber: "09125834567", name: "هاشم بیطرف" },
    { mobileNumber: "09125834567", name: "هاشم بیطرف" },
    { mobileNumber: "09125834567", name: "هاشم بیطرف" },
    { mobileNumber: "09125834567", name: "هاشم بیطرف" },
    { mobileNumber: "09125834568", name: "علی مرادی" },
    { mobileNumber: "09125834569", name: "سارا محمدی" },
    { mobileNumber: "09125834570", name: "زهرا سمیعی" },
    { mobileNumber: "09125834571", name: "مهدی کیانی" },
  ];

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center w-full">
      <div className="bg-white p-6 rounded-lg shadow-lg relative w-[80%]">
        <button
          className="absolute top-2 right-2 text-gray-600 hover:text-gray-800"
          onClick={onClose}
        >
          <IoMdClose size={24} />
        </button>
        <div className="mt-6">
          <h2 className="text-lg font-bold mb-4">گزارش فراخوان: {row.name}</h2>
          <DynamicTable data={tableData} />
          <button
            className="mt-4 bg-blue-500 text-white px-4 py-2 rounded"
            onClick={onClose}
          >
            بستن
          </button>
        </div>
      </div>
    </div>
  );
};

interface TableRow {
  name: string;
  mobileNumber: string;
}

interface DynamicTableProps {
  data: TableRow[];
}

const DynamicTable: React.FC<DynamicTableProps> = ({ data }) => {
  // State for pagination
  const [currentPage, setCurrentPage] = useState(1);
  const rowsPerPage = 5;

  // Calculate the total number of pages
  const totalPages = Math.ceil(data.length / rowsPerPage);

  // Determine the rows to display based on the current page
  const startIndex = (currentPage - 1) * rowsPerPage;
  const currentRows = data.slice(startIndex, startIndex + rowsPerPage);

  // Handle next and previous page buttons
  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  return (
    <div className="p-4 flex flex-col items-center">
      <table className="min-w-full bg-white border border-gray-300">
        <thead>
          <tr className="border-b">
            <th className="px-6 py-3 text-right">نام داوطلب</th>
            <th className="px-6 py-3 text-right">شماره تماس</th>
            <th className="px-6 py-3 text-right">رزومه</th>
          </tr>
        </thead>
        <tbody>
          {currentRows.map((row, index) => (
            <tr className="border-b" key={index}>
              <td className="px-6 py-4">{row.name}</td>
              <td className="px-6 py-4">{row.mobileNumber}</td>
              <td className="px-6 py-4">
                <button
                  className="mt-4 bg-slate-300 text-black px-4 py-2 rounded"
                  onClick={() => console.log("download")}
                >
                  دانلود
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {/* Pagination controls */}
      <div className="mt-4 flex justify-between items-center w-fit gap-2">
        <button
          className={`px-4 py-2 rounded ${
            currentPage === 1 ? "bg-gray-200" : "bg-blue-500 text-white"
          }`}
          onClick={handlePreviousPage}
          disabled={currentPage === 1}
        >
          قبلی
        </button>
        <span>
          صفحه {currentPage} از {totalPages}
        </span>
        <button
          className={`px-4 py-2 rounded ${
            currentPage === totalPages
              ? "bg-gray-200"
              : "bg-blue-500 text-white"
          }`}
          onClick={handleNextPage}
          disabled={currentPage === totalPages}
        >
          بعدی
        </button>
      </div>
    </div>
  );
};

export default Modal;
