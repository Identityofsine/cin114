import { notFound } from 'next/navigation';
import ScreeningCallToAction from '@/components/screening/ScreeningCallToAction';
import ScreeningDetails from '@/components/screening/ScreeningDetails';
import ScreeningHeader from '@/components/screening/ScreeningHeader';
import ScreeningMapPreview from '@/components/screening/ScreeningMapPreview';
import '@/components/screening/styles/screening.scss';
import { FilmBackground } from '@/template/FilmClient';
import { VideoMetadata } from '@/types/video';
import { getEvent, getVideo } from '@/api';

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const id = (await params).id;

  // Validate the ID is a number
  const eventId = Number(id);
  if (isNaN(eventId)) {
    notFound();
  }

  // Fetch the event
  const event = await getEvent(eventId);
  if (!event) {
    notFound();
  }

  // Fetch the video if event has a videoId
  let video: VideoMetadata | undefined = undefined;
  if (event.videoId) {
    video = await getVideo(event.videoId);
    if (!video) {
      notFound();
    }
  } else {
    notFound(); // If no video associated with the event
  }

  const headerImages = event.images?.filter(image => image.imageType === 'poster' || image.imageType === 'poster-mobile');
  const boxartVideo = event.images?.filter(image => image.imageType === 'video')?.[0]?.imageUrl ?? '';

  const boxart: VideoMetadata['boxart'] = {
    title: event.shortDescription || event.description || video.title || '',
    caption: event.shortDescription || event.description || video.description || '',
    img: headerImages?.find(image => image.imageType === 'poster')?.imageUrl || '',
    video: boxartVideo || '',
  }

  const mobileBoxart: VideoMetadata['mobileBoxart'] = {
    title: event.shortDescription || event.description || video.title || '',
    caption: event.shortDescription || event.description || video.description || '',
    img: headerImages?.find(image => image.imageType === 'poster-mobile')?.imageUrl || '',
    video: boxartVideo || '',
  }

  console.log(boxart, mobileBoxart);

  const location = event.locations?.[0];

  return (
    <section className="screening">
      <div className="screening-hero">
        <div className="screening-hero-content">
          <div className="screening-hero-header">
            <ScreeningHeader />
          </div>
          <div className="screening-hero-cta">
            <ScreeningCallToAction
              eventId={event.eventId || 0}
              expirationDate={event.expirationDate || new Date()}
            />
          </div>
        </div>
        <FilmBackground
          metadata={{
            ...video,
            boxart,
            mobileBoxart
          }}
        >
        </FilmBackground>
      </div>
      <div className="screening-content">
        <ScreeningMapPreview location={location} />
        <ScreeningDetails event={event} />
      </div>
    </section>
  )
}
