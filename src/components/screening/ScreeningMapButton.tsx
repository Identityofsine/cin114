'use client'

import Image from 'next/image';
import './styles/screeningmapbutton.scss';

export const MAP_PROVIDERS = {
  GOOGLE: 'Google Maps',
  APPLE: 'Apple Maps',
} as const;

export type MapProvider = (typeof MAP_PROVIDERS)[keyof typeof MAP_PROVIDERS];


type ScreeningMapButtonProps = {
  provider?: MapProvider;
  onClick?: (provider: MapProvider) => void;
}

function ScreeningMapButton({
  provider = MAP_PROVIDERS.GOOGLE,
  onClick = () => { },
}: ScreeningMapButtonProps) {
  return (
    <div className="screening-map-button" onClick={() => onClick(provider)}>
      <div className="screening-map-button__icon">
        <Image
          src={`/ui/map-${provider.toLowerCase().replace(' ', '-')}.svg`} alt={`${provider} icon`}
          width="37"
          height="37"
        />
      </div>
      <span>
        Open in {provider}
      </span>
    </div>
  )
}

export default ScreeningMapButton
