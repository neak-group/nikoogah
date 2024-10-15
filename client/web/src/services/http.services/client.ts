import axiosInstance from "./config";

interface ApiOptions {
  endpoint: string;
  method?: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  data?: any;
  tokenRequired?: boolean;
}

export const apiClient = async ({
  endpoint,
  method = "GET",
  data,
  tokenRequired = false,
}: ApiOptions) => {
  try {
    const response = await axiosInstance({
      url: endpoint,
      method,
      data,
      withCredentials: tokenRequired ? true : false,
    });

    return response.data;
  } catch (error) {
    // Handle error if needed
    console.error("API request failed", error);
    throw error;
  }
};
