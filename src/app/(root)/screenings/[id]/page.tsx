import ScreeningCallToAction from '@/components/screening/ScreeningCallToAction';
import ScreeningHeader from '@/components/screening/ScreeningHeader';
import '@/components/screening/styles/screening.scss';
import { FilmBackground } from '@/template/FilmClient';
import { Event } from '@/types/event';
import { VideoMetadata } from '@/types/video';

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const id = (await params).id;

  const { event, video } = await new Promise<{ event: Event | undefined, video: VideoMetadata | undefined }>(async (resolve) => {
    const event = await import('@/api').then(mod => mod.getEvent(Number(id)));
    if (!event) {
      throw new Error(`Event with ID ${id} not found`);
    }
    if (event?.videoId) {
      const video = await import('@/api').then(mod => mod.getVideo(event.videoId));
      resolve({ event, video: video });
    }
    resolve({ event, video: undefined });
  });

  if (!event) {
    return <div className="error">Event not found</div>;
  }

  if (!video) {
    return <div className="error">Video not found</div>;
  }

  const headerImages = event.images?.filter(image => image.imageType === 'poster' || image.imageType === 'poster-mobile');

  const boxart: VideoMetadata['boxart'] = {
    title: event.shortDescription || event.description || video.title || '',
    caption: event.shortDescription || event.description || video.description || '',
    img: headerImages?.find(image => image.imageType === 'poster')?.imageUrl || '',
    video: video.url || '',
  }


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
            boxart
          }}
        >
        </FilmBackground>
      </div>
      <div className="screening-content">
      </div>
    </section>
  )
}
