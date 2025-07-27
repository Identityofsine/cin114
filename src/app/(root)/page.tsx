import CinLogo from '@/components/CinLogo';
import '../styles/home.scss'
import { VideoFeature } from '@/components/home/VideoFeature';
import AboutCollage from '@/components/home/AboutCollage';
import ContactUs from '@/components/home/ContactUs';
import { VideoMetadata } from '@/types/video';
import { Films, MusicVideos, Videos } from '@/film.settings';


function repeat(n: number): VideoMetadata[] {
  const arr: VideoMetadata[] = [];
  let z = 0;
  const DUMMY_VIDEOS = Object.values(Films) as VideoMetadata[];
  for (let i = 0; i < n; i++) {
    if (z === DUMMY_VIDEOS.length) z = 0;
    arr.push(DUMMY_VIDEOS[z]);
    z++;
  }
  return arr;
}

export default async function Home() {

  return (
    <section className="home">
      <div className="home__cin">
        <CinLogo />
      </div>
      <div className="home__content">
        <VideoFeature videos={repeat(4)} style={'horiziontal'} title={'Cinema'} description={"CIN-114 makes films that are worth a trip to the theatre."} />
        <VideoFeature videos={[Videos['16round']]} style={'single'} title={'Videography'} description={`We won't just record your event, we'll make it a movie.
Our high end equipment and expertise will turn your wedding, tutorial series, or commercial into a filmlike experience.`} />
        <VideoFeature videos={[MusicVideos['mv1'], MusicVideos['mv2'], MusicVideos['mv3']]} style={'vertical'} title={'Music Videos'} description={`CIN-114 will work closely with you to plan, film, and edit your music video to bring your vision to life.`} />
        <AboutCollage />
        <div className="divider" />
        <ContactUs />

      </div>
    </section>
  );
}
