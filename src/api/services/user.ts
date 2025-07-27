import { authAxios } from "@/api/instance/instance";
import { LoginCredentials, LoginResponse, User, AuthError, AuthToken } from "@/types/user";
import { Storage, StorageLks } from "@/services/storage";

const storage = new Storage();

// Helper to log only in development
const devLog = (message: string, ...args: any[]) => {
  if (process.env.NODE_ENV === 'development') {
    console.log(message, ...args);
  }
};

const devError = (message: string, ...args: any[]) => {
  if (process.env.NODE_ENV === 'development') {
    console.error(message, ...args);
  } else {
    console.error(message.replace(/🔐|📡|📦|🔑|🗝️|🌐|💾|🔍|✅|❌|⚠️|🧹|📡|🚪|👤|🔄/g, '').trim());
  }
};

export async function login(credentials: LoginCredentials): Promise<LoginResponse> {
  try {
    devLog('🔐 user.ts: Attempting login with credentials:', { username: credentials.username });

    const response = await authAxios.post<LoginResponse>('/api/v1/auth/login/internal', credentials, {
      timeout: 5000, // 5 seconds timeout for authentication
    });

    devLog('📡 user.ts: Login response received:', response.data);

    // Store the auth token in storage
    if (response.data.access_token) {
      devLog('📦 user.ts: Access token found, attempting to store...');
      devLog('🔑 user.ts: Token preview:', response.data.access_token.substring(0, 20) + '...');
      devLog('🗝️ user.ts: Storage key:', StorageLks.AUTH);

      // Test localStorage is available
      if (typeof window === 'undefined') {
        devError('❌ user.ts: Window is undefined - cannot store token');
        throw new Error('Cannot store token: window is undefined');
      }

      devLog('🌐 user.ts: Window available, proceeding with storage...');

      // Try storing
      storage.safeSetItem(StorageLks.AUTH, response.data.access_token);
      devLog('💾 user.ts: safeSetItem called');

      // Immediate verification
      const immediateCheck = storage.safeGetItem(StorageLks.AUTH);
      devLog('🔍 user.ts: Immediate storage check:', !!immediateCheck);

      // Direct localStorage check
      const directCheck = localStorage.getItem(StorageLks.AUTH);
      devLog('🔍 user.ts: Direct localStorage check:', !!directCheck);

      if (!immediateCheck) {
        devError('❌ user.ts: Token storage failed - safeSetItem did not work');
        throw new Error('Token storage failed');
      }

      devLog('✅ user.ts: Token stored successfully');

      // Notify other components that auth state has changed
      devLog('📡 user.ts: Dispatching auth-state-changed event');
      if (typeof window !== 'undefined') {
        window.dispatchEvent(new Event('auth-state-changed'));
        devLog('✅ user.ts: Event dispatched');
      }
    } else {
      console.warn('No access_token in response:', response.data);
    }

    return response.data;
  } catch (error: any) {
    devError('❌ user.ts: Login failed:', error);

    // Handle different error responses
    if (error.response?.data) {
      const authError: AuthError = {
        error: error.response.data.error || 'Authentication failed',
        message: error.response.data.message || 'Invalid credentials',
        code: error.response.data.code || 'auth-failed',
        status: error.response.status || 401
      };
      throw authError;
    }

    // Fallback error
    const fallbackError: AuthError = {
      error: 'Network Error',
      message: 'Unable to connect to authentication server',
      code: 'network-error',
      status: 500
    };
    throw fallbackError;
  }
}

export async function logout(): Promise<void> {
  try {
    devLog('🚪 user.ts: Logging out...');
    // Call logout endpoint if it exists
    await authAxios.post('/api/v1/auth/logout', {}, {
      timeout: 3000,
    });
  } catch (error) {
    devLog('⚠️ user.ts: Logout endpoint failed:', error);
  } finally {
    // Always clear local storage regardless of server response
    devLog('🧹 user.ts: Clearing auth token from storage');
    storage.safeRemoveItem(StorageLks.AUTH);

    // Notify other components that auth state has changed
    if (typeof window !== 'undefined') {
      window.dispatchEvent(new Event('auth-state-changed'));
    }
  }
}

// Simple JWT decoder (for development - don't use in production without proper validation)
function decodeJWT(token: string): any {
  try {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
        .join('')
    );
    return JSON.parse(jsonPayload);
  } catch (error) {
    devError('Failed to decode JWT:', error);
    return null;
  }
}

export async function getCurrentUser(): Promise<User | null> {
  try {
    const token = storage.safeGetItem(StorageLks.AUTH);
    devLog('👤 user.ts: Getting current user, token exists:', !!token);

    if (!token) {
      devLog('❌ user.ts: No token found for getCurrentUser');
      return null;
    }

    // Since /api/v1/user/me doesn't exist, decode user info from JWT token
    devLog('🔍 user.ts: Decoding user info from JWT token...');
    const payload = decodeJWT(token);

    if (!payload || !payload.user_id) {
      devLog('❌ user.ts: Invalid token payload');
      storage.safeRemoveItem(StorageLks.AUTH);
      return null;
    }

    // Check if token is expired
    if (payload.exp && payload.exp < Date.now() / 1000) {
      devLog('❌ user.ts: Token is expired');
      storage.safeRemoveItem(StorageLks.AUTH);
      return null;
    }

    // Create user object from JWT payload
    const user: User = {
      id: payload.user_id,
      username: `user${payload.user_id}`, // We don't have username in JWT, use fallback
      email: undefined,
      createdAt: undefined,
      updatedAt: undefined
    };

    devLog('✅ user.ts: User extracted from JWT:', user);
    return user;
  } catch (error) {
    devLog('⚠️ user.ts: Failed to get current user:', error);
    // If user fetch fails, token might be invalid
    devLog('🧹 user.ts: Removing invalid token');
    storage.safeRemoveItem(StorageLks.AUTH);
    return null;
  }
}

export async function refreshToken(): Promise<AuthToken | null> {
  try {
    devLog('🔄 user.ts: Refreshing token...');
    const response = await authAxios.post<AuthToken>('/api/v1/auth/refresh', {}, {
      timeout: 3000,
    });

    // Update stored token
    if (response.data.access_token) {
      devLog('💾 user.ts: Updating stored token after refresh');
      storage.safeSetItem(StorageLks.AUTH, response.data.access_token);
    }

    return response.data;
  } catch (error) {
    devLog('⚠️ user.ts: Token refresh failed:', error);
    // Clear invalid token
    devLog('🧹 user.ts: Clearing invalid token after failed refresh');
    storage.safeRemoveItem(StorageLks.AUTH);
    return null;
  }
}

export function isAuthenticated(): boolean {
  // Only run on client side
  if (typeof window === 'undefined') {
    devLog('🌐 user.ts: isAuthenticated called on server side, returning false');
    return false;
  }

  const token = storage.safeGetItem(StorageLks.AUTH);
  devLog('🔍 user.ts: isAuthenticated check - token exists:', !!token);

  if (!token) {
    return false;
  }

  // Check if token is expired
  const payload = decodeJWT(token);
  if (payload && payload.exp && payload.exp < Date.now() / 1000) {
    devLog('🔍 user.ts: Token is expired, removing...');
    storage.safeRemoveItem(StorageLks.AUTH);
    return false;
  }

  return !!token;
} 
