import {writable} from "svelte/store";

export const viewerCardID = writable(0);

export const currentReviewID = writable("");
export const currentCardID = writable("");