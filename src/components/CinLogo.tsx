import Slideshow from './Slideshow'
import './styles/cinlogo.scss'

const slides = [
	{ img: '/home/collage/cimg-01.png', href: '#' },
	{ img: '/home/collage/cimg-02.png', href: '#' },
]


export default function CinLogo() {
	return (
		<div className="cin-logo">
			<div className="cin">
				<img src="/home/logo.svg" alt="logo_text" />
				<div className="cin__image">
					<div className="cin__image__box cin__image__box--1">
						<Slideshow slides={slides} duration={3000} />
					</div>
					<div className="cin__image__box cin__image__box--2"></div>
					<div className="cin__image__box cin__image__box--3"></div>
				</div>
			</div>
			<p className="cin-text">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
		</div>
	)
} 
