

/**
 * @function toUserTimeZone
 * @description Converts a date object to the user's local time zone.
 * @param {Date} date - The date object to convert.
 * @return {Date} - The date object in the user's local time zone.
 */
export function toUserTimeZone(date: Date): Date {
	return new Date(date)
}

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
		hour12: true,
	};

	return toUserTimeZone(date).toLocaleString('en-US', options);
}
