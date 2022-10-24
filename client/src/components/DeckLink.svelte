<script>
    import {onMount} from "svelte";
    import {Skeleton} from "flowbite-svelte";
    import {apiURL} from "../stuff.js";

    export let id;
    export let deck = {};

    async function f(){
        const res = await fetch(
            apiURL + '/decks/get?id=' + id
        )
        const data = await res.json();
        console.log(data);

        if(res.ok){
            deck = data;
        }
        return deck;
    }
    let m = f();
    onMount(async () => {

    });
</script>
{#await m}
    {:then m}
    <p class="text-xl"><a href="/deck/{deck.id}"> {deck.name} </a> <br/></p>

{/await}