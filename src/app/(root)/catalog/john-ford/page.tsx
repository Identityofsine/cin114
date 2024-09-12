import '@/app/styles/johnford.scss';
import FilmTemplate from '@/template/Film';
import { VideoMetadata } from '@/types/video';

export function metadata() {
	return {
		title: 'Interviewing John Ford (2024)',
		description: 'A collection of films by John Ford',
	};
}

const johnfordmovie: VideoMetadata = {
	title: 'Interviewing John Ford',
	description: 'A collection of films by John Ford',
	boxart: {
		title: '',
		caption: 'Ask the right questions...',
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
	url: '/film/john-ford',
	img: '/film/john-ford/boxart.png',
	date: '2024',

}

export default function JohnFord() {

	return (
		<FilmTemplate metadata={johnfordmovie}>
		</FilmTemplate>
	)
}
