<script>
    import {onMount} from "svelte";
    import DeckLink from "../../components/DeckLink.svelte";
    import ReviewLink from "../../components/ReviewLink.svelte";
    import {Button} from "flowbite-svelte";

    /** @type {import('./$types').PageData} */
    export let data;
    onMount(async () => {
    });
</script>

<p class="font-black text-6xl mt-6">
Hi, {window.localStorage.getItem("username")}! 👋
</p>
<br/> <br/>
<div class="columns-1   md:columns-2">
    <div class="break-inside-avoid-column">
        <p class="font-black text-xl">Decks: <br/> </p>
        <Button class="mt-3 mb-3" href="/deck/create" gradient color="blue" >+ New Deck</Button>
        {#each Object.keys(data.user.decks || {}) || [] as deck}
            <DeckLink id={deck}/>
        {/each}
    </div>
    <div>
        <p class="font-bold text-xl">Reviews:</p>
        {#each Object.keys(data.user.reviews || {}) || [] as review}
            <ReviewLink id={review}/>
        {/each}
    </div>
</div>

