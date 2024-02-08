// Format DDDD, DD MMMM YYYY
export function formatDate(inputDate: Date): string {
  const dt: Date = new Date(inputDate);
  const day: string = dt.toLocaleDateString( 'default', {weekday: 'long'} );
  const date: number = dt.getDate();
  const month: string = dt.toLocaleDateString( 'default', {month: 'long'} );
  const year: number = dt.getFullYear();
  const formattedDate: string = `${day}, ${date} ${month} ${year}`;
  return formattedDate;
}

// Format YYYY-MM-DD
export function getCurrentDate(): string {
  const dt: Date = new Date();
  const year: number = dt.getFullYear();
  const month: string = String(dt.getMonth() + 1).padStart(2, '0');
  const day: string = String(dt.getDate()).padStart(2, '0');
  const formattedDate: string = `${year}-${month}-${day}`;
  return formattedDate;
}

// Format YYYY-MM-DD
export function getOneMonthPastDate(): string {
  const dt: Date = new Date();
  const month: number = dt.getMonth();
  const year: number = dt.getFullYear();
  let pastMonth: number = month - 1;
  let pastYear: number = year;
  if (pastMonth < 0) {
    pastMonth = 11;
    pastYear--;
  }
  const pastMonthDays: number = new Date(pastYear, pastMonth + 1, 0).getDate();
  const currentDay: number = dt.getDate();
  const pastDay: number = currentDay > pastMonthDays ? pastMonthDays : currentDay;
  const pastDate: Date = new Date(pastYear, pastMonth, pastDay);
  const formattedDate: string = pastDate.toISOString().slice(0, 10);
  return formattedDate;
}

// Format YYYY-MM-DD
export function getTwoWeeksPastDate(): string {
  const dt: Date = new Date();
  const pastDate: Date = new Date(dt.getTime() - (14 * 24 * 60 * 60 * 1000));
  const year: number = pastDate.getFullYear();
  const month: string = String(pastDate.getMonth() + 1).padStart(2, '0');
  const day: string = String(pastDate.getDate()).padStart(2, '0');
  const formattedDate: string = `${year}-${month}-${day}`;
  return formattedDate;
}