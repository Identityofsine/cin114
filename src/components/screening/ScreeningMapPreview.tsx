'use client'
import { EventLocation } from '@/types/event';
import './styles/screeningmappreview.scss'
import { MapContainer, TileLayer, Marker } from 'react-leaflet';


type ScreeningMapProps = {
  location?: EventLocation
}

function ScreeningMapPreview({ location }: ScreeningMapProps) {

  const { latitude, longitude } = location || {};

  if (!latitude || !longitude) {
    return (
      <div className="screening-map-preview">
        <div className="screening-map-preview__error">
          <p>Sorry, Location not available</p>
        </div>
      </div>
    )
  }

  return (
    <div className="screening-map-preview">
      <MapContainer className="screening-map-preview__map" center={[longitude, latitude]} zoom={24} scrollWheelZoom={false}>
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        <Marker position={[longitude, latitude]} />

      </MapContainer>
    </div>
  )
}

export default ScreeningMapPreview
