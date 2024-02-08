export function formatDate(inputDate: Date): string {
  const dt: Date = new Date(inputDate);

  const day: string = dt.toLocaleDateString( 'default', {weekday: 'long'} );
  const date: number = dt.getDate();
  const month: string = dt.toLocaleDateString( 'default', {month: 'long'} );
  const year: number = dt.getFullYear();

  const formattedDate: string = `${day}, ${date} ${month} ${year}`;

  return formattedDate;
}

export function getCurrentDate(): string {
  const dt: Date = new Date();

  const year: number = dt.getFullYear();
  const month: string = String(dt.getMonth() + 1).padStart(2, '0');
  const day: string = String(dt.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

export function getOneMonthPastDate() {
  const currentDate = new Date();
  const currentMonth = currentDate.getMonth(); // Get the current month (0-11)
  const currentYear = currentDate.getFullYear(); // Get the current year

  // Calculate the month and year one month ago
  let pastMonth = currentMonth - 1;
  let pastYear = currentYear;

  if (pastMonth < 0) {
      pastMonth = 11; // December (0-based index)
      pastYear--; // Year decreases if month goes back from January
  }

  // Calculate the number of days in the past month
  const pastMonthDays = new Date(pastYear, pastMonth + 1, 0).getDate();
  const currentDay = currentDate.getDate();

  // If current day is greater than the last day of the past month, set it to the last day of the past month
  const pastDay = currentDay > pastMonthDays ? pastMonthDays : currentDay;

  // Create a new date using the past year, month, and day
  const pastDate = new Date(pastYear, pastMonth, pastDay);

  return pastDate.toISOString().slice(0, 10); // Return the date in YYYY-MM-DD format
}