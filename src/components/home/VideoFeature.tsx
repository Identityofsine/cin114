import { Postpone } from 'next/dist/server/app-render/dynamic-rendering';
import '../styles/videofeature.scss';


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


export function VideoFeature() {
	return (
		<div className="video-feature">
			<h2>Cinema</h2>
			<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
			<div className="video-feature__videos">
				<VideoShowcase video={DUMMY_VIDEOS[0]} />
				<VideoShowcase video={[...DUMMY_VIDEOS[0], ...DUMMY_VIDEOS[1]]} />
				<VideoShowcase video={[...DUMMY_VIDEOS[1], ...DUMMY_VIDEOS[0]]} altgrow />
			</div>
		</div>
	);
}

type VideoMetadata = {
	title: string;
	url: string;
	date: string;
	img?: string;
}

type VideoShowcaseProps = {
	video: VideoMetadata[];
	altgrow?: boolean
}

function VideoShowcase({ video, altgrow = false }: VideoShowcaseProps) {

	if (video.length === 0) throw new Error('No videos provided');

	if (video.length > 2) {
		console.warn('Only the first 2 videos will be displayed');
	} else if (video.length <= 2) {
		return (
			<div className="videos-showcase">
				{video.map((video, index) => (
					<div className={`video-showcase video-${index + 1} ${index > 0 && altgrow && 'altgrow'}`} style={{ backgroundImage: `url('${video.img}')` }}>
						<div className="flex column video-showcase__title">
							<h3>{video.title}</h3>
							<span>{video.date}</span>
						</div>
					</div>
				))}
			</div>
		);
	}

}
