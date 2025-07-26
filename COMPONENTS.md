# UI Components and Authentication

This document describes the reusable UI components and authentication system added to the application.

## UI Components

### Input Component

A reusable input component that follows the app's styling patterns.

```tsx
import { Input } from '@/components/ui';

<Input
  label="Username"
  type="text"
  value={value}
  onChange={handleChange}
  error={error}
  placeholder="Enter username"
  variant="default" // or "outlined"
  helperText="Helper text"
/>
```

**Props:**
- `label?: string` - Label text displayed above the input
- `error?: string` - Error message displayed below the input
- `variant?: 'default' | 'outlined'` - Visual style variant
- `helperText?: string` - Helper text shown when no error
- All standard HTML input props are supported

### Button Component

A reusable button component with animated styling matching the app's design.

```tsx
import { Button } from '@/components/ui';

<Button
  variant="primary"
  size="large"
  isLoading={isLoading}
  fullWidth
  onClick={handleClick}
>
  Click Me
</Button>
```

**Props:**
- `variant?: 'primary' | 'secondary' | 'outline'` - Visual style
- `size?: 'small' | 'medium' | 'large'` - Button size
- `isLoading?: boolean` - Shows spinner and disables button
- `fullWidth?: boolean` - Makes button take full container width
- All standard HTML button props are supported

## Authentication System

### User Service

The user service handles authentication with the backend API.

```tsx
import { login, logout, getCurrentUser, isAuthenticated } from '@/api/services/user';

// Login
try {
  const response = await login({ username, password });
  console.log('Login successful:', response);
} catch (error) {
  console.error('Login failed:', error);
}

// Check if user is authenticated
const authenticated = isAuthenticated();

// Get current user
const user = await getCurrentUser();

// Logout
await logout();
```

**Available Functions:**
- `login(credentials: LoginCredentials): Promise<LoginResponse>`
- `logout(): Promise<void>`
- `getCurrentUser(): Promise<User | null>`
- `refreshToken(): Promise<AuthToken | null>`
- `isAuthenticated(): boolean`

### Route Guard

Protects routes that require authentication.

```tsx
import { RouteGuard } from '@/components/auth';

export default function ProtectedPage() {
  return (
    <RouteGuard
      fallback={<div>Checking authentication...</div>}
      redirectTo="/zcpkcrucpw"
    >
      <ProtectedContent />
    </RouteGuard>
  );
}
```

**Props:**
- `children: React.ReactNode` - Content to show when authenticated
- `fallback?: React.ReactNode` - Loading component (default: "Redirecting...")
- `redirectTo?: string` - Redirect path for unauthenticated users (default: "/zcpkcrucpw")

### Auth Hook

Use the `useAuth` hook to get current authentication state.

```tsx
import { useAuth } from '@/components/auth';

function MyComponent() {
  const { isLoading, isAuthenticated, user } = useAuth();

  if (isLoading) return <div>Loading...</div>;
  if (!isAuthenticated) return <div>Please log in</div>;
  
  return <div>Welcome, {user?.username}!</div>;
}
```

### Login Form

A complete login form component.

```tsx
import { LoginForm } from '@/components/auth';

export default function LoginPage() {
  return (
    <LoginForm 
      onSuccess={() => console.log('Login successful')}
      redirectTo="/dashboard"
    />
  );
}
```

**Props:**
- `onSuccess?: () => void` - Callback when login succeeds
- `redirectTo?: string` - Where to redirect after login (default: "/")

## API Endpoints

The authentication system expects these backend endpoints:

- `POST /api/v1/auth/login/internal` - Login with username/password
- `POST /api/v1/auth/logout` - Logout (optional)
- `GET /api/v1/user/me` - Get current user info
- `POST /api/v1/auth/refresh` - Refresh auth token

## Usage Examples

### Simple Login Page

```tsx
// src/app/(root)/zcpkcrucpw/page.tsx
import { LoginForm } from '@/components/auth';

export default function LoginPage() {
  return <LoginForm redirectTo="/dashboard" />;
}
```

### Protected Dashboard

```tsx
// src/app/(root)/dashboard/page.tsx
import { RouteGuard } from '@/components/auth';

export default function DashboardPage() {
  return (
    <RouteGuard>
      <DashboardContent />
    </RouteGuard>
  );
}
```

### Custom Form with UI Components

```tsx
import { Input, Button } from '@/components/ui';
import { useState } from 'react';

function CustomForm() {
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');

  return (
    <form>
      <Input
        label="Email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        error={error}
        placeholder="Enter your email"
      />
      <Button type="submit" variant="primary" size="large">
        Submit
      </Button>
    </form>
  );
}
```

## Styling

All components follow the app's existing design system:
- Dark theme with gray color palette
- Inter font family for inputs and forms
- Rubik font for buttons
- Responsive breakpoints
- Consistent border radius and spacing
- Animated button effects matching existing patterns
- **All styling is properly organized in SCSS files** - no inline styles

The components are fully responsive and work across all device sizes. 