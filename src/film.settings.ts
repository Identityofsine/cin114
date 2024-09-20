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
				role: "Starring",
				name: ['Kai Luckey', 'Antonio Venticinque']
			},
			{
				role: 'Director',
				name: 'Kai Luckey',
			},
			{
				role: 'Writers',
				name: 'Kai Luckey'
			},
			{
				role: 'Cinemotographer',
				name: 'Shane Keeley'
			},
			{
				role: 'Producers',
				name: ['Erin Hennig', 'Kai Luckey', 'Noah Fields']
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
			youtube: 'https://www.youtube.com/watch?v=usKIKXzoXyM&ab_channel=CIN114',
		},
		credits: [
			{
				role: "Starring",
				name: ['Kai Luckey', 'Shane Keeley', 'Noah Fields'],
			},
			{
				role: 'Director',
				name: 'Kristopher King',
			},
			{
				role: 'Writer',
				name: 'Kristopher King'
			},
			{
				role: 'Cinematographer',
				name: 'Kai Luckey'
			},
			{
				role: 'Producer',
				name: ['Kristopher King', 'Kai Luckey']
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
		},
		credits: [
			{
				role: "Starring",
				name: ['Kristopher King', 'Kai Luckey'],
			},
			{
				role: 'Director',
				name: 'Shane Keeley',
			},
			{
				role: 'Writer',
				name: 'Shane Keeley'
			},
			{
				role: 'Cinemotographer',
				name: 'Kai Luckey'
			},
			{
				role: 'Producers',
				name: ['Shane Keeley', 'Kai Luckey']
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
			youtube: 'https://www.youtube.com/watch?v=Gk0a63sfaF0&ab_channel=CIN114',
		},
		credits: [
			{
				role: "Starring",
				name: ['Louis Clarke', 'Meredith Reed', 'Gabriel Patrascu', 'Sebastian Caldwell', 'Kenrhon Anthony', 'Cayson Rhodes', 'Makayla Russo'],
			},
			{
				role: 'Director',
				name: 'Noah Fields',
			},
			{
				role: 'A.Camera',
				name: 'Kristopher King'
			},
			{
				role: 'Cinemotographer',
				name: 'Kai Luckey'
			},
			{
				role: 'Producer',
				name: ['Sebastian Caldwell', 'Kai Luckey']
			},
		],
		date: '2024',
		url: '/catalog/imaginary-rules-of-engagement',
	}
}

type VideoIndex = '16round'

export const Videos: Record<VideoIndex, VideoMetadata> = {
	'16round': {
		title: '16 ROUND Wood Fired Pizza',
		description: 'A promotional video for 16 ROUND Wood Fired Pizza',
		boxart: {
			title: '',
			caption: 'A promotional video for 16 ROUND Wood Fired Pizza',
			img: '/video/16round/image1.png',
			video: '/video/16round/video.mp4',
		},
		useboxartaspreview: true,
		url: '#',
		date: '2024',
		img: '/video/16round/image1.png',
	}
}


type MusicVideoIndex = 'mv1' | 'mv2' | 'mv3'

export const MusicVideos: Record<MusicVideoIndex, VideoMetadata> = {
	'mv1': {
		title: '',
		description: '',
		boxart: {
			title: '',
			caption: '',
			img: '/mv/mv1/image1.png',
			video: '/mv/mv1video.mp4',
		},
		useboxartaspreview: false,
		url: '#',
		date: '',
		img: '/mv/mv1/image1.png',
	},
	'mv2': {
		title: '',
		description: '',
		boxart: {
			title: '',
			caption: '',
			img: '/mv/mv2/image1.png',
			video: '/mv/mv2/video.mp4',
		},
		useboxartaspreview: false,
		url: '#',
		date: '',
		img: '/mv/mv2/image1.png',
	},
	'mv3': {
		title: '',
		description: '',
		boxart: {
			title: '',
			caption: '',
			img: '/mv/mv3/image1.png',
			video: '/mv/mv3/video.mp4',
		},
		style: {
			backgroundPosition: '25% center',
		},
		useboxartaspreview: false,
		url: '',
		date: '',
		img: '/mv/mv3/image1.png',
	},

}
