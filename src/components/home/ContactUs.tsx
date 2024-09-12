import { BrandSettings } from '@/brand.settings';
import '../styles/contactus.scss';

function ContactUs() {

	return (
		<div className="contactus">
			<div className="contactus__bg"></div>
			<div className="contactus__content">

				<div className="contactus__bottom">
					<p className="contactus__text">
						Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus  luctus dapibus risus, ac tincidunt purus auctor id. Phasellus pulvinar  felis at sollicitudin iaculis. Quisque sed efficitur tortor, id porta  ex.
					</p>
					<h2>{BrandSettings.contact.email}</h2>
					<button className="contactus__button">
						<div className="container">
							Get In Touch
							<img src="/ui/share.svg" alt="share" />
						</div>
					</button>
				</div>
			</div>
		</div>
	)

}
export default ContactUs;
