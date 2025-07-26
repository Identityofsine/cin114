'use client';
import React, { useState, useEffect } from 'react';
import { RouteGuard } from '@/components/auth';
import TicketsTable from '@/components/tickets/table/TicketsTable';
import { getAllTickets } from '@/api/services/tickets';
import { Ticket } from '@/types/ticket';

function TicketsPage() {
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadTickets = async () => {
      try {
        setIsLoading(true);
        setError(null);
        const ticketsData = await getAllTickets();
        setTickets(ticketsData);
      } catch (err) {
        console.error('Error loading tickets:', err);
        setError('Failed to load tickets. Please try again later.');
      } finally {
        setIsLoading(false);
      }
    };

    loadTickets();
  }, []);

  if (error) {
    return (
      <RouteGuard>
        <div style={{
          minHeight: '100vh',
          backgroundColor: '#0E0E0E',
          color: '#ffffff',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          flexDirection: 'column',
          padding: '2rem'
        }}>
          <div style={{
            textAlign: 'center',
            background: 'linear-gradient(145deg, #2c2c2c 0%, #161616 100%)',
            padding: '3rem',
            borderRadius: '20px',
            boxShadow: '0 20px 60px rgba(0, 0, 0, 0.4)',
          }}>
            <div style={{ fontSize: '4rem', marginBottom: '1rem' }}>⚠️</div>
            <h2 style={{ 
              fontSize: '1.5rem', 
              marginBottom: '1rem',
              color: '#ffffff'
            }}>
              Error Loading Tickets
            </h2>
            <p style={{ 
              color: '#a6a6a6',
              marginBottom: '2rem'
            }}>
              {error}
            </p>
            <button 
              onClick={() => window.location.reload()}
              style={{
                background: 'linear-gradient(135deg, #ffffff 0%, #e6e6e6 100%)',
                color: '#000000',
                border: 'none',
                padding: '0.75rem 1.5rem',
                borderRadius: '8px',
                fontWeight: '600',
                cursor: 'pointer',
                fontSize: '0.9rem',
                textTransform: 'uppercase',
                letterSpacing: '0.5px'
              }}
            >
              Try Again
            </button>
          </div>
        </div>
      </RouteGuard>
    );
  }

  return (
    <RouteGuard>
      <TicketsTable tickets={tickets} isLoading={isLoading} />
    </RouteGuard>
  );
}

export default TicketsPage; 