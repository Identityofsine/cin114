import { Storage, StorageLks } from "@/services/storage";
import axios, { InternalAxiosRequestConfig } from "axios";

// Get base URL dynamically based on environment
const getBaseURL = () => {
  // In browser environment, use current origin
  if (typeof window !== 'undefined') {
    const currentOrigin = window.location.origin.toLowerCase(); // Make case-insensitive
    
    console.log('Current origin:', currentOrigin); // Debug log
    
    // If we're on dev.cin114.net, use dev API
    if (currentOrigin.includes('dev.cin114.net')) {
      console.log('Using dev API');
      return 'https://api.dev.cin114.net';
    }

    // If we're on production cin114.net or cin114films.com, use prod API
    if (currentOrigin.includes('cin114.net') || currentOrigin.includes('cin114films.com')) {
      console.log('Using prod API');
      return 'https://api.cin114.net';
    }

    // If we're on localhost, use dev API
    if (currentOrigin.includes('localhost')) {
      console.log('Using localhost/dev fallback API');
      return 'https://api.dev.cin114.net';
    }

    // For localhost or other development environments, use dev API as fallback
    console.log('Using prod API');
    return 'https://api.cin114.net';
  }

  // For SSR environment, check environment variables
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (apiBaseUrl) {
    return apiBaseUrl;
  }

  // Default fallback for development
  return 'https://api.cin114.net';
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
