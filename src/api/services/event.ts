import { authAxios } from "@/api/instance/instance";
import { Checkout, CheckoutApi } from "@/types/checkout";
import { Event, EventApi } from "@/types/event";

export async function getActiveEvents(): Promise<Event[]> {
  try {
    // Use a much shorter timeout for SSR to avoid long waits
    const response = await authAxios.get<EventApi[]>('/api/v1/events/active', {
      timeout: 2000, // 2 seconds instead of 10
    });
    // Convert API response to the desired format
    const events = response.data.map(populateEventFromBackend);

    return events;
  } catch (error) {
    // During SSR or when backend is unavailable, return an empty array instead of throwing
    console.warn('Health check failed:', error instanceof Error ? error.message : 'Unknown error');
    return [];
  }
}

export async function getEvents(): Promise<Event[]> {
  try {
    // Use a much shorter timeout for SSR to avoid long waits
    const response = await authAxios.get<EventApi[]>('/api/v1/events', {
      timeout: 2000, // 2 seconds instead of 10
    });
    // Convert API response to the desired format
    const events = response.data.map(populateEventFromBackend);

    return events;
  } catch (error) {
    // During SSR or when backend is unavailable, return an empty array instead of throwing
    console.warn('Health check failed:', error instanceof Error ? error.message : 'Unknown error');
    return [];
  }
}

export async function getEvent(eventId: number): Promise<Event | null> {
  try {
    // Use a much shorter timeout for SSR to avoid long waits
    const response = await authAxios.get<EventApi>('/api/v1/events/' + eventId, {
      timeout: 2000, // 2 seconds instead of 10
    });
    // Convert API response to the desired format if necessary
    const event = populateEventFromBackend(response.data);

    if (event.expirationDate && event.expirationDate < new Date()) {
      return null;
    }

    return event;
  } catch (error) {
    // During SSR or when backend is unavailable, return null instead of throwing
    console.warn('Health check failed:', error instanceof Error ? error.message : 'Unknown error');
    return null;
  }
}

export async function createCheckout(eventId: number): Promise<Checkout> {
  try {
    const response = await authAxios.post<CheckoutApi>('api/v1/events/' + eventId + '/checkout', {
      quantity: 1 // Assuming a default quantity of 1 for the checkout
    });
    // Convert API response to the desired format
    const checkout = populateCheckoutFromBackend(response.data);
    return checkout;
  } catch (error) {
    console.error('Error creating checkout:', error);
    throw error; // Re-throw the error for further handling
  }
}

function populateCheckoutFromBackend(checkout: CheckoutApi): Checkout {
  return {
    checkoutUrl: checkout.checkout_url,
    sessionId: checkout.session_id
  }
}

function populateEventFromBackend(event: EventApi): Event {
  // Helper function to parse date string as-is without timezone conversions
  const parseAsIs = (dateString: string): Date => {
    // Remove 'Z' if present to avoid UTC parsing
    const cleanDateString = dateString.replace('Z', '');
    // Parse as local time - treats the timestamp as-is
    return new Date(cleanDateString);
  };

  return {
    eventId: event.event_id,
    videoId: event.video_id,
    description: event.description,
    shortDescription: event.short_description,
    expirationDate: event.expiration_date ? parseAsIs(event.expiration_date) : undefined,
    locations: event.locations?.map(location => ({
      eventId: location.event_id,
      locationName: location.location_name,
      locationDescription: location.location_description,
      latitude: location.latitude,
      longitude: location.longitude,
      address: location.location_address,
      createdAt: location.created_at ? parseAsIs(location.created_at) : undefined,
      updatedAt: location.updated_at ? parseAsIs(location.updated_at) : undefined
    })),
    images: event.images?.map(image => ({
      eventId: image.event_id,
      imageUrl: image.image_url,
      imageType: image.image_type, // Assuming image_type is already in the correct format
      createdAt: image.created_at ? parseAsIs(image.created_at) : undefined,
      updatedAt: image.updated_at ? parseAsIs(image.updated_at) : undefined
    })),
    createdAt: event.created_at ? parseAsIs(event.created_at) : undefined,
    updatedAt: event.updated_at ? parseAsIs(event.updated_at) : undefined
  }
}
