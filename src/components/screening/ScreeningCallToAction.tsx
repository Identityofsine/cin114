import { useState } from 'react';
import './styles/calltoaction.scss';
import { createCheckout } from '@/api';

type ScreeningCallToActionProps = {
  eventId: number;
  expirationDate: string;
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

    </div>
  )
}

export default ScreeningCallToAction
