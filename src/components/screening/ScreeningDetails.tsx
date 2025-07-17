import { useMemo } from 'react';
import './styles/screeningdetails.scss';
import { Event } from "@/types/event"
import { displayDateForScreening } from '@/util/date';

type ScreeningDetailsProps = {
  event: Event
}

function ScreeningDetails({ event }: ScreeningDetailsProps) {

  const description = useMemo(() => event?.description?.replaceAll('\\n', ' - '), [event.description]);


  const location = useMemo(() => {
    if (!event.locations || event.locations.length === 0) {
      return;
    }
    return event.locations?.[0];
  }, [event.locations]);

  return (
    <div className="screening-details">
      <div className="flex column screening-details-header">
        <h2>Now Showing</h2>
        <p>{description}</p>
      </div>
      <div className="flex column screening-details-info">
        <h3>{location?.locationName}</h3>
        {event.expirationDate && (
          <p>Screening at <strong>{displayDateForScreening(event.expirationDate)}</strong></p>
        )}
      </div>
    </div>
  )
}

export default ScreeningDetails
