/**
 * Just trying to do some stuff that would block the main thread from running to test out web workers
 * 
 * Eventually want to clean up and have a sorting function or something that is expensive enought and can be replicated in Go, to show off WebAssembly
 * 
 */

const startTime = new Date().toLocaleTimeString();
console.log("PARTYTOWN", "starting at " + startTime);
const data = fetch("https://pokeapi.co/api/v2/pokemon")
  .then((resp) => {
    return resp.json();
  })
  .then((data) => {
    console.log("PARTYTOWN", data.results);
    const results = data.results;

    let duplicated = [];
    let counter = 0;
    while (counter < 3000) {
      duplicated = [...duplicated, ...results];
      counter++;
    }
    return duplicated;
  })
  .then((duplicated) => {
    console.log(
      "DUPLICATED 2 at ",
      new Date().toLocaleTimeString(),
      duplicated
    );
    sortPokemon(duplicated);
  })
  .then(() => {
    const endTime = new Date().toLocaleTimeString();
    console.log("PARTYTOWN", "done at " + endTime);
  });

function sortPokemon(pokemon) {
  return pokemon.sort(function (a, b) {
    if (a.name < b.name) {
      return -1;
    }
    if (a.name > b.name) {
      return 1;
    }
    return 0;
  });
}
