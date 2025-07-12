import '@/components/screening/styles/screening.scss';
import { Films } from '@/film.settings';
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
    console.log('Event:', event);
    if (event?.videoId) {
      const video = await import('@/api').then(mod => mod.getVideo(event.videoId));
      resolve({ event, video: video });
    }
    resolve({ event, video: undefined });
  });

  if (!video) {
    return <div className="error">Video not found</div>;
  }


  return (
    <section className="screening">
      <div className="screening-hero">
        <FilmBackground
          metadata={{
            ...video,
            boxart: {
              ...video.boxart,
            },
          }}
        >
        </FilmBackground>
      </div>
      <div className="screening-content">
      </div>
    </section>
  )
}
