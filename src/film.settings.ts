import { VideoMetadata } from "./types/video";

type FilmIndex = 'johnford' | 'uno' | 'justgivemetonight' | 'uno' | 'theimaginaryrulesofengagement'

export const Films: Record<FilmIndex, VideoMetadata> = {
	'johnford': {
		title: 'Interviewing John Ford',
		description: 'A collection of films by John Ford',
		boxart: {
			title: '',
			caption: 'Ask the right questions...',
			img: '/film/john-ford/boxart.png',
		},
		links: {
			vimeo: 'https://vimeo.com/123456789',
			youtube: 'https://www.youtube.com/watch?v=J6hMl5Y1BkE',
		},
		credits: [
			{
				role: 'Director',
				name: 'Kai Luckey',
			},
			{
				role: 'A.Director',
				name: 'Noah Fields'
			},
			{
				role: 'Writers',
				name: 'Kai Luckey'
			},
			{
				role: 'Cinemotographer',
				name: 'Kai Luckey'
			},
			{
				role: 'Producer',
				name: 'Erin Hennig, Kai Luckey, Noah Fields'
			},
			{
				role: 'Producer',
				name: 'Erin Hennig, Kai Luckey, Noah Fields'
			}
		],
		url: '/catalog/john-ford',
		img: '/film/john-ford/image1.png',
		date: '02/15/2024',
	},
	'uno': {
		title: 'UNO',
		description: 'Three friends play a game of UNO',
		boxart: {
			title: 'DRAW...',
			caption: 'Three friends get togehter to play a game of UNO until things turn sour.',
			img: '/film/uno/boxart.png',
		},
		links: {
			vimeo: 'https://vimeo.com/123456789',
			youtube: 'https://www.youtube.com/watch?v=J6hMl5Y1BkE',
		},
		credits: [
			{
				role: 'Director',
				name: 'Kai Luckey',
			},
			{
				role: 'A.Director',
				name: 'Noah Fields'
			},
			{
				role: 'Writers',
				name: 'Kai Luckey'
			},
			{
				role: 'Cinemotographer',
				name: 'Kai Luckey'
			},
			{
				role: 'Producer',
				name: 'Erin Hennig, Kai Luckey, Noah Fields'
			},
		],
		date: '2024',
		img: '/film/uno/image1.png',
		url: '/catalog/uno',
	},


}
