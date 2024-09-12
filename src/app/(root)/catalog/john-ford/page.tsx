import '@/app/styles/johnford.scss';
import { Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';
import { VideoMetadata } from '@/types/video';

export function metadata() {
	return {
		title: 'Interviewing John Ford (2024)',
		description: 'A collection of films by John Ford',
	};
}

export default function JohnFord() {

	return (
		<FilmTemplate metadata={Films.johnford}>
			<div className="johnford">
				<img className="logo" src="/film/john-ford/logo.svg" alt="John Ford" />
				<div className="links flex">
					{Films.johnford.links &&
						(
							(Films.johnford.links.youtube &&
								<img src="/ui/youtube.svg" alt="YouTube" />)
							||
							(Films.johnford.links.vimeo &&
								<img src="/ui/vimeo.svg" alt="Vimeo" />))
					}
				</div>
			</div>
		</FilmTemplate>
	)
}
