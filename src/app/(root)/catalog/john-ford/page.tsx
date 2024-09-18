import '@/app/styles/johnford.scss';
import { TemplateMetadata } from '@/app/template_metadata';
import { BrandSettings } from '@/brand.settings';
import { Films } from '@/film.settings';
import FilmTemplate from '@/template/Film';

export function metadata() {
	return {
		...TemplateMetadata,
		title: 'Interviewing John Ford (2024)',
		description: Films.johnford.description,
		twitter: {
			...TemplateMetadata.twitter,
			title: 'Interviewing John Ford (2024)',
			description: Films.johnford.description,
		},
		openGraph: {
			...TemplateMetadata.openGraph,
			title: 'Interviewing John Ford (2024)',
			description: Films.johnford.description,
			images: [
				{
					url: `${BrandSettings.url}/film/john-ford/boxart.png`,
					width: 1920,
					height: 1080,
					alt: 'Interviewing John Ford (2024)',
				},
				{
					url: `${BrandSettings.url}/film/john-ford/image1.png`,
					width: 1920,
					height: 1080,
					alt: 'Interviewing John Ford (2024)',
				}
			],
		}
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
								<a href={Films.johnford.links.youtube} ><img src="/ui/youtube.svg" alt="YouTube" /></a>)
							||
							(Films.johnford.links.vimeo &&
								<img src="/ui/vimeo.svg" alt="Vimeo" />)
						)
					}
				</div>
			</div>
		</FilmTemplate>
	)
}
