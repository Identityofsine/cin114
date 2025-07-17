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
      Open in {provider}
    </div>
  )
}

export default ScreeningMapButton
