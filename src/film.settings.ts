import { VideoMetadata } from "./types/video";

export type FilmIndex = 'johnford' | 'uno' | 'justgivemetonight' | 'uno' | 'theimaginaryrulesofengagement'

export const Films: Record<FilmIndex, VideoMetadata> = {
	'johnford': {
		title: 'Interviewing John Ford',
		description: 'A young interviewer gets the opportunity to interview famed film director John Ford.',
		boxart: {
			title: '',
			caption: 'Ask the right questions...',
			img: '/film/john-ford/boxart.png',
			video: '/film/john-ford/video.mp4',
		},
		links: {
			youtube: 'https://www.youtube.com/watch?v=gZaosS7-l5w&ab_channel=CIN114',
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
			video: '/film/uno/video.mp4',
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
	'justgivemetonight': {
		title: 'Just Give Me The Night',
		description: '',
		boxart: {
			title: "It's a simple task...",
			caption: 'A man ventures into the night to buy food for for his cat, but complications soon arise.',
			img: '/film/just-give-me-tonight/boxart.png',
			video: '/film/just-give-me-tonight/video.mp4',
		},
		links: {
			vimeo: 'https://vimeo.com/123456789',
			youtube: 'https://www.youtube.com/watch?v=J6hMl5Y1BkE',
		},
		credits: [
			{
				role: 'Director',
				name: 'Shane Keeley',
			},
			{
				role: 'A.Director',
				name: 'Noah Fields, Kai Luckey'
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
				name: 'Shane Keeley, Kai Luckey'
			},
		],
		date: '2024',
		img: '/film/just-give-me-tonight/image1.png',
		url: '/catalog/just-give-me-the-night',
	},
	'theimaginaryrulesofengagement': {
		title: 'The Imaginary Rules of Engagement',
		description: '',
		boxart: {
			title: '',
			caption: "Two kids are outside playing with toys while using their imaginations. What's the worst that can happen?",
			img: '/film/the-rules-of-engagement/boxart.png',
			video: '/film/the-rules-of-engagement/video.mp4',
		},
		img: '/film/the-rules-of-engagement/image1.png',
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
		url: '/catalog/imaginary-rules-of-engagement',
	}


}
