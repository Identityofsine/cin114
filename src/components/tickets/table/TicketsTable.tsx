'use client';
import React, { useState, useMemo } from 'react';
import { Ticket } from '@/types/ticket';
import { formatTicketId, parseTicketId } from '@/api/services/tickets';
import '../styles/table/tickets-table.scss';

type TicketsTableProps = {
  tickets?: Ticket[];
  isLoading?: boolean;
}

function TicketsTable({ tickets = [], isLoading = false }: TicketsTableProps) {
  const [searchQuery, setSearchQuery] = useState('');

  // Filter tickets based on search query (ticket ID only)
  const filteredTickets = useMemo(() => {
    if (!searchQuery.trim()) return tickets;

    const query = searchQuery.toLowerCase().trim();
    
    return tickets.filter(ticket => {
      const formattedId = formatTicketId(ticket.ticket_id).toLowerCase();
      const ticketId = ticket.ticket_id.toString();
      
      return (
        formattedId.includes(query) ||
        ticketId.includes(query)
      );
    });
  }, [tickets, searchQuery]);

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  const truncateText = (text: string, maxLength: number = 60) => {
    if (text.length <= maxLength) return text;
    return text.substring(0, maxLength) + '...';
  };

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchQuery(e.target.value);
  };

  const clearSearch = () => {
    setSearchQuery('');
  };

  if (isLoading) {
    return (
      <div className="tickets-container">
        <div className="tickets-header">
          <h1>Tickets</h1>
          <p className="tickets-subtitle">Loading tickets...</p>
        </div>
        <div className="tickets-table-wrapper">
          <div className="loading-state">
            <div className="loading-spinner"></div>
            <div className="loading-message">Loading tickets...</div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="tickets-container">
      <div className="tickets-header">
        <h1>Tickets</h1>
        <p className="tickets-subtitle">Search tickets by ID</p>
      </div>

      {/* Search Box */}
      <div className="search-section">
        <div className="search-box">
          <div className="search-input-wrapper">
            <input
              type="text"
              placeholder="Search by ticket ID (T-123)..."
              value={searchQuery}
              onChange={handleSearchChange}
              className="search-input"
            />
            {searchQuery && (
              <button onClick={clearSearch} className="clear-search-btn">
                ✕
              </button>
            )}
          </div>
          <div className="search-icon">🔍</div>
        </div>
        
        {searchQuery && (
          <div className="search-results-info">
            Found {filteredTickets.length} ticket{filteredTickets.length !== 1 ? 's' : ''} 
            {searchQuery && ` matching ticket ID "${searchQuery}"`}
          </div>
        )}
      </div>
      
      <div className="tickets-table-wrapper">
        {filteredTickets.length === 0 ? (
          <div className="empty-state">
            {searchQuery ? (
              <>
                <div className="empty-icon">🔍</div>
                <div className="empty-message">No tickets found</div>
                <div className="empty-subtitle">
                  Try searching with a different ticket ID or{' '}
                  <button onClick={clearSearch} className="clear-search-link">
                    clear your search
                  </button>
                </div>
              </>
            ) : (
              <>
                <div className="empty-icon">🎫</div>
                <div className="empty-message">No tickets available</div>
                <div className="empty-subtitle">
                  Tickets will appear here once they are purchased
                </div>
              </>
            )}
          </div>
        ) : (
          <table className="tickets-table">
            <thead>
              <tr>
                <th>Ticket ID</th>
                <th>Event</th>
                <th>Customer Email</th>
                <th>Purchase Date</th>
                <th>Last Updated</th>
              </tr>
            </thead>
            <tbody>
              {filteredTickets.map(ticket => (
                <tr key={ticket.ticket_id}>
                  <td>
                    <span className="ticket-id">{formatTicketId(ticket.ticket_id)}</span>
                  </td>
                  <td>
                    <div className="event-info">
                      <span className="event-title" title={ticket.event_description}>
                        {truncateText(ticket.event_description)}
                      </span>
                      <span className="event-id-small">Event #{ticket.event_id}</span>
                    </div>
                  </td>
                  <td>
                    <span className="customer-email">
                      {ticket.stripe_receipt_email || 'N/A'}
                    </span>
                  </td>
                  <td>
                    <span className="purchase-date">
                      {formatDate(ticket.created_at)}
                    </span>
                  </td>
                  <td>
                    <span className="updated-date">
                      {formatDate(ticket.updated_at)}
                    </span>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
}

export default TicketsTable; 