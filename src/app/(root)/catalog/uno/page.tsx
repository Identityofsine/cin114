import '@/app/styles/uno.scss';
import { Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';
import { VideoMetadata } from '@/types/video';

export function metadata() {
	return {
		title: 'UNO (2024)',
		description: Films.uno.description,
	};
}

const film = 'uno'

export default function JohnFord() {

	return (
		<FilmTemplate metadata={Films[film]}>
			<div className="uno">
				<img className="logo" src="/film/uno/logo.svg" alt="UNO" />
				<div className="links flex">
					{Films[film].links &&
						(
							(Films[film].links.youtube &&
								<img src="/ui/youtube.svg" alt="YouTube" />)
							||
							(Films[film].links.vimeo &&
								<img src="/ui/vimeo.svg" alt="Vimeo" />))
					}
				</div>
			</div>
		</FilmTemplate>
	)
}
