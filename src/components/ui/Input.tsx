'use client';

import React, { forwardRef } from 'react';
import './styles/input.scss';

export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  variant?: 'default' | 'outlined';
  helperText?: string;
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, variant = 'default', helperText, className, ...props }, ref) => {
    const inputClasses = [
      'input',
      `input--${variant}`,
      error && 'input--error',
      className
    ].filter(Boolean).join(' ');

    return (
      <div className="input-wrapper">
        {label && (
          <label className="input-label" htmlFor={props.id}>
            {label}
          </label>
        )}
        <input
          ref={ref}
          className={inputClasses}
          {...props}
        />
        {error && (
          <span className="input-error">{error}</span>
        )}
        {helperText && !error && (
          <span className="input-helper">{helperText}</span>
        )}
      </div>
    );
  }
);

Input.displayName = 'Input';

export default Input; 