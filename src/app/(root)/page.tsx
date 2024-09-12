import CinLogo from '@/components/CinLogo';
import '../styles/home.scss'
import { VideoFeature } from '@/components/home/VideoFeature';
import AboutCollage from '@/components/home/AboutCollage';
import ContactUs from '@/components/home/ContactUs';

type Credit = {
	name: string;
	role: string;
}

export type VideoMetadata = {
	title: string;
	description?: string;
	credits?: Credit[];
	url: string;
	date: string;
	img?: string;
}

const DUMMY_VIDEOS: VideoMetadata[][] = [[
	{
		title: 'Running From',
		url: 'https://www.youtube.com/watch?v=0fKg7e37bQE',
		date: '12/06/2023',
		img: '/film/runningfrom/runningfrom.jpg'
	}
],
[
	{
		title: 'Running From',
		url: 'https://www.youtube.com/watch?v=0fKg7e37bQE',
		date: '12/06/2023',
		img: '/film/runningfrom/runningfrom2.jpg'
	},
]
]

function repeat(n: number) {
	const arr = [];
	let z = 0;
	for (let i = 0; i < n; i++) {
		if (z === DUMMY_VIDEOS.length) z = 0;
		arr.push(DUMMY_VIDEOS[z][0]);
		z++;
	}
	return arr;
}

export default function Home() {
	return (
		<section className="home">
			<div className="home__cin">
				<CinLogo />
			</div>
			<div className="home__content">
				<VideoFeature videos={repeat(5)} style={'horiziontal'} title={'Cinema'} />
				<VideoFeature videos={repeat(4)} style={'mix'} title={'Videography'} />
				<VideoFeature videos={repeat(3)} style={'vertical'} title={'Music Videos'} />
				<AboutCollage />
				<div className="divider" />
				<ContactUs />

			</div>
		</section>
	);
}
