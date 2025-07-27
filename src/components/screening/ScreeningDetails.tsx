'use client';
import { useCallback, useMemo } from 'react';
import './styles/screeningdetails.scss';
import { Event } from "@/types/event"
import { displayDateForScreening } from '@/util/date';
import ScreeningMapButton, { MAP_PROVIDERS, MapProvider } from './ScreeningMapButton';

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

  console.log(location)

  const openMap = useCallback((provider: MapProvider) => {
    if (!location) {
      console.warn('Location not available');
      return;
    }

    // Prefer address when available, fallback to coordinates
    let query: string;
    if (location.address?.trim()) {
      query = encodeURIComponent(location.address);
    } else if (location.latitude && location.longitude) {
      query = `${location.latitude},${location.longitude}`;
    } else {
      console.warn('Neither address nor coordinates available for location');
      return;
    }

    let mapUrl: string;

    switch (provider) {
      case MAP_PROVIDERS.GOOGLE:
        mapUrl = `https://www.google.com/maps?q=${query}`;
        break;
      case MAP_PROVIDERS.APPLE:
        mapUrl = `https://maps.apple.com/?q=${query}`;
        break;
      default:
        console.warn('Unknown map provider:', provider);
        return;
    }

    window.open(mapUrl, '_blank', 'noopener,noreferrer');
  }, [location]);

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
      <div className="flex column screening-details-buttons">
        {Object.values(MAP_PROVIDERS).map((provider) => (
          <ScreeningMapButton
            key={provider}
            provider={provider as MapProvider}
            onClick={openMap}
          />
        ))}
      </div>
    </div>
  )
}

export default ScreeningDetails
