"use client";

import { VideoMetadata } from '@/types/video';
import '../styles/videofeature.scss';
import React, { useRef } from 'react';




type VideoFeatureProp = {
  title: string
  description?: string
  style: 'horiziontal' | 'vertical' | 'mix' | 'single'
  videos: VideoMetadata[]
}




export function VideoFeature({ title, style, description, videos }: VideoFeatureProp) {

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
      <p>{description}</p>
      <div className={`video-feature__videos video-feature__videos__${style}`}>
        {renderVideos()}
      </div>
    </div>
  );
}


type VideoShowcaseProps = {
  video: VideoMetadata[];
  style: 'horiziontal' | 'vertical' | 'mix' | 'single'
  className?: string
  altgrow?: boolean
}

type VideoProps = {
  video: VideoMetadata
  style: 'horiziontal' | 'vertical' | 'mix' | 'single'
  className?: string
  altgrow?: boolean
  index: number
}

function Video({ video, style, altgrow = false, index }: VideoProps) {
  const video_ref = useRef<HTMLVideoElement>(null);
  const [onHover, setOnHover] = React.useState(false);
  const [trigger, setTrigger] = React.useState<boolean>(false);

  React.useEffect(() => {
    if (video_ref.current) {
      if (onHover) {
        video_ref.current.play();
      }
      else {

      }
    }
  }, [onHover, video_ref]);

  React.useEffect(() => {
    if (!trigger) {
      if (video_ref.current) {
        video_ref.current.pause();
        video_ref.current.currentTime = 0;
      }
    }
  }, [trigger, video_ref]);

  return (
    <a href={video.url} className={`video-showcase video-${style} video-${index + 1} ${index > 0 && altgrow && 'altgrow'}`} style={{ backgroundImage: `url('${video.img}')`, ...video.style }} onMouseEnter={() => setOnHover(true)} onMouseLeave={() => setOnHover(false)} >
      {video.useboxartaspreview && <video ref={video_ref} className={`${onHover && 'hover' || ''}`} loop muted playsInline src={video.boxart.video} onTransitionEnd={() => setTrigger(onHover)} />}
      <div className="flex column video-showcase__title">
        <h3>{video.title}</h3>
        <span>{video.date}</span>
      </div>
    </a>
  )
}



function VideoShowcase({ video, style, className = "", altgrow = false }: VideoShowcaseProps) {



  if (video.length === 0) return (<></>);
  else if (video.length <= 2) {
    return (
      <div className={`videos-showcase videos-showcase-${style} ${className}`}>
        {video.map((video, index) => (
          index > 1 && <></> ||
          (<Video key={video.url + index} video={video} style={style} index={index} altgrow={altgrow} />)
        ))}
      </div>
    );
  }

}
