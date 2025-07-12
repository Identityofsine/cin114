'use client';

import { useState } from 'react';
import './styles/calltoaction.scss';
import { createCheckout } from '@/api';

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

  //HOUR(MINUTE?) MONTH DAY
  const displayDate = (date: Date) => {
    const options: Intl.DateTimeFormatOptions = {
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: date?.getMinutes() > 0 ? '2-digit' : undefined,
    };
    return date.toLocaleString('en-US', options);
  }

  return (
    <div className={`screening_call_to_action ${isLoading ? 'loading' : ''}`} onClick={onClick}>
      <div className="button">
        PURCHASE TICKET
      </div>
      <div className="expiration">
        <span>{displayDate(expirationDate)}</span>
      </div>
    </div>
  )
}

export default ScreeningCallToAction
