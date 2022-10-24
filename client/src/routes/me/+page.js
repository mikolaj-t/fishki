import {currentReviewID} from "../../stores.js";
import {browser} from "$app/environment";
import {apiURL} from "../../stuff.js";

export const ssr = false;

/** @type {import('./$types').PageLoad} */
export async function load({cookies}){
    //console.log(cookies.get('sessionID'));
    const res = await fetch(
        apiURL + '/users/get?uname=' + window.localStorage.getItem("username"), {
            credentials: 'include',
            method: 'GET'
        }
    );
    //alert('aaa');
    if(!res.ok) return;
    const data = await res.json();
    console.log(data);
    if(res.ok) {

    }
    return {
        user: data
    }
}