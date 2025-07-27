'use client';

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getCurrentUser, isAuthenticated } from '@/api/services/user';
import { User } from '@/types/user';
import { Storage, StorageLks } from '@/services/storage';
import './styles/routeguard.scss';

interface RouteGuardProps {
  children: React.ReactNode;
  fallback?: React.ReactNode;
  redirectTo?: string;
}

interface AuthState {
  isLoading: boolean;
  isAuthenticated: boolean;
  user: User | null;
  isHydrated: boolean;
}

const storage = new Storage();

// Helper to log only in development
const devLog = (message: string, ...args: unknown[]) => {
  if (process.env.NODE_ENV === 'development') {
    console.log(message, ...args);
  }
};

export default function RouteGuard({
  children,
  fallback = <div>Redirecting...</div>,
  redirectTo = '/zcpkcrucpw'
}: RouteGuardProps) {
  const router = useRouter();
  const [authState, setAuthState] = useState<AuthState>({
    isLoading: true,
    isAuthenticated: false,
    user: null,
    isHydrated: false
  });

  const checkAuth = async () => {
    try {
      devLog('RouteGuard: Checking authentication...');

      // Force re-read from localStorage
      const token = storage.safeGetItem(StorageLks.AUTH);
      devLog('RouteGuard: Raw token from localStorage:', !!token);

      // First check if there's a token in storage
      const hasToken = isAuthenticated();
      devLog('RouteGuard: Has token:', hasToken);

      if (!hasToken) {
        devLog('RouteGuard: No token found, redirecting to login');
        setAuthState(prev => ({
          ...prev,
          isLoading: false,
          isAuthenticated: false,
          user: null
        }));
        router.push(redirectTo);
        return;
      }

      // Then verify the token with the server
      devLog('RouteGuard: Verifying token...');
      const user = await getCurrentUser();
      devLog('RouteGuard: Current user:', !!user);

      if (user) {
        devLog('RouteGuard: User authenticated successfully');
        setAuthState(prev => ({
          ...prev,
          isLoading: false,
          isAuthenticated: true,
          user
        }));
      } else {
        devLog('RouteGuard: Token invalid, redirecting to login');
        setAuthState(prev => ({
          ...prev,
          isLoading: false,
          isAuthenticated: false,
          user: null
        }));
        router.push(redirectTo);
      }
    } catch (error) {
      console.error('RouteGuard: Auth check failed:', error);
      setAuthState(prev => ({
        ...prev,
        isLoading: false,
        isAuthenticated: false,
        user: null
      }));
      router.push(redirectTo);
    }
  };

  useEffect(() => {
    // Mark as hydrated and check auth immediately
    devLog('RouteGuard: Component hydrated, checking localStorage...');
    setAuthState(prev => ({ ...prev, isHydrated: true }));

    // Small delay to ensure localStorage is fully available
    setTimeout(() => {
      devLog('RouteGuard: Running delayed auth check...');
      checkAuth();
    }, 100);
  }, []);

  useEffect(() => {
    // Listen for auth state changes (like after login)
    const handleAuthChange = () => {
      devLog('RouteGuard: Auth state change event received');
      setTimeout(() => checkAuth(), 100); // Small delay to ensure localStorage is updated
    };

    const handleStorageChange = () => {
      devLog('RouteGuard: localStorage change detected');
      checkAuth();
    };

    if (typeof window !== 'undefined') {
      window.addEventListener('auth-state-changed', handleAuthChange);
      window.addEventListener('storage', handleStorageChange);
      return () => {
        window.removeEventListener('auth-state-changed', handleAuthChange);
        window.removeEventListener('storage', handleStorageChange);
      };
    }
  }, []);

  // Show loading state while checking authentication or during hydration
  if (!authState.isHydrated || authState.isLoading) {
    return (
      <div className="route-guard__loading">
        {fallback}
      </div>
    );
  }

  // Show protected content only if authenticated
  if (authState.isAuthenticated) {
    return <>{children}</>;
  }

  // This shouldn't be reached due to router.push, but just in case
  return (
    <div className="route-guard__fallback">
      {fallback}
    </div>
  );
}

// Hook to get current auth state
export function useAuth() {
  const [authState, setAuthState] = useState<AuthState>({
    isLoading: true,
    isAuthenticated: false,
    user: null,
    isHydrated: false
  });

  const checkAuth = async () => {
    try {
      devLog('useAuth: Checking authentication...');

      // Force re-read from localStorage
      const token = storage.safeGetItem(StorageLks.AUTH);
      devLog('useAuth: Raw token from localStorage:', !!token);

      const hasToken = isAuthenticated();
      devLog('useAuth: Has token:', hasToken);

      if (!hasToken) {
        setAuthState(prev => ({
          ...prev,
          isLoading: false,
          isAuthenticated: false,
          user: null
        }));
        return;
      }

      const user = await getCurrentUser();
      devLog('useAuth: Current user:', !!user);

      setAuthState(prev => ({
        ...prev,
        isLoading: false,
        isAuthenticated: !!user,
        user
      }));
    } catch (error) {
      console.error('useAuth: Auth check failed:', error);
      setAuthState(prev => ({
        ...prev,
        isLoading: false,
        isAuthenticated: false,
        user: null
      }));
    }
  };

  useEffect(() => {
    // Mark as hydrated and check auth immediately
    devLog('useAuth: Component hydrated, checking localStorage...');
    setAuthState(prev => ({ ...prev, isHydrated: true }));

    // Small delay to ensure localStorage is fully available
    setTimeout(() => {
      devLog('useAuth: Running delayed auth check...');
      checkAuth();
    }, 100);
  }, []);

  useEffect(() => {
    // Listen for auth state changes
    const handleAuthChange = () => {
      devLog('useAuth: Auth state change event received');
      setTimeout(() => checkAuth(), 100); // Small delay to ensure localStorage is updated
    };

    const handleStorageChange = () => {
      devLog('useAuth: localStorage change detected');
      checkAuth();
    };

    if (typeof window !== 'undefined') {
      window.addEventListener('auth-state-changed', handleAuthChange);
      window.addEventListener('storage', handleStorageChange);
      return () => {
        window.removeEventListener('auth-state-changed', handleAuthChange);
        window.removeEventListener('storage', handleStorageChange);
      };
    }
  }, []);

  return {
    isLoading: authState.isLoading || !authState.isHydrated,
    isAuthenticated: authState.isAuthenticated,
    user: authState.user
  };
} 
