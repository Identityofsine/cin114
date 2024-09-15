'use client'
import './styles/slideshow.scss';

import React, { CSSProperties, useContext } from "react";

export type Slide = {
	img: string;
	href: string;
	style?: CSSProperties;
}

type SlideshowProps = {
	slides: Slide[];
	duration: number;
}

type SlideAnimation = 'leaving' | 'entering' | '';

function Slideshow({ slides, duration }: SlideshowProps) {

	const [animationState, setAnimationState] = React.useState<SlideAnimation>('entering');
	const [curSlide, setCurSlide] = React.useState(0);
	const speed = React.useRef(duration / 10000);

	React.useEffect(() => {
		const interval = setInterval(() => {
			startTransition();
		}, duration);
		return () => clearInterval(interval);
	}, [slides.length, duration]);

	React.useEffect(() => {
		setAnimationState('');
		return () => {
			setAnimationState('entering');
		}
	}, [curSlide]);

	function startTransition() {
		setAnimationState('leaving');
	}

	function nextSlide() {
		setCurSlide((prev) => (prev + 1) % slides.length);
	}

	const getNext = React.useCallback(() => {
		return (curSlide + 1) % slides.length;
	}, [curSlide, slides.length]);

	const navigate = React.useCallback(() => {
		window.location.href = slides[curSlide].href;
	}, [curSlide]);

	return (
		<div className="slideshow" style={{ '--anim-speed': `${speed.current}s` } as React.CSSProperties}>
			<div className="slideshow__slide">
				<img src={slides[curSlide].img} alt="Slide" onClick={navigate} className={`slide slide-main slide-${animationState}`} onAnimationEnd={nextSlide} style={slides[curSlide].style} />
				<img src={slides[getNext()].img} alt="Slide" className={`slide slide-fadein slide-${animationState === 'leaving' ? 'entering' : ''}`} style={slides[getNext()].style} />
			</div>
		</div>
	)
}

export default Slideshow;
