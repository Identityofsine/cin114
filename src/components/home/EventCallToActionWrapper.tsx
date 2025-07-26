'use client';

import { usePathname } from 'next/navigation';
import { useEffect } from 'react';
import { Event } from '@/types/event';
import EventCallToAction from './EventCallToAction';

type EventCallToActionWrapperProps = {
  event: Event | null;
}

function EventCallToActionWrapper({ event }: EventCallToActionWrapperProps) {
  const pathname = usePathname();
  
  const shouldShowBanner = event && !pathname.startsWith('/screenings/');
  
  useEffect(() => {
    // Add/remove body class based on banner visibility
    if (shouldShowBanner) {
      document.body.classList.add('has-event-banner');
    } else {
      document.body.classList.remove('has-event-banner');
    }
    
    // Cleanup function to remove class when component unmounts
    return () => {
      document.body.classList.remove('has-event-banner');
    };
  }, [shouldShowBanner]);
  
  if (!shouldShowBanner) {
    return null;
  }

  return <EventCallToAction event={event} />;
}

export default EventCallToActionWrapper; 