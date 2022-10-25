<script>

    import { onMount } from 'svelte';
    import Sticky from "./Sticky.svelte";
    import {currentCardID} from "../stores.js";
    import {apiURL} from "../stuff.js";

    export let flashcard = { prompt: "", answer:""};
    export let id;

    onMount(async () => {
        const res = await fetch(
            apiURL + '/cards/get?id=' + id
        )
        const data = await res.json();
        console.log(data);

        if(res.ok){
            flashcard = data;
        }
    })
</script>

{flashcard.prompt} - {flashcard.answer} <a href="/card/{id}/edit" class="font-bold text-blue-400">Edit</a><br/>
