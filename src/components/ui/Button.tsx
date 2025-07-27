'use client';

import React, { forwardRef } from 'react';
import './styles/button.scss';

export interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'outline';
  size?: 'small' | 'medium' | 'large';
  isLoading?: boolean;
  fullWidth?: boolean;
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ 
    children, 
    variant = 'primary', 
    size = 'medium', 
    isLoading = false,
    fullWidth = false,
    className, 
    disabled,
    ...props 
  }, ref) => {
    const buttonClasses = [
      'btn',
      `btn--${variant}`,
      `btn--${size}`,
      fullWidth && 'btn--full-width',
      isLoading && 'btn--loading',
      className
    ].filter(Boolean).join(' ');

    return (
      <button
        ref={ref}
        className={buttonClasses}
        disabled={disabled || isLoading}
        {...props}
      >
        <span className="btn__content">
          {isLoading && <span className="btn__spinner" />}
          {children}
        </span>
      </button>
    );
  }
);

Button.displayName = 'Button';

export default Button; 