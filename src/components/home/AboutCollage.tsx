import { BrandSettings } from '@/brand.settings';
import '../styles/aboutcollage.scss';

function AboutCollage() {

	return (
		<div className="about-collage">
			<h2>{BrandSettings.brandName.display}</h2>
			<div className="about-collage__content flex">
				<div className="about-collage__collage">

				</div>
				<div className="about-collage__text flex column">
					<p>
						Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris mollis eget nunc non tincidunt. Ut consectetur bibendum velit, eget lacinia est vestibulum at. Suspendisse tempor lacus id nulla auctor pulvinar a faucibus ipsum. Duis rhoncus fermentum dui ac commodo. Integer at leo sed lacus consequat dapibus vel at metus.
					</p>
					<p>
						Proin ac mi accumsan, pharetra odio eget, aliquam arcu. Cras sagittis urna purus. Integer tincidunt cursus quam nec consequat. Sed viverra egestas odio a commodo.
					</p>
				</div>
			</div>

		</div>
	)

}

export default AboutCollage;
