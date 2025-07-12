import { VideoMetadata } from '@/types/video';
import { authAxios } from '../instance/instance';

export async function getVideo(videoId?: number): Promise<VideoMetadata | undefined> {

  if (!videoId) {
    return;
  }

  try {
    const r = (await authAxios.get<VideoMetadata>(`/video/${videoId}`))
    if (r.status !== 200 && r.status !== 404) {
      throw new Error(`Failed to fetch video with ID ${videoId}`);
    } else if (r.status === 404) {
      return;
    }
    return r.data;
  } catch (error) {
    if (error instanceof Error && error.message.includes('timeout')) {
      console.warn(`Timeout fetching video with ID ${videoId}. Returning undefined.`);
      return;
    }
    console.error(`Error fetching video with ID ${videoId}:`, error);
    return;
  }

}
