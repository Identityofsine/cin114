import '@/app/styles/imaginary.scss';
import { FilmIndex, Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';
import { VideoMetadata } from '@/types/video';

export function metadata() {
	return {
		title: 'Imaginary Rules of Engagement',
		description: Films.uno.description,
	};
}

const film: FilmIndex = 'theimaginaryrulesofengagement'

export default function TIRE() {

	return (
		<FilmTemplate metadata={Films[film]}>
			<div className="tire">
				<img className="logo" src="/film/the-rules-of-engagement/logo.svg" alt="The Imaginary Rules of Engagement" />
				<div className="links flex">

					{
						Films[film]?.links && Films[film].links.youtube &&
						<img src="/ui/youtube.svg" alt="YouTube" />
					}
					{
						Films[film]?.links && Films[film].links.vimeo &&
						<img src="/ui/vimeo.svg" alt="Vimeo" />
					}
				</div>
			</div>
		</FilmTemplate>
	)
}
