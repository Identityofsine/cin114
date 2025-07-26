'use client';

import { useEffect, useState } from 'react';
import { usePathname } from 'next/navigation';
import { createPortal } from 'react-dom';
import { Event } from '@/types/event';
import { getActiveEvents } from '@/api';
import EventCallToAction from './EventCallToAction';

function EventCallToActionClient() {
  const [event, setEvent] = useState<Event | null>(null);
  const [mounted, setMounted] = useState(false);
  const pathname = usePathname();

  useEffect(() => {
    setMounted(true);
    
    // Fetch events client-side
    getActiveEvents().then(activeEvents => {
      if (activeEvents.length > 0) {
        setEvent(activeEvents[0]);
      }
    }).catch(error => {
      console.warn('Failed to fetch active events:', error);
    });
  }, []);

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

  if (!mounted || !shouldShowBanner) {
    return null;
  }

  const bannerRoot = document.getElementById('event-banner-root');
  if (!bannerRoot) {
    return null;
  }

  return createPortal(
    <EventCallToAction event={event} />,
    bannerRoot
  );
}

export default EventCallToActionClient; 