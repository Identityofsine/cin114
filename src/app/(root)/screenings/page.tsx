import { redirect } from 'next/navigation';
import { getActiveEvents } from '@/api';

export default async function ScreeningsPage() {
  const activeEvents = await getActiveEvents();
  
  if (activeEvents.length > 0) {
    // Redirect to the first (most active) event
    const mostActiveEvent = activeEvents[0];
    redirect(`/screenings/${mostActiveEvent.eventId}`);
  }
  
  // If no active events, show a message
  return (
    <section className="screenings-empty">
      <div className="screenings-empty__content">
        <h1>No Active Screenings</h1>
        <p>There are currently no active screenings available.</p>
        <p>Check back soon for upcoming events!</p>
      </div>
      <style jsx>{`
        .screenings-empty {
          min-height: 100vh;
          display: flex;
          align-items: center;
          justify-content: center;
          text-align: center;
          padding: 2rem;
        }
        
        .screenings-empty__content h1 {
          margin-bottom: 1rem;
          font-size: 2.5rem;
          font-weight: 600;
        }
        
        .screenings-empty__content p {
          margin-bottom: 0.5rem;
          font-size: 1.1rem;
          opacity: 0.8;
        }
      `}</style>
    </section>
  );
}
