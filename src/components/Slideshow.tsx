'use client'

import React from "react";

type Slide = {
	img: string;
	href: string;
}

type SlideshowProps = {
	slides: Slide[];
	duration: number;
}

type SlideAnimation = 'leaving' | 'entering';

function Slideshow({ slides, duration }: SlideshowProps) {

	const [animationState, setAnimationState] = React.useState<SlideAnimation>('entering');
	const [curSlide, setCurSlide] = React.useState(0);

	React.useEffect(() => {
		const interval = setInterval(() => {
			setCurSlide((prev) => (prev + 1) % slides.length);
		}, duration);
		return () => clearInterval(interval);
	}, [slides.length, duration]);

	React.useEffect(() => {
		setAnimationState('entering');
	}, [curSlide]);

	function startTransition() {
		setAnimationState('leaving');
	}

	function nextSlide() {
		setCurSlide((prev) => (prev + 1) % slides.length);
	}

	function getNext() {
		return (curSlide + 1) % slides.length;
	}

	return (
		<div className="slideshow">
			<div className="slideshow__slide">
				<img src={slides[curSlide].img} alt="Slide" className={`slide slide-${animationState}`} onAnimationEnd={nextSlide} />
				<img src={slides[getNext()].img} alt="Slide" className={`slide slide-fadein slide-${animationState === 'leaving' ? 'entering' : ''}`} />
			</div>
		</div>
	)
}

export default Slideshow;
