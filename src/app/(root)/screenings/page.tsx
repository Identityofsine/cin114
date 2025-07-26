import { redirect } from 'next/navigation';
import '@/components/screening/styles/screening.scss';
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

    </section>
  );
}
