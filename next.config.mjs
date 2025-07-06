/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    // For local development, proxy to localhost:3030
    // For production/deployment, proxy to the backend service in Docker
    const backendURL = process.env.NODE_ENV
      ? 'http://api:3030' // Docker service name and internal port
      : 'http://localhost:3030';

    return [
      {
        source: '/api/:path*',
        destination: `${backendURL}/api/:path*` // Proxy to Backend
      }
    ]
  }
};

export default nextConfig;
