import {currentCardID, currentReviewID} from "./stores.js";

export const apiURL = 'http://localhost:8080'

let reviewID;
let cardID;

currentCardID.subscribe(value => { cardID = value });
currentReviewID.subscribe(value => { reviewID = value });

export async function SubmitAnswer(corr){
    const answer = {
        review: reviewID,
        // todo support modes
        modeID : 1,
        card: cardID,
        correct: corr,
        duration: 1.0
    }

    console.log(answer);

    const res = await fetch(apiURL + '/answer/submit', {
        method: 'POST',
        body: JSON.stringify(answer),
        credentials: "include"
    });

    console.log(res);
}