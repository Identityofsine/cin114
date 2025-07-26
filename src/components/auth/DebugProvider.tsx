'use client';

import AuthDebug from './AuthDebug';

export default function DebugProvider() {
  // Only render in development
  if (process.env.NODE_ENV !== 'development') {
    return null;
  }

  return <AuthDebug />;
} 