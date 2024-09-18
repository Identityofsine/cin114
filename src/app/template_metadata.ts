import { BrandSettings } from "@/brand.settings";
import { Metadata } from "next";

export const TemplateMetadata: Metadata = {
	title: BrandSettings.brandName.displayShort + ' - Home',
	description: 'CIN-114 is a production company built by artists, for artists.',
	twitter: {
		title: BrandSettings.brandName.displayShort + ' - Home',
		description: 'CIN-114 is a production company built by artists, for artists.',
		card: 'summary',
		site: '@cin114',
	},
	openGraph: {
		title: BrandSettings.brandName.displayShort + ' - Home',
		description: 'CIN-114 is a production company built by artists, for artists.',
		type: 'website',
		locale: 'en_US',
		images: [
			{
				url: `${BrandSettings.url}/home/bottomstill.png`,
				width: 1920,
				height: 1080,
				alt: 'Collage',
			},
		]
	},

};

