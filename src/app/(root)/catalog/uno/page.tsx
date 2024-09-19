import '@/app/styles/uno.scss';
import { TemplateMetadata } from '@/app/template_metadata';
import { BrandSettings } from '@/brand.settings';
import { Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';

export function metadata() {
	return {
		...TemplateMetadata,
		title: 'UNO (2024)',
		description: Films.uno.description,
		twitter: {
			...TemplateMetadata.twitter,
			title: 'UNO (2024)',
			description: Films.uno.description,
		},
		openGraph: {
			...TemplateMetadata.openGraph,
			title: 'UNO (2024)',
			description: Films.uno.description,
			images: [
				{
					url: `${BrandSettings.url}/film/uno/boxart.png`,
					width: 1920,
					height: 1080,
					alt: 'UNO (2024)',
				},
				{
					url: `${BrandSettings.url}/film/uno/image1.png`,
					width: 1920,
					height: 1080,
					alt: 'UNO (2024)',
				}
			],
		}
	};
}

const film = 'uno'

export default function UNO() {

	return (
		<FilmTemplate metadata={Films[film]}>
			<div className="uno">
				<img className="logo" src="/film/uno/logo.svg" alt="UNO" />
				<div className="links flex">
					{Films[film].links &&
						(
							(Films[film].links.youtube &&
								<a href={Films[film].links.youtube}><img src="/ui/youtube.svg" alt="YouTube" /></a>)
							||
							(Films[film].links.vimeo &&
								<img src="/ui/vimeo.svg" alt="Vimeo" />))
					}
				</div>
			</div>
		</FilmTemplate>
	)
}
