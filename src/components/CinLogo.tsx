import './styles/cinlogo.scss'

export default function CinLogo() {
	return (
		<div className="cin-logo">
			<div className="cin">
				<img src="/home/logo.svg" alt="logo_text" />
				<div className="cin__image">
					<div className="cin__image__box cin__image__box--1"></div>
					<div className="cin__image__box cin__image__box--2"></div>
					<div className="cin__image__box cin__image__box--3"></div>
				</div>
			</div>
			<p className="cin-text">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
		</div>
	)
} 
