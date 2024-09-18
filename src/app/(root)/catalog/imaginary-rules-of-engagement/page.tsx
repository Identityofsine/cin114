import '@/app/styles/imaginary.scss';
import { TemplateMetadata } from '@/app/template_metadata';
import { BrandSettings } from '@/brand.settings';
import { FilmIndex, Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';

export function metadata() {
	return {
		...TemplateMetadata,
		title: 'Imaginary Rules of Engagement',
		description: Films.uno.description,
		twitter: {
			...TemplateMetadata.twitter,
			title: 'Imaginary Rules of Engagement',
			description: Films.uno.description,
		},
		openGraph: {
			...TemplateMetadata.openGraph,
			title: 'Imaginary Rules of Engagement',
			description: Films.uno.description,
			images: [
				{
					url: `${BrandSettings.url}/film/the-imaginary-rules-of-engagement/boxart.png`,
					width: 1920,
					height: 1080,
					alt: 'Imaginary Rules of Engagement',
				},
				{
					url: `${BrandSettings.url}/film/the-imaginary-rules-of-engagement/image1.png`,
					width: 1920,
					height: 1080,
					alt: 'Imaginary Rules of Engagement',
				}
			],
		}
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
