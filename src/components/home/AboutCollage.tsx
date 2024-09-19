import { BrandSettings } from '@/brand.settings';
import '../styles/aboutcollage.scss';

function AboutCollage() {

	return (
		<div className="about-collage" id="about">
			<h2>{BrandSettings.brandName.displayShort}</h2>
			<div className="about-collage__content flex">
				<div className="about-collage__collage">
					<img src="/home/collage/c-full.png" className="cimg" alt="Collage" />
					{/* 
					<img src="/home/collage/cimg-01.png" className="c1" alt="Collage" />
					<img src="/home/collage/cimg-02.png" className="c2" alt="Collage" />
					<img src="/home/collage/cimg-03.png" className="c3" alt="Collage" />
					<img src="/home/collage/cimg-04.png" className="c4" alt="Collage" />
					<img src="/home/collage/cimg-05.png" className="c5" alt="Collage" />
					<img src="/home/collage/cimg-06.png" className="c6" alt="Collage" />
					<img src="/home/collage/cimg-07.png" className="c7" alt="Collage" />
					*/}
				</div>
				<div className="flex column text-block">
					<p>
						When we all met in a college class, coincidentally titled CIN-114, we had no plans to form a company. However, determined to collaborate, we started on our first film of many only a week after graduation. We haven’t been able to stop making movies since then.
					</p>
					<p>
						Our films have played at AMC Theatres, NYLIFF and LIYFF, along with streaming to thousands on YouTube and Vimeo. Our goal is to give tools and community to local artists so that they can see their vision come to life on the silver screen.
					</p>
					<p>
						We also provide videography services for weddings, birthdays, courses, and anything you can imagine. We believe that you are the star of your special day, which means it deserves to be captured in movie quality. Contact us for quotes.
					</p>
				</div>
			</div>

		</div >
	)

}

export default AboutCollage;
