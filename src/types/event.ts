export interface EventApi {
  event_id?: number; // Unique identifier for the event
  description?: string; // Description of the event
  short_description?: string; // Optional short description of the event
  expiration_date?: string; // Optional expiration date of the event
  locations?: EventLocationApi[]; // List of locations associated with the event
  images?: EventImageApi[]; // List of images associated with the event
  created_at?: string; // Creation date of the event
  updated_at?: string; // Last update date of the event
}

export interface Event {
  eventId?: number; // Unique identifier for the event
  description?: string; // Description of the event
  shortDescription?: string; // Optional short description of the event
  expirationDate?: Date; // Optional expiration date of the event
  locations?: EventLocation[]; // List of locations associated with the event
  images?: EventImage[]; // List of images associated with the event
  createdAt?: Date; // Creation date of the event
  updatedAt?: Date; // Last update date of the event
}

export interface EventLocationApi {
  event_id?: number; // Unique identifier for the event
  location_name?: string; // Name of the location
  location_description?: string; // Optional description of the location
  latitude?: number; // Latitude of the location
  longitude?: number; // Longitude of the location
  created_at?: string; // Creation date of the location
  updated_at?: string; // Last update date of the location
}

export interface EventLocation {
  eventId?: number;
  locationName?: string; // Name of the location
  locationDescription?: string; // Optional description of the location
  latitude?: number; // Latitude of the location
  longitude?: number; // Longitude of the location
  createdAt?: Date; // Creation date of the location
  updatedAt?: Date; // Last update date of the location
}

const EventImageType = {
  POSTER: 'poster',
  MOBILE_POSTER: 'poster-mobile'
}

type EventImageType = (typeof EventImageType)[keyof typeof EventImageType];

export interface EventImageApi {
  event_id?: number; // Unique identifier for the event
  image_url?: string; // URL of the image
  image_type: EventImageType; // Type of the image (e.g., poster, mobile poster)
  created_at?: string; // Creation date of the image
  updated_at?: string; // Last update date of the image
}

export interface EventImage {
  eventId?: number; // Unique identifier for the event
  imageUrl?: string; // URL of the image
  imageType: EventImageType; // Type of the image (e.g., poster, mobile poster)
  createdAt?: Date; // Creation date of the image
  updatedAt?: Date; // Last update date of the image
}

/**type Event struct {
  EventId          int64      `json:"event_id"`
  Description      string     `json:"description"`
  ShortDescription *string    `json:"short_description"`
  ExpirationDate   *time.Time `json:"expiration_date"`
  CreatedAt        time.Time  `json:"created_at"`
  UpdatedAt        time.Time  `json:"updated_at"`
}*/
