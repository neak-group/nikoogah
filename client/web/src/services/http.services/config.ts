import axios, { AxiosInstance, AxiosResponse, AxiosError } from "axios";

// Base URL for API requests
const BASE_URL = "base_url";

// Create an Axios instance with default configuration
const axiosInstance: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add a request interceptor to include authentication tokens if needed
// axiosInstance.interceptors.request.use(
//   (config) => {
//     // Retrieve token from local storage or context if needed
//     const token = localStorage.getItem("authToken");

//     if (token && config.headers) {
//       config.headers.Authorization = `Bearer ${token}`;
//     }

//     return config;
//   },
//   (error: AxiosError) => {
//     // Handle the request error
//     return Promise.reject(error);
//   }
// );

// Add a response interceptor for global error handling
axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError) => {
    // Handle errors globally here
    if (error.response) {
      console.error("API Error:", error.response.data);
      switch (error.response.status) {
        case 401:
          console.error("Unauthorized access - perhaps redirect to login");
          break;
        case 403:
          console.error("Forbidden - insufficient permissions");
          break;
        case 500:
          console.error("Server error - try again later");
          break;
        default:
          console.error("An unknown error occurred");
      }
    } else {
      console.error("Network error or no response");
    }
    return Promise.reject(error);
  }
);

export default axiosInstance;
