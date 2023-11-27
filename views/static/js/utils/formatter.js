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