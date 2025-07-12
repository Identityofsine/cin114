import { VideoFeature } from '@/components/home/VideoFeature';

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const id = (await params).id;
  const event = await import('@/api').then(mod => mod.getEvent(Number(id)));

  return (
    <section>
      <div>
      </div>
    </section>
  )
}
