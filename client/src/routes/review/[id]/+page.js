/** @type {import('./$types').PageLoad} */
import {currentReviewID} from "../../../stores.js";
import {apiURL} from "../../../stuff.js";

export const ssr = false;

export let idGlob;

/** @type {import('./$types').PageLoad} */
export async function load({params, cookies}){
    return fetchData(params.id)
}

export async function fetchData(id){
    idGlob = id;
//console.log(cookies.get('sessionID'));
    const res = await fetch(
        apiURL + '/reviews/get?id=' + id, {
            credentials: 'include'
        }
    );
    //alert('aaa');
    if(!res.ok) return;
    const data = await res.json();
    let goodCards = []

    if(res.ok) {
        // todo mode checking

        // who in their right mind thought that having a month indexed from 0 is a good idea???
        let dateRef = new Date(2022, 1, 2);
        let currentDay = new Date(Date.now());
        let testDif = differenceInDays(dateRef, new Date(2022, 1, 3))
        let currentDayDif = differenceInDays(dateRef, currentDay);
        Object.entries(data.mode.dates).forEach((entry) => {
            if(entry[1] <= currentDayDif){
                console.log("pushed");
                goodCards.push(entry[0]);
            }
        });
        console.log(goodCards);

        currentReviewID.set(id);
        console.log("set to ", id);

    }
    return {
        cards: shuffle(goodCards)
    }
}

function shuffle(array) {
    let currentIndex = array.length,  randomIndex;

    // While there remain elements to shuffle.
    while (currentIndex != 0) {

        // Pick a remaining element.
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex--;

        // And swap it with the current element.
        [array[currentIndex], array[randomIndex]] = [
            array[randomIndex], array[currentIndex]];
    }

    return array;
}

export function differenceInDays(date1, date2){
    let  dateDif = date2 - date1;
    let dayDif = Math.floor(dateDif / (1000 * 3600 * 24));
    //console.log(date1, "->", date2, dayDif);
    return dayDif;
}