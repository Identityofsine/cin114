import { Storage, StorageLks } from "@/services/storage";
import axios, { InternalAxiosRequestConfig } from "axios";

// Get base URL from environment variable, default to localhost:3000
const getBaseURL = () => {
  return 'http://localhost:3000';
};

// Create axios instance with base URL from environment
const axiosInstance = axios.create({
  baseURL: getBaseURL(),
  timeout: 10000,
});

const storage = new Storage();

// middleware to handle auth token from localStorage
axiosInstance.interceptors.request.use(
  (config) => {
    const AUTH_TOKEN = storage.safeGetItem(StorageLks.AUTH);
    if (!AUTH_TOKEN) {
      return config;
    }
    return {
      ...config,
      headers: {
        ...config.headers,
        Authorization: `Bearer ${AUTH_TOKEN}`
      }
    } as InternalAxiosRequestConfig
  }
)

export const authAxios = axiosInstance;
