'use client';

import { VideoMetadata } from "@/types/video";
import React from "react";

type FilmClientProps = {
	metadata: VideoMetadata;
}

export function FilmBackground({ metadata }: FilmClientProps) {

	const [loaded, setLoaded] = React.useState(false);
	const [playing, setPlaying] = React.useState(false);
	const [progress, setProgress] = React.useState(0);
	const ref = React.useRef<HTMLVideoElement>(null);

	React.useEffect(() => {
		if (loaded || playing) return;
		const timeout = setTimeout(() => {
			if (ref.current && metadata.boxart?.video) {
				const xhr = new XMLHttpRequest();
				xhr.open('GET', metadata.boxart.video, true);
				xhr.responseType = 'blob';
				xhr.onload = () => {
					setProgress(0);
					const video = URL.createObjectURL(xhr.response);
					ref.current!.src = video;
					ref.current!.playbackRate = .95;
					let progressUpdate: NodeJS.Timeout | undefined = undefined;
					ref.current!.onplaying = () => {
						if (progressUpdate) {
							return;
						}
						progressUpdate = setInterval(() => {
							const progress = (ref.current!.currentTime / ref.current!.duration) * 100;
							setProgress(progress);
						}, 100);
					}
					setLoaded(true);
				}
				xhr.send();
			}
		}, 6000)

		return () => {
			clearTimeout(timeout);
		}
	}, [ref.current, loaded, playing])

	function start(e: AnimationEvent) {
		if (e.animationName !== 'waiting') return;
		if (playing) return;
		if (ref.current && loaded && !playing) {
			setPlaying(true);
			ref.current.play();
			ref.current.autoplay = true;
			setLoaded(false);
		}
	}

	return (
		<>
			{metadata.boxart?.video &&
				(
					<div className={`film_loadbar ${!playing && !loaded && 'loading'} ${!playing && loaded && 'waiting'} ${playing && 'playing'}`} onAnimationEnd={(e) => start(e as unknown as AnimationEvent)}>
						<div className="film_loadbar__progress" style={{ width: `${progress}%` }}>
						</div>
					</div>
				)
			}
			<div className="film__bg" style={{ backgroundImage: `url("${metadata.boxart.img}")` }} />
			<div className={`film__bg_video ${!loaded && 'unloaded' || ''} ${playing && 'playing' || ''}`}>
				<video muted loop ref={ref} className={playing && 'playing' || ''}>
				</video>
			</div>
		</>
	)
}
