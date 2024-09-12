import CinLogo from '@/components/CinLogo';
import '../styles/home.scss'
import { VideoFeature } from '@/components/home/VideoFeature';
import AboutCollage from '@/components/home/AboutCollage';
import ContactUs from '@/components/home/ContactUs';
import { VideoMetadata } from '@/types/video';
import { BrandSettings } from '@/brand.settings';
import { Films } from '@/film.settings';


export function metadata() {
	return {
		title: 'Home - ' + BrandSettings.brandName.displayShort,
	};
}

function repeat(n: number) {
	const arr = [];
	let z = 0;
	const DUMMY_VIDEOS = Object.values(Films) as VideoMetadata[];
	for (let i = 0; i < n; i++) {
		if (z === DUMMY_VIDEOS.length) z = 0;
		arr.push(DUMMY_VIDEOS[z]);
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
