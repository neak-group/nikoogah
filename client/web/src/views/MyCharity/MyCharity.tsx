import Cookies from "js-cookie";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { FaUserAlt, FaRegFileAlt, FaTimes } from "react-icons/fa";
import Modal from "./MyCharityModal";

export default function MyCharity() {
  const isAuthenticated = !!Cookies.get("isAuthenticated");

  const navigate = useNavigate();
  useEffect(() => {
    if (!isAuthenticated) navigate("/signin");
  }, []);
  return <ProfilePage />;
}

const ProfilePage: React.FC = () => {
  const tableData: TableRow[] = [
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
    {
      capacity: "23/200",
      deadline: "سه روز",
      holders: "32",
      name: "کمک به ایتام",
    },
  ];

  return (
    <div className="pb-16 flex flex-col items-center p-8">
      {/* Profile Header */}
      <div className="w-full max-w-md bg-white shadow-md rounded-md p-6">
        {/* Avatar and Username */}
        <div className="flex items-center justify-between mb-4">
          <div>
            <h2 className="text-lg font-semibold">اسم خیریه</h2>
            <p className="text-sm text-gray-500">مدیر عامل: ممد قلی</p>
          </div>
          <div className="rounded-full h-16 w-16 bg-gray-300 flex items-center justify-center">
            {/* Placeholder for avatar */}
            <span>آواتار</span>
          </div>
        </div>

        {/* Stats */}
        <div className="bg-gray-50 p-4 rounded-md mb-4">
          <p className="text-sm font-semibold">اعتبار خیریه: ۲۰۰۰۰</p>
          <p className="text-sm font-semibold">تعداد فراخوان: ۳۲</p>

          {/* سهم کمک‌ها */}
        </div>
      </div>

      {/* Resume Upload Section */}
      <div className="mt-8 w-full">
        <DynamicTable data={tableData} />
      </div>
    </div>
  );
};

interface TableRow {
  name: string;
  holders: string;
  capacity: string;
  deadline: string;
}

interface DynamicTableProps {
  data: TableRow[];
}

const DynamicTable: React.FC<DynamicTableProps> = ({ data }) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedRow, setSelectedRow] = useState<TableRow | null>(null);
  const [currentPage, setCurrentPage] = useState(1); // Pagination state
  const rowsPerPage = 5; // Number of rows per page

  const totalPages = Math.ceil(data.length / rowsPerPage); // Total number of pages
  const navigate = useNavigate();

  const handleModalOpen = (row: TableRow) => {
    setSelectedRow(row);
    setIsModalOpen(true);
  };

  const handleModalClose = () => {
    setIsModalOpen(false);
    setSelectedRow(null);
  };

  // Paginate the data
  const paginatedData = data.slice(
    (currentPage - 1) * rowsPerPage,
    currentPage * rowsPerPage
  );

  const handlePrevPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  return (
    <div className="p-4">
      <button
        onClick={() => {
          navigate("/new-rally");
        }}
        className="mb-4 bg-blue-500 text-white px-4 py-2 rounded"
      >
        + فراخوان جدید
      </button>

      <table className="min-w-full bg-white border border-gray-300">
        <thead>
          <tr className="border-b">
            <th className="px-6 py-3 text-right">نام فراخوان</th>
            <th className="px-6 py-3 text-right">تعداد دارندگان</th>
            <th className="px-6 py-3 text-right">ظرفیت فعلی</th>
            <th className="px-6 py-3 text-right">مهلت</th>
            <th className="px-6 py-3 text-right">عملیات</th>
          </tr>
        </thead>
        <tbody>
          {paginatedData.map((row, index) => (
            <tr className="border-b" key={index}>
              <td className="px-6 py-4">{row.name}</td>
              <td className="px-6 py-4">{row.holders}</td>
              <td className="px-6 py-4">{row.capacity}</td>
              <td className="px-6 py-4">{row.deadline}</td>
              <td className="px-6 py-4 flex items-center justify-end">
                <button className="p-2 mx-1 bg-gray-200 rounded hover:bg-gray-300">
                  <FaTimes />
                </button>
                <button
                  onClick={() => handleModalOpen(row)}
                  className="p-2 mx-1 bg-gray-200 rounded hover:bg-gray-300"
                >
                  <FaUserAlt />
                </button>
                <button className="p-2 mx-1 bg-gray-200 rounded hover:bg-gray-300">
                  <FaRegFileAlt />
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {/* Pagination Controls */}
      <div className="flex justify-center mt-4">
        <button
          onClick={handlePrevPage}
          disabled={currentPage === 1}
          className={`px-4 py-2 rounded ${
            currentPage === 1 ? "bg-gray-200" : "bg-blue-500 text-white"
          }`}
        >
          صفحه قبل
        </button>
        <span className="px-3 py-1 mx-1">
          {currentPage} / {totalPages}
        </span>
        <button
          onClick={handleNextPage}
          disabled={currentPage === totalPages}
          className={`px-4 py-2 rounded ${
            currentPage === totalPages
              ? "bg-gray-200"
              : "bg-blue-500 text-white"
          }`}
        >
          صفحه بعد
        </button>
      </div>

      {isModalOpen && <Modal onClose={handleModalClose} row={selectedRow} />}
    </div>
  );
};
