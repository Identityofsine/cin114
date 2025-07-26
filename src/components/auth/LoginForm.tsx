'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Input, Button } from '@/components/ui';
import { login, isAuthenticated } from '@/api/services/user';
import { LoginCredentials, AuthError } from '@/types/user';
import { Storage, StorageLks } from '@/services/storage';
import DebugProvider from './DebugProvider';
import './styles/loginform.scss';

interface LoginFormProps {
  onSuccess?: () => void;
  redirectTo?: string;
}

const storage = new Storage();

export default function LoginForm({ onSuccess, redirectTo = '/' }: LoginFormProps) {
  const router = useRouter();
  const [credentials, setCredentials] = useState<LoginCredentials>({
    username: '',
    password: ''
  });
  const [errors, setErrors] = useState<Partial<LoginCredentials & { general: string }>>({});
  const [isLoading, setIsLoading] = useState(false);

  const validateForm = (): boolean => {
    const newErrors: Partial<LoginCredentials & { general: string }> = {};

    if (!credentials.username.trim()) {
      newErrors.username = 'Username is required';
    }

    if (!credentials.password.trim()) {
      newErrors.password = 'Password is required';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validateForm()) {
      return;
    }

    setIsLoading(true);
    setErrors({});

    try {
      const response = await login(credentials);
      
      // Add a small delay and check if token is actually stored
      setTimeout(() => {
        const storedToken = storage.safeGetItem(StorageLks.AUTH);
        const authCheck = isAuthenticated();
        
        if (storedToken && authCheck) {
          if (onSuccess) {
            onSuccess();
          } else {
            router.push(redirectTo);
          }
        } else {
          setErrors({
            general: 'Login succeeded but authentication state is invalid. Please try again.'
          });
        }
      }, 200); // Give a bit more time for storage
      
    } catch (error) {
      const authError = error as AuthError;
      setErrors({
        general: authError.message || 'Login failed. Please try again.'
      });
    } finally {
      setIsLoading(false);
    }
  };

  const handleInputChange = (field: keyof LoginCredentials) => (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setCredentials(prev => ({
      ...prev,
      [field]: e.target.value
    }));
    
    // Clear error when user starts typing
    if (errors[field]) {
      setErrors(prev => ({
        ...prev,
        [field]: undefined
      }));
    }
  };

  return (
    <>
      <DebugProvider />
      <div className="login-form">
        <div className="login-form__container">
          <div className="login-form__header">
            <h1>Sign In</h1>
            <p>Welcome back! Please sign in to your account.</p>
          </div>

          <form onSubmit={handleSubmit} className="login-form__form">
            {errors.general && (
              <div className="login-form__error">
                {errors.general}
              </div>
            )}

            <Input
              id="username"
              label="Username"
              type="text"
              value={credentials.username}
              onChange={handleInputChange('username')}
              error={errors.username}
              placeholder="Enter your username"
              autoComplete="username"
              disabled={isLoading}
            />

            <Input
              id="password"
              label="Password"
              type="password"
              value={credentials.password}
              onChange={handleInputChange('password')}
              error={errors.password}
              placeholder="Enter your password"
              autoComplete="current-password"
              disabled={isLoading}
            />

            <Button
              type="submit"
              variant="primary"
              size="large"
              fullWidth
              isLoading={isLoading}
              disabled={isLoading}
            >
              {isLoading ? 'Signing In...' : 'Sign In'}
            </Button>
          </form>
        </div>
      </div>
    </>
  );
} 