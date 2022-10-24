<script>
    import Flashcard from "../../../components/Flashcard.svelte";
    import Sticky from "../../../components/Sticky.svelte";

    /** @type {import('./$types').PageData} */
   export let data;

   export let flashcard = {
       //prompt: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
       prompt: 'Prompt',
       answer: 'Answer'
   }

</script>

<!--<div class="title">
    <p class="font-800">card:</p>
    <p class="font-200">{data.id}</p>
</div>-->

<Flashcard flashcard={flashcard} id={data.id}/>

<!--<Sticky/>-->

<style>
    .title {
        display: flex;
        justify-content: center;
        font-size: 2rem;
        width: 100vw;
        z-index: 10;
    }
</style>