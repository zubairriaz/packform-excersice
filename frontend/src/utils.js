import axios from "axios";

export const getColumns = () => [
  {
    headerName: "Order Name",
    field: "OrderName",
    sortable: true,
    filter: true,
  },
  { headerName: "Customer Company", field: "CompanyName" },
  { headerName: "Customer Name", field: "CustomerName" },
  {
    headerName: "Order Date",
    field: "OrderDate",
    sortable: true,
    filter: true,
    valueFormatter: formatTime,
  },
  {
    headerName: "Delivered Amount",
    field: "DeliveredAmount",
    valueFormatter: formatAmount,
  },
  {
    headerName: "Total Amount",
    field: "TotalAmount",
    valueFormatter: formatAmount,
  },
];

export const AXIOS = axios.create({
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
  },
  withCredentials: true,
  credentials: "same-origin",
  crossdomain: true,
});

export const RECORDS_PER_PAGE = 5;

export const formatTime = ({ value }) => {
  let dateTime = new Date(value);
  return dateTime.toLocaleTimeString([], {
    month:"short",
    day:"numeric",
    hour: "2-digit",
    minute: "2-digit",
    timeZone: "Australia/Melbourne",
  });
};

const formatter = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'AUD',
});

export const formatAmount = ({ value }) => {
  return value ? formatter.format(value) : "N/A";
};


