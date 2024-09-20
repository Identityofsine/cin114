import { BrandSettings } from '@/brand.settings';
import './styles/footer.scss';

function Footer() {
	return (
		<footer>
			<div className="flex">
				<div className="footer__links">
					<h1>© 2024 <b>{BrandSettings.brandName.displayShort}</b></h1>
					<div className="footer__links__contact">
						<a href={BrandSettings.socials.youtube}><img src="/ui/youtube.svg" alt="YouTube" /></a>
						<a href={BrandSettings.socials.instagram}><img src="/ui/instagram.svg" alt="instagram" /></a>
						<a href={BrandSettings.socials.twitter}><img src="/ui/x.svg" alt="Twitter" /></a>
						<a href={BrandSettings.socials.vimeo}><img src="/ui/vimeo.svg" alt="Vimeo" /></a>
						<a href={BrandSettings.socials.pateron}><img src="/ui/patreon.svg" alt="Patreon" /></a>
					</div>
				</div>
				<div className="flex owners row">
					<div className="flex column ">
						<span className="">Kai Luckey</span>
						<span>Erin Hennig</span>
					</div>
					<div className="flex column">
						<span>Shane Keely</span>
						<span>Noah Fields</span>
					</div>
					<div className="flex column">
						<span>Kristopher King</span>
					</div>
				</div>
			</div>
			<div className="footer__href">
				<a href="/#about">About</a>
				<a href="/#contact">Contact</a>
			</div>
		</footer>
	);
}

export default Footer;
