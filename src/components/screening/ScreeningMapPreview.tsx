'use client'
import { EventLocation } from '@/types/event';
import './styles/screeningmappreview.scss'
import { MapContainer, MapContainerProps, TileLayer, TileLayerProps, Marker } from 'react-leaflet';

interface MapContainerPropsFixed extends MapContainerProps {
  center: [number, number];
  zoom: number;
  scrollWheelZoom?: boolean;
}

const MapContainerWrapper = MapContainer as unknown as (props: MapContainerPropsFixed) => JSX.Element;

interface TileLayerPropsFixed extends TileLayerProps {
  url: string;
  attribution: string;
}

const TileLayerWrapper = TileLayer as unknown as (props: TileLayerPropsFixed) => JSX.Element;

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
      <MapContainerWrapper className="screening-map-preview__map" center={[longitude, latitude]} zoom={24} scrollWheelZoom={false}>
        <TileLayerWrapper
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />
        <Marker position={[longitude, latitude]} />

      </MapContainerWrapper>
    </div>
  )
}

export default ScreeningMapPreview
