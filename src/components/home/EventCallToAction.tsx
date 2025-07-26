'use client';

import Link from 'next/link';
import { Event } from '@/types/event';
import './styles/eventcalltoaction.scss';

type EventCallToActionProps = {
  event: Event;
}

function EventCallToAction({ event }: EventCallToActionProps) {
  if (!event.eventId) return null;

  return (
    <div className="event_cta_banner">
      <Link href={`/screenings/${event.eventId}`} className="event_cta_link">
        <div className="event_cta_content">
          <div className="event_cta_indicator">
            <div className="pulse_dot"></div>
            <span>NOW SHOWING</span>
          </div>
          <div className="event_cta_text">
            {event.shortDescription || event.description || 'Special Screening Available'}
          </div>
          <div className="event_cta_arrow">→</div>
        </div>
      </Link>
    </div>
  );
}

export default EventCallToAction;
