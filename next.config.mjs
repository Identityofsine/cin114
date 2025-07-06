/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    // For local development, proxy to localhost:3030
    // For production/deployment, proxy to the external API subdomain
    const isLocal = process.env.NODE_ENV === 'development' && !process.env.DOCKER_ENV;
    const isProduction = process.env.NEXT_PUBLIC_BRANCH === 'main' || process.env.NEXT_PUBLIC_BRANCH === 'prod';
    
    let backendURL;
    if (isLocal) {
      backendURL = 'http://localhost:3030'; // Local development
    } else if (isProduction) {
      backendURL = 'https://api.cin114.net'; // Production API subdomain
    } else {
      backendURL = 'https://api.dev.cin114.net'; // Dev API subdomain
    }

    return [
      {
        source: '/api/:path*',
        destination: `${backendURL}/api/:path*` // Proxy to Backend (local or external API subdomain)
      }
    ]
  }
};

export default nextConfig;
