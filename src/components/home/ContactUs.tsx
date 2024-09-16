import { BrandSettings } from '@/brand.settings';
import '../styles/contactus.scss';
import { TextChange } from './ContactUsClient';

function ContactUs() {

	return (
		<div className="contactus" id="contact">
			<div className="contactus__bg"></div>
			<div className="contactus__content">
				<div className="contactus__contact_text">
					<TextChange />
					<div className="contactus__contact_text_img">
						<a href="#"><img src="/ui/youtube.svg" alt="YouTube" /></a>
						<a href="#"><img src="/ui/instagram.svg" alt="instagram" /></a>
						<a href="#"><img src="/ui/x.svg" alt="Twitter" /></a>
						<a href="#"><img src="/ui/vimeo.svg" alt="Vimeo" /></a>
						<a href="#"><img src="/ui/patreon.svg" alt="Patreon" /></a>
					</div>
				</div>
				<div className="contactus__bottom">
					<p className="contactus__text mobile-disable">
						Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus  luctus dapibus risus, ac tincidunt purus auctor id. Phasellus pulvinar  felis at sollicitudin iaculis. Quisque sed efficitur tortor, id porta  ex.
					</p>
					<h2>{BrandSettings.contact.email}</h2>
					<button className="contactus__button">
						<div className="container">
							Get In Touch
							<img src="/ui/share.svg" alt="share" />
						</div>
					</button>
					<p className="contactus__text mobile-enable">
						Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus  luctus dapibus risus, ac tincidunt purus auctor id. Phasellus pulvinar  felis at sollicitudin iaculis. Quisque sed efficitur tortor, id porta  ex.
					</p>
				</div>
			</div>
		</div>
	)

}
export default ContactUs;
