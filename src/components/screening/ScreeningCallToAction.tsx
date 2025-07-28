'use client';

import { useState } from 'react';
import './styles/calltoaction.scss';
import { createCheckout } from '@/api';
import { displayDateForScreening } from '@/util/date';
import ScreeningTicketQuantity from './ScreeningTicketQuantity';

type ScreeningCallToActionProps = {
  eventId: number;
  expirationDate: Date;
}

function ScreeningCallToAction({
  eventId,
  expirationDate,
}: ScreeningCallToActionProps) {

  const [isLoading, setIsLoading] = useState(false);
  const [quantity, setQuantity] = useState(1);


  const onClick = () => {
    setIsLoading(true);
    createCheckout(eventId, quantity).then((checkout) => {
      window.location.href = checkout.checkoutUrl;
    }).finally(() => {
      setIsLoading(false);
    });
  }



  const handleQuantityContainerClick = (e: React.MouseEvent) => {
    e.stopPropagation();
  };

  return (
    <div className={`screening_call_to_action ${isLoading ? 'loading' : ''}`}>
      <div className="button_container">
        <div className="button" onClick={onClick}>
          PURCHASE TICKETS
        </div>
        <div onClick={handleQuantityContainerClick}>
          <ScreeningTicketQuantity
            quantity={quantity}
            setQuantity={setQuantity}
          />
        </div>
      </div>
      <div className="expiration">
        <span>{displayDateForScreening(expirationDate)}</span>
      </div>
    </div>
  )
}

export default ScreeningCallToAction
