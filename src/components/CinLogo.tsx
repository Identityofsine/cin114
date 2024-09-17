import Slideshow, { Slide } from './Slideshow'
import './styles/cinlogo.scss'

const slides_1: Slide[] = [
	{ img: '/home/logo/box1-1.png', href: '/catalog/john-ford', },
	{ img: '/home/logo/box1-2.png', href: '/catalog/john-ford' },
	{ img: '/home/logo/box1-3.png', href: '/catalog/john-ford' },
	{ img: '/home/logo/box1-4.png', href: '/catalog/imaginary-rules-of-engagement' },
	{ img: '/home/logo/box1-5.png', href: '/catalog/just-give-me-the-night' },
]
const slides_2: Slide[] = [
	{ img: '/home/logo/box2-1.png', href: '/catalog/john-ford' },
	{ img: '/home/logo/box2-2.png', href: '/catalog/john-ford' },
	{ img: '/home/logo/box2-3.png', href: '/catalog/john-ford' },
	{ img: '/home/logo/box2-4.jpg', href: '/catalog/imaginary-rules-of-engagement' },
]
const slides_3: Slide[] = [
	{ img: '/home/logo/box3-1.png', href: '/catalog/imaginary-rules-of-engagement' },
	{ img: '/home/logo/box3-2.png', href: '/catalog/just-give-me-the-night' },
	{ img: '/home/logo/box3-3.png', href: '/catalog/just-give-me-the-night' },
]


export default function CinLogo() {
	return (
		<div className="cin-logo">
			<div className="cin">
				<img className="logo" src="/home/logo.svg" alt="logo_text" />
				<div className="cin__image">
					<div className="cin__image__box cin__image__box--1">
						<Slideshow slides={slides_1} duration={7500} />
					</div>
					<div className="cin__image__box cin__image__box--2">
						<Slideshow slides={slides_2} duration={6000} />
					</div>
					<div className="cin__image__box cin__image__box--3">
						<Slideshow slides={slides_3} duration={9000} />
					</div>
				</div>
			</div>
			<p className="cin-text">
				Giving control back to creatives.
				A production company built by artists, for artists.
			</p>
		</div>
	)
} 
