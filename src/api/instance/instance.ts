import { Storage, StorageLks } from "@/services/storage";
import axios, { InternalAxiosRequestConfig } from "axios";

// Get base URL dynamically based on environment
const getBaseURL = () => {
  // In browser environment, use current origin
  if (typeof window !== 'undefined') {
    const currentOrigin = window.location.origin;

    // If we're on dev.cin114.net, use dev API
    if (currentOrigin.includes('dev.cin114.net')) {
      return 'https://api.dev.cin114.net';
    }

    // If we're on cin114.net (production), use prod API
    if (currentOrigin.includes('cin114.net') || currentOrigin.includes('cin114films.com')) {
      return 'https://api.cin114.net';
    }

    // For localhost or other development environments, use local API
    return 'http://localhost:3030';
  }

  // For SSR environment, check environment variables
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (apiBaseUrl) {
    return apiBaseUrl;
  }

  // Default fallback for development
  return 'http://localhost:3030';
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
