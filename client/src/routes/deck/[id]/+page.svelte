<script>
    import {onMount} from "svelte";
    import { page } from '$app/stores';
    import Flashcard from "../../../components/Flashcard.svelte";
    import CardViewer from "../../../components/CardViewer.svelte";
    import FlashcardSimple from "../../../components/FlashcardSimple.svelte";
    import {Button} from "flowbite-svelte";

    ///** @type {import('./$types').PageData} */
    //export let data;

        /** @type {import('./$types').PageData} */
        export let data;
        let cards = data.deck.cards;
        //$: if(data.deck === undefined){ alert("A") }
        console.log(data);
        /*data.deck.cards = []*/
    //let deck = data.deck;

</script>

<p class="font-bold text-4xl mt-6 mb-10"> {data.deck.name}</p>

<p class="font-semibold">Cards: </p> <br/>
{#each cards || [] as card}
    <FlashcardSimple id={card}/>
{/each}
<br/>
<Button gradient color="blue" href="{$page.params.id}/create/card">+ Add card</Button>
<Button gradient color="blue" href="{$page.params.id}/create/review">+ Start review</Button>