import { BrandSettings } from '@/brand.settings';
import '../styles/aboutcollage.scss';

function AboutCollage() {

	return (
		<div className="about-collage">
			<h2>{BrandSettings.brandName.display}</h2>
			<div className="about-collage__content flex">
				<div className="about-collage__collage">
					<img src="/home/collage/cimg-01.png" className="c1" alt="Collage" />
					<img src="/home/collage/cimg-02.png" className="c2" alt="Collage" />
					<img src="/home/collage/cimg-03.png" className="c3" alt="Collage" />
					<img src="/home/collage/cimg-04.png" className="c4" alt="Collage" />
					<img src="/home/collage/cimg-05.png" className="c5" alt="Collage" />
					<img src="/home/collage/cimg-06.png" className="c6" alt="Collage" />
					<img src="/home/collage/cimg-07.png" className="c7" alt="Collage" />
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
