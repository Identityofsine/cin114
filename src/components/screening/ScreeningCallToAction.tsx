'use client';

import { useState } from 'react';
import './styles/calltoaction.scss';
import { createCheckout } from '@/api';
import { displayDateForScreening } from '@/util/date';

type ScreeningCallToActionProps = {
  eventId: number;
  expirationDate: Date;
}

function ScreeningCallToAction({
  eventId,
  expirationDate,
}: ScreeningCallToActionProps) {

  const [isLoading, setIsLoading] = useState(false);


  const onClick = () => {
    setIsLoading(true);
    createCheckout(eventId).then((checkout) => {
      window.location.href = checkout.checkoutUrl;
    }).finally(() => {
      setIsLoading(false);
    });
  }



  return (
    <div className={`screening_call_to_action ${isLoading ? 'loading' : ''}`} onClick={onClick}>
      <div className="button">
        PURCHASE TICKETS
      </div>
      <div className="expiration">
        <span>{displayDateForScreening(expirationDate)}</span>
      </div>
    </div>
  )
}

export default ScreeningCallToAction
