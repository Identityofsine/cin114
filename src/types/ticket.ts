export interface Ticket {
  ticket_id: number;
  event_id: number;
  event_description: string;
  stripe_receipt_email?: string | null;
  created_at: string;
  updated_at: string;
}

export interface TicketsResponse {
  tickets: Ticket[];
  count: number;
}

export interface TicketResponse {
  ticket: Ticket;
}

export interface TicketExistsResponse {
  has_tickets: boolean;
  ticket_count: number;
} 