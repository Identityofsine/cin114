import { BrandSettings } from '@/brand.settings';
import './styles/footer.scss';

function Footer() {
	return (
		<footer>
			<div className="footer__links">
				<h1>© 2024 <b>{BrandSettings.brandName.displayShort}</b></h1>
				<div className="footer__links__contact">
					<a href="#"><img src="/ui/youtube.svg" alt="YouTube" /></a>
					<a href="#"><img src="/ui/instagram.svg" alt="instagram" /></a>
					<a href="#"><img src="/ui/x.svg" alt="Twitter" /></a>
					<a href="#"><img src="/ui/vimeo.svg" alt="Vimeo" /></a>
					<a href="#"><img src="/ui/patreon.svg" alt="Patreon" /></a>
				</div>
			</div>
			<div className="footer__href">
				<a href="#">About</a>
				<a href="#">Contact</a>
			</div>
		</footer>
	);
}

export default Footer;
