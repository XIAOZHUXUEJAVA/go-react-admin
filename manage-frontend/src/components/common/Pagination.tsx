import React from "react";
import { PaginationInfo } from "@/types/api";

interface PaginationProps {
  pagination: PaginationInfo;
  onPageChange: (page: number) => void;
  onPageSizeChange: (pageSize: number) => void;
  className?: string;
}

/**
 * 分页组件
 */
export const Pagination: React.FC<PaginationProps> = ({
  pagination,
  onPageChange,
  onPageSizeChange,
  className = "",
}) => {
  const {
    page,
    page_size: pageSize,
    total,
    total_pages: totalPages,
  } = pagination;

  const handlePrevious = () => {
    if (page > 1) {
      onPageChange(page - 1);
    }
  };

  const handleNext = () => {
    if (page < totalPages) {
      onPageChange(page + 1);
    }
  };

  const handlePageSizeChange = (
    event: React.ChangeEvent<HTMLSelectElement>
  ) => {
    const newPageSize = parseInt(event.target.value, 10);
    onPageSizeChange(newPageSize);
  };

  const getPageNumbers = (): number[] => {
    const pages: number[] = [];
    const maxVisible = 5;

    if (totalPages <= maxVisible) {
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      const start = Math.max(1, page - 2);
      const end = Math.min(totalPages, start + maxVisible - 1);

      for (let i = start; i <= end; i++) {
        pages.push(i);
      }
    }

    return pages;
  };

  return (
    <div className={`flex items-center justify-between ${className}`}>
      {/* 分页信息 */}
      <div className="text-sm text-gray-700">
        显示第 <span className="font-medium">{(page - 1) * pageSize + 1}</span>{" "}
        到{" "}
        <span className="font-medium">{Math.min(page * pageSize, total)}</span>{" "}
        条，共 <span className="font-medium">{total}</span> 条记录
      </div>

      <div className="flex items-center space-x-4">
        {/* 每页显示数量选择 */}
        <div className="flex items-center space-x-2">
          <label htmlFor="pageSize" className="text-sm text-gray-700">
            每页显示:
          </label>
          <select
            id="pageSize"
            value={pageSize}
            onChange={handlePageSizeChange}
            className="border border-gray-300 rounded px-2 py-1 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value={5}>5</option>
            <option value={10}>10</option>
            <option value={20}>20</option>
            <option value={50}>50</option>
          </select>
        </div>

        {/* 分页按钮 */}
        <div className="flex items-center space-x-1">
          <button
            onClick={handlePrevious}
            disabled={page <= 1}
            className="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            上一页
          </button>

          {getPageNumbers().map((pageNum) => (
            <button
              key={pageNum}
              onClick={() => onPageChange(pageNum)}
              className={`px-3 py-1 text-sm border rounded ${
                pageNum === page
                  ? "bg-blue-500 text-white border-blue-500"
                  : "border-gray-300 hover:bg-gray-50"
              }`}
            >
              {pageNum}
            </button>
          ))}

          <button
            onClick={handleNext}
            disabled={page >= totalPages}
            className="px-3 py-1 text-sm border border-gray-300 rounded hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            下一页
          </button>
        </div>
      </div>
    </div>
  );
};
