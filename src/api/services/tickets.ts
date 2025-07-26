import { authAxios } from '../instance/instance';
import { Ticket, TicketsResponse, TicketResponse, TicketExistsResponse } from '@/types/ticket';

// Get all owned tickets
export const getAllTickets = async (): Promise<Ticket[]> => {
  try {
    const response = await authAxios.get<TicketsResponse>('/api/v1/tickets');
    return response.data.tickets;
  } catch (error) {
    console.error('Error fetching all tickets:', error);
    throw error;
  }
};

// Get ticket by ID
export const getTicketById = async (ticketId: string | number): Promise<Ticket> => {
  try {
    const response = await authAxios.get<TicketResponse>(`/api/v1/tickets/${ticketId}`);
    return response.data.ticket;
  } catch (error) {
    console.error('Error fetching ticket by ID:', error);
    throw error;
  }
};

// Get tickets by email
export const getTicketsByEmail = async (email: string): Promise<Ticket[]> => {
  try {
    const response = await authAxios.get<TicketsResponse>(`/api/v1/tickets/email/${email}`);
    return response.data.tickets;
  } catch (error) {
    console.error('Error fetching tickets by email:', error);
    throw error;
  }
};

// Get tickets by event ID
export const getTicketsByEventId = async (eventId: string | number): Promise<Ticket[]> => {
  try {
    const response = await authAxios.get<TicketsResponse>(`/api/v1/tickets/event/${eventId}`);
    return response.data.tickets;
  } catch (error) {
    console.error('Error fetching tickets by event ID:', error);
    throw error;
  }
};

// Check if event has tickets
export const checkEventHasTickets = async (eventId: string | number): Promise<TicketExistsResponse> => {
  try {
    const response = await authAxios.get<TicketExistsResponse>(`/tickets/event/${eventId}/exists`);
    return response.data;
  } catch (error) {
    console.error('Error checking if event has tickets:', error);
    throw error;
  }
};

// Helper function to format ticket ID as T-{number}
export const formatTicketId = (ticketId: number): string => {
  return `T-${ticketId}`;
};

// Helper function to parse formatted ticket ID back to number
export const parseTicketId = (formattedId: string): number | null => {
  const match = formattedId.match(/^T-(\d+)$/);
  return match ? parseInt(match[1], 10) : null;
}; 