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
						const video = URL.createObjectURL(xhr.response);
						ref.current!.src = video;
						setLoaded(true);
					}, 650);
				}
				xhr.send();
			}
		}, 2500)

		return () => {
			clearTimeout(timeout);
		}
	}, [ref.current])


	return (
		<>
			{metadata.boxart?.video && !loaded &&
				(
					<div className={`film_loadbar ${progress === 0 && 'waiting'}`}>
						<div className="film_loadbar__progress" style={{ width: `${progress}%` }}>
						</div>
					</div>
				)
			}
			<div className="film__bg" style={{ backgroundImage: `url("${metadata.boxart.img}")` }} />
			<div className="film__bg_video">
				<video autoPlay muted loop ref={ref} >
				</video>
			</div>
		</>
	)
}
