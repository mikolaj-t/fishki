/** @type {import('./$types').PageLoad} */
import {apiURL} from "../../../stuff.js";

export async function load({params}){
    const res = await fetch(
        apiURL + '/decks/get?id=' + params.id, {
        credentials: 'include'
    });
    const data = await res.json();
    let deck;
    if(res.ok) {
        deck = data;
    }
    return {
        deck: deck
    }
}