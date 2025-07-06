import { authAxios } from "@/api/instance/instance";
import { ServerHealth } from "@/types/health";

export async function getHealth(): Promise<ServerHealth | null> {
  try {
    // Use a much shorter timeout for SSR to avoid long waits
    const response = await authAxios.get<ServerHealth>('/api/v1/health', {
      timeout: 2000, // 2 seconds instead of 10
    });
    return response.data;
  } catch (error) {
    // During SSR or when backend is unavailable, return null instead of throwing
    console.warn('Health check failed:', error instanceof Error ? error.message : 'Unknown error');
    return null;
  }
}
