/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    // For local development, proxy to localhost:3030
    // For Docker deployment, use the literal API URL from environment
    const isLocal = process.env.NODE_ENV === 'development' && !process.env.DOCKER_ENV;
    
    let backendURL;
    if (isLocal) {
      backendURL = 'http://localhost:3030'; // Local development
    } else if (process.env.NEXT_PUBLIC_API_BASE_URL) {
      backendURL = process.env.NEXT_PUBLIC_API_BASE_URL; // Literal URL from Docker Compose
    } else {
      backendURL = 'https://api.dev.cin114.net'; // Fallback
    }

    return [
      {
        source: '/api/:path*',
        destination: `${backendURL}/api/:path*` // Proxy to Backend (local or literal API URL)
      }
    ]
  }
};

export default nextConfig;
