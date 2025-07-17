
export function toTimeZone(date: Date, timeZone: string): Date {
  const utcDate = date.toUTCString();
  const localDate = new Date(utcDate + ' GMT' + timeZone);
  return localDate;
}


/** 
 * @function displayDate
 * @description Formats a date object into a human-readable string.
 * @param {Date} date - The date object to format.
 * @return {string} - The formatted date string. in the format of DayName Month Day, Year HH:MM AM/PM 
 */
export function displayDate(date: Date): string {
  const options: Intl.DateTimeFormatOptions = {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: true,
  };

  return date.toLocaleString('en-US', options);
}

//assume the date is in UTC and convert it to New York timezone
const mutateDateIntoETC = (date: Date): Date => {
  const utcDate = date.toISOString();
  // minute off 
  const newYorkOffset = -5 * 60 + 1; // New York is UTC-5
  const localDate = new Date(new Date(utcDate).getTime() + newYorkOffset * 60 * 1000);
  return localDate;
}



export const displayDateForScreening = (date: Date) => {
  date = mutateDateIntoETC(date);
  // Get ordinal suffix for the day
  const getOrdinalSuffix = (day: number): string => {
    if (day > 3 && day < 21) return 'th';
    switch (day % 10) {
      case 1: return 'st';
      case 2: return 'nd';
      case 3: return 'rd';
      default: return 'th';
    }
  };

  const day = date.getDate();
  const suffix = getOrdinalSuffix(day);

  // Format time (without minutes if they're zero)
  const hour = date.getHours() % 12 || 12; // Convert to 12-hour format
  const minutes = date.getMinutes();
  const ampm = date.getHours() >= 12 ? 'PM' : 'AM';
  const time = minutes > 0 ? `${hour}:${minutes.toString().padStart(2, '0')}${ampm}` : `${hour}${ampm}`;

  // Format month
  const month = date.toLocaleString('en-US', { month: 'long' });

  return `${time} ${month} ${day}${suffix}`;
}
