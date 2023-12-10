// Utility functions
function unescapeHtmlEntities(text) {
  const element = document.createElement("div");
  element.innerHTML = text;
  return element.textContent;
}

function unescapeHtmlEntitiesToJSON(text) {
  const element = document.createElement("div");
  element.innerHTML = text;

  let tC = element.textContent;
  let j = JSON.stringify(tC);
  return j;
}

function formatDatetime( dateData ) {
  if( !dateData ) return "";
  const dt = new Date( dateData );
  // const date = dt.getDate();
  const month = dt.toLocaleDateString( "default", {month: "long"} );
  const year = dt.getFullYear();
  const formattedDate = `${month} ${year}`;
  return formattedDate;
}