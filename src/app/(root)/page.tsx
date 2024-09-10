import CinLogo from '@/components/CinLogo';
import '../styles/home.scss'
import { VideoFeature } from '@/components/home/VideoFeature';

export default function Home() {
	return (
		<section className="home">
			<div className="home__cin">
				<CinLogo />
			</div>
			<div className="home__content">
				<VideoFeature />
			</div>
		</section>
	);
}
