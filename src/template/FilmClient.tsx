'use client';

import { VideoMetadata } from "@/types/video";
import React from "react";

type FilmClientProps = {
	metadata: VideoMetadata;
}

export function FilmBackground({ metadata }: FilmClientProps) {

	const [loaded, setLoaded] = React.useState(false);
	const [progress, setProgress] = React.useState(0);
	const ref = React.useRef<HTMLVideoElement>(null);

	React.useEffect(() => {
		if (loaded) return;
		const timeout = setTimeout(() => {
			if (ref.current && metadata.boxart?.video) {
				const xhr = new XMLHttpRequest();
				xhr.open('GET', metadata.boxart.video, true);
				xhr.responseType = 'blob';
				xhr.onprogress = (e) => {
					if (e.lengthComputable) {
						setProgress((e.loaded / e.total) * 100);
					}
				}
				xhr.onload = () => {
					setTimeout(() => {
						setProgress(0);
						const video = URL.createObjectURL(xhr.response);
						ref.current!.src = video;
						ref.current!.playbackRate = .95;
						var progressUpdate: any = undefined;
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
					}, 650);
				}
				xhr.send();
			}
		}, 6000)

		return () => {
			clearTimeout(timeout);
		}
	}, [ref.current, loaded])


	function start() {

	};

	return (
		<>
			{metadata.boxart?.video &&
				(
					<div className={`film_loadbar ${progress === 0 && 'waiting'}`} onAnimationEnd={() => { start(); }}>
						<div className="film_loadbar__progress" style={{ width: `${progress}%` }}>
						</div>
					</div>
				)
			}
			<div className="film__bg" style={{ backgroundImage: `url("${metadata.boxart.img}")` }} />
			<div className={`film__bg_video ${!loaded && 'unloaded' || ''}`}>
				<video autoPlay muted loop ref={ref}>
				</video>
			</div>
		</>
	)
}
