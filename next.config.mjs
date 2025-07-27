/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    // For local development, proxy to localhost:3030
    // For Docker deployment, use the API base URL from environment
    let backendURL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:3030";

    return [
      {
        source: '/api/:path*',
        destination: `${backendURL}/api/:path*` // Proxy to Backend (local or literal API URL)
      }
    ]
  }
};

export default nextConfig;
