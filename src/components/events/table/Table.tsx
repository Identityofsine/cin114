'use client';
import { Event } from '@/types/event';
import React from 'react';
import { useRouter } from 'next/navigation';
import '../styles/table/table.scss';

type EventTableProps = {
  events?: Event[]
}

function EventTable({ events = [] }: EventTableProps) {
  const router = useRouter();

  const handleTicketPurchase = (eventId: number) => {
    router.push(`/events/tickets/${eventId}`);
  };

  const formatDate = (date: Date | string | undefined) => {
    if (!date) return 'TBD';
    const eventDate = new Date(date);
    return eventDate.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  const truncateDescription = (description: string, maxLength: number = 80) => {
    if (description.length <= maxLength) return description;
    return description.substring(0, maxLength) + '...';
  };

  if (events.length === 0) {
    return (
      <div className="events-container">
        <div className="events-header">
          <h1>Events</h1>
          <p className="events-subtitle">our upcoming cinema experiences</p>
        </div>
        <div className="events-table-wrapper">
          <div className="empty-state">
            <div className="empty-icon">🎬</div>
            <div className="empty-message">No events scheduled</div>
            <div className="empty-subtitle">Check back soon for upcoming screenings</div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="events-container">
      <div className="events-header">
        <h1>Events</h1>
        <p className="events-subtitle">Discover our upcoming cinema experiences</p>
      </div>
      
      <div className="events-table-wrapper">
        <table className="events-table">
          <thead>
            <tr>
              <th>Event ID</th>
              <th>Title</th>
              <th>Description</th>
              <th>Event Date</th>
              <th>Tickets</th>
            </tr>
          </thead>
          <tbody>
            {events.map(event => (
              <tr 
                key={event.eventId || Math.random()} 
                onClick={() => event.eventId && handleTicketPurchase(event.eventId)}
              >
                <td>
                  <span className="event-id">#{event.eventId || 'N/A'}</span>
                </td>
                <td>
                  <span className="event-title">{event.shortDescription || 'Untitled Event'}</span>
                </td>
                <td>
                  <span className="event-description" title={event.description}>
                    {truncateDescription(event.description || 'No description available')}
                  </span>
                </td>
                <td>
                  <span className="event-date">
                    {formatDate(event.expirationDate)}
                  </span>
                </td>
                <td>
                  <button 
                    className="ticket-btn"
                    onClick={(e) => {
                      e.stopPropagation();
                      if (event.eventId) {
                        handleTicketPurchase(event.eventId);
                      }
                    }}
                    disabled={!event.eventId}
                  >
                    Buy Tickets
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default EventTable
