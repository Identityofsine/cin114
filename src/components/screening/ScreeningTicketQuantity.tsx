import { useRef, useState, useEffect } from 'react';
import './styles/screeningticketquantity.scss';

type QuantityProps = {
  quantity: number;
  setQuantity: (quantity: number) => void;
}

function ScreeningTicketQuantity({
  quantity,
  setQuantity,
}: QuantityProps) {
  const inputRef = useRef<HTMLInputElement>(null);
  const [isInputFocused, setIsInputFocused] = useState(false);
  const [tempValue, setTempValue] = useState(quantity.toString());

  // Update temp value when quantity prop changes
  useEffect(() => {
    setTempValue(quantity.toString());
  }, [quantity]);

  const handleIncrement = () => {
    const newQuantity = Math.min(quantity + 1, 99); // Max 99 tickets
    setQuantity(newQuantity);
  };

  const handleDecrement = () => {
    const newQuantity = Math.max(quantity - 1, 1); // Min 1 ticket
    setQuantity(newQuantity);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    
    // Allow empty value temporarily while typing
    if (value === '') {
      setTempValue('');
      return;
    }
    
    // Only allow numbers
    if (!/^\d+$/.test(value)) return;
    
    const numValue = parseInt(value);
    
    // Don't allow values over 99
    if (numValue > 99) return;
    
    setTempValue(value);
  };

  const handleInputFocus = () => {
    setIsInputFocused(true);
  };

  const handleInputBlur = () => {
    setIsInputFocused(false);
    
    // Validate and set final value
    let finalValue = parseInt(tempValue) || 1;
    finalValue = Math.max(1, Math.min(finalValue, 99));
    
    setQuantity(finalValue);
    setTempValue(finalValue.toString());
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    // Allow backspace, delete, tab, escape, enter
    if ([8, 9, 27, 13, 46].indexOf(e.keyCode) !== -1 ||
        // Allow Ctrl+A, Ctrl+C, Ctrl+V, Ctrl+X
        (e.keyCode === 65 && e.ctrlKey === true) ||
        (e.keyCode === 67 && e.ctrlKey === true) ||
        (e.keyCode === 86 && e.ctrlKey === true) ||
        (e.keyCode === 88 && e.ctrlKey === true)) {
      return;
    }
    
    // Ensure that it is a number and stop the keypress
    if ((e.shiftKey || (e.keyCode < 48 || e.keyCode > 57)) && (e.keyCode < 96 || e.keyCode > 105)) {
      e.preventDefault();
    }
    
    // Handle arrow keys
    if (e.keyCode === 38) { // Up arrow
      e.preventDefault();
      handleIncrement();
    } else if (e.keyCode === 40) { // Down arrow
      e.preventDefault();
      handleDecrement();
    }
  };

  return (
    <div className="ticket_quantity">
      <div className="quantity_label">
        Quantity
      </div>
      <div className="quantity_controls">
        <button 
          type="button"
          className="quantity_btn quantity_btn--minus"
          onClick={handleDecrement}
          disabled={quantity <= 1}
          aria-label="Decrease quantity"
        >
          <span>−</span>
        </button>
        
        <div className={`quantity_input_container ${isInputFocused ? 'focused' : ''}`}>
          <input
            ref={inputRef}
            type="text"
            inputMode="numeric"
            pattern="[0-9]*"
            className="quantity_input"
            value={tempValue}
            onChange={handleInputChange}
            onFocus={handleInputFocus}
            onBlur={handleInputBlur}
            onKeyDown={handleKeyDown}
            min="1"
            max="99"
            aria-label={`Ticket quantity: ${quantity}`}
          />
        </div>
        
        <button 
          type="button"
          className="quantity_btn quantity_btn--plus"
          onClick={handleIncrement}
          disabled={quantity >= 99}
          aria-label="Increase quantity"
        >
          <span>+</span>
        </button>
      </div>
    </div>
  )
}

export default ScreeningTicketQuantity
