'use client';

import React, { useState, useEffect } from 'react';
import { useAuth } from './RouteGuard';
import { isAuthenticated, logout } from '@/api/services/user';
import { Storage, StorageLks } from '@/services/storage';

const storage = new Storage();

export default function AuthDebug() {
  const { isLoading, isAuthenticated: authHookAuthenticated, user } = useAuth();
  const [rawToken, setRawToken] = useState<string | null>(null);
  const [localStorageTest, setLocalStorageTest] = useState<string>('');

  // Only show in development environment
  if (process.env.NODE_ENV === 'production') {
    return null;
  }

  const refreshTokenDisplay = () => {
    const token = storage.safeGetItem(StorageLks.AUTH);
    setRawToken(token);
    console.log('Manual token check:', token);
    
    // Also try direct localStorage access
    if (typeof window !== 'undefined') {
      const directToken = localStorage.getItem(StorageLks.AUTH);
      console.log('Direct localStorage check:', directToken);
    }
  };

  const testLocalStorage = () => {
    console.log('Testing localStorage...');
    const testKey = 'test-key';
    const testValue = 'test-value-' + Date.now();
    
    try {
      // Test setting
      localStorage.setItem(testKey, testValue);
      console.log('localStorage SET test:', testValue);
      
      // Test getting
      const retrieved = localStorage.getItem(testKey);
      console.log('localStorage GET test:', retrieved);
      
      // Test via storage service
      storage.safeSetItem(testKey as any, testValue);
      const serviceRetrieved = storage.safeGetItem(testKey as any);
      console.log('Storage service test:', serviceRetrieved);
      
      // Clean up
      localStorage.removeItem(testKey);
      
      setLocalStorageTest(`Direct: ${retrieved === testValue ? 'PASS' : 'FAIL'}, Service: ${serviceRetrieved === testValue ? 'PASS' : 'FAIL'}`);
    } catch (error) {
      console.error('localStorage test failed:', error);
      setLocalStorageTest('ERROR: ' + error);
    }
  };

  const debugAuthToken = () => {
    console.log('=== AUTH TOKEN DEBUG ===');
    console.log('StorageLks.AUTH key:', StorageLks.AUTH);
    console.log('window.localStorage:', typeof window !== 'undefined' ? window.localStorage : 'undefined');
    
    if (typeof window !== 'undefined') {
      console.log('All localStorage keys:', Object.keys(localStorage));
      console.log('All localStorage items:');
      for (let i = 0; i < localStorage.length; i++) {
        const key = localStorage.key(i);
        const value = localStorage.getItem(key!);
        console.log(`  ${key}: ${value}`);
      }
    }
      };

  useEffect(() => {
    // Check token on mount
    refreshTokenDisplay();
    testLocalStorage();
    
    // Check token every 2 seconds to see if it changes
    const interval = setInterval(refreshTokenDisplay, 2000);
    return () => clearInterval(interval);
  }, []);

  const handleLogout = async () => {
    await logout();
    window.location.reload(); // Force reload to reset auth state
  };

  const checkToken = () => {
    const token = storage.safeGetItem(StorageLks.AUTH);
    console.log('Raw token from storage:', token);
    console.log('isAuthenticated() result:', isAuthenticated());
    refreshTokenDisplay();
  };

  const forceRefresh = () => {
    console.log('Forcing auth state refresh...');
    window.dispatchEvent(new Event('auth-state-changed'));
  };

  return (
    <div style={{ 
      position: 'fixed', 
      top: '10px', 
      right: '10px', 
      background: '#333', 
      color: 'white', 
      padding: '1rem', 
      borderRadius: '8px',
      fontSize: '12px',
      zIndex: 9999,
      maxWidth: '320px'
    }}>
      <h4>Auth Debug (DEV)</h4>
      <p>Loading: {isLoading ? 'Yes' : 'No'}</p>
      <p>Authenticated: {authHookAuthenticated ? 'Yes' : 'No'}</p>
      <p>User: {user ? user.username : 'None'}</p>
      <p>Raw Token: {rawToken ? 'EXISTS' : 'NULL'}</p>
      <p>Storage Test: {localStorageTest}</p>
      <p style={{ fontSize: '10px', wordBreak: 'break-all' }}>
        Token: {rawToken ? rawToken.substring(0, 20) + '...' : 'None'}
      </p>
      
      <div style={{ display: 'flex', flexWrap: 'wrap', gap: '4px' }}>
        <button 
          onClick={checkToken}
          style={{ 
            background: '#666', 
            color: 'white', 
            border: 'none', 
            padding: '4px 8px', 
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '11px'
          }}
        >
          Check Token
        </button>
        
        <button 
          onClick={testLocalStorage}
          style={{ 
            background: '#960', 
            color: 'white', 
            border: 'none', 
            padding: '4px 8px', 
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '11px'
          }}
        >
          Test Storage
        </button>
        
        <button 
          onClick={debugAuthToken}
          style={{ 
            background: '#609', 
            color: 'white', 
            border: 'none', 
            padding: '4px 8px', 
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '11px'
          }}
        >
          Debug Storage
        </button>
        
        <button 
          onClick={forceRefresh}
          style={{ 
            background: '#06d', 
            color: 'white', 
            border: 'none', 
            padding: '4px 8px', 
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '11px'
          }}
        >
          Force Refresh
        </button>
        
        {authHookAuthenticated && (
          <button 
            onClick={handleLogout}
            style={{ 
              background: '#d44', 
              color: 'white', 
              border: 'none', 
              padding: '4px 8px', 
              borderRadius: '4px',
              cursor: 'pointer',
              fontSize: '11px'
            }}
          >
            Logout
          </button>
        )}
      </div>
    </div>
  );
} 