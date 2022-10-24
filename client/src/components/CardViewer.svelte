<script>
    import Flashcard from "./Flashcard.svelte";
    import Sticky from "./Sticky.svelte";
    import {currentCardID, viewerCardID} from "../stores.js";
    import {fetchData, idGlob} from "../routes/review/[id]/+page.js";
    import {Button} from "flowbite-svelte";

    export let IDs = ["a", "b", "c"];
    let currentIndex = 0;

    viewerCardID.subscribe(async value => {
        currentIndex = value;
        currentCardID.set(IDs[value]);


        if(currentIndex >= IDs.length){
            console.log("too big index, try to fetch again")
            let newIDs = await fetchData(idGlob);
            console.log("currentIds", IDs, "new IDs", newIDs.cards)
            IDs = newIDs.cards;
            viewerCardID.set(0);
            currentIndex = 0
        }
    });
</script>

<!--<button on:click="{() => nextCard() }">Click me</button>-->
{#if IDs.length > 0}
    {currentIndex + 1} / {IDs.length}
    <Flashcard id={IDs[currentIndex]} sticky="{true}">
</Flashcard>
    {/if}
{#if IDs.length === 0}
    <p class="text-3xl font-bold">That's it for today!</p>
{/if}