import '@/app/styles/jgmtd.scss';
import { Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';


export function metadata() {
	return {
		title: 'Just Give Me The Night (2024)',
		description: Films[film].description,
	};
}

const film = 'justgivemetonight'

export default function Page() {
	return (
		<FilmTemplate metadata={Films[film]}>
			<div className="jgmtd">
				<img className="logo" src="/film/just-give-me-tonight/logo.svg" alt="Just Give Me Tonight" />
				<div className="links flex">
					{Films[film].links &&
						(
							(Films[film].links.youtube &&
								<img src="/ui/youtube.svg" alt="YouTube" />)
							||
							(Films[film].links.vimeo &&
								<img src="/ui/vimeo.svg" />))
					}
				</div>
			</div>
		</FilmTemplate>
	)
}
