'use client';
import './styles/film.scss';

import { VideoMetadata } from "@/types/video";
import React, { CSSProperties } from "react";

//@FILEDESC FilmClient is a client-side file that contains various Components strictly used in the `Film.tsx` file. These shouldn't be used outside of the `Film.tsx` file.

type FilmClientProps = {
  metadata: VideoMetadata;
}

//@BLOCK
//@TITLE FilmBackground
//@DESC FilmBackground is a component that renders the film's background video and image. This component first fetches the video file and then waits a few moments before playing the video. This is to ensure a valuable user experience so that the still image is shown first before the video starts playing. While this isn't a perfect solution, it keeps a good check on React Renders and ensures that the video is loaded before playing.
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
      ref.current.play().then(() => {
        setPlaying(true);
        setLoaded(false);
      }).catch((error) => {
        setPlaying(false);
        console.error('Error playing video:', error);
      });
      ref.current.autoplay = true;
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
//@END

//@BLOCK
//@TITLE FilmCredit
//@DESC FilmCredit is a component that renders the film's credits. It takes in the metadata and renders the credits in a horizontal scrollable container. This component is used to display the credits of the film. This component is adjusted to have a horiziontal scrollbar that shows the credits of the film if any overflow occurs.
export function FilmCredit({ metadata }: FilmClientProps) {

  const scroll_ref = React.useRef<HTMLDivElement>(null);
  const [scroll, setScroll] = React.useState(0);

  function onScroll(e: React.UIEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;
    const scroll = e.currentTarget.scrollLeft / (target.scrollWidth - target.clientWidth) * 100;
    const clamp = (v: number, min: number, max: number) => Math.min(Math.max(v, min), max);

    setScroll(clamp(scroll, 0, 100));
  }

  return (
    <div className="film__credits">
      <div className="scrollbar" style={{ '--scrollbar-width': `${scroll}%` } as CSSProperties} />
      <div ref={scroll_ref} className="film__credits__container" onScroll={onScroll}>
        {metadata?.credits && metadata.credits.map((credit, index) => (
          <div className="film__credits__credit" key={credit.role + credit.name + index}>
            <div className="film__credits__credit__title" key={credit.role + index}>
              <p>
                {credit.role}
              </p>
            </div>
            <div className="film__credits__credit__name">
              {!Array.isArray(credit.name) ? <p>{credit.name}</p> :
                credit.name.map((name, index) => (
                  <p key={name + index}>{name}</p>
                ))
              }
            </div>
          </div>
        ))}
      </div>
    </div>)
}
//@END
