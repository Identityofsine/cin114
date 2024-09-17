import { BrandSettings } from '@/brand.settings';
import '../styles/contactus.scss';
import { TextChange } from './ContactUsClient';

const text = "Contact us to pitch ideas, book our crew for hands on set, and get quotes for videography. We’re always just an email away.";

function ContactUs() {

	return (
		<div className="contactus" id="contact">
			<div className="contactus__bg">
				<video autoPlay muted loop>
					<source src="/home/bgcontact.mp4" type="video/mp4" />
				</video>
			</div>
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
						{text}
					</p>
					<h2>{BrandSettings.contact.email}</h2>
					<button className="contactus__button">
						<div className="container">
							Get In Touch
							<img src="/ui/share.svg" alt="share" />
						</div>
					</button>
					<p className="contactus__text mobile-enable">
						{text}
					</p>
				</div>
			</div>
		</div>
	)

}
export default ContactUs;
