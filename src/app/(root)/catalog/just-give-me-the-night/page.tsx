import '@/app/styles/jgmtd.scss';
import { TemplateMetadata } from '@/app/template_metadata';
import { BrandSettings } from '@/brand.settings';
import { Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';


export function metadata() {
	return {
		...TemplateMetadata,
		title: 'Just Give Me The Night (2024)',
		description: Films[film].description,
		twitter: {
			...TemplateMetadata.twitter,
			title: 'Just Give Me The Night (2024)',
			description: Films[film].description,
		},
		openGraph: {
			...TemplateMetadata.openGraph,
			title: 'Just Give Me The Night (2024)',
			description: Films[film].description,
			images: [
				{
					url: `${BrandSettings.url}/film/just-give-me-tonight/boxart.png`,
					width: 1920,
					height: 1080,
					alt: 'Just Give Me The Night (2024)',
				},
				{
					url: `${BrandSettings.url}/film/just-give-me-tonight/image1.png`,
					width: 1920,
					height: 1080,
					alt: 'Just Give Me The Night (2024)',
				}
			],
		}
	};
}

const film = 'justgivemetonight'

export default function Page() {
	return (
		<FilmTemplate metadata={Films[film]}>
			<div className="jgmtd">
				<img className="logo" src="/film/just-give-me-tonight/logo.svg" alt="Just Give Me Tonight" />
				<div className="links flex">
					{(Films[film].links &&
						(
							(Films[film].links.youtube &&
								<img src="/ui/youtube.svg" alt="YouTube" />)
							||
							(Films[film].links.vimeo &&
								<img src="/ui/vimeo.svg" />))
					) || <p>Coming Soon</p>}
				</div>
			</div>
		</FilmTemplate>
	)
}
