"use client";

import '../styles/videofeature.scss';
import { useRef } from 'react';




type VideoFeatureProp = {
	title: string
	style: 'horiziontal' | 'vertical' | 'mix'
	videos: VideoMetadata[]
}




export function VideoFeature({ title, style, videos }: VideoFeatureProp) {

	const maxVideos = useRef(videos.length > 5 ? 5 : videos.length);

	function renderVideos() {
		if (style === 'horiziontal') {
			return (
				<>
					<VideoShowcase video={videos.slice(0, 1)} style={style} />
					{maxVideos.current > 1 && <VideoShowcase video={videos.slice(1, 3)} style={style} />}
					{(maxVideos.current > 3 && <VideoShowcase video={videos.slice(3, 5)} style={style} altgrow />)}
				</>
			)
		} else {
			return (
				<>
					{videos.map((video, index) => (
						<VideoShowcase key={video.url + index} className={`cvideo-${index}`} video={[video]} style={style} />
					))}
				</>
			)
		}
	}

	return (
		<div className="video-feature">
			<h2>{title}</h2>
			<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
			<div className={`video-feature__videos video-feature__videos__${style}`}>
				{renderVideos()}
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
	style: 'horiziontal' | 'vertical' | 'mix'
	className?: string
	altgrow?: boolean
}

function VideoShowcase({ video, style, className = "", altgrow = false }: VideoShowcaseProps) {

	if (video.length === 0) throw new Error('No videos provided');

	if (video.length > 2) {
		console.warn('Only the first 2 videos will be displayed');
	} else if (video.length <= 2) {
		return (
			<div className={`videos-showcase videos-showcase-${style} ${className}`}>
				{video.map((video, index) => (
					<div key={video?.date + index} className={`video-showcase video-${style} video-${index + 1} ${index > 0 && altgrow && 'altgrow'}`} style={{ backgroundImage: `url('${video.img}')` }}>
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
